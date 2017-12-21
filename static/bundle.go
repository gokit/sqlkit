package static

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

type fileData struct {
	path string
	root string
	data []byte
}

var (
	assets = map[string][]string{

		".tml": []string{ // all .tml assets.

			"sql-api-backend.tml",

			"sql-api-json.tml",

			"sql-api-migration.tml",

			"sql-api-readme.tml",

			"sql-api-test.tml",

			"sql-api.tml",
		},
	}

	assetFiles = map[string]fileData{

		"sql-api-backend.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xc1\x6e\xa3\x30\x10\x86\xef\x3c\xc5\x1c\x13\x09\xc1\x2b\xec\x12\xb4\xd1\x5e\x36\x91\x56\x3d\x55\x3d\x18\xf3\x07\xdc\x18\xdb\xb2\x27\x69\x50\xc4\xbb\x57\x01\x94\xa2\x2a\x54\x4d\x7a\x42\x0c\x33\xf3\x7d\xff\x88\x34\xa5\xf3\x39\xf9\xcf\xfe\x20\x39\xd9\x14\xaf\x90\x9c\xfc\x13\x0d\xba\x2e\xcf\x32\x21\xf7\x30\x25\x95\xd8\x29\x83\x40\x82\x8a\xb1\xf2\x56\x2b\x59\x93\x87\xf3\x08\x30\x1c\x88\x6b\x50\xa5\x8e\xca\x54\x51\x9a\x52\x03\xae\x6d\x19\x08\x27\x67\x03\x4a\x2a\xda\xbe\x21\xcf\x48\x35\x4e\xa3\x81\x61\xc1\xca\x1a\xda\x59\x3f\x19\x25\x6e\x1d\xe6\x74\x92\xcb\xe2\x5f\xd7\xf9\xe8\xab\xde\x0f\x75\x65\x18\x7e\x27\x24\xce\x11\x51\x0e\x0d\xc6\x42\xf2\x89\xa4\x35\x8c\x13\x27\xab\xe1\x19\x93\x3b\x14\x5a\xc9\xbf\x39\x05\xf6\xca\x54\x4b\x82\xf7\xd6\x47\x44\x2b\x0f\x31\x37\x04\x8d\x66\xe2\xb0\x15\x72\x2f\xaa\x8b\xeb\x8c\xd7\xb8\x95\x22\xa2\x35\xf8\x9b\x22\xb4\xb8\x83\x10\xd3\x80\x58\x46\x44\x4f\xae\x9c\x35\xff\x44\x79\x38\xca\x90\xe4\xb7\xd6\x59\xbb\xf1\x25\xfc\x6d\x9a\xbd\x7c\xba\xa2\xfa\xb7\xac\x9d\x04\x7c\x7e\x79\x30\xe2\x1a\x9c\xb5\x7f\x14\x74\x79\x1b\xbc\x47\x7b\xc5\x1e\x85\x3e\x60\xf2\x43\x74\x3f\x38\xed\x90\xf9\xfe\xb0\x31\x39\x51\xf5\x12\x31\x79\x04\x67\x4d\xc0\x16\x7e\x3b\x16\x1f\xb9\x45\xbf\x6b\x14\xeb\xde\x03\x00\x00\xff\xff\xe4\xd6\x4e\xc2\xce\x03\x00\x00"),
			path: "sql-api-backend.tml",
			root: "sql-api-backend.tml",
		},

		"sql-api-json.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\xcf\x4a\xfc\x30\x10\x80\xcf\x0d\xe4\x1d\xe6\xd7\xc3\x8f\x16\x4a\xf6\xae\xf4\x05\x44\x76\xc5\xc5\x93\x08\x3b\x9b\x9d\xae\x5d\xdb\x89\x26\xa9\x7f\x08\x79\x77\x49\xba\x82\x07\x41\x7b\xc9\x61\x48\xbe\xef\x63\x22\xc5\x6a\x05\x27\x67\x18\xba\xfe\xdd\x4f\x96\x1c\x28\xa5\xa4\x78\x45\x0b\x95\x14\x10\x82\xda\x7a\x3b\x69\xaf\x36\xfb\x13\x69\xaf\xd6\x38\x52\x3e\x62\xbc\xda\x6e\xd6\xd0\xc2\x2e\x04\x18\xf1\xf9\x16\xf9\x60\xc6\x3c\x3b\x3f\x81\xd2\xbd\x0c\x25\x94\x09\x5f\x42\x8c\x3b\x29\x6a\x29\xb2\xf1\xda\xe0\xe1\x57\xb4\x25\x3f\x59\x76\x80\xc0\xf4\x06\x3d\x3b\x8f\xac\x09\x4c\x07\xf8\xad\xeb\x06\xf5\x13\x1e\x29\x46\xf5\x23\x30\x46\x25\x45\x37\xb1\xfe\x93\xb3\xd2\x86\x3d\xb1\x07\xe7\x6d\xcf\xc7\x1a\xaa\x05\xa2\x06\xc8\x5a\x63\x6b\x08\x52\x14\x69\x81\x34\xd0\xb8\xa4\x34\x2d\xa7\xe8\xbb\x84\x81\x8b\x36\x7f\x8b\xba\xe3\x11\xad\x7b\xc4\xa1\xba\x7f\xd8\x7f\x78\xfa\x2a\xac\x1b\xf8\x9f\xf8\xf5\x65\xbe\xfe\xaf\x05\xee\x87\x6c\x2e\xe6\xbd\x2d\x11\x87\xb9\x5d\x8a\x62\x6e\x38\x13\x12\xbf\x49\x5c\x29\xf2\xfc\x33\x00\x00\xff\xff\xff\x44\x96\x9f\x2d\x02\x00\x00"),
			path: "sql-api-json.tml",
			root: "sql-api-json.tml",
		},

		"sql-api-migration.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xcf\x4b\xfb\x40\x10\xc5\xef\xfd\x2b\x1e\x3d\x7d\xbf\x50\xd3\x7b\xc0\x43\xd5\x0a\x42\xb5\xa0\xf1\x24\x22\xd3\xec\xb4\x59\xe9\xfe\x70\x77\x6a\x94\x90\xff\x5d\x76\x25\xc9\x41\xcd\x69\xb2\xf3\x3e\x6f\xde\x7b\xa7\x80\x7f\x33\x00\x58\x2e\xd1\x75\xc5\x83\x84\x53\x2d\xc5\x76\xf7\xca\xb5\x14\x77\x64\xb8\xef\x2b\xda\x1d\x19\x8a\xf7\xda\x72\x84\x34\x0c\xc9\x2f\x46\x1f\x02\x89\x76\x16\x86\x85\xce\x14\x09\xa1\x6d\x74\xdd\x60\x30\xd4\x11\xa7\xc8\x0a\xe2\x70\x60\xcb\x81\x84\x33\x4f\xde\x07\xe7\x83\x4e\xff\x97\xf7\xeb\x55\xb5\x46\xb5\xba\xd8\xac\x11\xdf\x8e\x88\x42\xc2\x86\xad\x60\xef\xc2\xc0\x69\x7b\x18\x5d\x29\x46\x57\x27\x56\x65\x7d\x0e\x53\x0c\xcb\x6a\x7b\xb5\x2d\xf1\xe8\x55\xf2\x6e\xb5\x34\xd8\x6b\x3e\x2a\x58\x32\x0c\xf9\xf4\x1c\x41\x56\x41\x5b\xc5\x1f\xdf\xd0\x1f\xad\x6f\xc7\x76\xe7\x18\xe7\x2e\x13\xe9\xab\xb4\xe1\x28\x64\x3c\xab\x12\x12\x4e\xbc\xc0\xb4\x4b\x89\x92\x4b\x89\x79\xd7\x89\xdb\xb8\x96\x03\x7e\xbd\xf2\x92\xd3\xcf\x17\x23\x7b\x93\x82\x71\x2c\xf1\xf4\x9c\xc7\xe9\x74\x3f\x89\xae\x53\xa5\xac\xc9\xd3\x0f\x4d\x3f\xfb\x3f\xfb\x0a\x00\x00\xff\xff\xde\xf6\xc3\xf5\xda\x01\x00\x00"),
			path: "sql-api-migration.tml",
			root: "sql-api-migration.tml",
		},

		"sql-api-readme.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\x4f\x6f\xda\x40\x10\xc5\xef\x48\x7c\x87\xa9\xb8\x98\x8a\x2e\xf7\x48\x3d\x00\x6e\xa3\x48\x51\x4b\x42\x73\x42\x91\xbc\xac\x07\x7b\xcb\x7a\xc7\xdd\x1d\xb7\x8e\x90\xbf\x7b\xe5\x3f\x29\x90\x3a\x4d\x50\xe3\xcb\xee\x3e\x79\xe6\xfd\xe6\xcd\x7e\x2f\x56\xec\x0a\xc5\xe2\xeb\xe6\x3b\x2a\x16\x5f\x64\x86\x55\x05\xab\x9b\xeb\x70\x0e\xb3\xe5\xd5\x70\xf0\xf1\xe5\x6f\x38\x18\x0e\xd6\xef\xd6\x97\x04\xb7\x98\x93\x63\x58\x48\x17\xdf\x07\x29\x73\xee\x2f\xa6\xd3\x84\x5c\x23\x2b\xe9\x62\xa1\x28\x9b\x6e\x64\x9c\xe0\x74\xbf\x17\x4b\xa9\x76\x32\xc1\xa5\xe4\xb4\xaa\xc6\xff\xa8\x68\x9f\x7f\x97\xd4\xce\x2f\xce\x00\xda\x83\x04\x59\x30\x7d\x48\xd0\xa2\x93\x8c\x31\x2c\x6e\xef\x42\xd0\x59\x6e\x30\x43\xcb\x92\x35\x59\xd8\x92\x03\x4e\x11\xa2\xde\x96\x5d\xdf\x08\xb4\x85\xbc\xa5\x68\xfe\x5c\xee\x12\xd1\xe2\x44\xa2\xe6\xf9\x96\x22\x6c\xc9\x18\xfa\xa5\x6d\x02\x19\x72\x4a\x31\x60\xa9\x3d\xfb\xc6\x41\x15\x9e\x29\x03\xca\x6b\x12\x4d\xd6\x5f\xd4\x55\xa3\x11\x7c\x2a\x51\xd5\xd7\x28\x8a\x12\x1a\x0e\xea\x67\xa0\xb8\x04\x45\x96\xb1\x64\xb1\x68\xcf\x09\x6c\x4b\xd8\x16\x56\x05\xef\xfd\x0f\x23\x56\x37\xd7\x13\xa8\x2f\xe1\x7c\x0c\xe8\x1c\xb9\xee\x68\x1a\x3d\x07\xe4\x1f\x89\xb4\x6d\x46\x3e\x04\x53\x07\x26\x3d\xe4\xe8\x58\x6a\x5b\x57\x30\x35\x69\x3d\x62\x2e\x1c\x4a\xc6\x23\xd0\x56\xe8\x47\x45\x83\x19\x1c\xe2\xec\xb6\x57\x55\xe2\x99\xad\x3d\x65\x1f\x8d\xe0\x12\x79\xfe\xf0\x59\xa3\x89\x8f\x3c\x0f\x62\xbf\xef\x0e\x1f\xc0\xb3\xd3\x36\x99\xc0\x4f\x69\x0a\xec\x5e\x63\x80\xe0\x0c\x9c\x09\x74\x91\x3e\x01\x3a\x25\xe9\x47\xc8\x8b\x8d\xd1\xea\x2a\xfc\xe3\x7c\x9e\x71\xaf\x2f\xcc\x8c\x39\xf5\x9e\x19\xd3\x67\x3f\x86\x60\x7d\xff\x9f\x7e\x77\x79\x7c\xba\xe8\x56\x78\xd5\xb4\x6f\xb2\xf9\x10\x0d\x9e\x00\xb4\xc2\x2b\xe3\x3e\x6e\xf7\x3b\x00\x00\xff\xff\xd0\xde\x8b\x84\xe8\x04\x00\x00"),
			path: "sql-api-readme.tml",
			root: "sql-api-readme.tml",
		},

		"sql-api-test.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xdf\x6f\xe2\x38\x10\x7e\x47\xe2\x7f\x98\x8b\xee\xa4\xb0\x97\xf3\x6e\x5f\x39\xf1\x50\x9a\xdd\x6a\x4f\x2b\xca\x35\xf4\xf6\x11\xb9\xf1\x84\x7a\xd7\xc4\xd4\x76\xa0\x55\x94\xff\xfd\x64\xe7\x07\xb4\x94\x2e\xa1\xdb\x7b\xb8\x52\xa9\x55\x49\x66\xe6\xf3\x7c\x9e\xef\x23\x71\x9e\xc3\xaf\x8b\x84\xa3\x60\xd0\x1f\x80\xfb\x67\x78\x3f\xa2\x73\x04\x12\x19\x95\xc5\x06\xbc\x71\x76\x2d\x78\xfc\x39\xf4\xa0\x28\xba\x1d\x97\x60\xe8\xcc\x86\xcf\xd0\x4c\xe8\xac\x29\xe0\xe9\x5b\xe1\x81\xf7\x4d\xcb\xb4\x8c\xed\x76\x96\x54\x81\xdf\xed\x00\x00\xe0\x12\x53\xa3\x61\x00\x73\x34\x8a\xc7\x9a\x8c\x70\xe5\xc7\x99\x36\x72\x4e\x22\x43\xe3\xef\x21\xd7\x0b\x41\xef\x7d\xa9\x49\x64\x98\xcc\x4c\xaf\x67\x4b\xd8\xdc\x58\xa6\x09\x9f\xc1\x00\xf4\xad\x20\x67\xee\x43\x5e\xde\xb1\x3f\xe1\xd0\x2e\xb8\x0f\x52\x93\x73\x34\x98\x2e\x7d\x2f\xcf\xab\xe5\x93\x31\x8d\xbf\xd3\x19\x16\xc5\x34\xfa\xfb\xcb\x74\xf2\x31\x9a\x4c\xc3\xa1\xd7\x0b\xd6\xe9\x57\x1a\xd5\xfe\xc9\x57\xd1\xc7\xcb\x07\xe9\xe1\xf0\xf3\x78\xff\xf4\xd3\x30\x7c\x9c\x3e\x96\xca\xec\x5f\x60\x7c\x71\x39\x79\x54\x20\x54\x7c\xd9\xa6\x85\x32\x7e\x8b\x83\x31\xd5\x7a\x25\x15\x6b\xb1\x96\xd3\x28\xfa\x7a\x71\x19\x36\xa5\x8a\x7a\xc3\x0c\x6a\x73\x26\x05\x0c\xc0\xcb\x73\x21\x57\xa8\xea\x79\x22\x17\xd7\xdf\x30\x36\xc4\x6e\x99\xfb\x53\x14\x53\x1b\xed\x75\x3b\x6e\xbb\xdf\xbf\x87\x09\x6a\x73\x8e\x66\x0d\xbd\x91\x52\x14\xb0\xa4\x82\x33\x6a\x50\x83\xb9\x41\x50\x76\x9a\x70\x49\x05\xc8\x04\x28\xec\x48\x72\x75\x15\xc6\x52\x31\x48\x94\x9c\x03\xb5\xa3\xc4\xae\x49\xb7\x93\x64\x69\xfc\x03\x48\xdf\xc0\x3b\xbb\x46\x9e\xce\xc8\xa4\x57\x4d\x1e\x5d\x70\xab\x81\xb2\x8c\x9d\xe5\xaa\xe7\xa0\x9a\xf4\xc0\x0d\xeb\x08\x57\xe1\xd0\x2f\xe7\xb7\xbe\xb3\x31\xd7\xe6\xce\xd6\x88\x65\x6a\xf0\xce\x90\xaf\xdc\xdc\x4c\xf8\x1c\x65\x66\xfc\xfa\xda\x08\x57\xff\x50\x91\xe1\x90\xce\xfc\x5e\x00\x27\x1f\xe0\x1d\x18\x3e\x47\x12\x61\x2c\x53\xd6\x94\x42\x81\xf3\x00\x50\x29\x5b\xf0\x8b\xa4\xec\xc9\x5e\xaa\x86\xfe\x8a\x2e\x46\xfe\x8f\x37\xe6\x4c\x21\x35\x68\x83\x7b\x25\x08\x4f\x1c\xc2\x2f\x03\x48\xb9\x80\x0d\x09\xda\xde\x35\xf9\x44\xb9\x40\xe6\x7b\x51\x16\xc7\xa8\x75\x92\x09\x71\x0f\x42\x52\x86\x0c\x6c\x15\x48\xa4\xda\xb5\x43\xd5\xee\xf4\xe1\xb7\xdf\x6f\x89\xe7\x3a\xe9\x35\x43\xb5\x86\xb0\x43\xfa\x42\x08\xaf\xe1\x8c\x61\x82\xca\x6e\x24\x09\x51\xa0\x41\x3f\x36\x77\x81\x63\x92\xd4\xa6\xd7\xc4\x56\xad\xf7\x07\x2e\xbe\x64\x66\x1d\xdf\xfb\xb3\x3d\x31\x94\xd9\x45\xd7\x43\xf9\xcc\xb2\x79\x6a\x24\xb0\xeb\x83\xa8\x69\x0b\x42\xd6\xec\x4c\xcb\x71\x2a\x1b\x3e\x47\xf3\x34\x3b\x87\x8e\x45\xa5\x5b\x64\xa0\x8d\x54\xfb\xad\xd1\x29\xf7\x40\x22\x5e\x80\xe7\x38\x29\x1e\x9a\xd3\xa9\x10\x87\xf8\x93\x10\x2f\x74\xa8\xdd\xb8\x47\x93\x3a\x9a\xd4\xdb\x33\xa9\x32\x49\x07\xb5\x5b\xf5\x1b\xbb\x3a\x15\xa2\x6c\xdd\xa3\x3a\xf6\x02\x2f\xcf\xdd\x73\x2b\x71\x83\x5a\x14\x5e\x00\x7f\x9c\xd8\xdf\x9f\xe2\x61\x56\xd8\xd5\x52\xfe\x03\x07\x6b\x89\xb6\x41\x17\x4f\x40\x60\xea\x57\xc9\x3d\x18\x0c\xe0\x43\xfb\x66\x8d\x40\xaa\x0d\x9c\xb4\xf5\xd0\xd6\x7d\x1e\x0a\xb4\x69\xd6\x3b\xe2\x4b\x71\x3c\x32\xeb\xd8\x5e\xe4\x32\xdd\xfb\x59\x72\xc5\xcd\xcd\x53\x4e\xfd\x2c\xe8\xd1\xa9\x8f\x4e\xfd\x96\x9c\x7a\x0f\x35\x5e\x2d\xd8\xb6\x1a\xb3\xf2\xe2\x2b\x69\xb1\x84\x3c\x6a\xf1\xa8\xc5\xb7\xa4\xc5\xb2\x6a\x9e\x83\xa2\x29\xfb\xe4\x8e\xed\x9a\x83\x3e\xdb\x6d\x73\x7e\xe7\x2d\x1c\x4d\x53\xce\xea\xb3\xbc\x6d\xaa\x2a\x09\x6d\x53\x7b\x38\x73\xa5\xe6\x5f\x9d\xbb\xf6\x30\xc7\x17\xe3\xc3\x9f\xb5\x4a\x09\x6e\xbd\x18\xcf\xe5\xcf\x38\xb6\x7b\x16\xf3\x68\xef\xff\x4f\x7b\x7f\x7b\x96\xfd\xb0\xe3\xdd\x5f\x6a\x07\xb4\xee\x94\xf8\xfa\xcd\xb7\x87\xa9\xfd\xe5\xdf\x00\x00\x00\xff\xff\x19\xb7\xa0\x1f\xa5\x1a\x00\x00"),
			path: "sql-api-test.tml",
			root: "sql-api-test.tml",
		},

		"sql-api.tml": { // all .tml assets.
			data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x4b\x8f\xdb\x38\x12\x3e\xc7\x80\xff\x43\xad\xb0\x18\xc8\x03\x8d\x82\xbd\xf6\x6c\x1f\xa6\x1f\x19\x04\xd8\xe9\x64\x93\xcc\xee\x61\x30\x08\x68\xb1\x64\x33\xd1\xc3\x21\xa9\x6e\x1b\x86\xff\xfb\xa2\x48\xea\x69\xcb\xb6\xdc\x4e\x36\x9d\x49\x03\x89\x5b\x16\x59\x8f\xaf\x8a\xf5\x15\x45\xf5\xf3\xe7\x80\x52\xe6\x52\x41\x18\x86\xe3\xd1\x3d\x93\xe0\x8f\x47\x00\x00\xb7\x52\xde\xe5\xfa\x45\x5e\x64\x1c\x2e\xdd\xa0\xf0\x0e\x1f\x7c\x4f\x62\x94\x4b\x0e\x59\xae\x21\xa6\xdb\xde\x64\x3c\x9a\x8c\x47\xe3\xd1\xf3\xe7\x90\xb2\xc5\x0b\x81\x09\x57\xc0\x31\x16\x19\x2a\x60\xa0\x57\x0b\x84\x38\x97\xc0\xe8\x36\xe8\x39\xd3\x80\xcb\x45\xae\xcc\x5d\x3b\xdc\x9f\x40\x8a\x7a\x9e\xf3\x70\x3c\x32\xe3\x6b\x41\x29\x5b\xfc\xa1\xb4\x14\xd9\xec\x4f\x91\x69\x94\x31\x8b\x70\xbd\x71\xfa\xdc\x18\x89\xba\x90\x99\x02\x3d\x37\x33\x41\x68\x85\x49\x0c\x2c\xe3\xb0\x90\xf9\xbd\xe0\x46\x95\xd5\x00\x3a\x87\x94\xe9\x68\x6e\x46\xab\x4f\x49\xf8\x8e\x4d\x13\x34\x92\xa0\xd2\x10\x8e\x47\x71\x91\x45\xe0\xa7\xb5\x29\x93\xda\x58\x7f\xb7\x55\x81\x05\x6a\x02\x6b\x0b\xa2\xb5\x0b\xd2\x00\x32\x91\x8c\x47\xa5\xd5\xeb\x75\xf8\x56\xcb\x22\xd2\xe1\xab\xe9\x07\x8c\x74\x78\xc7\x52\xdc\x6c\xba\xc8\x65\xb5\x39\xf0\x30\x17\xd1\xbc\x82\xad\xf6\xc4\x69\xb0\xd0\xe6\x31\xb0\x24\x31\x2a\x98\xd6\x52\x4c\x0b\x4d\x72\x94\xca\x23\xc1\x34\x72\x78\x10\xda\xba\x6d\x75\x70\x50\xc6\x8c\x42\x22\x30\x52\x1c\x09\x8e\x1c\xa6\x2b\x0b\x4d\x79\xaf\x8c\xc9\x7e\xb3\x1b\xc6\xae\xc7\xa3\x67\xc7\x42\x75\x10\x95\xeb\x3c\x53\x45\x8a\x72\x1f\x2e\x2c\x8a\x70\xa1\x55\x0d\x03\x67\x9a\xb9\x7b\x0f\x22\x49\x60\x8a\x10\x59\x39\xdc\xe8\x12\x99\xce\x8d\x93\x33\x71\x2f\xb2\x19\x88\x74\x91\x60\x8a\x99\xa6\x8b\x73\x80\x52\x59\xdd\x46\xc5\x7d\xdd\x83\xc9\xc4\x62\x72\x10\x92\x9b\xab\xc6\xf2\xaa\xad\xb5\xfe\xba\x84\x87\x9b\x2b\xb8\x7e\xf3\xfb\x0d\xe4\x0b\x94\x4c\x8b\x3c\x53\x46\x64\xa1\x8c\x8b\x9f\x12\x72\x8e\x3c\x2a\x32\x8e\x32\x11\x19\x02\x9f\x1e\x70\xea\xe6\xca\x69\x33\xe9\x1d\xe5\x09\x58\x1f\xe8\x4a\x2d\xcd\x5a\xba\xb9\xa2\x0b\xbe\x84\x1f\xe9\xea\xed\xbf\xff\x45\x97\x29\x6a\x29\x22\x55\x7e\x86\xbf\xd9\x4f\xba\xa5\x69\xed\x91\x66\xb3\x08\x5f\x72\x0a\x81\x5e\xd5\x08\xdc\xe1\x43\xb5\xba\x19\x64\xf8\x00\x22\x53\x9a\x65\x11\x52\x98\x7b\xed\x2c\x97\x2f\x15\x2c\xab\xc2\x5a\x1a\x40\xda\xb5\x22\xa8\x4d\x0f\x40\xa7\x54\x08\xcd\x0c\x65\x4d\xfa\x4d\xcc\x2c\x7c\x13\xf8\xb1\x57\x9d\x5b\xee\x7c\x09\x17\x97\x46\x16\xe9\x4d\x49\x32\x89\x0c\xc3\xd0\xd4\xc7\x46\x45\xf8\xe1\x90\x28\xfa\x51\xcb\x0b\x92\x50\x7f\xc1\x97\x17\xc0\x9b\x5f\x44\x79\x72\x61\x21\x6c\x7c\xe9\xfc\xbb\x80\xb4\xf1\xa5\x19\x74\x51\x01\x4d\xda\xd6\xf4\x9f\x9b\xbe\x71\x43\x37\x35\xf2\x37\x98\xa0\x46\x2a\x24\x98\xd2\xda\x32\xc5\x26\xcd\xef\xd1\x64\x8d\x23\x81\x58\xe6\xa9\x2d\x28\x53\x97\x59\x74\xe1\x52\x90\xc3\xa2\x98\x26\x22\x7a\x79\x13\x1a\x89\x6f\xcc\x1c\x55\x0d\x14\x8a\x92\x34\x2d\x94\x86\x39\xbb\x47\x60\x6e\x3c\x08\x0e\xf7\x2c\x29\x30\xa0\x82\x27\x51\x29\xe4\x80\x42\xcf\x51\xd2\x32\x64\x30\x55\x79\x06\xb9\x84\x0f\xf4\xa9\xd9\xcc\x48\xa7\x5f\xed\x82\xc6\xac\x91\x18\xaf\x59\xf4\x91\xcd\x70\xb3\x09\x7b\x20\x77\x29\x5d\x17\x7c\x3e\xdd\x13\xe9\x89\xc3\xc5\x8f\xf4\x92\x4a\x8a\xc6\xa5\x0e\xaf\xed\x67\x50\xf9\xeb\xb2\xcd\xad\xe7\x92\x0d\x52\xca\x8e\x32\xfb\xee\xf0\xe1\x9d\x64\x11\xfa\x5e\x7f\x12\x5b\x55\xc4\xaf\x26\xfa\x18\xa3\x84\x94\x4f\xc3\x52\xc6\x6d\x2a\xb4\x5f\x5e\xbc\xcc\xe2\xfc\x18\x61\x41\x39\xe1\xbf\x42\xcf\x7d\xaf\x34\xd9\xab\xad\x6f\x0f\xb1\x56\xa6\xe1\x6d\xc6\xfd\xc9\xa4\x4a\x64\x11\x83\x50\xce\xf1\xdb\xe5\x42\x48\xe4\x84\x49\x45\x7d\xf4\x83\x52\x92\xcb\x71\xaa\xc3\x5b\x02\x22\xf6\x3d\x37\x03\xe6\x4c\x51\x70\x69\x5a\xe9\x9f\x81\xa8\xcf\xbb\x72\xfe\x0b\x26\x12\x34\xcc\xc7\x6d\x7e\xda\x44\x1c\xee\x96\xef\x99\xc4\xf7\x02\xa3\x33\xca\x93\xee\x6d\x13\x3a\xcf\xd0\x94\xd5\x6e\xbd\x2f\x2d\x75\x0b\x19\xa5\xac\x56\x4e\x05\x8c\xf3\x9b\x04\xf3\xa5\x03\x9e\xf2\xca\x16\x96\x00\x9c\x75\xef\x05\x6f\x9a\xf7\xb3\x99\xf7\xb7\x4b\xea\x15\x9a\x28\x9e\x0d\x93\x03\x0e\x1f\x01\xd9\x01\x4c\x9c\xeb\x97\x97\xc0\xa9\x08\x9a\x16\xf2\x4d\xfe\xa0\x9a\xde\x34\xb0\x6b\xb4\x98\xf5\xfd\xcd\x11\x10\x1f\x58\x03\x16\x70\x7e\x46\x18\xba\xe5\xbb\xd5\xcd\x5d\x4b\x64\x9d\x42\xc9\x38\x6f\x56\xc9\xaa\xdb\xd8\x5d\x25\x9b\x8c\xa6\xe7\xd8\x61\xfe\x83\x05\xec\xff\x58\x5c\x1f\x55\x48\x2d\x6e\xbb\x0b\x29\x26\x98\x0e\xc1\xe0\xb1\x95\xd6\xda\x72\xa6\x4a\x5b\x0a\xeb\x4f\x2c\x72\x2f\x7c\xfd\xa4\xca\x6d\x64\xb3\xfc\x70\xb9\xdd\xe3\xdb\xa9\x35\xf7\xd9\xb3\x67\xcf\x7a\x8b\xc1\x7a\x4d\xe0\xf8\x73\xa6\x5e\x50\xda\xb9\xa0\x80\x67\x37\x20\xde\x04\x36\x9b\xad\x0a\x55\x17\xe7\xb7\xec\xbe\x55\x9a\xc9\xfa\xde\x4a\xbc\x13\xb2\xf6\xed\x46\x03\xd6\x0f\x62\xdf\x02\xaa\xc0\xed\x97\xd9\x03\xe3\xa1\x09\xe4\x96\x8b\xcd\x11\x83\x77\x44\xa1\x33\xa9\xca\xcb\xbe\x5a\xdd\x0a\xd1\x11\x35\xfb\xba\x93\x5e\x5b\xed\xec\x61\xcf\xfb\x3d\x9e\x94\x99\x82\x89\xc2\x66\x3a\x98\xba\x93\x69\x93\x0f\xe5\x2e\xdf\x5f\xaf\xe9\x22\x2f\xb4\xdb\xd8\xba\xf5\xfc\x4b\x44\xbb\x00\xb0\x82\xc1\xa3\x32\xe9\x81\xf7\xc1\x7c\x6c\x36\xdb\x2c\xd8\x9f\x63\x4e\xeb\xb0\x34\x7b\x7c\x52\x0d\x5a\x89\x9f\x0a\x94\x2b\xaf\xb6\x75\x58\x23\xf0\x39\x32\xe2\x51\xe6\xd7\x19\x90\xf1\xcd\x66\x1f\x9b\xff\x8a\xfa\x97\x24\xb9\x5a\xbd\x92\x1c\x65\x67\xf7\xa3\xa5\x40\xa2\xd2\x24\x31\x79\x85\x99\x56\x76\x07\x54\xf1\x7a\xc5\xe9\xb9\x9d\x9e\xb9\xdf\xa6\x2b\x30\xd2\x0d\xf3\xaa\xa3\x29\xb2\x65\xcc\x6e\xa6\xb4\x8a\xca\xdd\x6d\xa9\xac\xdc\x7f\x80\xff\xc7\x9f\x03\x68\x74\xfb\xf9\x95\x0a\xde\x07\xcd\x74\xb6\x16\x91\x29\x4e\x59\xa5\x33\x80\x9f\xfe\x41\xff\x26\x2d\x6c\x49\x82\xcd\x81\x36\xc0\x15\x9a\xca\xc0\x29\x5d\xf7\xd2\xdc\x4f\x12\x78\xf5\xd6\x5f\x25\xa2\xbb\xe7\x3f\xbc\xb5\xd3\xab\x05\x3e\xd5\xf6\xa8\x46\x7a\x58\xd0\x03\x58\xb0\x19\x52\xcf\x19\x10\xfa\x8b\x3c\x53\xf8\x1a\xe5\x6b\xf7\xe5\x09\x39\x61\x44\xb5\x13\x63\x70\x83\x65\xbd\x39\x53\x83\x55\x0a\x7b\x12\x8d\x53\x55\x36\x4e\xab\x64\x3d\x95\x76\x6b\x97\x94\x89\xc4\xae\xc0\xed\x0e\xe9\x9e\x49\xd0\xb9\x66\x09\x45\xb2\xfe\x4a\x0a\x8d\xa9\x82\x41\xc9\x70\xa8\xe9\x72\xcf\x37\xbb\x5d\x17\x29\x0a\xac\x0d\x41\x87\x1c\x6d\x2c\x5d\x82\x36\x59\xb2\x5b\x5d\x28\xab\xb7\x32\x7a\xf2\xf3\x16\xef\x3e\x9a\x4f\x5b\x75\xbe\x2c\x4c\xfd\x0f\x1b\xdd\xf9\x86\xa5\x81\x33\x45\xb7\x69\xf9\xb1\x9b\xea\x6e\x2e\x38\xb4\x77\x6e\xb2\xa1\xbd\xd1\xee\x99\xda\x4f\xdd\x71\x2e\xe1\x7d\x60\xe2\x4a\xa1\x94\x2c\xa3\xe2\x62\xd2\xa9\x63\x15\x25\xda\xd0\xcd\xdc\x4e\xff\x2f\x2e\xed\xce\xa2\x7c\x82\x4e\xda\xf6\x76\x50\x03\xa3\xee\x8e\x07\xca\x2d\xbb\x39\x40\x20\x37\xfb\xa2\x7e\x62\xc0\x49\xae\x67\x91\x3b\x21\x17\xf6\xae\xf7\xbe\xc0\xda\xb8\x5c\x02\x5b\x2c\x30\xe3\xbe\x74\xab\xd1\x74\xc6\x5b\xcf\x5d\xb6\x7b\xe4\x05\xcb\x44\xd4\x53\x8d\x4b\x30\x0c\x8f\x56\x67\x28\xbb\x31\xeb\x1c\x8c\x78\x7b\x5b\x32\xd9\x2e\x19\xdd\x0e\xed\x6a\x65\x8f\xeb\x1a\x4d\xc4\xfe\x27\xd2\xb1\x19\xfe\x11\x57\xa6\xab\xb0\x04\x6f\x84\x35\x4f\x0e\xff\x42\x9d\x85\x03\x70\x77\x77\x41\x30\x95\xbd\x84\xb1\x18\x5a\x67\x54\xe0\x0f\xea\x1e\xce\xd0\x39\x9c\xaf\x6d\xe8\x2e\xd7\x8f\xb8\x72\x3e\x7e\x13\xcd\x44\x9f\x3b\x67\xeb\x32\x06\x44\xde\x1d\xf1\xee\xe8\x45\x86\xf3\xc1\x49\x4d\x47\x77\x1b\xfe\x2b\xea\x66\x7f\xf1\x03\x99\x61\x92\xbd\x84\xec\x4c\xfb\xf1\x41\xfd\x43\x4f\x13\xd1\x51\x3e\xe4\xe1\xcf\xc0\xc0\xc2\x69\x2d\xc6\xc0\x44\xd8\xd7\x85\xec\xec\x43\x4e\x4d\xb4\xaf\x8f\xca\x6c\x9a\x75\x19\xec\x78\xea\xaa\xce\x14\x9b\xfb\xe1\xbf\x18\x5d\x1d\x79\xda\xfa\x2d\x31\xd3\x49\x07\xb3\x4f\x83\xa5\xbe\xd8\xe1\xec\xd7\xc0\x56\xd0\x4b\x57\x00\x6d\xc2\x5a\xaf\xe1\xef\x0b\xdb\xa8\x52\x64\xe8\x97\xab\x15\xc9\xac\xa7\x96\xc7\x2a\xde\xf6\x44\xcd\x66\x34\x6d\x86\xfa\x1d\x9b\x55\x82\x3c\xf5\x29\x29\x9f\x52\x6f\x5a\xb5\xf6\x48\x7a\xf4\xd6\x6b\x23\x3b\xfc\x0f\xd5\x82\xcd\xe6\x98\x03\xeb\xef\x7c\xf9\x59\xf9\xf2\xdb\x63\xcb\x43\x7c\xf9\xfb\x82\x33\x8d\x50\xa8\xef\x6c\x79\x90\x2d\x2d\x56\x47\x11\xe6\x97\x3f\x66\xb7\xc6\x9d\x89\x34\x4b\x61\xdf\xca\x0b\x4d\xb1\xc8\x84\x9a\x07\x65\xd8\xda\x12\x4f\x79\x87\xa5\xf3\x7e\xd1\x99\xde\x6f\x3a\x89\xa6\x06\x53\x54\x2f\x7b\xba\x63\xfd\x2e\x79\x6e\xd1\x99\x5b\x07\x9d\xa3\xfd\x2f\x42\x68\x85\x2d\x57\x03\x0f\x64\xcb\x23\x4b\x7b\x46\x3f\x28\xdc\x7b\x5c\xfa\x22\xcf\x9b\x07\xb2\xd5\x71\xf4\x63\x9e\xc0\x7e\xde\x83\xf9\xed\x1c\x21\xa5\xdf\x73\xe4\xeb\xcd\x91\x76\xc7\x70\x80\x3a\x5c\xdb\x70\xda\xb9\xdb\xf1\x78\xb5\x82\xb2\xf7\x4d\xc1\xdb\x25\x46\xcd\x3f\x48\x21\x46\x37\xb9\xeb\xfe\x8a\x21\x49\x08\x43\x6a\x09\x70\x89\x51\x61\x6e\xe5\x31\x30\x88\x0a\xa5\xf3\xb4\x1e\xcf\x66\x4c\x64\x4a\x9b\xa1\xc6\x8f\xa3\xdb\x03\x32\x61\x77\x73\x10\x2f\x8d\x7c\xbf\x7c\x71\x3f\x70\x2f\xc5\x4f\xca\x6d\xf2\xe3\xa8\x9f\x14\x9f\x89\xf8\xad\xa8\x27\xc1\xe9\x36\x8e\x58\xff\x11\xc6\x99\x0f\x08\x8f\x78\x0b\x39\x5e\xfa\xb6\xdc\x59\x1d\x6a\x79\x8e\x17\x8d\x3f\xb7\x5f\x5f\xc9\x9b\xc4\xaf\x4a\xf7\x4a\x87\x0f\x97\x91\xfe\x02\x60\x16\xe8\xae\x74\xec\xae\xc5\x09\x4c\xf3\xbc\x8a\x8c\xc2\x04\xdd\x5f\xdc\xd8\x9f\x88\x29\x84\x7f\xfe\x14\xe9\x65\x78\x93\x67\xe8\x4f\x2e\x76\x62\xa2\x65\x81\x0d\x36\xc5\x98\x15\x89\xde\x3d\x34\x66\x89\xc2\x0a\x99\xcd\x78\xf4\xbf\x00\x00\x00\xff\xff\x38\xe2\xd5\x89\x17\x38\x00\x00"),
			path: "sql-api.tml",
			root: "sql-api.tml",
		},
	}
)

