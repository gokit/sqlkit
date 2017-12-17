package static


import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

type fileData struct{
  path string
  root string
  data []byte
}

var (
  assets = map[string][]string{
    
      ".tml": []string{  // all .tml assets.
        
          "sql-api-json.tml",
        
          "sql-api-migration.tml",
        
          "sql-api-readme.tml",
        
          "sql-api-test.tml",
        
          "sql-api.tml",
        
      },
    
  }

  assetFiles = map[string]fileData{
    
      
        "sql-api-json.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\x41\x6b\xc2\x40\x10\x85\xcf\x09\xe4\x3f\xbc\x7a\x28\x06\x24\xde\x5b\x3c\x94\x42\x0f\x3d\x68\x41\x3c\x95\x82\xeb\x3a\xd1\xd8\x64\x57\x26\xa3\x52\x96\xfd\xef\x65\x37\x4a\x3d\xb4\xa5\xa1\x97\x84\x9d\x19\xde\xf7\xf1\x8e\x8a\xe1\x5c\x6d\x4f\xc4\x28\xe6\xc2\x07\x2d\xc5\x6c\xb5\x23\x2d\xc5\x54\x35\x14\x3f\xde\x3f\xcf\x67\x53\x4c\xb0\x74\x0e\x8d\xda\xc7\xd7\xf9\x18\x83\x55\x6b\xcd\x00\x83\x5d\xfc\x79\xbf\xcc\xd2\x2c\xfd\x5b\xea\x23\x93\x12\xfa\x2e\xbb\xdb\x3c\x68\xa9\xac\xf9\x0f\x61\xb1\x5f\xff\x40\xe8\x36\xbf\x11\xc6\x63\xd4\x56\xad\xc3\xf9\x93\x65\x30\xc9\x81\x4d\x0b\x05\x43\x27\x54\xa6\x15\x65\x34\xc1\x96\x50\x70\xee\xa2\xf0\xa2\xf4\xbb\xda\x90\xf7\xc5\xd7\xec\x4a\xcb\x7b\x94\x6c\x1b\xc8\x96\xb0\x67\x7b\xac\xd6\x84\x00\x85\xb6\x46\xc8\x48\x91\xa5\xe5\xc1\xe8\x6b\xf0\xf0\xbc\x42\x2b\x5c\x99\x4d\x8e\x61\x0f\xda\x08\xc4\x6c\x39\x87\xcb\xd2\x24\x54\x46\x35\x35\x7d\x74\x43\x11\x49\x55\x86\x18\xdc\x4d\xa2\x6b\xb1\x30\x8d\xe2\x76\xab\xea\xe1\xeb\xdb\xea\x43\xe8\x62\x98\x8f\x70\x1b\xf2\xf3\xfb\x78\x7e\x33\x81\xa9\xea\x48\x4e\xba\xf2\xfa\x80\x5d\xe7\x9e\xa5\x49\xe7\x70\x4e\x08\xf9\xa3\x90\x9b\xa5\x71\xfe\x19\x00\x00\xff\xff\x20\x59\xff\x66\xbf\x02\x00\x00"),
          path: "sql-api-json.tml",
          root: "sql-api-json.tml",
        },
      
        "sql-api-migration.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\xcf\x4b\xfb\x40\x10\xc5\xef\xfd\x2b\x1e\x3d\x7d\xbf\x50\xd3\x7b\xc0\x43\xd5\x0a\x42\xb5\xa0\xf1\x24\x22\xd3\xec\xb4\x59\xe9\xfe\x70\x77\x6a\x94\x90\xff\x5d\x76\x25\xc9\x41\xcd\x69\xb2\xf3\x3e\x6f\xde\x7b\xa7\x80\x7f\x33\x00\x58\x2e\xd1\x75\xc5\x83\x84\x53\x2d\xc5\x76\xf7\xca\xb5\x14\x77\x64\xb8\xef\x2b\xda\x1d\x19\x8a\xf7\xda\x72\x84\x34\x0c\xc9\x2f\x46\x1f\x02\x89\x76\x16\x86\x85\xce\x14\x09\xa1\x6d\x74\xdd\x60\x30\xd4\x11\xa7\xc8\x0a\xe2\x70\x60\xcb\x81\x84\x33\x4f\xde\x07\xe7\x83\x4e\xff\x97\xf7\xeb\x55\xb5\x46\xb5\xba\xd8\xac\x11\xdf\x8e\x88\x42\xc2\x86\xad\x60\xef\xc2\xc0\x69\x7b\x18\x5d\x29\x46\x57\x27\x56\x65\x7d\x0e\x53\x0c\xcb\x6a\x7b\xb5\x2d\xf1\xe8\x55\xf2\x6e\xb5\x34\xd8\x6b\x3e\x2a\x58\x32\x0c\xf9\xf4\x1c\x41\x56\x41\x5b\xc5\x1f\xdf\xd0\x1f\xad\x6f\xc7\x76\xe7\x18\xe7\x2e\x13\xe9\xab\xb4\xe1\x28\x64\x3c\xab\x12\x12\x4e\xbc\xc0\xb4\x4b\x89\x92\x4b\x89\x79\xd7\x89\xdb\xb8\x96\x03\x7e\xbd\xf2\x92\xd3\xcf\x17\x23\x7b\x93\x82\x71\x2c\xf1\xf4\x9c\xc7\xe9\x74\x3f\x89\xae\x53\xa5\xac\xc9\xd3\x0f\x4d\x3f\xfb\x3f\xfb\x0a\x00\x00\xff\xff\xde\xf6\xc3\xf5\xda\x01\x00\x00"),
          path: "sql-api-migration.tml",
          root: "sql-api-migration.tml",
        },
      
        "sql-api-readme.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\xd1\x6a\xdb\x30\x14\x86\xef\x03\x79\x87\x1f\x72\x93\x8c\x4c\x0f\x30\xd8\x45\x12\x6f\xa3\x50\xb6\x74\x59\xaf\xc6\xc0\xaa\x72\x62\x6b\x91\x25\x4f\x3a\xde\x5c\x8a\xdf\x7d\xc8\x76\x97\xb8\xb8\xb4\x65\xf5\x8d\xa4\x83\x75\xbe\x4f\xff\xb9\xbb\x13\x3b\xf6\x95\x62\xf1\xe5\xe6\x27\x29\x16\x9f\x65\x41\x4d\x83\xdd\xd5\x65\xb2\xc6\x6a\x7b\x31\x9d\xbc\x7f\xfa\x9b\x4e\xa6\x93\x27\x3b\x41\x07\x48\xc8\x8a\xdd\xdb\x8c\x2c\x79\xc9\xb4\xc7\xe6\xeb\x75\x02\x5d\x94\x86\x0a\xb2\x2c\x59\x3b\x8b\x83\xf3\xe0\x9c\x90\x8e\xb6\xec\xfb\xa6\xd0\x16\xa5\x54\x47\x99\x75\x7f\x6e\x8f\x99\xd8\x4a\xce\x9b\x26\x15\xd1\xe7\x5b\x4e\x38\x38\x63\xdc\x1f\x6d\x33\x14\xc4\xb9\xdb\x83\x6a\x1d\x38\xb4\x04\x55\x05\x76\x05\x5c\x19\x4d\xb4\xb3\xe1\x5d\xbc\x35\x9b\xe1\x43\x4d\x2a\x6e\xd3\x34\xcd\xdc\x74\x12\x8f\x73\xc5\x35\x94\xb3\x4c\x35\x8b\x4d\xb7\x2e\x71\xa8\x71\xa8\xac\x9a\xbf\x09\xbf\x8c\xd8\x5d\x5d\x2e\x11\x37\xc9\x7a\x01\xf2\xde\xf9\x7e\x69\x1b\x3d\x26\x14\xee\x8d\xb4\x6d\x9f\x7c\x0a\x26\x06\x26\x03\x4a\xf2\x2c\xb5\x8d\x37\xd8\xb5\x69\xdd\x6b\x6e\x3c\x49\xa6\x33\xd1\xae\x30\xae\x4a\x86\x0a\x9c\xe2\xdc\x76\xb9\x35\x8d\x78\x64\x6a\x0f\xdd\x67\x33\x7c\x22\x5e\xdf\x7e\xd4\x64\xf6\x67\xcc\x53\x71\x9c\x7b\xa4\x5b\x04\xf6\xda\x66\x4b\xfc\x96\xa6\xa2\xfe\xb4\x00\xe6\x2f\xd0\x59\xa2\x8f\xf4\x81\xd0\xd0\x64\x5c\xa1\xac\x6e\x8c\x56\x17\xc9\x3f\xf2\xcb\xc0\xa3\x5c\xac\x8c\x19\xb2\x57\xc6\x8c\xe1\x17\x98\x7f\xff\xf1\x9f\xbc\xeb\x72\x3f\x1c\x74\x57\x78\xd6\x6b\x5f\x65\xf2\x09\x19\x1a\x08\x74\x85\x67\xc6\x7d\xde\xee\x6f\x00\x00\x00\xff\xff\x66\x01\x04\xbe\x6e\x04\x00\x00"),
          path: "sql-api-readme.tml",
          root: "sql-api-readme.tml",
        },
      
        "sql-api-test.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xdd\x4f\xe3\xc6\x17\x7d\x8f\x94\xff\xe1\xfe\xac\x5f\x25\x7b\xeb\x0e\xcb\x6b\xaa\x3c\x10\xdc\x45\x5b\x55\x49\x8a\x43\xf7\x11\x0d\xf6\x75\x98\xdd\xb1\x07\x66\xae\x13\x90\xe5\xff\xbd\x1a\x8f\xf3\xc1\x47\xd8\x24\x2c\x95\x4a\x83\x04\x02\xfb\x7e\x9e\x39\xe7\x60\xa7\xaa\xe0\xff\x37\x99\x40\x99\x42\xaf\x0f\xcd\x2f\x83\xfb\x21\xcf\x11\x58\x4c\xba\x4c\x08\xbc\x71\x79\x25\x45\xf2\x39\xf2\xa0\xae\xbb\x9d\x26\x81\xf8\xd4\x86\x4f\x91\x26\x7c\xba\x2c\xe0\x99\x5b\xe9\x81\xf7\xd5\xa8\xc2\xc5\x76\x3b\x33\xae\xc1\xef\x76\x00\x00\x70\x86\x05\x19\xe8\x43\x8e\xa4\x45\x62\xd8\x10\xe7\x7e\x52\x1a\x52\x39\x8b\x89\x27\xdf\x22\x61\x6e\x24\xbf\xf7\x95\x61\x31\xa5\xaa\xa4\x20\xb0\x25\x6c\x6e\xa2\x8a\x4c\x4c\xa1\x0f\xe6\x56\xb2\xd3\xe6\x8f\xca\xdd\xb1\x5f\xd1\xc0\x0e\xdc\x03\x65\xd8\x19\x12\x16\x33\xdf\xab\xaa\x76\x7c\x36\xe6\xc9\x37\x3e\xc5\xba\xbe\x8c\xff\xfc\xe3\x32\x1a\x78\x41\xb8\xca\xbc\x30\xa8\xb7\xca\xbb\x88\x7f\x3b\x7f\x90\x19\x0d\x3e\x8f\xb7\xca\x3c\x89\xa2\xc7\x99\x63\xa5\x69\xab\xdc\xf1\xe8\x7c\xf2\x28\x37\xd2\x62\xb6\xe5\xcc\x2e\xf4\xc9\xbe\x63\x6e\xcc\x5c\xe9\x74\xbb\x09\x4e\xe2\xf8\xcb\xe8\x3c\x5a\x56\xa9\x17\x47\x42\x68\xe8\x54\x49\xe8\x83\x57\x55\x52\xcd\x51\x2f\x18\xc3\x46\x57\x5f\x31\x21\x66\x0f\xa5\xf9\x51\xd7\x97\x36\xda\xeb\x76\x9a\x03\x3d\x3a\x82\x09\x1a\x3a\x43\x5a\x75\x5d\x4b\xa9\x6b\x98\x71\x29\x52\x4e\x68\x80\xae\x11\xb4\xe5\x0b\xce\xb8\x04\x95\x01\x87\x0d\x49\x4d\x5d\x8d\x89\xd2\x29\x64\x5a\xe5\xc0\x2d\x59\xd2\x2b\xd6\xed\x64\x65\x91\x7c\xa7\xa5\x4f\xf0\xc1\xce\x28\x8a\x29\x9b\x04\x2d\xb7\xf8\x8d\xb0\x2c\x77\x65\x2c\x5b\xdb\x9d\xc3\x96\xcb\x61\x43\xc7\x21\xce\xa3\x81\xef\x18\xba\xb8\xb3\xc6\x5c\xba\xb3\x35\x12\x55\x10\xde\x11\xfb\x22\xe8\x7a\x22\x72\x54\x25\xf9\x8b\x6b\x43\x9c\xff\xc5\x65\x89\x03\x3e\xf5\x83\x10\x8e\x3f\xc2\x07\x20\x91\x23\x8b\x31\x51\x45\xba\x2c\x85\x12\xf3\x10\x50\x6b\x5b\x50\x2a\x9e\xfe\x1e\x8f\x86\x9f\x94\xf6\xbf\x8f\xff\xa9\x46\x4e\x68\xe3\x03\x57\x4b\x64\x4d\xa1\xff\xf5\xa1\x10\x12\xd6\xb4\x64\x57\x34\xec\x13\x17\x12\x53\xdf\x8b\xcb\x24\x41\x63\xb2\x52\xca\xfb\xa6\x25\xa6\x60\xab\x40\xa6\xf4\xa6\x83\x68\x0f\xa1\x07\x3f\xfd\x7c\xcb\xbc\x66\xe0\x60\xc9\x9d\x55\x0b\x4b\xc3\x57\xb6\xf0\x96\xd0\xa4\x98\xa1\xb6\xe7\xc5\x22\x94\x48\xe8\x27\x74\x17\x36\x80\xb1\x85\x7b\x2d\x63\xdb\xd5\x7b\xfd\x26\xde\x21\xb3\x8a\x0f\x7e\xdd\x1d\x18\x9e\xda\xa1\x17\xdc\x7b\x61\x6c\x51\x90\x82\xf4\x6a\x2f\x68\x76\x6d\xc2\x56\xe8\x5c\x3a\xd6\xb8\x85\xcf\x90\x9e\x47\x67\x5f\x5a\xb4\xf2\xc4\x14\x0c\x29\xbd\xdd\x8c\x8d\x40\xf7\x04\xe2\x15\xfd\x1a\x4c\xea\x87\x1e\x74\x22\xe5\x3e\x36\x24\xe5\x2b\x8d\x68\x73\xdf\x83\x17\x1d\xbc\xe8\xdd\x7a\x91\x4b\x32\xe1\xc2\x94\x7a\x4b\x57\x3a\x91\xd2\xad\xee\x71\x93\x78\xa1\x57\x55\xcd\x73\x26\x6b\xf8\x58\xd7\x5e\x08\xbf\x1c\xdb\xef\x1f\x62\x55\x56\xbf\xed\x28\xff\x80\x51\xed\xd8\x6d\x0d\x2e\x91\x81\xc4\xc2\x6f\x93\x03\xe8\xf7\xe1\xe3\xee\xcb\x92\x44\x6e\x08\x8e\x77\xb5\xca\x9d\xf7\xdc\xb7\xd1\xba\x27\x6f\x88\x77\xe2\x78\xe4\xc9\x89\xbd\x28\x54\xb1\xf5\x93\xe1\x5c\xd0\xf5\x73\x86\xfc\x62\xd3\x83\x21\x1f\x0c\xf9\x1d\x1a\xf2\x16\xa2\xbb\xb8\x49\x9f\x8a\xae\x74\x17\xdf\x48\x72\xae\xe5\x41\x72\x07\xc9\xbd\x43\xc9\xb9\xaa\x47\x47\x93\x51\x34\xea\x41\xab\x2e\xa3\x72\xa4\x6b\xcb\xf4\xe7\x21\x69\x15\xf1\x14\xc2\xfd\x11\x72\x12\x7e\x73\x8c\x76\x6f\x73\x78\x6b\xdd\xff\x09\xc9\x49\xed\xc9\x5b\x6b\xae\x7e\xc4\x47\x67\x2f\xf6\x3c\xb8\xf5\xbf\xda\xad\xff\x7b\x0e\xfc\x70\xe3\xcd\xff\xa3\xf6\x58\xbd\x11\xdc\xdb\x2f\xbf\x7b\x9b\x85\x8d\xfc\x1d\x00\x00\xff\xff\x08\x8f\xe4\xc4\xf2\x19\x00\x00"),
          path: "sql-api-test.tml",
          root: "sql-api-test.tml",
        },
      
        "sql-api.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\xdd\x6f\xdb\x38\x12\x7f\x4e\x80\xfc\x0f\xb3\xc2\x61\x21\x2d\x74\x0a\xee\x35\x41\x1f\x36\x1f\x3d\x14\xb8\x4d\x7b\x6d\xf7\xee\x61\xb1\x58\xd0\xe2\xd8\x66\xab\x0f\x97\xa4\x63\x1b\x86\xff\xf7\xc3\x90\xd4\x67\x2c\x5b\xb6\x75\xd9\xa6\x5d\x3f\xc4\x91\x44\x0e\x67\x7e\xf3\xc9\x11\x7d\x79\x09\x29\x9b\xbd\x16\x98\x70\x05\x1c\xc7\x22\x43\x05\x0c\xf4\x6a\x86\x30\xce\x25\x30\x7a\x0c\x7a\xca\x34\xe0\x72\x96\x2b\xf3\xd4\x0e\xf7\x03\x48\x51\x4f\x73\x1e\x5d\x9c\x9b\xf1\x15\xa1\x94\xcd\x7e\x53\x5a\x8a\x6c\xf2\xbb\xc8\x34\xca\x31\x8b\x71\xbd\xb9\x38\xbf\x38\xbf\xbc\x74\xb3\x41\xa2\x9e\xcb\x4c\x81\x9e\x9a\x99\x20\xb4\xc2\x64\x0c\x2c\xe3\x30\x93\xf9\xa3\xe0\x66\x29\xbb\x02\xe8\x1c\x52\xa6\xe3\xa9\x19\xad\xbe\x24\xd1\x47\x36\x4a\xd0\x50\x82\x72\x85\xe8\xe2\x7c\x3c\xcf\x62\xf0\xd3\x8a\x95\xa0\x62\xd6\xdf\xce\x55\x08\x28\x65\x2e\x03\x58\x5f\x9c\x03\x80\xe3\x0b\xd2\x10\x32\x91\x5c\x9c\x17\x5c\xaf\xd7\xd1\x07\x2d\xe7\xb1\x8e\xde\x8e\x3e\x61\xac\xa3\x07\x96\xe2\x66\xd3\x46\x2e\xab\xd8\x81\xc5\x54\xc4\xd3\x12\xb6\x4a\x12\xb7\x82\x85\x36\x1f\x03\x4b\x12\xb3\x04\xd3\x5a\x8a\xd1\x5c\x13\x1d\xa5\xf2\x58\x30\x8d\x1c\x16\x42\x5b\xb1\xed\x1a\x1c\x94\x61\x63\x2e\x11\x18\x2d\x1c\x0b\x8e\x1c\x46\x2b\x0b\x4d\xf1\xac\xd0\xc9\x6e\xb6\x6b\xcc\xae\x2f\xce\xcf\xfa\x42\xb5\x17\x95\xdb\x3c\x53\xf3\x14\xe5\x2e\x5c\x58\x1c\xe3\x4c\xab\x0a\x06\xce\x34\x73\xcf\x16\x22\x49\x60\x84\x10\x5b\x3a\xdc\xac\x25\x32\x9d\x1b\x21\x27\xe2\x51\x64\x13\x10\xe9\x2c\xc1\x14\x33\x4d\x17\x43\x80\x52\x72\xdd\x44\xc5\xdd\xee\xc0\x24\xb0\x98\xec\x85\xe4\xee\xa6\xe6\x5e\x15\xb7\x56\x5e\x67\xf0\x70\x77\x03\xb7\xef\x7f\xbd\x83\x7c\x86\x92\x69\x91\x67\xca\x90\x9c\x2b\x23\xe2\x97\x84\x84\x23\x89\xe6\x19\x47\x99\x88\x0c\x81\x8f\xf6\x08\x75\x77\xe3\x56\x33\xe6\x1d\xe7\x09\x58\x19\xe8\x4a\x2d\x8d\x2f\xdd\xdd\xd0\x05\x5f\xc2\x4f\x74\xf5\xe1\xdf\xff\xa2\xcb\x14\xb5\x14\xb1\x2a\xbe\xa3\x5f\xec\x37\x3d\xd2\xe4\x7b\xb4\xb2\x71\xc2\x37\x9c\x54\xa0\x57\x15\x02\x0f\xb8\x28\xbd\x9b\x41\x86\x0b\x10\x99\xd2\x2c\x8b\x91\xd4\xdc\xc9\x67\xe1\xbe\x0f\xb8\xf0\xed\x12\x96\xd3\x10\xd2\x36\x17\x61\xc5\x7a\x08\x3a\x85\x28\x8a\xcc\x0c\x65\x59\xfa\x45\x4c\x2c\x7c\x01\xfc\xd4\xb9\x9c\x73\x77\xbe\x84\xab\x57\x86\x16\xad\x9b\x12\x65\x22\x19\x45\x51\x40\xe2\xd4\x22\xc2\x8f\xfb\x48\xd1\x47\x2d\xaf\x88\x42\x75\x83\x2f\xaf\x80\xd7\x6f\xc4\x79\x72\x65\x21\xac\xdd\x74\xf2\x5d\x41\x5a\xbb\x69\x06\x5d\x95\x40\xd3\x6a\x6b\xfa\xe3\xa6\x6f\xdc\xd0\x4d\x85\xfc\x1d\x26\xa8\x91\x02\x09\xa6\xe4\x5b\x26\xd8\xa4\xf9\x23\x1a\xab\x91\x18\xe7\x92\xc3\x58\xe6\xa9\x0d\x28\x23\x67\x59\x74\xe1\x4c\x90\xc3\x6c\x3e\x4a\x44\xfc\xe6\x2e\x32\x14\xdf\x9b\x39\xaa\x1c\x28\x14\x19\x69\x3a\x57\x1a\xa6\xec\x11\x81\xb9\xf1\x20\x38\x3c\xb2\x64\x8e\x21\x05\x3c\x89\x4a\x21\x07\x14\x7a\x8a\x92\xdc\x90\xc1\x48\xe5\x19\xe4\x12\x3e\xd1\xb7\x66\x13\x43\x9d\xfe\xb5\x0e\x8d\x59\xcd\x30\xde\xb1\xf8\x33\x9b\xe0\x66\x13\x75\x40\xee\x4c\xba\x0a\xf8\x7c\xb4\x43\xd3\x81\xc3\xc5\x8f\xf5\x92\x42\x8a\xc6\xa5\x8e\x6e\xed\x77\x58\xca\xeb\xac\xcd\xf9\x73\x91\x0d\x52\xb2\x8e\xc2\xfa\x1e\x70\xf1\x51\xb2\x18\x7d\xaf\xdb\x88\xed\x52\x5e\xe0\xac\x0b\xc7\x28\x21\xe5\xa3\xa8\xa0\x71\x9f\x0a\xed\x17\x17\x6f\xb2\x71\xde\x87\x58\x58\x4c\xf8\xaf\xd0\x53\xdf\x2b\x58\xf6\x2a\xee\x9b\x43\x2c\x97\x69\x74\x9f\x71\x3f\x08\x4a\x43\x16\xe3\x52\xfc\x37\xea\x7e\x39\x13\x12\x39\x81\x52\xe6\x3e\xfa\xa0\x94\x24\xf3\x38\xd5\xd1\x3d\x21\x31\xf6\x3d\x87\x15\x4c\x99\x22\xed\xd2\xb4\x42\x40\x83\x51\x97\x78\xc5\xfc\xd7\x4c\x24\x68\x52\x1f\xb7\x06\x6a\x2d\xf1\x70\xb9\x7c\xcf\x58\xbe\x17\x9a\x35\xe3\x3c\x69\x3f\x36\xba\xf3\x4c\x9e\xb2\xab\x5b\xf1\x0b\x4e\x9d\x27\xa3\x94\xa5\xeb\x94\xc8\x38\xb9\x89\x30\x5f\x3a\xe4\xc9\xb0\x6c\x64\x09\xc1\x71\xf7\x87\xe0\x75\xf6\xae\xcd\xbc\x1f\x5e\x51\xb1\x50\x47\x71\x30\x4c\xf6\x08\xdc\x03\xb2\x63\x31\xd9\x63\xb5\x16\x21\x3e\x20\xdf\xed\x80\xdb\xa8\xbf\x6e\x25\xb2\x56\x68\x63\x9c\xd7\xe3\x5a\x59\x1f\x6c\x8f\x6b\xf5\x1c\xa4\xa7\xd8\xca\xd5\x7b\x43\xce\x9f\x18\x0e\x4f\x0a\x7d\x16\xb7\xed\xa1\x0f\x13\x4c\x0f\xc1\xe0\xd4\xd8\x68\x79\x19\x28\x36\x16\xc4\xba\x0d\x8b\xc4\x8b\xde\xbd\xac\x00\x19\x5b\x33\xdf\x1f\x20\x77\x08\x77\x6c\x94\x3c\x3b\x3b\x3b\xdb\x15\x21\xc7\xb4\x3d\x40\x19\x42\xfe\x99\x00\xa8\x15\xc1\x3e\x71\x13\x44\xfe\xce\xdd\x46\x70\x4d\x13\x6b\x40\x3e\x09\xba\x1f\xd8\x63\x23\xe4\xba\x05\x3b\x83\xec\x56\x6c\x9b\x8f\x6b\xc5\x55\x37\xda\x5d\xae\x56\x6a\xa1\x9b\x66\x07\xde\xfb\x26\x10\x5e\x4e\x89\x3d\x06\x6f\x51\x57\x6b\x52\x69\xc1\x5d\x51\xbd\xa1\xcb\x1e\xd1\xfd\xb6\x65\x87\x4f\x4a\xd5\xfd\x92\x77\x4b\xdc\xe0\xb6\x1e\xec\xdb\x16\x37\xa4\x33\xbe\x38\x77\xec\x52\x65\x09\x90\x41\x27\xd3\xc6\x81\x8a\x7e\x87\xbf\x5e\xd3\x45\x3e\xd7\x6e\x8b\xef\xe2\xe4\xcf\x31\xed\x87\xc0\xaa\x01\x3c\x4a\x3f\x1e\x78\x9f\xcc\xd7\x66\x13\x74\xd6\x41\x6d\x97\x74\x6b\x0e\x51\xf7\xf4\x75\xbe\x83\xb0\xfc\x32\x47\xb9\xf2\x2a\x3e\x9f\xa7\x16\x6a\x7b\xcb\x49\x2c\x07\xdd\x45\xd0\x3f\x51\xff\x9c\x24\xf4\x48\x0a\x7c\x44\x05\xcc\x5c\xd9\x9a\xa4\xbe\xaf\x63\x19\xaf\x6d\xc1\x55\x22\xda\x7b\xef\xfd\x5b\x2c\xbd\x9a\xe1\x4b\x2d\x7a\x2c\x4e\xdb\x8b\x9e\x5c\x72\x94\x65\x6b\xc1\x5c\x8d\x56\xe5\xf5\x8c\x4d\x90\x32\x5b\x08\x12\xd5\x2c\xcf\x14\xbe\x43\xf9\xce\xdd\x0c\x00\xfc\xdf\x7e\x3f\x00\xc4\xd0\x92\x6a\x36\x18\x0f\x2e\x9b\xac\x34\x03\x95\x4d\x05\xb1\x97\x51\x0e\x15\x96\x7e\xa4\x6b\x75\xb8\xfb\xb6\xe4\x13\xc2\xdf\xff\x11\x6e\xf1\xfc\xf5\x9a\x60\xf0\x49\x96\xd7\x64\x7b\x0e\x57\xf0\x5c\x5b\xd0\x03\x08\x60\xb3\xa9\x48\x3e\x32\x09\x52\x68\x4c\x15\x1c\x64\x2b\x75\xb6\xcc\xf4\x10\x74\xae\x59\x12\xb6\x22\xb2\x55\xa0\xb3\xca\x7a\x68\x36\xa6\x5c\x5a\xb4\x35\xe5\x27\x66\x1c\x5c\x3f\xa9\xbf\x7a\xd6\x55\x3d\x94\x54\x8f\x46\xdd\x9d\x3e\xf7\x72\x81\xa2\x15\x1f\x0d\xae\xd1\xb6\x56\x2b\x0c\x3b\x0a\xa1\x71\x2e\xe1\x8f\xd0\x40\x4e\x28\x4b\x96\x91\xb3\x1b\xfd\xb5\x20\x21\xcd\x1e\xba\x65\x6a\xb3\x56\x25\x58\x53\x31\x14\xbd\x65\x5a\x6f\x67\x99\x7b\x68\x6e\xb5\x74\x8b\xad\xb1\x69\xad\x93\xa0\x5d\x2a\x39\x52\x1b\x44\xd7\xb3\xd8\x1d\xa5\xa8\x9d\x2e\xb8\x55\x5b\x66\x86\xd5\xce\x2b\x60\xb3\x19\x66\xdc\x97\xce\x5d\x4c\x6d\xd9\xa1\x65\xb7\x8a\x6c\x7a\x56\x59\x74\xae\xd7\x80\x89\xc2\xc1\x1d\x79\xbf\x07\xdf\xac\x0e\xf6\xe1\x10\x28\x0b\xfa\x32\x5f\x28\xd3\xb2\x5f\x46\xef\xf3\x85\x6a\x6d\xce\xeb\xe6\x4d\x23\xa3\x07\x5c\x6a\x3f\xd8\x66\x57\xa5\xa0\xa7\x18\x76\xd3\xb8\xcd\x8a\x76\xda\x87\x98\x65\xfe\x8f\xb2\x97\x85\xf7\xb1\x72\x94\xb2\x69\x6b\xae\xf0\x2d\x6e\x99\xcb\x0e\xda\xf4\x21\xcb\xf4\xae\x6c\xe2\xda\xb2\xf5\x2a\x87\x59\xd3\xbf\x82\x52\x3b\xb6\x0b\xef\x6f\xdb\xaf\x19\x73\xdb\x6e\xde\xd0\xb1\x19\x2b\xa7\x6d\xf5\x89\x0e\x0b\xb7\x28\xee\x73\x8f\xd6\x86\xca\x32\x77\xdd\xc8\x2f\x03\x45\xfe\x67\x0a\x33\xc3\x07\xfd\x1e\xe1\x20\xe3\x26\x1a\x98\x69\x55\xf9\x7d\xb3\xb2\x2f\x9d\x6b\x25\xf8\xee\xf7\x2a\xa6\x87\x01\x9f\x71\x65\x6a\x72\x5b\x1e\x1b\x62\xf5\xf7\xdf\xdf\x51\x5d\xee\x00\xdc\x5e\x9b\x13\x4c\x45\x25\x6e\x38\x2e\x5f\xca\x80\x7f\x50\xd9\x3d\x40\xc9\x3d\x5c\xbd\xdd\x36\xf7\xcf\xb8\x72\xe2\x7d\x13\x55\xb8\x8b\xbf\x5d\x52\x0d\x56\xa4\x1f\x60\x00\xee\xa8\x42\xbb\xa1\xa4\x31\xdd\x92\x8b\xeb\xf9\xb7\x26\xc4\x75\xa3\x23\x72\x52\xa7\xe3\xab\x2a\x90\x87\x42\xf3\xd8\x6a\xf8\xec\x4c\x8c\xa9\xe0\xb6\x4d\x64\x53\x05\x77\xf5\x8d\x8b\x03\x19\xc1\x35\xfc\x60\x5b\xc7\xb5\x0e\x75\xe2\x94\x99\x4b\xe3\xd2\xbe\xf7\x36\x4b\x56\x7b\x8f\x76\xb0\x24\xc9\x17\xd6\x11\xce\xce\xb6\xbc\xfe\xeb\x5f\x95\x7f\x7d\x69\xf2\x94\x6a\x7c\x28\xab\x68\x68\xa7\xdd\xbb\xea\x9f\x35\xcb\x97\xf2\xf5\x46\xd6\x77\x96\x29\x7b\x1e\x57\xf8\x96\x32\xe3\x51\x27\x1b\x5e\x46\x96\x7c\xbe\xd3\x0d\x27\xfb\xf0\xc1\x2d\xaf\x1e\x89\xb5\xe3\xf4\xc4\xf7\xd3\x8c\x3a\x56\x2b\x4f\xb6\x2e\xa7\x76\xa1\x8e\xed\x40\x7d\x5b\xf9\xee\xff\xa1\x99\xa7\xb9\x0f\x5a\xad\xa5\xe1\xb4\xb8\x5e\xc3\xdf\x66\x76\x6f\x49\x01\x8d\xfe\xb9\x59\xd1\xb0\xca\x5b\x8b\xb7\x91\x5e\xc3\x57\xcd\x44\xcd\x26\x34\x6d\x82\xfa\x23\x9b\x94\x84\x3c\xf5\x25\x29\x5e\x0a\x76\xd9\x4c\xe5\xdd\xcd\xc6\x55\xd1\x91\xaa\x1a\x52\x1d\xfd\xa8\x46\x8f\xa8\xd1\x22\x32\xed\xbb\x53\x7b\xa0\xa7\x74\x87\x7a\x74\x86\x0e\xea\x0a\x75\x74\x84\xba\xba\x41\xfd\x1a\x39\x21\x78\xeb\xb5\xd1\x60\xf4\x1f\xaa\x54\x36\x9b\x3e\xc7\xd1\xfa\x40\xf7\x75\xf7\x75\x9e\xc1\x4b\x9b\x1d\x9f\xcb\x4b\xf8\x75\xc6\x99\x46\x98\xab\xbf\x0a\xd6\xbd\x05\xab\xc5\xaa\x57\xcd\xfa\xfc\x07\xcf\x2c\x73\x03\xd5\xad\x05\xb1\x6f\xe6\x50\xee\x58\x64\x42\x4d\xc3\x82\x8d\x26\xc5\x63\x8e\x75\xb6\xaa\xbc\x81\xce\x60\x50\xca\x34\x91\x68\xe7\x29\xb3\xf6\x79\x35\x35\xd8\x71\x35\xb3\xb6\xcb\xb5\xca\xf1\x46\x99\xfd\xe9\x90\xd3\x4f\xe3\x40\xab\x1e\x3f\x3c\xd7\x1f\x9c\xe7\x9f\xe4\x78\xe7\xd1\xb5\x24\x4f\xe2\x1d\x97\x7f\x0e\xb0\xc6\xb9\x0d\xba\x07\x1e\x0b\x2a\x0e\xd1\x10\x8b\x07\xda\xec\x0e\x71\x9e\xe7\xf4\x90\x4b\x33\xc7\x1d\x71\xe8\xcf\x7d\x03\xa2\x9d\x67\xad\xef\x97\x18\xd7\x7f\x84\x47\x19\xc0\x18\xac\xfb\xe5\x56\x92\xe4\x0b\x9b\xdd\x70\x89\xf1\xdc\x3c\xca\xc7\xc0\x20\x9e\x2b\x9d\xa7\xd5\x78\x36\x61\x22\x53\xda\x0c\xb5\xd5\x52\xdf\x74\x42\x2c\x6c\x4f\x26\xe3\xa5\x2d\x36\x8b\x1f\x2b\x85\xee\x87\x40\x41\xd1\xd9\x38\x2d\x55\xd0\xc2\x03\x25\x0a\x4b\xea\x65\xe4\x00\xab\x48\xac\x7e\x79\x36\x50\x2d\xb7\xd7\x37\xaa\xb0\x33\x5e\xfa\x36\xf2\xd8\x35\xd4\x72\x88\x68\xf2\xa7\xc9\xb5\xc7\x6a\xde\x16\xfc\x14\x1c\xee\x77\xfc\x2e\x97\xfd\x5f\x00\x00\x00\xff\xff\x80\x2e\x76\x26\xc3\x3b\x00\x00"),
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
func FindDecompressedGzippedFile(path string) (io.Reader, int64, error){
	return FindFile(path, true)
}

// MustFindDecompressedGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindDecompressedGzippedFile(path string) (io.Reader, int64){
	reader, size, err := FindDecompressedGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindGzippedFile(path string) (io.Reader, int64, error){
	return FindFile(path, false)
}

// MustFindGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindGzippedFile(path string) (io.Reader, int64){
	reader, size, err := FindGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFile returns a io.Reader by seeking the giving file path if it exists.
func FindFile(path string, doGzip bool) (io.Reader, int64, error){
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
func MustFindFileReader(path string) (*bytes.Reader, int64){
	reader, size, err := FindFileReader(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFileReader returns a io.Reader by seeking the giving file path if it exists.
func FindFileReader(path string) (*bytes.Reader, int64, error){
  item, ok := assetFiles[path]
  if !ok {
    return nil,0, fmt.Errorf("File %q not found in file system", path)
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
func ReadFile(path string, doGzip bool) (string, error){
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
func ReadFileByte(path string, doGzip bool) ([]byte, error){
  reader, _, err := FindFile(path, doGzip)
  if err != nil {
    return nil, err
  }

  if closer, ok := reader.(io.Closer); ok {
    defer closer.Close()
  }

  var bu bytes.Buffer

  _, err = io.Copy(&bu, reader);
  if err != nil && err != io.EOF {
   return nil, fmt.Errorf("File %q failed to be read: %+q", path, err)
  }

  return bu.Bytes(), nil
}
