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
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xc1\x6a\xc2\x40\x10\x86\xef\x79\x8a\x39\x2a\x84\xe4\x15\xda\x28\x95\x5e\xaa\x50\x7a\x2a\x3d\x6c\x36\xbf\xc9\xd6\xcd\xee\xb2\x3b\x5a\x83\xe4\xdd\x8b\x26\x95\xb4\x18\x69\xed\x69\x61\x32\x33\xff\xf7\x0d\x49\x53\x3a\x1c\x92\x67\xf6\x5b\xc9\xc9\x32\x7f\x87\xe4\xe4\x49\xd4\x68\xdb\x4c\xc8\x0d\x4c\x41\x05\xd6\xca\x20\x90\xa0\xbc\xaf\x7c\x54\x4a\x56\xe4\xe1\x3c\x02\x0c\x07\xe2\x0a\x54\xaa\x9d\x32\x25\x45\x69\x4a\x35\xb8\xb2\x45\x20\xec\x9d\x0d\x28\x28\x6f\x4e\x1d\xf3\x8c\x54\xed\x34\x6a\x18\x16\xac\xac\xa1\xb5\xf5\xc3\x59\x6e\x1c\xc6\x68\x92\xe3\xe2\xbb\xf3\x7c\x74\xad\xf7\x8b\x5c\x19\x86\x5f\x0b\x89\x43\x44\x34\x87\x06\x63\x22\x79\x4f\xd2\x1a\xc6\x9e\x93\x59\xf7\xc6\xe4\xb6\xb9\x56\xf2\x71\x4e\x81\xbd\x32\xe5\x94\xe0\xbd\xf5\x11\xd1\xcc\x43\x8c\x0d\x41\xa3\x1e\x10\xac\x84\xdc\x88\xf2\x48\x3a\x42\xd5\x6f\xa5\x88\x68\x01\xfe\x25\x08\x4d\xfe\x90\x10\x53\x17\x31\x8d\x88\x5e\x5c\x31\x4a\xfe\x23\xe5\x66\x95\xce\xe4\x5e\xeb\xac\x59\xfa\x02\xfe\x72\x9a\x3d\x7e\xea\x9f\xac\x19\x98\xbd\xbe\xdd\xe8\xb6\x00\x67\xcd\x83\x82\x2e\x2e\x27\x6e\xd0\x9c\xd5\x76\x42\x6f\x31\xf8\x13\xda\x7f\xdc\xb4\x93\xbd\x62\x79\x4e\xfd\x2e\x1b\x93\x13\xe5\x09\x22\x26\x8f\xe0\xac\x09\x58\xc1\xaf\xfa\xe2\x2d\xb7\x38\xed\xea\xc1\xda\xcf\x00\x00\x00\xff\xff\x0b\xd5\xde\x04\xc4\x03\x00\x00"),
          path: "sql-api-backend.tml",
          root: "sql-api-backend.tml",
        },
      
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
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\x4f\x6f\xda\x40\x10\xc5\xef\x48\x7c\x87\xa9\xb8\x98\x8a\x2e\xf7\x48\x3d\x00\x6e\xa3\x48\x51\x4b\x42\x73\x42\x91\xbc\xac\x07\x7b\xcb\x7a\xc7\xdd\x1d\xb7\x8e\x90\xbf\x7b\xe5\x3f\x29\x90\x3a\x4d\x50\xe3\xcb\xee\x3e\x79\xe6\xfd\xe6\xcd\x7e\x2f\x56\xec\x0a\xc5\xe2\xeb\xe6\x3b\x2a\x16\x5f\x64\x86\x55\x05\xab\x9b\xeb\x70\x0e\xb3\xe5\xd5\x70\xf0\xf1\xe5\x6f\x38\x18\x0e\xd6\xef\xd6\x97\x04\xb7\x98\x93\x63\x58\x48\x17\xdf\x07\x29\x73\xee\x2f\xa6\xd3\x84\x5c\x23\x2b\xe9\x62\xa1\x28\x9b\x6e\x64\x9c\xe0\x74\xbf\x17\x4b\xa9\x76\x32\xc1\xa5\xe4\xb4\xaa\xc6\xff\xa8\x68\x9f\x7f\x97\xd4\xce\x2f\xce\x00\xda\x83\x04\x59\x30\x7d\x48\xd0\xa2\x93\x8c\x31\x2c\x6e\xef\x42\xd0\x59\x6e\x30\x43\xcb\x92\x35\x59\xd8\x92\x03\x4e\x11\xa2\xde\x96\x5d\xdf\x08\xb4\x85\xbc\xa5\x68\xfe\x5c\xee\x12\xd1\xe2\x44\xa2\xe6\xf9\x96\x22\x6c\xc9\x18\xfa\xa5\x6d\x02\x19\x72\x4a\x31\x60\xa9\x3d\xfb\xc6\x41\x15\x9e\x29\x03\xca\x6b\x12\x4d\xd6\x5f\xd4\x55\xa3\x11\x7c\x2a\x51\xd5\xd7\x28\x8a\x12\x1a\x0e\xea\x67\xa0\xb8\x04\x45\x96\xb1\x64\xb1\x68\xcf\x09\x6c\x4b\xd8\x16\x56\x05\xef\xfd\x0f\x23\x56\x37\xd7\x13\xa8\x2f\xe1\x7c\x0c\xe8\x1c\xb9\xee\x68\x1a\x3d\x07\xe4\x1f\x89\xb4\x6d\x46\x3e\x04\x53\x07\x26\x3d\xe4\xe8\x58\x6a\x5b\x57\x30\x35\x69\x3d\x62\x2e\x1c\x4a\xc6\x23\xd0\x56\xe8\x47\x45\x83\x19\x1c\xe2\xec\xb6\x57\x55\xe2\x99\xad\x3d\x65\x1f\x8d\xe0\x12\x79\xfe\xf0\x59\xa3\x89\x8f\x3c\x0f\x62\xbf\xef\x0e\x1f\xc0\xb3\xd3\x36\x99\xc0\x4f\x69\x0a\xec\x5e\x63\x80\xe0\x0c\x9c\x09\x74\x91\x3e\x01\x3a\x25\xe9\x47\xc8\x8b\x8d\xd1\xea\x2a\xfc\xe3\x7c\x9e\x71\xaf\x2f\xcc\x8c\x39\xf5\x9e\x19\xd3\x67\x3f\x86\x60\x7d\xff\x9f\x7e\x77\x79\x7c\xba\xe8\x56\x78\xd5\xb4\x6f\xb2\xf9\x10\x0d\x9e\x00\xb4\xc2\x2b\xe3\x3e\x6e\xf7\x3b\x00\x00\xff\xff\xd0\xde\x8b\x84\xe8\x04\x00\x00"),
          path: "sql-api-readme.tml",
          root: "sql-api-readme.tml",
        },
      
        "sql-api-test.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x5d\x4f\xe3\x46\x14\x7d\x8f\x94\xff\x70\x6b\xb5\x92\xbd\x75\x87\xe5\x35\x55\x1e\x08\xde\x45\x5b\x55\x49\x8a\x43\xf7\x31\x1a\xec\xeb\x30\xbb\x63\x0f\xcc\x5c\x27\x20\xcb\xff\xbd\x1a\xdb\xf9\xe0\x23\x6c\x1c\xa0\x52\x69\x90\x40\x60\xdf\xcf\x33\xe7\x1c\xec\x14\x05\xfc\x7c\x9d\x08\x94\x31\xf4\xfa\x50\xfd\x32\xb8\x1b\xf2\x14\x81\x85\xa4\xf3\x88\xc0\x19\xe7\x97\x52\x44\x5f\x02\x07\xca\xb2\xdb\xa9\x12\x88\xcf\x6c\xf8\x0c\x69\xc2\x67\xab\x02\x8e\xb9\x91\x0e\x38\xdf\x8c\xca\xea\xd8\x6e\x67\xce\x35\xb8\xdd\x0e\x00\x00\xce\x31\x23\x03\x7d\x48\x91\xb4\x88\x0c\x1b\xe2\xc2\x8d\x72\x43\x2a\x65\x21\xf1\xe8\x7b\x20\xcc\xb5\xe4\x77\xae\x32\x2c\xa4\x58\xe5\xe4\x79\xb6\x84\xcd\x8d\x54\x96\x88\x19\xf4\xc1\xdc\x48\x76\x5a\xfd\x51\xd4\x77\xec\x57\x30\xb0\x03\xf7\x40\x19\x76\x86\x84\xd9\xdc\x75\x8a\xa2\x19\x9f\x8d\x79\xf4\x9d\xcf\xb0\x2c\xa7\xe1\x5f\x7f\x4e\x27\x9f\xc2\xc9\x34\x18\x38\x9e\xbf\x4e\xbf\x30\xa8\x77\x4f\xbe\x08\x3f\x9d\xdf\x4b\x0f\x06\x5f\xc6\xbb\xa7\x9f\x04\xc1\xc3\xf4\xb1\xd2\xb4\x7b\x81\xf1\xe8\x7c\xf2\xa0\x40\xa0\xc5\xbc\xcd\x0a\x75\xfc\x23\x0c\xc6\xdc\x98\x85\xd2\x71\x8b\x59\x4e\xc2\xf0\xeb\xe8\x3c\x58\x95\x2a\x97\x07\x46\x68\xe8\x54\x49\xe8\x83\x53\x14\x52\x2d\x50\x2f\xf9\xc4\x46\x97\xdf\x30\x22\x66\x8f\xac\xfa\x51\x96\x53\x1b\xed\x74\x3b\xd5\x71\x1f\x1d\xc1\x04\x0d\x9d\x21\xad\x5b\x6f\xa4\x94\x25\xcc\xb9\x14\x31\x27\x34\x40\x57\x08\xda\xb2\x09\xe7\x5c\x82\x4a\x80\xc3\x96\xa4\xaa\xae\xc6\x48\xe9\x18\x12\xad\x52\xe0\x96\x4a\xf1\x25\xeb\x76\x92\x3c\x8b\x7e\xd0\xd2\x25\xf8\x60\x67\x14\xd9\x8c\x4d\xbc\x86\x79\xfc\x5a\x58\x0d\xd4\x65\x2c\x97\x9b\x9d\xfd\x86\xe9\x7e\x45\xd6\x21\x2e\x82\x81\x5b\xf3\x77\x79\x67\x83\xd7\x74\x6b\x6b\x44\x2a\x23\xbc\x25\xf6\x55\xd0\xd5\x44\xa4\xa8\x72\x72\x97\xd7\x86\xb8\xf8\x9b\xcb\x1c\x07\x7c\xe6\x7a\x3e\x1c\x7f\x84\x0f\x40\x22\x45\x16\x62\xa4\xb2\x78\x55\x0a\x25\xa6\x3e\xa0\xd6\xb6\xa0\x54\x3c\xfe\x23\x1c\x0d\x3f\x2b\xed\xfe\x18\xff\x53\x8d\x9c\xd0\xc6\x7b\x75\x2d\x91\x54\x85\x7e\xea\x43\x26\x24\x6c\x28\xcd\xae\x68\xd8\x67\x2e\x24\xc6\xae\x13\xe6\x51\x84\xc6\x24\xb9\x94\x77\x55\x4b\x8c\xc1\x56\x81\x44\xe9\x6d\x07\xd1\x1c\x42\x0f\x7e\xf9\xf5\x86\x39\xd5\xc0\xde\x8a\x3b\xeb\x16\x96\x8b\x2f\x6c\xe1\xac\xa0\x89\x31\x41\x6d\xcf\x8b\x05\x28\x91\xd0\x8d\xe8\xd6\xaf\x00\x63\x4b\x6f\x5b\xc5\x36\xab\xf7\xfa\x55\x7c\x8d\xcc\x3a\xde\xfb\xbd\x3d\x30\x3c\xb6\x43\x2f\xb9\xf7\xcc\xd8\x22\x23\x05\xf1\xe5\x5e\xd0\xb4\x6d\xc2\xd6\xe8\x4c\x6b\xd6\xd4\x0b\x9f\x21\x3d\x8d\xce\xbe\xb4\x68\xe4\x89\x31\x18\x52\x7a\xb7\x19\x2b\x81\xee\x09\xc4\x0b\xfa\x55\x98\x94\xf7\x3d\xe8\x44\xca\x7d\x6c\x48\xca\x17\x1a\xd1\xf6\xbe\x07\x2f\x3a\x78\xd1\xbb\xf5\xa2\x3a\xc9\xf8\x4b\x53\xea\xad\x5c\xe9\x44\xca\x7a\x75\x87\x9b\xc8\xf1\x9d\xa2\xa8\x9e\x42\x59\xc5\xc7\xb2\x74\x7c\xf8\xed\xd8\x7e\xbf\x8a\x55\x59\xfd\x36\xa3\xfc\x0b\x46\xd5\xb2\xdb\x06\x5c\x22\x01\x89\x99\xdb\x24\x7b\xd0\xef\xc3\xc7\xf6\xcb\x92\x44\x6e\x08\x8e\xdb\x5a\x65\xeb\x3d\xf7\x6d\xb4\xe9\xc9\x5b\xe2\x6b\x71\x3c\xf0\xe4\xc8\x5e\x14\x2a\xdb\xf9\xc9\x70\x21\xe8\xea\x29\x43\x7e\xb6\xe9\xc1\x90\x0f\x86\xfc\x0e\x0d\x79\x07\xd1\x5d\x5c\xc7\x8f\x45\x97\xd7\x17\xdf\x48\x72\x75\xcb\x83\xe4\x0e\x92\x7b\x87\x92\xab\xab\x1e\x1d\x4d\x46\xc1\xa8\x07\x8d\xba\x8c\x4a\x91\xae\x2c\xd3\x9f\x86\xa4\x51\xc4\x63\x08\xf7\x47\xa8\x96\xf0\x9b\x63\xd4\xbe\xcd\xe1\xad\x75\xff\x27\xa4\x5a\x6a\x8f\xde\x5a\x53\xf5\x1a\x1f\x9d\x3d\xdb\xf3\xe0\xd6\xff\x69\xb7\xfe\xff\x39\xf0\xfd\x8d\xb7\xff\x8f\xda\x63\xf5\x4a\x70\x6f\xbf\x7c\xfb\x36\x4b\x1b\xf9\x27\x00\x00\xff\xff\xae\x6c\x2c\xd9\x10\x1a\x00\x00"),
          path: "sql-api-test.tml",
          root: "sql-api-test.tml",
        },
      
        "sql-api.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x59\x6f\xdb\x48\x12\x7e\xb6\x01\xff\x87\x5a\x62\x31\xa0\x06\x5c\x06\xfb\xea\x81\x1f\xc6\x47\x06\x01\x76\x3c\xd9\x24\xb3\xfb\x30\x18\x04\x2d\x76\x51\xea\x84\x87\xd2\xdd\xb2\x25\x08\xfa\xef\x8b\xea\x83\x97\x44\x49\x94\x95\x6c\x9c\x89\x81\x44\xa6\xd8\x5d\xc7\x57\xd5\xf5\x55\xb3\xe9\x17\x2f\x00\xa5\x2c\xa5\x82\x38\x8e\x2f\xce\x1f\x98\x84\xf0\xe2\x1c\x00\xe0\x4e\xca\xfb\x52\xbf\x2c\xe7\x05\x87\x2b\x37\x28\xbe\xc7\xc7\x30\x90\x98\x94\x92\x43\x51\x6a\x48\xe9\x76\x30\xba\x38\x1f\x5d\x9c\x5f\x9c\xbf\x78\x01\x39\x9b\xbd\x14\x98\x71\x05\x1c\x53\x51\xa0\x02\x06\x7a\x39\x43\x48\x4b\x09\x8c\x6e\x83\x9e\x32\x0d\xb8\x98\x95\xca\xdc\xb5\xc3\xc3\x11\xe4\xa8\xa7\x25\x8f\x2f\xce\xcd\xf8\x5a\x50\xce\x66\x7f\x28\x2d\x45\x31\xf9\x53\x14\x1a\x65\xca\x12\x5c\xad\x9d\x3e\x37\x46\xa2\x9e\xcb\x42\x81\x9e\x9a\x99\x20\xb4\xc2\x2c\x05\x56\x70\x98\xc9\xf2\x41\x70\xa3\xca\x6a\x00\x5d\x42\xce\x74\x32\x35\xa3\xd5\xa7\x2c\x7e\xc7\xc6\x19\x1a\x49\x50\x69\x88\x2f\xce\xd3\x79\x91\x40\x98\xd7\xa6\x8c\x6a\x63\xc3\xed\x56\x45\x16\xa8\x11\xac\x2c\x88\xd6\x2e\xc8\x23\x28\x44\x76\x71\xee\xad\x5e\xad\xe2\xb7\x5a\xce\x13\x1d\xff\x36\xfe\x80\x89\x8e\xef\x59\x8e\xeb\x75\x17\xb9\xa2\x36\x07\x1e\xa7\x22\x99\x56\xb0\xd5\x9e\x38\x0d\x16\xda\x32\x05\x96\x65\x46\x05\xd3\x5a\x8a\xf1\x5c\x93\x1c\xa5\xca\x44\x30\x8d\x1c\x1e\x85\xb6\x6e\x5b\x1d\x1c\x94\x31\x63\x2e\x11\x18\x29\x4e\x04\x47\x0e\xe3\xa5\x85\xc6\xdf\xf3\x31\xd9\x6d\x76\xc3\xd8\xd5\xc5\xf9\xd9\xa1\x50\xed\x45\xe5\xa6\x2c\xd4\x3c\x47\xb9\x0b\x17\x96\x24\x38\xd3\xaa\x86\x81\x33\xcd\xdc\xbd\x47\x91\x65\x30\x46\x48\xac\x1c\x6e\x74\x89\x42\x97\xc6\xc9\x89\x78\x10\xc5\x04\x44\x3e\xcb\x30\xc7\x42\xd3\xc5\x29\x40\xa9\xac\x6e\xa3\xe2\xbe\xee\xc1\x64\x64\x31\xd9\x0b\xc9\xed\x75\x63\x79\xd5\xd6\x5a\x7f\x5d\xc2\xc3\xed\x35\xdc\xbc\xf9\xfd\x16\xca\x19\x4a\xa6\x45\x59\x28\x23\x72\xae\x8c\x8b\x9f\x32\x72\x8e\x3c\x9a\x17\x1c\x65\x26\x0a\x04\x3e\xde\xe3\xd4\xed\xb5\xd3\x66\xd2\x3b\x29\x33\xb0\x3e\xd0\x95\x5a\x98\xb5\x74\x7b\x4d\x17\x7c\x01\x3f\xd2\xd5\xdb\x7f\xff\x8b\x2e\x73\xd4\x52\x24\xca\x7f\xc6\xbf\xda\x4f\xba\xa5\x69\xed\x91\x66\xb3\x08\x5f\x71\x0a\x81\x5e\xd6\x08\xdc\xe3\x63\xb5\xba\x19\x14\xf8\x08\xa2\x50\x9a\x15\x09\x52\x98\x7b\xed\xf4\xcb\x97\x0a\x96\x55\x61\x2d\x8d\x20\xef\x5a\x11\xd5\xa6\x47\xa0\x73\x2a\x84\x66\x86\xb2\x26\xfd\x2a\x26\x16\xbe\x11\xfc\xd8\xab\xce\x2d\x77\xbe\x80\xcb\x2b\x23\x8b\xf4\xe6\x24\x99\x44\xc6\x71\x6c\xea\x63\xa3\x22\xfc\xb0\x4f\x14\xfd\xa8\xc5\x25\x49\xa8\xbf\xe0\x8b\x4b\xe0\xcd\x2f\x92\x32\xbb\xb4\x10\x36\xbe\x74\xfe\x5d\x42\xde\xf8\xd2\x0c\xba\xac\x80\x26\x6d\x2b\xfa\xcf\x4d\x5f\xbb\xa1\xeb\x1a\xf9\x5b\xcc\x50\x23\x15\x12\xcc\x69\x6d\x99\x62\x93\x97\x0f\x68\xb2\xc6\x91\x40\x2a\xcb\xdc\x16\x94\xb1\xcb\x2c\xba\x70\x29\xc8\x61\x36\x1f\x67\x22\x79\x75\x1b\x1b\x89\x6f\xcc\x1c\x55\x0d\x14\x8a\x92\x34\x9f\x2b\x0d\x53\xf6\x80\xc0\xdc\x78\x10\x1c\x1e\x58\x36\xc7\x88\x0a\x9e\x44\xa5\x90\x03\x0a\x3d\x45\x49\xcb\x90\xc1\x58\x95\x05\x94\x12\x3e\xd0\xa7\x66\x13\x23\x9d\x7e\xb5\x0b\x1a\x8b\x46\x62\xbc\x66\xc9\x47\x36\xc1\xf5\x3a\xee\x81\xdc\xa5\x74\x5d\xf0\xf9\x78\x47\xa4\x47\x0e\x97\x30\xd1\x0b\x2a\x29\x1a\x17\x3a\xbe\xb1\x9f\x51\xe5\xaf\xcb\x36\xb7\x9e\x3d\x1b\xe4\x94\x1d\x3e\xfb\xee\xf1\xf1\x9d\x64\x09\x86\x41\x7f\x12\x5b\x55\xc4\xaf\x26\xfa\x98\xa2\x84\x9c\x8f\x63\x2f\xe3\x2e\x17\x3a\xf4\x17\xaf\x8a\xb4\x3c\x44\x58\xe4\x27\xfc\x57\xe8\x69\x18\x78\x93\x83\xda\xfa\xf6\x10\x6b\x65\x1e\xdf\x15\x3c\x1c\x8d\xaa\x44\x16\x69\xe5\xfe\x2b\x75\xb7\x98\x09\x89\x9c\x40\xa9\xb8\x8f\x7e\x50\x4a\xf2\x39\xcd\x75\x7c\x47\x48\xa4\x61\xe0\xb0\x82\x29\x53\x14\x5d\x9a\xe6\x1d\x34\x18\xf5\xb9\xe7\xe7\xbf\x64\x22\x43\x43\x7d\xdc\x26\xa8\xcd\xc4\xe1\x7e\x85\x81\xc9\xfc\x20\x32\x3a\x93\x32\xeb\xde\x36\xb1\x0b\x0c\x4f\x59\xed\xd6\x7d\x6f\xa9\x5b\xc9\x28\x65\xb5\x74\x2a\x64\x9c\xdf\x24\x98\x2f\x1c\xf2\x94\x58\xb6\xb2\x44\xe0\xac\x7b\x2f\x78\xd3\xbc\x9f\xcc\xbc\xbf\x5d\x51\xb3\xd0\x44\xf1\x64\x98\xec\x71\xf8\x00\xc8\xf6\x60\xe2\x5c\xbf\xba\x02\x4e\x55\xd0\xf4\x90\x6f\xca\x47\xd5\xf4\xa6\x81\x5d\xa3\xc7\xac\xef\xaf\x0f\x80\x78\xcf\x22\xb0\x80\xf3\x13\xc2\xd0\xad\xdf\xad\x76\xee\x46\x22\xeb\x54\x4a\xc6\x79\xb3\x4c\x56\xed\xc6\xf6\x32\xd9\xa4\x34\x3d\xc5\x0e\xf5\xef\xad\x60\xff\xc7\xea\xfa\xa4\x4a\x6a\x71\xdb\x5e\x49\x31\xc3\x7c\x08\x06\x4f\x2d\xb5\xd6\x96\x13\x95\x5a\x2f\xac\x3f\xb1\xc8\xbd\xf8\xf5\xf3\xaa\xb7\x89\x4d\xf3\xfd\xf5\x76\x87\x73\xc7\x16\xdd\xb3\xb3\xb3\xb3\xde\x6a\xb0\x5a\x11\x3a\xe1\x94\xa9\x97\x94\x77\x2e\x2a\x10\xd8\x2d\x48\x30\x82\xf5\x7a\xa3\x44\xd5\xd5\xf9\x2d\x7b\x68\xd5\x66\xb2\xbe\xb7\x14\x6f\x85\xac\x7d\xbb\xd1\x82\xf5\x83\xd8\xb7\x82\x2a\x70\xfb\x65\xf6\xc0\xb8\x6f\x02\xb9\xe5\x62\x73\xc0\xe0\x2d\x51\xe8\x4c\xaa\x12\xb3\xaf\x58\xb7\x42\x74\x40\xd1\xbe\xe9\xa4\xd7\x46\x43\xbb\xdf\xf3\x7e\x8f\x47\x3e\x53\x30\x53\xd8\x4c\x07\xb3\xa6\x0a\x6d\xf2\xc1\xef\xf3\xc3\xd5\x8a\x2e\xca\xb9\x76\x5b\x5b\xb7\xa0\x7f\x4e\x68\x1f\x00\x56\x30\x04\x54\x27\x03\x08\x3e\x98\x8f\xf5\x7a\x93\x06\xfb\x73\xcc\x69\x1d\x96\x66\x4f\x4f\xaa\x41\x2b\xf1\xd3\x1c\xe5\x32\xa8\x6d\x1d\xd6\x09\x7c\x8e\x8c\x78\x92\xf9\x75\x06\x14\x7c\xbd\xde\x45\xe7\xbf\xa0\xfe\x39\xcb\xae\x97\xbf\x49\x8e\xb2\xb3\xff\xd1\x52\x20\x71\x69\x96\x99\xbc\xc2\x42\x2b\xbb\x07\xaa\x88\xbd\x22\xf5\xd2\x4e\x2f\xdc\x6f\xe3\x25\x18\xe9\x86\x7a\xd5\xc1\x1c\xd9\x32\x66\x3b\x55\x5a\x45\x7e\x7f\xeb\x95\xf9\x1d\x08\x84\x7f\xfc\x39\x80\x47\x37\x9f\x60\xa9\xe8\x7d\xd4\x4c\x67\x6b\x11\x99\xe2\x94\x55\x3a\x23\xf8\xc7\x3f\xe9\xdf\xa8\x85\x2d\x49\xb0\x39\xd0\x06\xb8\x42\x53\x19\x38\xa5\x6b\x5f\x9a\x3b\x4a\x02\xaf\xde\xfc\xab\x4c\x74\x77\xfd\xfb\x37\x77\x7a\x39\xc3\xe7\xda\x1f\xd5\x48\x0f\x0b\x7a\x04\x33\x36\x41\x6a\x3a\x23\x42\x7f\x56\x16\x0a\x5f\xa3\x7c\xed\xbe\x3c\x22\x27\x8c\xa8\x76\x62\x0c\xee\xb0\xac\x37\x27\xea\xb0\xbc\xb0\xe7\xd1\x39\x55\x75\xe3\xb8\x52\xd6\x53\x6a\x37\xf6\x49\x85\xc8\xec\x12\xdc\x6c\x91\x1e\x98\x04\x5d\x6a\x96\x51\x28\xeb\xaf\xa4\xd0\x98\x2b\x18\x94\x0d\xfb\xba\x2e\xf7\x88\xb3\xdb\x76\x91\xa2\xc8\xda\x10\x75\xd8\xd1\x06\xd3\x65\x68\x93\x26\xbb\xe5\x85\xd2\x7a\x23\xa5\x47\x3f\x6d\x10\xef\x93\x09\xb5\x55\xe8\x7d\x65\xea\x7f\xde\xe8\x8e\x38\x2c\x0f\x9c\x28\xba\x4d\xcb\x0f\xdd\x56\x77\x73\xc1\xa1\xbd\x75\x9b\x0d\xed\xad\x76\xcf\xd4\x7e\xee\x4e\x4b\x09\xef\x23\x13\x57\x0a\xa5\x64\x05\x55\x17\x93\x4e\x1d\xab\x28\xd1\x86\x6e\xe7\xb6\xfa\x7f\x79\x65\xb7\x16\xfe\x21\x3a\x69\xdb\xd9\x42\x0d\x8c\xba\x3b\x21\xf0\x9b\x76\x73\x86\x40\x6e\xf6\x45\xfd\xc8\x80\x93\xdc\xc0\x22\x77\x44\x2e\xec\x5c\xef\x7d\x81\xb5\x71\xb9\x02\x36\x9b\x61\xc1\x43\xe9\x56\xa3\x69\x8d\x37\x9e\xbc\x6c\x36\xc9\x33\x56\x88\xa4\xa7\x1c\x7b\x30\x0c\x91\x56\xc7\x28\xdb\x31\xeb\x9c\x8d\x04\x3b\x7b\x32\xd9\x2e\x19\xdd\x16\xed\x7a\x69\x4f\xec\x1a\x5d\xc4\xee\x87\xd2\xa9\x19\xfe\x11\x97\xa6\xad\xb0\x0c\x6f\x84\x35\x0f\x0f\xff\x42\xad\x85\x03\x70\x7b\x7b\x41\x30\xf9\x66\xc2\x58\x0c\xad\x63\x2a\x08\x07\xb5\x0f\x27\x68\x1d\x4e\xd7\x37\x74\x97\xeb\x47\x5c\x3a\x1f\xbf\x8d\x6e\xa2\xcf\x9f\x93\xb5\x19\x03\x42\xef\x8e\x79\xb7\x34\x23\xc3\x09\xe1\xa8\xae\xa3\xbb\x11\xff\x05\x75\xb3\xc1\xf8\x81\xcc\x30\xd9\xee\x21\x3b\xd1\x8e\x7c\x50\x03\xd1\xd3\x45\x74\x94\x0f\x79\xfc\x33\x30\xb0\x70\x5c\x8f\x31\x30\x11\x76\xb5\x21\x5b\x1b\x91\x63\x13\xed\xeb\xe3\x32\x9b\x66\x5d\x0a\x3b\x9c\xbb\xaa\x73\xc5\xe6\x8e\xf8\x2f\xc6\x57\x07\x9e\xb8\x7e\x4b\xd4\x74\xd4\xe1\xec\x33\xa1\xa9\x2f\x76\x40\xfb\x35\xd0\x15\xf4\xf2\x15\x40\x9b\xb1\x56\x2b\xf8\xfb\xcc\xb6\xaa\x14\x19\xfa\xe5\x7a\x49\x32\xeb\xa9\xfe\x64\x25\xd8\x9c\xa8\xd9\x84\xa6\x4d\x50\xbf\x63\x93\x4a\x50\xa0\x3e\x65\xfe\x41\xf5\xba\x55\x6c\x0f\xe4\xc7\x60\xb5\x32\xb2\xe3\xff\x50\x31\x58\xaf\x0f\x39\xb4\xfe\x4e\x98\x9f\x95\x30\xbf\x3d\xba\xdc\x47\x98\xbf\xcf\x38\xd3\x08\x73\xf5\x9d\x2e\xf7\xd2\xa5\xc5\xea\x20\xc6\xfc\xf2\x47\xed\xd6\xb8\x13\xb1\xa6\x17\xf6\xcd\xbc\xd5\x94\x8a\x42\xa8\x69\xe4\xcd\x68\x4b\x3c\xe6\x45\x96\xce\x4b\x46\x27\x7a\xc9\xe9\x28\x9e\x1a\xcc\x51\xbd\xf4\xe9\x8e\xf6\xbb\xec\xb9\xc1\x67\x6e\x21\x74\x8e\xf7\xbf\x08\xa3\xcd\x6d\xbd\x1a\x78\x28\xeb\x8f\x2d\xed\x39\xfd\xa0\x70\xef\x70\xe9\x8b\x3c\x72\x1e\x48\x57\x87\xf1\x8f\x79\x08\xfb\x79\x0f\xe7\x37\x73\x84\x94\x7e\xcf\x91\xaf\x37\x47\xda\x2d\xc3\x1e\xee\x70\x7d\xc3\x71\x47\x6f\x87\xe3\xd5\x0a\xca\xce\xd7\x05\xef\x16\x98\x34\xff\x2c\x85\x28\xdd\xe4\xae\xfb\x5b\x86\x2c\x23\x0c\xa9\x27\xc0\x05\x26\x73\x73\xab\x4c\x81\x41\x32\x57\xba\xcc\xeb\xf1\x6c\xc2\x44\xa1\xb4\x19\x6a\xfc\x38\xb8\x3f\x20\x13\xb6\x77\x07\xe9\xc2\xc8\x0f\xfd\xeb\xfb\x91\x7b\x35\x7e\xe4\x37\xca\x4f\xe3\x7e\x52\x7c\x22\xe6\xb7\xa2\x9e\x07\xa9\xdb\x40\x62\xfd\xb7\x18\x27\x3e\x24\x3c\xe0\x5d\xe4\x74\x11\xda\x7a\x67\x75\xa8\xc5\x29\x5e\x37\xfe\xdc\x7e\x7d\x25\xef\x13\xff\xe6\xdd\xf3\x0e\xef\xaf\x23\x7d\x15\xe0\x7f\x01\x00\x00\xff\xff\xa9\x30\x87\x3a\x71\x37\x00\x00"),
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
