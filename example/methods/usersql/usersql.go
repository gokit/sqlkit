package usersql

import (
	"context"

	"fmt"

	"errors"

	"time"

	"strconv"

	"strings"

	"encoding/json"

	dsql "database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/influx6/faux/metrics"

	"github.com/influx6/faux/db/sql/tables"

	"github.com/gokit/sqlkit/example/methods"
)

// errors ...
var (
	ErrNotFound             = errors.New("record not found")
	ErrExpiredContext       = errors.New("context has expired")
	ErrExpectedFieldType    = errors.New("only support types matching UserFields interface")
	ErrExpectedConsumerType = errors.New("only support types matching UserConsumer interface")
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
func NewDB(config Config) SqlDB {
	return sqlDB{
		config: config,
	}
}

// sqlDB defines a structure which returns a new db connection for
// use in creating new sql db instances for db ops.
type sqlDB struct {
	config Config
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
		return nil, err
	}

	return db, nil
}

//**********************************************************
// SqlDB Types
//**********************************************************

// Validation defines an interface which expose a method to validate a giving type.
type Validation interface {
	Validate() error
}

// UserFields defines an interface which exposes method to return a map of all
// attributes associated with the defined structure as decided by the structure.
type UserFields interface {
	Fields() (map[string]interface{}, error)
}

// UserConsumer defines an interface which accepts a map of data which will be consumed
// into the giving implementing structure as decided by the structure.
type UserConsumer interface {
	Consume(map[string]interface{}) error
}

//**********************************************************
// SqlDB Methods
//**********************************************************

// contains templates of sql statement for use in operations.
const (
	countTemplate            = "SELECT count(*) FROM %s"
	selectAllTemplate        = "SELECT * FROM %s"
	selectAllByOrderTemplate = "SELECT * FROM %s ORDER BY %s %s"
	selectLimitedTemplate    = "SELECT * FROM %s ORDER BY %s %s LIMIT %d OFFSET %d"
	selectItemTemplate       = "SELECT * FROM %s WHERE %s=%s"
	insertTemplate           = "INSERT INTO %s %s VALUES %s"
	updateTemplate           = "UPDATE %s SET %s WHERE %s=%s"
	deleteTemplate           = "DELETE FROM %s WHERE %s=%s"
)

// MigrateTables takes the giving TableMigration data and generates necessary queries which
// then gets executed within database.
func MigrateTables(ctx context.Context, m metrics.Metrics, dbx SqlDB, migrations ...tables.TableMigration) error {
	defer m.CollectMetrics("DB.MigrateTables")

	if len(migrations) == 0 {
		return nil
	}

	if isContextExpired(ctx) {
		m.Emit(metrics.Error(ErrExpiredContext))
		return ErrExpiredContext
	}

	dbi, err := dbx.New()
	if err != nil {
		m.Emit(metrics.Error(err))
		return err
	}

	defer dbi.Close()

	for _, table := range migrations {
		m.Emit(metrics.Info("Executing Migration"), metrics.WithFields(metrics.Field{
			"query": table.String(),
			"table": table.TableName,
		}))

		if _, err := dbi.Exec(table.String()); err != nil {
			m.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{"query": table.String(), "table": table.TableName}))
			return err
		}
	}

	return nil
}

