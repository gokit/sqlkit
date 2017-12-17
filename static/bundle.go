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
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x5b\x6f\xdb\x46\x16\x7e\xb6\x01\xff\x87\x53\x62\x11\x90\x05\x97\xc1\xbe\xda\xc8\x43\x7d\x49\x11\x60\xeb\x64\x93\x74\xf7\xa1\x28\x8a\x11\xe7\x50\x9a\x84\x17\x65\x66\x64\x49\x10\xf4\xdf\x17\x67\x66\x78\x11\x4d\x8a\x92\xcd\xda\x70\x5a\x3f\x44\xa1\x38\x73\x2e\xdf\xb9\xce\x45\xaf\x5f\x43\xc6\xe6\x6f\x05\xa6\x5c\x01\xc7\x44\xe4\xa8\x80\x81\x5e\xcf\x11\x92\x42\x02\xa3\xd7\xa0\x67\x4c\x03\xae\xe6\x85\x32\x6f\xed\x70\x3f\x80\x0c\xf5\xac\xe0\xd1\xd9\xa9\x19\x5f\x13\xca\xd8\xfc\x37\xa5\xa5\xc8\xa7\xbf\x8b\x5c\xa3\x4c\x58\x8c\x9b\xed\xd9\xe9\xd9\xe9\xeb\xd7\x6e\x36\x48\xd4\x0b\x99\x2b\xd0\x33\x33\x13\x84\x56\x98\x26\xc0\x72\x0e\x73\x59\xdc\x09\x6e\x58\x59\x0e\xa0\x0b\xc8\x98\x8e\x67\x66\xb4\xfa\x96\x46\x9f\xd9\x24\x45\x43\x09\x2a\x0e\xd1\xd9\x69\xb2\xc8\x63\xf0\xb3\x5a\x94\xa0\x16\xd6\xef\x96\x2a\x04\x94\xb2\x90\x01\x6c\xce\x4e\x01\xc0\xc9\x05\x59\x08\xb9\x48\xcf\x4e\x4b\xa9\x37\x9b\xe8\x93\x96\x8b\x58\x47\xef\x27\x5f\x30\xd6\xd1\x2d\xcb\x70\xbb\x6d\x23\x97\xd7\xe2\xc0\x72\x26\xe2\x59\x05\x5b\xad\x89\xe3\x60\xa1\x2d\x12\x60\x69\x6a\x58\x30\xad\xa5\x98\x2c\x34\xd1\x51\xaa\x88\x05\xd3\xc8\x61\x29\xb4\x55\xdb\xf2\xe0\xa0\x8c\x18\x0b\x89\xc0\x88\x71\x2c\x38\x72\x98\xac\x2d\x34\xe5\xbb\xd2\x26\xfb\xc5\x6e\x08\xbb\x39\x3b\x3d\x39\x14\xaa\x41\x54\xae\x8a\x5c\x2d\x32\x94\xfb\x70\x61\x71\x8c\x73\xad\x6a\x18\x38\xd3\xcc\xbd\x5b\x8a\x34\x85\x09\x42\x6c\xe9\x70\xc3\x4b\xe4\xba\x30\x4a\x4e\xc5\x9d\xc8\xa7\x20\xb2\x79\x8a\x19\xe6\x9a\x1e\xc6\x00\xa5\x92\x7a\x17\x15\xf7\x75\x0f\x26\x81\xc5\x64\x10\x92\xeb\xcb\x46\x78\xd5\xd2\x5a\x7d\x9d\xc3\xc3\xf5\x25\x5c\x7d\xfc\xf5\x1a\x8a\x39\x4a\xa6\x45\x91\x2b\x43\x72\xa1\x8c\x8a\xdf\x52\x52\x8e\x34\x5a\xe4\x1c\x65\x2a\x72\x04\x3e\x19\x50\xea\xfa\xd2\x71\x33\xee\x1d\x17\x29\x58\x1d\xe8\x49\xad\x4c\x2c\x5d\x5f\xd2\x03\x5f\xc1\x8f\xf4\xf4\xe9\x3f\xff\xa6\xc7\x0c\xb5\x14\xb1\x2a\x3f\xa3\x5f\xec\x27\xbd\xd2\x14\x7b\xc4\xd9\x04\xe1\x3b\x4e\x26\xd0\xeb\x1a\x81\x5b\x5c\x56\xd1\xcd\x20\xc7\x25\x88\x5c\x69\x96\xc7\x48\x66\xee\x95\xb3\x0c\xdf\x5b\x5c\xfa\x96\x85\x95\x34\x84\xac\x2d\x45\x58\x8b\x1e\x82\xce\x20\x8a\x22\x33\x43\x59\x91\x7e\x11\x53\x0b\x5f\x00\x3f\xf6\xb2\x73\xe1\xce\x57\x70\xfe\xc6\xd0\x22\xbe\x19\x51\x26\x92\x51\x14\x05\xa4\x4e\x23\x23\xbc\x1a\x22\x45\x7f\x6a\x75\x4e\x14\xea\x2f\xf8\xea\x1c\x78\xf3\x8b\xb8\x48\xcf\x2d\x84\x8d\x2f\x9d\x7e\xe7\x90\x35\xbe\x34\x83\xce\x2b\xa0\x89\xdb\x86\xfe\x71\xd3\xb7\x6e\xe8\xb6\x46\xfe\x1a\x53\xd4\x48\x89\x04\x33\x8a\x2d\x93\x6c\xb2\xe2\x0e\x8d\xd7\x48\x8c\x0b\xc9\x21\x91\x45\x66\x13\xca\xc4\x79\x16\x3d\x38\x17\xe4\x30\x5f\x4c\x52\x11\xbf\xbb\x8e\x0c\xc5\x8f\x66\x8e\xaa\x06\x0a\x45\x4e\x9a\x2d\x94\x86\x19\xbb\x43\x60\x6e\x3c\x08\x0e\x77\x2c\x5d\x60\x48\x09\x4f\xa2\x52\xc8\x01\x85\x9e\xa1\xa4\x30\x64\x30\x51\x45\x0e\x85\x84\x2f\xf4\xa9\xd9\xd4\x50\xa7\xff\xda\x80\xc6\xbc\xe1\x18\x1f\x58\xfc\x95\x4d\x71\xbb\x8d\x7a\x20\x77\x2e\x5d\x27\x7c\x3e\xd9\x63\xe9\xc0\xe1\xe2\xc7\x7a\x45\x29\x45\xe3\x4a\x47\x57\xf6\x33\xac\xf4\x75\xde\xe6\xe2\xb9\xac\x06\x19\x79\x47\xe9\x7d\xb7\xb8\xfc\x2c\x59\x8c\xbe\xd7\xef\xc4\x96\x95\x17\x38\xef\xc2\x04\x25\x64\x7c\x12\x95\x34\x6e\x32\xa1\xfd\xf2\xe1\x5d\x9e\x14\x87\x10\x0b\xcb\x09\xff\x13\x7a\xe6\x7b\xa5\xc8\x5e\x2d\xfd\xee\x10\x2b\x65\x16\xdd\xe4\xdc\x0f\x82\xca\x91\x45\x52\xa9\xff\x4e\xdd\xac\xe6\x42\x22\x27\x50\xaa\xda\x47\x7f\x28\x25\xe9\x9c\x64\x3a\xba\x21\x24\x12\xdf\x73\x58\xc1\x8c\x29\xb2\x2e\x4d\x2b\x15\x34\x18\xf5\xa9\x57\xce\x7f\xcb\x44\x8a\xa6\xf4\x71\xeb\xa0\xd6\x13\x8f\xd7\xcb\xf7\x8c\xe7\x7b\xa1\xe1\x19\x17\x69\xfb\xb5\xb1\x9d\x67\xea\x94\xe5\x6e\xd5\x2f\x25\x75\x91\x8c\x52\x56\xa1\x53\x21\xe3\xf4\x26\xc2\x7c\xe5\x90\x27\xc7\xb2\x99\x25\x04\x27\xdd\x1f\x82\x37\xc5\xbb\x30\xf3\x7e\x78\x43\xcd\x42\x13\xc5\xd1\x30\x19\x50\xf8\x00\xc8\x1e\x8a\xc9\x80\xd7\x5a\x84\xf8\x88\x72\xb7\x13\xee\x4e\xff\x75\x25\x91\xb5\x52\x1b\xe3\xbc\x99\xd7\xaa\xfe\xa0\x3b\xaf\x35\x6b\x90\x9e\x61\xab\x56\x0f\xa6\x9c\x67\x4c\x87\x8f\x4a\x7d\x16\xb7\xee\xd4\x87\x29\x66\xc7\x60\xf0\xd8\xdc\x68\x65\x19\x29\x37\x96\xc4\xfa\x1d\x8b\xd4\x8b\x3e\xbc\xac\x04\x19\x5b\x37\x1f\x4e\x90\x7b\x94\x7b\x68\x96\x3c\x39\x39\x39\xd9\x97\x21\x13\x5a\x1e\xa0\x0c\xa1\xf8\x4a\x00\x34\x9a\x60\x9f\xa4\x09\x22\x7f\xef\x6a\x23\xb8\xa0\x89\x0d\x20\xef\x25\xdd\x4f\xec\x6e\x27\xe5\x3a\x86\xbd\x49\xb6\x13\xdb\xdd\xd7\x8d\xe6\xaa\x1f\xed\xbe\x50\xab\xac\xd0\x4f\xb3\x07\xef\xa1\x09\x84\x97\x33\xe2\x01\x83\x3b\xcc\xd5\x9a\x54\x79\x70\x5f\x56\xdf\xb1\xe5\x01\xd9\xfd\xaa\xe5\x87\xf7\x5a\xd5\x61\xcd\xfb\x35\xde\x91\xb6\x99\xec\xdb\x1e\x37\x66\x30\xbe\xb8\x70\xec\x33\x65\x05\x90\x41\x27\xd7\x26\x80\xca\xfd\x0e\x7f\xb3\xa1\x87\x62\xa1\xdd\x12\xdf\xe5\xc9\x9f\x62\x5a\x0f\x81\x35\x03\x78\x54\x7e\x3c\xf0\xbe\x98\x8f\xed\x36\xe8\xe9\x82\xda\x01\xe9\x38\x8e\xd1\xf5\x1c\x1a\x7a\x47\x21\xf9\x6d\x81\x72\xed\xd5\x72\x3e\x4d\x27\xd4\x8e\x95\x47\x89\x1c\xf4\xb7\x40\x3f\xa3\xfe\x29\x4d\xe9\x95\x14\x78\x87\x0a\x98\x79\xb2\x1d\x49\x73\x55\xc7\x72\xde\x58\x80\xab\x54\xb4\x57\xde\xc3\x0b\x2c\xbd\x9e\xe3\x4b\x6d\x79\x2c\x4e\xdd\x2d\x4f\x21\x39\xca\x6a\x63\xc1\x3c\x4d\xd6\xd5\xf3\x9c\x4d\x91\xea\x5a\x08\x12\xd5\xbc\xc8\x15\x7e\x40\xf9\xc1\x7d\x19\x00\xf8\xbf\xfd\x7e\x04\x88\xa1\x25\xb5\xbb\xbd\x78\x74\xd3\x64\xb5\x19\xa9\x69\x2a\x89\xbd\x8c\x66\xa8\xf4\xf4\x07\x86\x56\x4f\xb8\x77\x95\x9e\x10\xfe\xf9\xaf\xb0\x23\xf2\x37\x1b\x82\xc1\x27\x5d\xde\x92\xef\x39\x5c\xc1\x73\x9b\x82\x1e\x40\x00\xdb\x6d\x4d\xf2\x8e\x49\x90\x42\x63\xa6\xe0\x28\x5f\x69\x8a\x65\xa6\x87\xa0\x0b\xcd\xd2\xb0\x95\x91\xad\x01\x9d\x57\x36\x53\xb3\x71\xe5\xca\xa3\xad\x2b\xdf\x73\xe3\xe0\xe2\x5e\xf7\x75\x60\x57\x75\x80\x91\x9a\xd9\xa8\x7f\x9f\xcf\x1d\x2d\x50\xb6\xe2\x93\xd1\x2d\xda\xb6\x6a\x8d\x61\x4f\x1b\x94\x14\x12\xfe\x08\x0d\xe4\x84\xb2\x64\x39\x05\xbb\xb1\x5f\x0b\x12\xb2\xec\xb1\x0b\xa6\xb6\x68\x75\x81\x35\xfd\x42\xb9\xb3\x4c\xfc\xf6\x36\xb9\xc7\xd6\x56\x4b\xb7\x5c\x18\x9b\x8d\x75\x52\xb4\xcf\x24\x0f\xb4\x06\xd1\xf5\x2c\x76\x0f\x32\xd4\xde\x10\xec\xb4\x96\x99\x61\xad\xf3\x06\xd8\x7c\x8e\x39\xf7\xa5\x0b\x17\xd3\x59\xf6\x58\xd9\x71\x91\xbb\x91\x55\xb5\x9c\x9b\x0d\x60\xaa\x70\xf4\x40\x1e\x8e\xe0\xcb\xf5\xd1\x31\x1c\x02\x55\x41\x5f\x16\x4b\x65\x36\xec\x57\xd1\xc7\x62\xa9\x5a\x4b\xf3\xa6\x7b\xd3\xc8\xe8\x16\x57\xda\x0f\xba\xfc\xaa\x52\xf4\x31\x8e\xbd\xeb\xdc\x86\xa3\x9d\xf6\x29\x66\xb9\xff\x4a\x1e\xe4\xe1\xc3\x5e\xee\xa3\x94\xbb\xae\xe6\xba\xde\xf2\x2b\xf3\xd8\x43\x9a\xfe\xc8\x31\xbd\x73\x5b\xb7\x3a\xd6\x5d\xd5\x30\xeb\xf9\xe7\x50\x19\xc7\x6e\xc1\xfb\x5d\x8b\x35\xe3\x6d\xdd\xde\x0d\x3d\x2b\xb1\x6a\x5a\x67\x48\xf4\x38\xb8\x05\x71\x28\x3a\x5a\xab\x29\x2b\xdc\xc5\x4e\x79\x19\x29\xf1\x3f\x51\x96\x19\x3f\xe7\x1f\x90\x0d\x72\x6e\x92\x41\xa3\xf3\xbe\x5c\xdb\xd3\xe6\x46\xf7\xbd\xff\x40\xc5\x6c\x5e\xc0\x57\x5c\x9b\x76\xdc\x76\xc6\x86\x58\xf3\xe0\xfb\x2f\xd4\x92\x3b\x00\xbb\xdb\x72\x82\xa9\x6c\xc2\x8d\xc4\xd5\x69\x0c\xf8\x47\x75\xdc\x23\x74\xdb\xe3\xb5\xda\x6d\x57\xff\x8a\x6b\xa7\xde\xf7\xd1\x80\xf7\xe9\x33\x5a\x67\x7e\x84\xe9\xdd\xed\x84\x56\xff\xfe\x98\xa6\x4d\x24\xd4\x19\x1e\xbb\xd7\x59\x5e\x22\x08\x2e\xe0\x87\xdd\xed\xce\xb2\x10\xa4\x98\xb9\x45\xa1\x71\x49\xdf\x7b\x9f\xa7\xeb\xee\x40\x6b\x5f\x4c\x60\x69\x5a\x2c\x6b\x6b\xf6\x1f\x5f\xfd\x8c\xba\xd9\x58\xbc\xb2\x4c\x1b\xf6\x1a\x61\x07\xe7\xa8\xc6\xbf\xa7\xfb\x7f\xfc\xb6\xe2\x53\x79\xd2\x8e\xf5\xda\xdb\x32\x87\x57\x85\xea\xb4\xb9\xb9\x47\xf3\x17\xab\x04\x07\x9e\xc3\x7f\x4f\x99\xff\x41\x47\xf6\x2f\xa4\x0a\x3c\xd9\xb1\xfd\x33\x55\x83\x17\x56\x0b\x36\x1b\xf8\xc7\xdc\x36\x9f\xe4\x11\xf4\x9f\xcb\x35\xcd\xad\x37\xaf\xca\x73\x0a\xaf\x5a\xf1\x9a\x49\x9a\x4d\x69\xca\x14\xf5\x67\x36\xad\x88\x78\xea\x5b\x5a\x1e\x15\x6c\x8f\xad\x38\xde\x66\x63\xe8\x46\xff\xa5\xe4\xb3\xdd\x8e\x7c\x75\xe2\xef\x12\x64\x4b\xd0\xaf\x73\xce\x34\xc2\x42\xfd\x5d\x80\x06\x0b\x90\xc5\xea\xa0\x1a\xf4\xf4\x37\x24\xac\x70\x23\xd5\xa1\x92\xd8\x77\x73\x7b\x2c\x11\xb9\x50\xb3\xb0\x14\x63\x97\xe2\x43\xee\x1f\xb5\x2e\x73\x8d\x74\x5c\x48\x65\xc6\xec\x86\xec\xbd\x0e\xd1\xbe\x58\xa1\x46\xbb\x57\x61\x78\xbb\xd4\xaf\x9c\x6c\x98\x2a\xbc\x3f\x64\x9c\x63\xe3\xe7\x2e\x3d\x2e\xa2\x1b\xd5\x87\xd4\xfb\xd3\x8b\xcf\xc2\x26\xdd\x23\x4f\xb0\xcb\xf3\x5e\x12\xf1\x48\x9f\xdd\xa3\xce\xd3\x1c\x74\xbb\x32\xf3\xb0\xd3\xb8\xc3\xa5\xdf\x81\x68\xef\xa5\xc0\x9b\x15\xc6\xcd\x5f\x8b\x50\x05\x30\x0e\xeb\x7e\x62\x40\x2d\x92\xad\x6e\xb8\xc2\x78\x61\x5e\x15\x09\x30\x88\x17\x4a\x17\x59\x3d\x9e\x4d\x99\xc8\x95\x36\x43\xed\x7e\xef\xa1\xe5\x84\x44\xe8\x2e\x26\xc9\xca\xee\xd4\x97\xb7\xea\x43\x77\x63\x3d\x28\x57\x2a\x8f\x2b\x15\xc4\x78\xa4\x42\x61\x49\xbd\x8c\x1a\x60\x0d\x89\xf5\x4f\x24\x46\xda\x4f\x1e\x8c\x8d\x3a\xed\x24\x2b\xdf\x66\x1e\xcb\x43\xad\xc6\xc8\x26\xcf\xa6\xd7\x80\xd7\xbc\x2f\xe5\x29\x25\x1c\x0e\xfc\xbe\x90\xfd\x7f\x00\x00\x00\xff\xff\x54\xca\xc0\xa9\x6c\x36\x00\x00"),
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
