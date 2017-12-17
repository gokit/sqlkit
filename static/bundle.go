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
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x49\x8f\xdb\x46\x16\x3e\x77\x03\xfd\x1f\xde\x10\x83\x80\x0a\x38\x34\xe6\xda\x81\x0f\xe9\xc5\x81\x81\x89\xe3\x89\x9d\x99\x43\x10\x04\x25\xd6\xa3\x54\x36\xb7\xae\x2a\x76\x4b\x10\xf4\xdf\x83\x57\x55\x5c\xc4\x16\x45\x51\xcd\x38\x69\x2f\x07\xab\xb9\xd4\x5b\xbe\xb7\x7c\xb5\xf0\xc5\x0b\x48\x59\xf1\x4a\x60\xc2\x15\x70\x8c\x45\x86\x0a\x18\xe8\x75\x81\x10\xe7\x12\x18\x3d\x06\xbd\x64\x1a\x70\x55\xe4\xca\x3c\xb5\xaf\xfb\x33\x48\x51\x2f\x73\x1e\x5e\x9c\x9b\xf7\x1b\x41\x29\x2b\x7e\x55\x5a\x8a\x6c\xf1\x9b\xc8\x34\xca\x98\x45\xb8\xd9\x5e\x9c\x5f\x9c\xbf\x78\xe1\x46\x83\x44\x5d\xca\x4c\x81\x5e\x9a\x91\x20\xb4\xc2\x24\x06\x96\x71\x28\x64\x7e\x2f\xb8\x51\x65\x35\x80\xce\x21\x65\x3a\x5a\x9a\xb7\xd5\x5d\x12\xbe\x67\xf3\x04\x8d\x24\xa8\x35\x84\x17\xe7\x71\x99\x45\xe0\xa7\x8d\x29\xb3\xc6\x58\x7f\xbf\x55\x01\xa0\x94\xb9\x9c\xc1\xe6\xe2\x1c\x00\x9c\x5d\x90\x06\x90\x89\xe4\xe2\xbc\xb2\x7a\xb3\x09\xdf\x69\x59\x46\x3a\xfc\x69\xfe\x01\x23\x1d\xbe\x61\x29\x6e\xb7\x5d\xe4\xb2\xc6\x1c\x78\x58\x8a\x68\x59\xc3\xd6\x78\xe2\x34\x58\x68\xf3\x18\x58\x92\x18\x15\x4c\x6b\x29\xe6\xa5\x26\x39\x4a\xe5\x91\x60\x1a\x39\x3c\x08\x6d\xdd\xb6\x3a\x38\x28\x63\x46\x29\x11\x18\x29\x8e\x04\x47\x0e\xf3\xb5\x85\xa6\x7a\x56\xc5\xe4\xb0\xd9\x2d\x63\x37\x17\xe7\x67\xc7\x42\x35\x88\xca\x75\x9e\xa9\x32\x45\x79\x08\x17\x16\x45\x58\x68\xd5\xc0\xc0\x99\x66\xee\xd9\x83\x48\x12\x98\x23\x44\x56\x0e\x37\xba\x44\xa6\x73\xe3\xe4\x42\xdc\x8b\x6c\x01\x22\x2d\x12\x4c\x31\xd3\x74\x31\x05\x28\xb5\xd5\xbb\xa8\xb8\xdb\x3d\x98\xcc\x2c\x26\x83\x90\xdc\x5c\xb5\xca\xab\xb1\xd6\xfa\xeb\x12\x1e\x6e\xae\xe0\xfa\xe7\x5f\x6e\x20\x2f\x50\x32\x2d\xf2\x4c\x19\x91\xa5\x32\x2e\xde\x25\xe4\x1c\x79\x54\x66\x1c\x65\x22\x32\x04\x3e\x1f\x70\xea\xe6\xca\x69\x33\xe9\x1d\xe5\x09\x58\x1f\xe8\x4a\xad\x4c\x2d\xdd\x5c\xd1\x05\x5f\xc1\xb7\x74\xf5\xee\xbf\xff\xa1\xcb\x14\xb5\x14\x91\xaa\x7e\xc3\x1f\xed\x2f\x3d\xd2\x54\x7b\xa4\xd9\x14\xe1\x6b\x4e\x21\xd0\xeb\x06\x81\x37\xf8\x50\x57\x37\x83\x0c\x1f\x40\x64\x4a\xb3\x2c\x42\x0a\x73\xaf\x9d\x55\xf9\xbe\xc1\x07\xdf\xaa\xb0\x96\x06\x90\x76\xad\x08\x1a\xd3\x03\xd0\x29\x84\x61\x68\x46\x28\x6b\xd2\x8f\x62\x61\xe1\x9b\xc1\xb7\xbd\xea\x5c\xb9\xf3\x15\x5c\xbe\x34\xb2\x48\x6f\x4a\x92\x49\x64\x18\x86\x33\x72\xa7\xd5\x11\xbe\x19\x12\x45\xff\xd4\xea\x92\x24\x34\x37\xf8\xea\x12\x78\xfb\x46\x94\x27\x97\x16\xc2\xd6\x4d\xe7\xdf\x25\xa4\xad\x9b\xe6\xa5\xcb\x1a\x68\xd2\xb6\xa1\xff\xdc\xf0\xad\x7b\x75\xdb\x20\x7f\x83\x09\x6a\xa4\x46\x82\x29\xd5\x96\x69\x36\x69\x7e\x8f\x26\x6b\x24\x46\xb9\xe4\x10\xcb\x3c\xb5\x0d\x65\xee\x32\x8b\x2e\x5c\x0a\x72\x28\xca\x79\x22\xa2\xd7\x37\xa1\x91\xf8\xb3\x19\xa3\xea\x17\x85\xa2\x24\x4d\x4b\xa5\x61\xc9\xee\x11\x98\x7b\x1f\x04\x87\x7b\x96\x94\x18\x50\xc3\x93\xa8\x14\x72\x40\xa1\x97\x28\xa9\x0c\x19\xcc\x55\x9e\x41\x2e\xe1\x03\xfd\x6a\xb6\x30\xd2\xe9\x4f\x5b\xd0\x98\xb5\x12\xe3\x2d\x8b\x3e\xb2\x05\x6e\xb7\x61\x0f\xe4\x2e\xa5\x9b\x86\xcf\xe7\x07\x22\x3d\x73\xb8\xf8\x91\x5e\x51\x4b\xd1\xb8\xd2\xe1\xb5\xfd\x0d\x6a\x7f\x5d\xb6\xb9\x7a\xae\xd8\x20\xa5\xec\xa8\xb2\xef\x0d\x3e\xbc\x97\x2c\x42\xdf\xeb\x4f\x62\xab\xca\x9b\xb9\xec\xc2\x18\x25\xa4\x7c\x1e\x56\x32\x6e\x53\xa1\xfd\xea\xe2\x75\x16\xe7\xc7\x08\x0b\xaa\x01\xff\x17\x7a\xe9\x7b\x95\xc9\x5e\x63\xfd\xee\x2b\xd6\xca\x34\xbc\xcd\xb8\x3f\x9b\xd5\x89\x2c\xe2\xda\xfd\xd7\xea\x76\x55\x08\x89\x9c\x40\xa9\xb9\x8f\xfe\xa1\x94\xe4\x73\x9c\xea\xf0\x96\x90\x88\x7d\xcf\x61\x05\x4b\xa6\x28\xba\x34\xac\x72\xd0\x60\xd4\xe7\x5e\x35\xfe\x15\x13\x09\x1a\xea\xe3\x36\x41\x6d\x26\x8e\xf7\xcb\xf7\x4c\xe6\x7b\x81\xd1\x19\xe5\x49\xf7\xb1\x89\x9d\x67\x78\xca\x6a\xb7\xee\x57\x96\xba\x4a\x46\x29\xeb\xd2\xa9\x91\x71\x7e\x93\x60\xbe\x72\xc8\x53\x62\xd9\xce\x12\x80\xb3\xee\x77\xc1\xdb\xe6\x7d\x67\xc6\xfd\xe3\x25\x4d\x16\xda\x28\x4e\x86\xc9\x80\xc3\x47\x40\x76\x2a\x26\x03\x59\x6b\x11\xe2\x13\xda\xdd\x6d\xb8\x3b\xf3\xaf\x6b\x89\xac\xd3\xda\x18\xe7\xed\xbe\x56\xcf\x0f\xf6\xf7\xb5\x36\x07\xe9\x25\x76\xb8\x7a\xb0\xe5\xfc\x85\xed\xf0\x49\xad\xcf\xe2\xb6\xbf\xf5\x61\x82\xe9\x18\x0c\x9e\xda\x1b\xad\x2d\x13\xf5\xc6\x4a\x58\x7f\x62\x91\x7b\xe1\xdb\xe7\xd5\x20\x23\x9b\xe6\xc3\x0d\xf2\x80\x73\xa7\x76\xc9\xb3\xb3\xb3\xb3\xde\x6e\xb0\xd9\x10\x3a\xfe\x92\xa9\x57\x94\x77\x2e\x2a\xe0\xd9\x35\x83\x37\x83\xed\xb6\x71\xf8\x51\x3b\x7d\xc7\xee\x77\x9a\x29\x59\xdf\xdb\x3b\xf7\x42\xb6\xfb\xb8\x35\x67\xea\x07\xb1\xaf\x82\x6a\x70\xfb\x65\xf6\xc0\x38\x34\x80\xdc\x72\xb1\x39\xe2\xe5\x3d\x51\xe8\x0c\xaa\x13\xb3\xaf\x59\xef\x84\xe8\x88\xa6\x7d\xdd\x49\xaf\x47\x33\xd0\x61\xcf\xfb\x3d\x9e\x55\x99\x82\x89\xc2\x76\x3a\x98\x9a\xca\xb4\xc9\x87\x6a\x61\xee\x6f\x36\x74\x91\x97\xda\xad\x45\x5d\x41\x7f\x1f\xd1\xc4\x1d\xac\x60\xf0\xa8\x4f\x7a\xe0\x7d\x30\x3f\xdb\xed\x6c\x44\x8e\x39\xad\xe3\xd2\xec\xe9\x49\x35\xaa\x12\xef\x4a\x94\x6b\xaf\xb1\x75\x1c\x75\xff\x19\x19\xf1\x24\xf3\x9b\x0c\xc8\xf8\x76\x7b\x88\xce\x7f\x40\xfd\x7d\x92\xd0\x23\x29\xf0\x9e\x16\xc5\xe6\xca\xb2\x6b\x7b\x85\xc2\x32\xde\x5a\x4c\xaa\x44\x74\x57\x91\xc3\x8b\x05\x5a\x1e\x3f\x57\xfa\xb6\x38\xed\xa7\xef\x5c\x72\x94\xf5\x22\xd9\x5c\xcd\xd7\xf5\x75\xc1\x16\x48\x73\xa2\x00\x24\xaa\x22\xcf\x14\xbe\x45\xf9\xd6\xdd\x9c\x01\xf8\xbf\xfe\x36\x02\xc4\xc0\x8a\xda\xdd\x2a\x1b\x3d\x01\xb0\xde\x4c\x34\x01\xa8\x84\x3d\x0f\x62\xaf\x32\xfd\xc4\x4a\xeb\xe9\x04\x8f\xa6\xf1\x99\x48\x02\xf8\xd7\xbf\x83\x3d\x0c\x7e\xcf\x24\xe8\x5c\xb3\x84\x42\xd9\xdc\x92\x42\x63\xaa\x60\x54\x36\x0c\x4d\x0a\xdc\x96\x59\x77\x56\x40\x8a\x02\x6b\x43\xd0\x69\xde\x36\x98\x2e\x43\xdb\x5d\xdc\xa4\x75\x9d\xdd\x36\xad\x1f\xa5\xf4\xec\xbb\x47\xbc\xf0\xe4\x7e\x5f\x07\xac\xdd\x99\xfa\xf7\xaf\xdc\x96\x39\x75\x2e\x3e\x9f\x28\xba\x7b\xfa\xbc\x89\x6f\x83\x60\x4f\xcf\x8f\x73\x09\xbf\x07\x06\x70\xc2\x58\xb2\x8c\xca\xde\xc4\xb9\x03\x08\x65\xc0\xd8\x65\xc0\xae\x84\x86\x87\xcd\x94\xb4\xda\x2d\x25\x6d\x07\xa9\x77\x2c\xfd\x5a\xb9\xd5\x62\xcf\x6c\x16\x93\x9b\x7d\xe1\x38\x31\x12\x24\xd7\xb3\xc8\x9d\x10\xa4\x83\x85\xd8\x8a\xd4\xce\xfb\x36\x2e\x2f\x81\x15\x05\x66\xdc\x97\xae\x4c\xcc\x94\x6a\x27\xbe\xb0\x77\x72\x55\xb0\x4c\x44\x3d\x7d\xb2\x02\xc3\x30\x5c\xbd\x5f\xbe\x1f\xb3\xce\x26\xb8\x77\x90\xcb\xe5\x6e\x2d\x77\xa9\xfd\x6a\x6d\x8f\x66\x5a\xf4\x7e\x78\xf7\x31\x36\xaf\x7f\xc4\xb5\xe1\x7b\x4b\xbd\x46\x58\xfb\x94\xe8\x0b\xe2\x7c\x07\xe0\x7e\xde\x27\x98\x2a\x96\x37\x16\xd7\x5b\x97\xe0\x8f\xa2\xf4\x09\xe8\x7c\x3a\x2e\xef\x56\xea\x47\x5c\x3b\xf7\x3e\x0f\x86\xef\xf3\x67\x32\xea\x1f\x11\x7a\x77\x94\xb7\x67\x82\x30\x9e\x0b\x4e\x9a\x09\x74\xd7\x6e\x3f\xa0\x6e\x93\xfe\x37\x64\x86\x49\xf4\x0a\xb2\x89\x16\x71\xa3\x48\xbd\x87\xd9\x3b\xca\xc7\xec\x18\x8c\x0c\xec\x04\xc1\xfd\xfb\x51\x87\x0d\x6d\x97\x31\x8e\xa7\x8a\xfa\xbc\xa6\xbd\x32\xfc\xc2\xe8\xe1\xc8\x93\xac\xcf\x89\x0e\x4e\x3a\xf4\x7a\x26\xd4\xf0\xc9\x0e\xbe\xfe\x0e\x14\x01\xbd\x1c\x01\xb0\xcb\x12\x9b\x0d\xfc\xb3\xb0\x33\x43\x8a\x0c\xfd\x71\xb5\x26\x99\xcd\xd0\x6a\x03\xdc\x7b\x3c\x50\xb3\x05\x0d\x5b\xa0\x7e\xcf\x16\xb5\x20\x4f\xdd\x25\xd5\x7e\xe2\x76\x67\xcd\x74\x24\x27\x79\x9b\x8d\x91\x1d\xfe\x8f\x9a\xc1\x76\x7b\xcc\x61\xe0\x57\x92\x7a\x16\x24\x35\x44\x53\xbf\x14\x9c\x69\x84\x52\x7d\x25\xa9\x41\x92\xb2\x58\x1d\xc5\x53\x9f\xfe\x1c\xd2\x1a\x37\x11\x57\x55\xc2\x3e\x9b\x6f\x34\x62\x91\x09\xb5\x0c\x2a\x33\x76\x25\x9e\x72\xca\xdf\xf9\x64\x62\xa2\xcf\x13\x4e\x62\x87\xd1\xcc\xd0\x4b\x5a\xee\xdc\xb3\xcb\x59\x8f\x58\xc4\x15\x42\xe7\xec\xf3\x93\xf0\x48\x69\xfb\xd5\xc8\x13\xab\xea\x4c\x87\x33\xcd\x46\x86\xfb\x80\x4b\x93\x1d\x6c\xed\x27\x09\xb3\x1d\xf8\xe7\x1e\x2f\x3e\x0e\x24\x29\xfd\x1a\xc8\x27\x06\x72\x97\x7c\x07\xba\xb0\x63\xe0\xd3\x8e\x50\x8e\x77\x6a\x07\xb9\x83\x5f\x25\xdd\xae\x30\x6a\x7f\xae\x4e\xe4\x68\x12\xcc\x7d\xe3\x9c\x24\xf9\x83\x25\x7e\x5c\x61\x54\x9a\x47\x79\x0c\x0c\xa2\x52\xe9\x3c\x6d\xde\x67\x0b\x26\x32\xa5\xcd\xab\xc6\x8f\xa3\x99\x96\x4c\xd8\xcf\xb3\xf1\xca\xc8\xf7\xab\xcf\x7a\x03\xf7\xc9\xec\xac\x5a\xe8\x3d\x8d\x45\x49\xf1\x44\x1c\x6a\x45\x3d\x0f\x7a\xb4\x81\xc4\xe6\x1b\xed\x89\x0f\x7b\x8e\xf8\x46\x31\x5e\xf9\xb6\x29\x59\x1d\x6a\x35\xc5\x67\x88\x7f\x99\x5f\x03\x59\xf3\x53\x65\x4f\x65\xe1\x70\xe1\xf7\x95\xec\x1f\x01\x00\x00\xff\xff\x30\x11\xcd\x52\xed\x32\x00\x00"),
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