// Create takes the giving table name with the giving fields and attempts to save this giving
// data appropriately into the giving db.
func Create(ctx context.Context, m metrics.Metrics, dbx SqlDB, table string, elem methods.User) error {
	defer m.CollectMetrics("DB.Create")
	defer m.Emit(metrics.Info("Create User in db"), metrics.With("table", table))

	if isContextExpired(ctx) {
		m.Emit(metrics.Error(ErrExpiredContext))
		return ErrExpiredContext
	}

	if validator, ok := interface{}(elem).(Validation); ok {
		if err := validator.Validate(); err != nil {
			m.Emit(metrics.Errorf("Failed validate record"),
				metrics.With("table", table),
				metrics.With("error", err.Error()))
			return err
		}
	}

	if _, ok := interface{}(elem).(UserFields); !ok {
		m.Emit(metrics.Error(ErrExpectedFieldType))
		return ErrExpectedFieldType
	}

	m.Emit(metrics.Info("Create User record"),
		metrics.With("table", table),
		metrics.With("elem", elem))

	fields, err := elem.Fields()
	if err != nil {
		m.Emit(metrics.Error(err), metrics.With("table", table), metrics.With("elem", elem))
		return err
	}

	fieldNames := fieldNames(fields)
	values := fieldValues(fieldNames, fields)

	query := fmt.Sprintf(insertTemplate, table, fieldNameMarkers(fieldNames), fieldMarkers(len(fieldNames)))
	m.Emit(metrics.Info("DB:"), metrics.With("query", query))

	db, err := dbx.New()
	if err != nil {
		m.Emit(metrics.Error(err), metrics.With("table", table), metrics.With("elem", elem))
		return err
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		m.Emit(metrics.Error(err), metrics.With("table", table), metrics.With("elem", elem))
		return err
	}

	if _, err := db.Exec(query, values...); err != nil {
		m.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"err":   err,
			"query": query,
			"table": table,
			"elem":  elem,
		}))
		return err
	}

	if err := tx.Commit(); err != nil {
		m.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"err":   err,
			"query": query,
			"table": table,
			"elem":  elem,
		}))
		return err
	}

	return nil
}

// Update attempts to update giving record by using provided id provided.
func Update(ctx context.Context, m metrics.Metrics, dbx SqlDB, table string, id string, elem methods.User) error {
	defer m.CollectMetrics("DB.Update")

	return UpdateByField(ctx, m, dbx, table, "public_id", id, elem)
}

// UpdateByField takes the giving table name with the giving fields and attempts to update this giving
// data appropriately into the giving db.
// index - defines the string which should identify the key to be retrieved from the fields to target the
// data to be updated in the db.
func UpdateByField(ctx context.Context, m metrics.Metrics, dbx SqlDB, table string, index string, indexValue interface{}, elem methods.User) error {
	defer m.CollectMetrics("DB.UpdateByField")
	defer m.Emit(metrics.Info("Update User in DB"), metrics.With("table", table))

	if isContextExpired(ctx) {
		m.Emit(metrics.Error(ErrExpiredContext))
		return ErrExpiredContext
	}

	if validator, ok := interface{}(elem).(Validation); ok {
		if err := validator.Validate(); err != nil {
			m.Emit(metrics.Errorf("Failed validate record"),
				metrics.With("table", table),
				metrics.With("error", err.Error()))
			return err
		}
	}

	if _, ok := interface{}(elem).(UserFields); !ok {
		m.Emit(metrics.Error(ErrExpectedFieldType))
		return ErrExpectedFieldType
	}

	m.Emit(metrics.Info("Update User record"),
		metrics.With("table", table), metrics.With("elem", elem))

	fields, err := elem.Fields()
	if err != nil {
		m.Emit(metrics.Error(err), metrics.With("table", table))
		return err
	}

	indexValueString, err := ToValueString(indexValue)
	if err != nil {
		m.Emit(metrics.Error(err), metrics.With("table", table))
		return err
	}

	sets, err := setValues(fields)
	if err != nil {
		m.Emit(metrics.Error(err), metrics.With("table", table))
		return err
	}

	db, err := dbx.New()
	if err != nil {
		m.Emit(metrics.Error(err), metrics.With("table", table))
		return err
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		m.Emit(metrics.Error(err), metrics.With("table", table))
		return err
	}

	query := fmt.Sprintf(updateTemplate, table, sets, index, indexValueString)
	m.Emit(metrics.Info("DB:"), metrics.With("query", query))

	if _, err := db.Exec(query); err != nil {
		m.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"err":   err,
			"query": query,
			"table": table,
			"elem":  elem,
		}))
		return err
	}

	if err := tx.Commit(); err != nil {
		m.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"err":   err,
			"query": query,
			"table": table,
			"elem":  elem,
		}))
		return err
	}

	return nil
}

