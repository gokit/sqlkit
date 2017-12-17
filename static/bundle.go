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
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xc1\x6e\xe2\x30\x10\x86\xef\x79\x8a\x39\x82\x14\x25\xcf\xb0\x01\x2d\xda\xcb\x82\xb4\xda\x53\xd5\x83\x63\xff\x24\x2e\x8e\x6d\xd9\x03\x25\x42\x79\xf7\x0a\x92\xa2\xb4\x22\xa8\xa5\x27\x4b\x33\x99\xf9\xbf\x6f\x94\x3c\xa7\xd3\x29\xfb\xc7\x61\x2f\x39\x5b\x97\x2f\x90\x9c\xfd\x15\x0d\xba\xae\x10\x72\x07\xab\x48\x61\xab\x2d\x22\x09\x2a\x87\xca\x6b\xad\x65\x4d\x01\x3e\x20\xc2\x72\x24\xae\x41\x95\x3e\x68\x5b\x51\x92\xe7\xd4\x80\x6b\xa7\x22\xe1\xe8\x5d\x84\xa2\xb2\xbd\x7c\xb1\x2c\x48\x37\xde\xa0\x81\x65\xc1\xda\x59\xda\xba\x30\x9e\xe5\xd6\x63\x8a\x26\x4b\xee\x75\xdf\x59\xb5\x65\x84\xad\x90\x38\x25\x44\x4b\x18\x30\x66\x92\x8f\x24\x9d\x65\x1c\x39\x5b\xf4\x6f\x4a\x7e\x5f\x1a\x2d\xff\x2c\x29\x72\xd0\xb6\x9a\x13\x42\x70\x21\x21\x5a\x04\x88\xa9\x21\x18\x34\x23\x82\x8d\x90\x3b\x51\x9d\xd9\x26\xa8\x86\xad\x94\x10\xad\xc0\x5f\x04\xa1\xd9\x37\x12\x52\xea\x23\xe6\x09\xd1\x7f\xaf\x26\xc9\x3f\xa5\x3c\xac\xd2\x9b\xfc\x32\xa6\x68\xd7\x41\x21\xdc\x4e\x73\xe7\xd6\xf0\x14\xed\xc8\xec\xe9\xf9\x41\xb7\x15\xb8\x68\x7f\x6b\x18\x75\x3b\x71\x87\xf6\xaa\x76\x10\x66\x8f\xd1\x9f\xd0\xfd\xe0\xa6\xbd\xec\x1d\xcb\x6b\xea\x47\xd9\x94\xbc\xa8\x2e\x10\x29\x05\x44\xef\x6c\xc4\x06\x61\x33\x14\x1f\xb9\xc5\x65\xd7\x00\xd6\xbd\x05\x00\x00\xff\xff\x4a\x68\x4a\xbf\xb6\x03\x00\x00"),
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
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\xdd\x8f\xdb\x36\x12\x7f\xce\x02\xfb\x3f\xcc\x09\x87\xc2\x2e\x74\x0a\xee\x75\x8b\x3c\x74\x3f\x52\x04\xb8\x26\xb9\x26\xbd\x7b\x28\x8a\x82\x16\x47\x36\x13\x7d\x85\xa4\x77\x6d\x18\xfe\xdf\x8b\xe1\x87\x24\xcb\x96\x65\x79\xd5\xb4\x9b\x26\x40\xeb\x95\x44\xce\xc7\x6f\x86\xf3\x1b\x8a\x7a\xfe\x1c\x32\x56\xbe\x14\x98\x72\x05\x1c\x13\x91\xa3\x02\x06\x7a\x5d\x22\x24\x85\x04\x46\x8f\x41\x2f\x98\x06\x5c\x95\x85\x32\x4f\xed\xf0\xc9\x14\x32\xd4\x8b\x82\x47\x97\x17\x66\x7c\x2d\x28\x63\xe5\x2f\x4a\x4b\x91\xcf\x7f\x15\xb9\x46\x99\xb0\x18\x37\xdb\xcb\x8b\xcb\x8b\xe7\xcf\xdd\x6c\x90\xa8\x97\x32\x57\xa0\x17\x66\x26\x08\xad\x30\x4d\x80\xe5\x1c\x4a\x59\xdc\x0b\x6e\x54\x59\x0d\xa0\x0b\xc8\x98\x8e\x17\x66\xb4\xfa\x94\x46\xef\xd9\x2c\x45\x23\x09\x2a\x0d\xd1\xe5\x45\xb2\xcc\x63\x98\x64\xb5\x29\xd3\xda\xd8\xc9\x61\xab\x42\x40\x29\x0b\x39\x85\xcd\xe5\x05\x00\x38\xbb\x20\x0b\x21\x17\xe9\xe5\x85\xb7\x7a\xb3\x89\xde\x69\xb9\x8c\x75\xf4\x66\xf6\x01\x63\x1d\xbd\x66\x19\x6e\xb7\x6d\xe4\xf2\xda\x1c\x78\x58\x88\x78\x51\xc1\x56\x7b\xe2\x34\x58\x68\x8b\x04\x58\x9a\x1a\x15\x4c\x6b\x29\x66\x4b\x4d\x72\x94\x2a\x62\xc1\x34\x72\x78\x10\xda\xba\x6d\x75\x70\x50\xc6\x8c\xa5\x44\x60\xa4\x38\x16\x1c\x39\xcc\xd6\x16\x1a\xff\xcc\xc7\xe4\xb8\xd9\x0d\x63\x37\x97\x17\xcf\x4e\x85\xaa\x17\x95\x9b\x22\x57\xcb\x0c\xe5\x31\x5c\x58\x1c\x63\xa9\x55\x0d\x03\x67\x9a\xb9\x67\x0f\x22\x4d\x61\x86\x10\x5b\x39\xdc\xe8\x12\xb9\x2e\x8c\x93\x73\x71\x2f\xf2\x39\x88\xac\x4c\x31\xc3\x5c\xd3\xc5\x18\xa0\x54\x56\xef\xa2\xe2\x6e\x77\x60\x32\xb5\x98\xf4\x42\x72\x7b\xdd\x58\x5e\xb5\xb5\xd6\x5f\x97\xf0\x70\x7b\x0d\x37\x3f\xfd\x7c\x0b\x45\x89\x92\x69\x51\xe4\xca\x88\x5c\x2a\xe3\xe2\xa7\x94\x9c\x23\x8f\x96\x39\x47\x99\x8a\x1c\x81\xcf\x7a\x9c\xba\xbd\x76\xda\x4c\x7a\xc7\x45\x0a\xd6\x07\xba\x52\x2b\xb3\x96\x6e\xaf\xe9\x82\xaf\xe0\x5b\xba\x7a\xf7\xdf\xff\xd0\x65\x86\x5a\x8a\x58\xf9\xdf\xe8\x47\xfb\x4b\x8f\x34\xad\x3d\xd2\x6c\x16\xe1\x2b\x4e\x21\xd0\xeb\x1a\x81\xd7\xf8\x50\xad\x6e\x06\x39\x3e\x80\xc8\x95\x66\x79\x8c\x14\xe6\x4e\x3b\xfd\xf2\x7d\x8d\x0f\x13\xab\xc2\x5a\x1a\x42\xd6\xb6\x22\xac\x4d\x0f\x41\x67\x10\x45\x91\x99\xa1\xac\x49\x3f\x8a\xb9\x85\x6f\x0a\xdf\x76\xaa\x73\xcb\x9d\xaf\xe0\xea\x85\x91\x45\x7a\x33\x92\x4c\x22\xa3\x28\x9a\x92\x3b\x8d\x8a\xf0\x4d\x9f\x28\xfa\xa7\x56\x57\x24\xa1\xbe\xc1\x57\x57\xc0\x9b\x37\xe2\x22\xbd\xb2\x10\x36\x6e\x3a\xff\xae\x20\x6b\xdc\x34\x83\xae\x2a\xa0\x49\xdb\x86\xfe\xe7\xa6\x6f\xdd\xd0\x6d\x8d\xfc\x2d\xa6\xa8\x91\x0a\x09\x66\xb4\xb6\x4c\xb1\xc9\x8a\x7b\x34\x59\x23\x31\x2e\x24\x87\x44\x16\x99\x2d\x28\x33\x97\x59\x74\xe1\x52\x90\x43\xb9\x9c\xa5\x22\x7e\x75\x1b\x19\x89\x3f\x99\x39\xaa\x1a\x28\x14\x25\x69\xb6\x54\x1a\x16\xec\x1e\x81\xb9\xf1\x20\x38\xdc\xb3\x74\x89\x21\x15\x3c\x89\x4a\x21\x07\x14\x7a\x81\x92\x96\x21\x83\x99\x2a\x72\x28\x24\x7c\xa0\x5f\xcd\xe6\x46\x3a\xfd\x69\x17\x34\xe6\x8d\xc4\x78\xcb\xe2\x8f\x6c\x8e\xdb\x6d\xd4\x01\xb9\x4b\xe9\xba\xe0\xf3\xd9\x91\x48\x4f\x1d\x2e\x93\x58\xaf\xa8\xa4\x68\x5c\xe9\xe8\xc6\xfe\x86\x95\xbf\x2e\xdb\xdc\x7a\xf6\x6c\x90\x51\x76\xf8\xec\x7b\x8d\x0f\xef\x25\x8b\x71\x12\x74\x27\xb1\x55\x15\x4c\x5d\x76\x61\x82\x12\x32\x3e\x8b\xbc\x8c\xbb\x4c\xe8\x89\xbf\x78\x95\x27\xc5\x29\xc2\x42\x3f\xe1\xff\x42\x2f\x26\x81\x37\x39\xa8\xad\xdf\x1d\x62\xad\xcc\xa2\xbb\x9c\x4f\xa6\xd3\x2a\x91\x45\x52\xb9\xff\x4a\xdd\xad\x4a\x21\x91\x13\x28\x15\xf7\xd1\x3f\x94\x92\x7c\x4e\x32\x1d\xdd\x11\x12\xc9\x24\x70\x58\xc1\x82\x29\x8a\x2e\x4d\xf3\x0e\x1a\x8c\xba\xdc\xf3\xf3\x5f\x32\x91\xa2\xa1\x3e\x6e\x13\xd4\x66\xe2\x70\xbf\x26\x81\xc9\xfc\x20\x34\x3a\xe3\x22\x6d\x3f\x36\xb1\x0b\x0c\x4f\x59\xed\xd6\x7d\x6f\xa9\x5b\xc9\x28\x65\xb5\x74\x2a\x64\x9c\xdf\x24\x98\xaf\x1c\xf2\x94\x58\xb6\xb2\x84\xe0\xac\xfb\x4d\xf0\xa6\x79\xdf\x99\x79\xff\x78\x41\xcd\x42\x13\xc5\xd1\x30\xe9\x71\xf8\x04\xc8\xce\xc5\xa4\x27\x6b\x2d\x42\x7c\x44\xbb\xdb\x05\x77\xa7\xff\xba\x91\xc8\x5a\xa5\x8d\x71\xde\xac\x6b\x55\x7f\x70\xb8\xae\x35\x39\x48\x2f\xb0\xc5\xd5\xbd\x25\xe7\x4f\x2c\x87\x8f\x2a\x7d\x16\xb7\xc3\xa5\x0f\x53\xcc\x86\x60\xf0\xd8\xda\x68\x6d\x19\xa9\x36\x7a\x61\xdd\x89\x45\xee\x45\x6f\x9f\x56\x81\x8c\x6d\x9a\xf7\x17\xc8\x23\xce\x9d\x5b\x25\x9f\x3d\x7b\xf6\xac\xb3\x1a\x6c\x36\x84\xce\x64\xc1\xd4\x4b\xca\x3b\x17\x15\x08\xec\x9e\x21\x98\xc2\x76\x5b\x3b\xbc\x57\x4e\xdf\xb1\xfb\x9d\x62\x4a\xd6\x77\xd6\xce\x83\x90\xed\x3e\x6e\xf4\x4c\xdd\x20\x76\xad\xa0\x0a\xdc\x6e\x99\x1d\x30\xf6\x4d\x20\xb7\x5c\x6c\x4e\x18\x7c\x20\x0a\xad\x49\x55\x62\x76\x15\xeb\x9d\x10\x9d\x50\xb4\x6f\x5a\xe9\xb5\xd7\x81\xf6\x7b\xde\xed\xf1\xd4\x67\x0a\xa6\x0a\x9b\xe9\x60\xd6\x54\xae\x4d\x3e\xf8\x8d\xf9\x64\xb3\xa1\x8b\x62\xa9\xdd\x5e\xd4\x2d\xe8\xef\x63\x6a\xdc\xc1\x0a\x86\x80\xea\x64\x00\xc1\x07\xf3\xb3\xdd\x4e\x07\xe4\x98\xd3\x3a\x2c\xcd\x1e\x9f\x54\x83\x56\xe2\xa7\x25\xca\x75\x50\xdb\x3a\x8c\xba\xff\x88\x8c\x78\x94\xf9\x75\x06\xe4\x7c\xbb\x3d\x46\xe7\x3f\xa0\xfe\x3e\x4d\xaf\xd7\x6f\x24\x47\xd9\xda\xb0\x68\x29\x90\xb8\x34\x4d\x4d\x5e\x61\xae\x95\xdd\xb4\x54\xc4\x5e\x91\x7a\x61\xa7\xe7\xee\xaf\xd9\x1a\x8c\x74\x43\xbd\xea\x64\x8e\xdc\x31\xe6\x30\x55\x5a\x45\x7e\x43\xea\x95\xf9\x2d\x03\x4c\x7e\xf9\x75\x00\x8f\x86\xd4\xab\x74\xbd\x77\xe2\xb3\xc8\xda\x43\x86\x38\x55\x95\xc6\x10\xfe\xf5\x6f\xfa\x6f\xda\x46\xb2\x82\x4d\x19\xdc\xa4\xeb\x53\x9a\x7b\x3d\x42\xa9\xde\x96\xab\x54\xb4\xf7\xe3\xfd\xdb\x2e\xbd\x2e\xf1\xa9\x36\x42\x35\xa8\xc3\xa2\x1b\x42\xc9\xe6\x68\x23\x26\x51\x95\x45\xae\xf0\x2d\xca\xb7\xee\xe6\x38\xc1\x1f\xdc\x4a\x59\x6f\x46\x6a\xa5\xbc\xb0\xa7\xd1\x22\x55\x05\xe2\xbc\x9a\xd5\x51\x53\xf7\x36\x44\xb9\x48\xed\x6a\xdb\xef\x85\xee\x99\x04\x5d\x68\x96\x52\x28\xeb\x5b\x52\x68\xcc\x14\x0c\xca\x86\xbe\xf6\xca\xbd\x7c\x6c\xf7\x57\xa4\x28\xb4\x36\x84\x2d\x1a\xb4\xc1\x74\x19\xda\xe4\xc3\x76\x25\xa1\xb4\xde\x4b\xe9\xe9\x77\x7b\x0c\xfb\x68\xe6\xdc\xa9\xe8\xbe\x32\x75\xbf\x09\x74\x87\x0f\xb6\xe0\x8f\x14\xdd\x03\x8c\x69\xe2\x5b\x23\xd8\xc1\x9e\x49\x21\xe1\xb7\xd0\x00\x4e\x18\x4b\x96\xd3\xb2\x37\x71\x6e\x01\x42\x19\x30\x74\x43\xb5\x2b\xa1\xee\x68\x4c\x73\xef\xdf\x3b\x93\xb6\xa3\x4d\xcc\xd0\x46\xc6\xca\xf5\xdb\x66\xf3\xda\x9d\xdc\xec\x0a\xc7\x99\x91\x20\xb9\x81\x45\xee\x8c\x20\x1d\x5d\x88\x8d\x48\xed\x8c\xb7\x71\x79\x01\xac\x2c\x31\xe7\x13\xe9\x96\x89\x69\x4e\x77\xe2\x0b\x07\xdb\xd4\x92\xe5\x22\xee\xa8\x93\x1e\x0c\xc3\x70\xd5\xc9\xc3\x61\xcc\x5a\xc7\x09\xc1\xd1\xae\x48\xee\xae\xe5\x76\x93\x74\xbd\xb6\x87\x5c\x0d\x7a\x3f\xfe\x1e\x37\x31\xc3\x3f\xe2\xda\xf0\xbd\xa5\x5e\x23\xac\x79\xde\xf6\x37\xe2\x7c\x07\xe0\x61\xde\x27\x98\x3c\xcb\x1b\x8b\x1b\x1d\xdd\x20\x4a\x1f\x81\xce\xc7\xe3\xf2\xf6\x4a\xfd\x88\x6b\xe7\xde\x97\xc1\xf0\x5d\xfe\x8c\x46\xfd\x03\x42\xef\x0e\x45\x0f\x34\x08\xc3\xb9\xe0\xac\x4e\xa0\xbd\x0b\xfe\x01\x75\x93\xf4\xbf\x21\x33\x4c\xa2\x7b\xc8\x46\xda\x0e\x0f\x22\xf5\x0e\x66\x6f\x29\x1f\xf2\xee\x65\x60\x60\x47\x08\xee\x5f\x8f\x3a\x6c\x68\xdb\x8c\x71\x3a\x55\x54\x27\x5f\xcd\x9d\xe1\xdf\x8c\x1e\x4e\x3c\x13\xfc\x92\xe8\xe0\xac\xe3\xc3\x27\x42\x0d\x9f\xed\x08\xf1\xaf\x40\x11\xd0\xc9\x11\x00\xbb\x2c\xb1\xd9\xc0\x3f\x4b\xdb\x19\x52\x64\xe8\x8f\xeb\x35\xc9\xac\xa7\xfa\xa3\x84\x60\x7f\xa2\x66\x73\x9a\x36\x47\xfd\x9e\xcd\x2b\x41\x81\xfa\x94\xfa\x37\xb3\xdb\x9d\x3d\xd3\x89\x9c\x14\x6c\x36\x46\x76\xf4\x3f\x2a\x06\xdb\xed\x29\xc7\xaa\x5f\x49\xea\x49\x90\x54\x1f\x4d\xfd\x5c\x72\xa6\x11\x96\xea\x2b\x49\xf5\x92\x94\xc5\xea\x24\x9e\xfa\xfc\x27\xba\xd6\xb8\x91\xb8\xca\x0b\xfb\x62\xbe\x76\x49\x44\x2e\xd4\x22\xf4\x66\xec\x4a\x3c\xe7\x7b\x89\xd6\xc7\x27\x23\x7d\xe8\x71\x16\x3b\x0c\x66\x86\x4e\xd2\x72\x27\xc8\x6d\xce\xda\x63\x11\xb7\x10\x5a\xa7\xc8\x9f\x85\x47\x96\xb6\x5e\x0d\x3c\xfb\xf3\xa7\x63\xf6\x38\x78\x50\xb8\x8f\xb8\x34\xda\x11\xe1\x61\x92\x30\xaf\x03\xff\xd8\x83\xda\xfd\x40\x92\xd2\xaf\x81\x7c\x64\x20\x77\xc9\xb7\xa7\x0a\x3b\x06\x3e\xef\x08\xe5\x74\xa7\x76\x90\x3b\xfa\x7d\xd7\xdd\x0a\xe3\xe6\x87\xff\x44\x8e\x26\xc1\xdc\xd7\xe2\x69\x5a\x3c\x58\xe2\xc7\x15\xc6\x4b\xf3\xa8\x48\x80\x41\xbc\x54\xba\xc8\xea\xf1\x6c\xce\x44\xae\xb4\x19\x6a\xfc\x38\x99\x69\xc9\x84\xc3\x3c\x9b\xac\x8c\xfc\x89\xff\x40\x3a\x74\x1f\x1f\x4f\xfd\x46\xef\x71\x2c\x4a\x8a\x47\xe2\x50\x2b\xea\x69\xd0\xa3\x0d\x24\xd6\x5f\xbb\x8f\x7c\xd8\x73\xc2\xd7\x9e\xc9\x6a\x62\x8b\x92\xd5\xa1\x56\x63\x7c\xd0\xf9\xa7\xf9\xd5\x93\x35\x6f\xbc\x3d\xde\xc2\xfe\x85\xdf\xb5\x64\x7f\x0f\x00\x00\xff\xff\xb3\xb1\x5d\xa7\x37\x34\x00\x00"),
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