//==============================================================================

// FilesFor returns all files that use the provided extension, returning a
// empty/nil slice if none is found.
func FilesFor(ext string) []string {
	return assets[ext]
}

// MustFindFile calls FindFile to retrieve file reader with path else panics.
func MustFindFile(path string, doGzip bool) (io.Reader, int64) {
	reader, size, err := FindFile(path, doGzip)
	if err != nil {
		panic(err)
	}

	return reader, size
}

// FindDecompressedGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindDecompressedGzippedFile(path string) (io.Reader, int64, error) {
	return FindFile(path, true)
}

// MustFindDecompressedGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindDecompressedGzippedFile(path string) (io.Reader, int64) {
	reader, size, err := FindDecompressedGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindGzippedFile(path string) (io.Reader, int64, error) {
	return FindFile(path, false)
}

// MustFindGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindGzippedFile(path string) (io.Reader, int64) {
	reader, size, err := FindGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFile returns a io.Reader by seeking the giving file path if it exists.
func FindFile(path string, doGzip bool) (io.Reader, int64, error) {
	reader, size, err := FindFileReader(path)
	if err != nil {
		return nil, size, err
	}

	if !doGzip {
		return reader, size, nil
	}

	gzr, err := gzip.NewReader(reader)
	return gzr, size, err
}