// GetAll retrieves the all records from the specific db table.
func GetAll(ctx context.Context, m metrics.Metrics, dbx SqlDB, table string) ([]methods.User, error) {
	defer m.CollectMetrics("DB.GetAll")
	defer m.Emit(metrics.Info("Retrieve all records from DB"), metrics.With("table", table))

	if _, ok := (interface{}(&methods.User{})).(UserConsumer); !ok {
		m.Emit(metrics.Error(ErrExpectedConsumerType))
		return nil, ErrExpectedConsumerType
	}

	if isContextExpired(ctx) {
		m.Emit(metrics.Error(ErrExpiredContext))
		return nil, ErrExpiredContext
	}

	db, err := dbx.New()
	if err != nil {
		m.Emit(metrics.Error(err))
		return nil, err
	}

	defer db.Close()

	query := fmt.Sprintf(selectAllTemplate, table)
	m.Emit(metrics.Info("DB:GetAll"), metrics.With("query", query))

	rows, err := db.Queryx(query)
	if err != nil {
		m.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"err":   err,
			"query": query,
			"table": table,
		}))
		return nil, err
	}

	var records []methods.User

	for rows.Next() {
		mo := make(map[string]interface{})
		if err := rows.MapScan(mo); err != nil {
			m.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
				"err":   err,
				"query": query,
				"table": table,
			}))

			return nil, err
		}

		var record methods.User
		if err := record.Consume(naturalizeMap(mo)); err != nil {
			m.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
				"err":   err,
				"query": query,
				"table": table,
			}))

			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}

// GetAllPerPage retrieves the giving data from the specific db table with the specific index and value.
func GetAllPerPage(ctx context.Context, m metrics.Metrics, dbx SqlDB, table string, order string, orderBy string, page int, responsePerPage int) ([]methods.User, int, error) {
	defer m.CollectMetrics("DB.GetAllPerPage")

	defer m.Emit(metrics.Info("Retrieve all User records from DB"),
		metrics.With("table", table), metrics.WithFields(metrics.Field{
			"page":            page,
			"order":           order,
			"orderBy":         orderBy,
			"responsePerPage": responsePerPage,
		}))

	if _, ok := (interface{}(&methods.User{})).(UserConsumer); !ok {
		m.Emit(metrics.Error(ErrExpectedConsumerType))
		return nil, -1, ErrExpectedConsumerType
	}

	if isContextExpired(ctx) {
		m.Emit(metrics.Error(ErrExpiredContext))
		return nil, -1, ErrExpiredContext
	}

	db, err := dbx.New()
	if err != nil {
		m.Emit(metrics.Error(err))
		return nil, -1, err
	}

	defer db.Close()

	if page <= 0 && responsePerPage <= 0 {
		records, err := GetAllByOrder(ctx, m, dbx, table, order, orderBy)
		if err != nil {
			m.Emit(metrics.Error(err))
		}
		return records, len(records), err
	}

	// Get total number of records.
	totalRecords, err := Count(ctx, m, dbx, table)
	if err != nil {
		m.Emit(metrics.Error(err))
		return nil, -1, err
	}

	var totalWanted, indexToStart int

	if page <= 1 && responsePerPage > 0 {
		totalWanted = responsePerPage
		indexToStart = 0
	} else {
		totalWanted = responsePerPage * page
		indexToStart = totalWanted / 2

		if page > 1 {
			indexToStart++
		}
	}

	m.Emit(metrics.Info("DB:GetAllPerPage"), metrics.WithFields(metrics.Field{
		"starting_index":       indexToStart,
		"total_records_wanted": totalWanted,
		"order":                order,
		"page":                 page,
		"responsePerPage":      responsePerPage,
	}))

	// If we are passed the total, just return nil records and total without error.
	if indexToStart > totalRecords {
		return nil, totalRecords, nil
	}

	query := fmt.Sprintf(selectLimitedTemplate, table, orderBy, order, totalWanted, indexToStart)
	m.Emit(metrics.Info("DB:GetAllPerPage"), metrics.With("query", query))

	rows, err := db.Queryx(query)
	if err != nil {
		m.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"err":   err,
			"query": query,
			"table": table,
		}))
		return nil, -1, err
	}

	var records []methods.User

	for rows.Next() {
		mo := make(map[string]interface{})
		if err := rows.MapScan(mo); err != nil {
			m.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
				"err":   err,
				"query": query,
				"table": table,
			}))

			return nil, -1, err
		}

		var record methods.User
		if err := record.Consume(naturalizeMap(mo)); err != nil {
			m.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
				"err":   err,
				"query": query,
				"table": table,
			}))

			return nil, -1, err
		}

		records = append(records, record)
	}

	return records, totalRecords, nil
}

