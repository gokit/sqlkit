package sqldb

import (
	"context"

	"fmt"

	"errors"

	dsql "database/sql"

	"github.com/influx6/faux/metrics"

	"github.com/jmoiron/sqlx"
)

// errors ...
var (
	ErrNotFound       = errors.New("record not found")
	ErrExpiredContext = errors.New("context has expired")
)

//**********************************************************
// SqlDB Config and Setup
//**********************************************************

// Config is a configuration struct for the DB connection for DBMaker.
type Config struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Port     string `json:"port"`
	Host     string `json:"ip"`
	DB       string `json:"name"`
	Driver   string `json:"driver"`
}

// SqlDB defines an interface which exposes a method to return a new
// underline sql.Database.
type SqlDB interface {
	New() (*sqlx.DB, error)
}

// NewDB returns a new instance of a DB.
func NewDB(config Config, metrics metrics.Metrics) SqlDB {
	return &sqlDB{
		config: config,
		log:    metrics,
	}
}

// sqlDB defines a structure which returns a new db connection for
// use in creating new sql db instances for db ops.
type sqlDB struct {
	config Config
	log    metrics.Metrics
}

// New returns a new instance of a sqlx.DB connected to the db with the provided
// credentials pulled from the host environment.
func (dl sqlDB) New() (*sqlx.DB, error) {
	if dl.config.Host == "" {
		dl.config.Host = "0.0.0.0"
	}

	addr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dl.config.User, dl.config.Password, dl.config.Host, dl.config.Port, dl.config.DB)
	db, err := sqlx.Connect(dl.config.Driver, addr)
	if err != nil {
		dl.log.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"ip":     dl.config.Host,
			"port":   dl.config.Port,
			"dbName": dl.config.DB,
			"driver": dl.config.Driver,
		}))

		return nil, err
	}

	return db, nil
}

//**********************************************************
// SqlDB Functions
//**********************************************************

// Exec provides a function which allows the execution of a custom function against a sql db.
func Exec(ctx context.Context, m metrics.Metrics, db SqlDB, fx func(*sqlx.DB) error) error {
	defer m.CollectMetrics("DB.Exec")

	if isContextExpired(ctx) {
		err := ErrExpiredContext
		m.Emit(metrics.Errorf("Failed to execute operation"), metrics.With("error", err.Error()))
		return err
	}

	sx, err := db.New()
	if err != nil {
		m.Emit(metrics.Errorf("Failed to execute operation"), metrics.With("error", err.Error()))
		if err == dsql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	defer sx.Close()

	if err := fx(sx); err != nil {
		m.Emit(metrics.Errorf("Failed to execute operation"), metrics.With("error", err.Error()))
		if err == dsql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	m.Emit(metrics.Info("Operation executed"))

	return nil
}

func isContextExpired(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