// MustFindFileReader returns bytes.Reader for path else panics.
func MustFindFileReader(path string) (*bytes.Reader, int64) {
	reader, size, err := FindFileReader(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFileReader returns a io.Reader by seeking the giving file path if it exists.
func FindFileReader(path string) (*bytes.Reader, int64, error) {
	item, ok := assetFiles[path]
	if !ok {
		return nil, 0, fmt.Errorf("File %q not found in file system", path)
	}

	return bytes.NewReader(item.data), int64(len(item.data)), nil
}

// MustReadFile calls ReadFile to retrieve file content with path else panics.
func MustReadFile(path string, doGzip bool) string {
	body, err := ReadFile(path, doGzip)
	if err != nil {
		panic(err)
	}

	return body
}

// ReadFile attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFile(path string, doGzip bool) (string, error) {
	body, err := ReadFileByte(path, doGzip)
	return string(body), err
}

// MustReadFileByte calls ReadFile to retrieve file content with path else panics.
func MustReadFileByte(path string, doGzip bool) []byte {
	body, err := ReadFileByte(path, doGzip)
	if err != nil {
		panic(err)
	}

	return body
}

// ReadFileByte attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFileByte(path string, doGzip bool) ([]byte, error) {
	reader, _, err := FindFile(path, doGzip)
	if err != nil {
		return nil, err
	}

	if closer, ok := reader.(io.Closer); ok {
		defer closer.Close()
	}

	var bu bytes.Buffer

	_, err = io.Copy(&bu, reader)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("File %q failed to be read: %+q", path, err)
	}

	return bu.Bytes(), nil
}