// GetAllByOrder retrieves the all data from the db with the specific order based on giving key in orderBy.
func GetAllByOrder(ctx context.Context, m metrics.Metrics, dbx SqlDB, table string, order string, orderBy string) ([]methods.User, error) {
	defer m.CollectMetrics("DB.GetAllByOrder")
	defer m.Emit(metrics.Info("Retrieve all records from DB"), metrics.With("table", table))

	if _, ok := (interface{}(&methods.User{})).(UserConsumer); !ok {
		m.Emit(metrics.Error(ErrExpectedConsumerType))
		return nil, ErrExpectedConsumerType
	}

	if isContextExpired(ctx) {
		m.Emit(metrics.Error(ErrExpiredContext))
		return nil, ErrExpiredContext
	}

	db, err := dbx.New()
	if err != nil {
		m.Emit(metrics.Error(err))
		return nil, err
	}

	defer db.Close()

	switch strings.ToLower(order) {
	case "asc":
		order = "ASC"
	case "dsc", "desc":
		order = "DESC"
	default:
		order = "ASC"
	}

	query := fmt.Sprintf(selectAllByOrderTemplate, table, orderBy, order)
	m.Emit(metrics.Info("DB:GetAll"), metrics.With("query", query))

	rows, err := db.Queryx(query)
	if err != nil {
		m.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"err":   err,
			"query": query,
			"table": table,
		}))
		return nil, err
	}

	var records []methods.User

	for rows.Next() {
		mo := make(map[string]interface{})
		if err := rows.MapScan(mo); err != nil {
			m.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
				"err":   err,
				"query": query,
				"table": table,
			}))

			return nil, err
		}

		var record methods.User
		if err := record.Consume(naturalizeMap(mo)); err != nil {
			m.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
				"err":   err,
				"query": query,
				"table": table,
			}))

			return nil, err
		}

		records = append(records, record)
	}

	return records, nil
}

// Get returns giving record by retrieving based on it's public_id.
func Get(ctx context.Context, m metrics.Metrics, dbx SqlDB, table string, id string) (methods.User, error) {
	defer m.CollectMetrics("DB.Get")

	return GetByField(ctx, m, dbx, table, "public_id", id)
}

// GetByField retrieves the giving data from the specific db with the specific index and value.
func GetByField(ctx context.Context, m metrics.Metrics, dbx SqlDB, table string, index string, indexValue interface{}) (methods.User, error) {
	defer m.CollectMetrics("DB.GetByField")
	defer m.Emit(metrics.Info("Get record from DB"), metrics.WithFields(metrics.Field{
		"table":      table,
		"index":      index,
		"indexValue": indexValue,
	}))

	if _, ok := (interface{}(&methods.User{})).(UserConsumer); !ok {
		m.Emit(metrics.Error(ErrExpectedConsumerType))
		return methods.User{}, ErrExpectedConsumerType
	}

	if isContextExpired(ctx) {
		m.Emit(metrics.Error(ErrExpiredContext))
		return methods.User{}, ErrExpiredContext
	}

	db, err := dbx.New()
	if err != nil {
		m.Emit(metrics.Error(err))
		return methods.User{}, err
	}

	defer db.Close()

	indexValueString, err := ToValueString(indexValue)
	if err != nil {
		m.Emit(metrics.Errorf("Get User record"), metrics.WithFields(metrics.Field{
			"err":   err,
			"table": table,
		}))
		return methods.User{}, err
	}

	query := fmt.Sprintf(selectItemTemplate, table, index, indexValueString)
	m.Emit(metrics.Info("DB:"), metrics.With("query", query))

	row := db.QueryRowx(query)
	if err := row.Err(); err != nil {
		m.Emit(metrics.Errorf("Get User record"), metrics.WithFields(metrics.Field{
			"err":   err,
			"query": query,
			"table": table,
		}))
		return methods.User{}, err
	}

	mo := make(map[string]interface{})

	if err := row.MapScan(mo); err != nil {
		m.Emit(metrics.Errorf("Get User record"), metrics.WithFields(metrics.Field{
			"err":   err,
			"query": query,
			"table": table,
		}))

		return methods.User{}, err
	}

	m.Emit(metrics.Info("Consumer:Get:Fields"), metrics.WithFields(metrics.Field{
		"table":    table,
		"response": mo,
	}))

	var record methods.User
	if err := record.Consume(naturalizeMap(mo)); err != nil {
		m.Emit(metrics.Errorf("Get User record"), metrics.WithFields(metrics.Field{
			"err":   err,
			"query": query,
			"table": table,
		}))

		return methods.User{}, err
	}

	return record, nil
}

// Count retrieves the total number of records from the specific table from the db.
func Count(ctx context.Context, m metrics.Metrics, dbx SqlDB, table string) (int, error) {
	defer m.CollectMetrics("DB.Count")
	defer m.Emit(metrics.Info("Count record from DB"), metrics.WithFields(metrics.Field{
		"table": table,
	}))

	if isContextExpired(ctx) {
		m.Emit(metrics.Error(ErrExpiredContext))
		return 0, ErrExpiredContext
	}

	db, err := dbx.New()
	if err != nil {
		m.Emit(metrics.Error(err))
		return 0, err
	}

	defer db.Close()

	var records int

	query := fmt.Sprintf(countTemplate, table)
	m.Emit(metrics.Info("DB:"), metrics.With("query", query))

	if err := db.Get(&records, query); err != nil {
		m.Emit(metrics.Errorf("DB:"), metrics.WithFields(metrics.Field{
			"err":   err,
			"query": query,
			"table": table,
		}))
		return 0, err
	}

	return records, nil
}

// Delete removes the giving data from the specific db with the specific PublicID of record.
func Delete(ctx context.Context, m metrics.Metrics, dbx SqlDB, table string, id string) error {
	defer m.CollectMetrics("DB.Delete")

	return DeleteByField(ctx, m, dbx, table, "public_id", id)
}

// DeleteByField removes the giving data from the specific db with the specific index and value.
func DeleteByField(ctx context.Context, m metrics.Metrics, dbx SqlDB, table string, index string, indexValue interface{}) error {
	defer m.CollectMetrics("DB.DeleteByField")
	defer m.Emit(metrics.Info("Delete record from DB"), metrics.WithFields(metrics.Field{
		"table":      table,
		"index":      index,
		"indexValue": indexValue,
	}))

	if isContextExpired(ctx) {
		m.Emit(metrics.Error(ErrExpiredContext))
		return ErrExpiredContext
	}

	db, err := dbx.New()
	if err != nil {
		m.Emit(metrics.Error(err))
		return err
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		m.Emit(metrics.Error(err))
		return err
	}

	indexValueString, err := ToValueString(indexValue)
	if err != nil {
		m.Emit(metrics.Errorf("DB: %+q", err), metrics.WithFields(metrics.Field{
			"err":   err,
			"table": table,
		}))
		return err
	}

	query := fmt.Sprintf(deleteTemplate, table, index, indexValueString)
	m.Emit(metrics.Info("DB:"), metrics.With("query", query))

	if _, err := db.Exec(query); err != nil {
		m.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"err":   err,
			"query": query,
			"table": table,
		}))
		return err
	}

	if err := tx.Commit(); err != nil {
		m.Emit(metrics.Error(err), metrics.WithFields(metrics.Field{
			"err":   err,
			"query": query,
			"table": table,
		}))
		return err
	}

	return nil
}

// Exec provides a function which allows the execution of a custom function against a sql db.
func Exec(ctx context.Context, m metrics.Metrics, dbx SqlDB, fx func(*sqlx.DB) error) error {
	defer m.CollectMetrics("DB.Exec")

	if isContextExpired(ctx) {
		err := ErrExpiredContext
		m.Emit(metrics.Errorf("Failed to execute operation"), metrics.With("error", err.Error()))
		return err
	}

	sx, err := dbx.New()
	if err != nil {
		m.Emit(metrics.Errorf("Failed to execute operation"), metrics.With("error", err.Error()))
		if err == dsql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	defer sx.Close()

	if err := fx(sx); err != nil {
		m.Emit(metrics.Errorf("Failed to execute operation"),
			metrics.With("error", err.Error()))
		if err == dsql.ErrNoRows {
			return ErrNotFound
		}
		return err
	}

	m.Emit(metrics.Info("Operation executed"))

	return nil
}

//**********************************************************
// SqlDB internal functions
//**********************************************************

// fieldMarkers returns a (?,...,>) string which represents
// all filedNames extrated from the provided TableField.
func fieldMarkers(total int) string {
	var markers []string

	for i := 0; i < total; i++ {
		markers = append(markers, "?")
	}

	return "(" + strings.Join(markers, ",") + ")"
}

// fieldNameMarkers returns a (fieldName,...,fieldName) string which represents
// all filedNames extrated from the provided TableField.
func fieldNameMarkers(fields []string) string {
	return "(" + strings.Join(fields, ", ") + ")"
}

// fieldValues returns a (fieldName,...,fieldName) string which represents
// all filedNames extrated from the provided TableField.
func fieldValues(names []string, fields map[string]interface{}) []interface{} {
	var vals []interface{}

	for _, name := range names {
		vals = append(vals, fields[name])
	}

	return vals
}

func setValues(fields map[string]interface{}) (string, error) {
	var vals []string

	for name, val := range fields {
		rv, err := ToValueString(val)
		if err != nil {
			return "", err
		}

		vals = append(vals, fmt.Sprintf("%s=%s", name, rv))
	}

	return strings.Join(vals, ","), nil
}

// naturalizeMap returns a new map where all values of []bytes are converted to strings
func naturalizeMap(fields map[string]interface{}) map[string]interface{} {
	nz := map[string]interface{}{}

	for key, val := range fields {
		switch rv := val.(type) {
		case []byte:
			nz[key] = string(rv)
			continue
		default:
			nz[key] = val
			continue
		}
	}

	return nz
}

// fieldNamesFromMap returns a (fieldName,...,fieldName) string which represents
// all filedNames extrated from the provided TableField.
func fieldNamesFromMap(fields map[string]interface{}) []string {
	var names []string

	for key := range fields {
		names = append(names, key)
	}

	return names
}

// fieldNames returns a (fieldName,...,fieldName) string which represents
// all filedNames extrated from the provided TableField.
func fieldNames(fields map[string]interface{}) []string {
	var names []string

	for key := range fields {
		names = append(names, key)
	}

	return names
}

// ToValueString returns the string representation of a basic go core data type for usage in
// a db call.
func ToValueString(val interface{}) (string, error) {
	switch bo := val.(type) {
	case *time.Time:
		return bo.String(), nil
	case time.Time:
		return bo.String(), nil
	case error:
		return strconv.Quote(bo.Error()), nil
	case string:
		return strconv.Quote(bo), nil
	case int:
		return strconv.Itoa(bo), nil
	case []byte:
		return strconv.Quote(string(bo)), nil
	case int64:
		return strconv.Itoa(int(bo)), nil
	case rune:
		return strconv.QuoteRune(bo), nil
	case bool:
		return strconv.FormatBool(bo), nil
	case byte:
		return strconv.QuoteRune(rune(bo)), nil
	case float64:
		return strconv.FormatFloat(bo, 'f', 4, 64), nil
	case float32:
		return strconv.FormatFloat(float64(bo), 'f', 4, 64), nil
	}

	data, err := json.Marshal(val)
	return string(data), err
}

func isContextExpired(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
