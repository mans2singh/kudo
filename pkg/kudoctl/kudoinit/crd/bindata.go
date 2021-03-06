// Code generated for package crd by go-bindata DO NOT EDIT. (@generated)
// sources:
// config/crds/kudo.dev_instances.yaml
// config/crds/kudo.dev_operators.yaml
// config/crds/kudo.dev_operatorversions.yaml
package crd

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configCrdsKudoDev_instancesYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x5a\x5b\x6f\x1b\xb9\xf5\x7f\xd7\xa7\x38\xf0\x3e\x38\xc6\x5a\xa3\x7f\xfc\x07\x8a\x42\x28\x0a\xa4\x4e\xb6\x50\x9b\xda\x46\xe4\xa4\x58\x64\x53\x80\x22\x8f\x34\xac\x39\xe4\x84\x17\xc9\xea\x66\xbf\x7b\x71\xc8\x99\xd1\x65\x66\x74\xc9\x6e\xb6\x7c\xb1\x87\x43\x9e\xcb\xef\x5c\x78\x78\x34\x83\xe1\x70\x38\x60\xa5\xfc\x80\xd6\x49\xa3\xc7\xc0\x4a\x89\xcf\x1e\x35\x3d\xb9\xec\xe9\x8f\x2e\x93\x66\xb4\x7c\x39\x43\xcf\x5e\x0e\x9e\xa4\x16\x63\xb8\x0d\xce\x9b\xe2\x1d\x3a\x13\x2c\xc7\xd7\x38\x97\x5a\x7a\x69\xf4\xa0\x40\xcf\x04\xf3\x6c\x3c\x00\x60\x5a\x1b\xcf\x68\xda\xd1\x23\x00\x37\xda\x5b\xa3\x14\xda\xe1\x02\x75\xf6\x14\x66\x38\x0b\x52\x09\xb4\x91\x43\xcd\x7f\xf9\x7f\xd9\x4d\xf6\x87\x01\x00\xb7\x18\xb7\x3f\xca\x02\x9d\x67\x45\x39\x06\x1d\x94\x1a\x00\x68\x56\xe0\x18\xa4\x76\x9e\x69\x8e\x2e\x7b\x0a\xc2\x64\x02\x97\x03\x57\x22\x27\x66\x0b\x6b\x42\x39\x86\x66\x3e\x6d\xa9\xe4\x48\x3a\x4c\xaa\xdd\x71\x4a\x49\xe7\xff\xbe\x33\xfd\x56\x3a\x1f\x5f\x95\x2a\x58\xa6\xb6\xb8\xc5\x59\x27\xf5\x22\x28\x66\x37\xf3\x03\x00\xc7\x4d\x89\x63\xb8\x23\x56\x25\xe3\x28\x68\x2e\xcc\x6c\x85\x53\xc5\xde\x79\xe6\x83\x1b\xc3\xcf\xbf\x0c\x00\x96\x4c\x49\x11\xb5\x4c\x2f\x4d\x89\xfa\xd5\xc3\xe4\xc3\xff\x4f\x79\x8e\x05\x4b\x93\x00\x02\x1d\xb7\xb2\x8c\xeb\x1a\x11\x41\x3a\xf0\x39\x42\x5a\x0a\x73\x63\xe3\x63\x23\x28\xbc\x7a\x98\x64\x15\x81\xd2\x9a\x12\xad\x97\xb5\x10\x34\xb6\x8c\xde\xcc\xed\xb1\xba\x24\x59\xd2\x1a\x10\x64\x66\x4c\x2c\x2b\x63\xa1\x00\x97\x98\x9b\x39\xf8\x5c\x3a\xb0\x58\x5a\x74\xa8\x93\xe1\xb7\xc8\x02\x2d\x61\x1a\xcc\xec\xdf\xc8\x7d\x06\x53\xb4\x44\x04\x5c\x6e\x82\x12\xe4\x1b\x4b\xb4\x1e\x2c\x72\xb3\xd0\xf2\x3f\x0d\x65\x07\xde\x44\x96\x8a\x79\xac\x4c\x52\x0f\xa9\x3d\x5a\xcd\x14\xa1\x18\xf0\x1a\x98\x16\x50\xb0\x35\x58\x24\x1e\x10\xf4\x16\xb5\xb8\xc4\x65\xf0\x0f\x63\x09\xa2\xb9\x19\x43\xee\x7d\xe9\xc6\xa3\xd1\x42\xfa\xda\xcd\xb9\x29\x8a\xa0\xa5\x5f\x8f\xa2\xb3\xca\x59\xf0\xc6\xba\x91\xc0\x25\xaa\x91\x93\x8b\x21\xb3\x3c\x97\x1e\xb9\x0f\x16\x47\xac\x94\xc3\x28\xb8\x8e\x5e\x9e\x15\xe2\xbb\xc6\xd6\x97\x5b\x92\xfa\x35\xb9\x85\xf3\x56\xea\x45\x33\x1d\xbd\xb0\x17\x77\x72\x46\xb2\x2f\xab\xb6\x25\xf9\x37\xf0\xd2\x14\xa1\xf2\xee\xcd\xf4\x11\x6a\xa6\xd1\x04\xbb\x98\x47\xb4\x37\xdb\xdc\x06\x78\x02\x4a\xea\x39\xda\x64\xb8\xb9\x35\x45\xa4\x88\x5a\x94\x46\x6a\x1f\x1f\xb8\x92\xa8\x77\x41\x77\x61\x56\x48\x4f\x96\xfe\x1c\xd0\x79\xb2\x4f\x06\xb7\x31\xd8\x61\x86\x10\x4a\xc1\x3c\x8a\x0c\x26\x1a\x6e\x59\x81\xea\x96\x39\xfc\xe6\xb0\x13\xc2\x6e\x48\x90\x1e\x07\x7e\x3b\x47\xed\x2e\x4c\x68\x35\xd3\x75\x32\xe9\xb4\x50\x1d\x84\xd3\x12\xf9\x4e\x68\x08\x74\xd2\x92\xfb\x7a\xe6\x91\x9c\xbe\x5e\x99\x6d\x91\xea\x0a\xc7\x2a\xfc\x2d\xf3\xc6\x76\xc4\x65\x4b\x82\xfb\xdd\xb5\x51\x5c\x39\x97\x48\x4e\x63\x71\x8e\x16\x29\x47\x78\x43\x3e\x94\x5e\xf1\xfd\x3d\x7b\xe4\x6b\x7f\xc9\xf6\xe6\xfb\xa4\x85\xde\x24\xd2\x29\xf0\xab\x87\x49\x9d\x38\x52\xbe\xc0\x5a\xce\x16\xc7\x5e\xe3\xd5\x63\x2e\x51\x89\x07\xe6\xf3\xa3\x5c\x2f\x27\xf3\xc4\x26\x86\x51\x84\xa3\x94\xc8\x71\x27\x1f\xc5\xa4\x89\x4c\xa4\xc9\x0e\x92\x00\xe4\x6d\x16\xab\xf5\xd7\x29\x68\xaa\xd8\xdc\xe4\x30\xcf\xa4\x06\x96\xb2\x3a\xfc\x6d\x7a\x7f\x37\xfa\xab\x49\xb2\x76\xd2\x64\x9c\xa3\x73\xc9\x55\x0a\xd4\xfe\x1a\x5c\xe0\x39\x30\x57\x7b\xd1\x94\xde\x64\x05\xd3\x72\x8e\xce\x67\x15\x07\xb4\xee\xe3\xcd\xa7\x2e\xcc\x00\x7e\x30\x16\xf0\x99\x15\xa5\xc2\x6b\x90\x09\xe5\x26\x0b\xd4\x4e\x21\x5d\x02\xa2\xa1\x07\x2b\xe9\x73\xd9\xad\x38\x83\xd2\x88\x4a\xe1\x55\x54\xd4\xb3\x27\x04\x53\x29\x1a\x10\x94\x7c\xc2\x31\x5c\x90\x97\x6d\x89\xf8\x33\x1d\xb9\xbf\x5c\x74\xd2\x7c\xb1\xca\xd1\x22\x5c\xd0\x92\x8b\x24\x58\x93\xe8\x69\xae\xf6\x8f\x8d\x80\x3e\x67\x1e\xbc\x95\x8b\x05\x5a\xec\x46\x33\x66\x2f\xca\x0a\x57\x60\x2c\xe9\xae\xcd\x16\x81\x48\x96\x6c\x56\x85\x89\x68\x09\xfc\xf1\xe6\x53\x8f\xb4\xbb\x38\x81\xd4\x02\x9f\xe1\x06\xa4\x4e\xa8\x94\x46\x5c\x65\xf0\x18\x3d\x62\xad\x3d\x7b\x26\x3e\x3c\x37\x0e\x35\x18\xad\xd6\xdd\xd2\x1a\xc8\xd9\x12\xc1\x99\x02\x61\x85\x4a\x0d\x53\x16\x11\xb0\x62\x6b\xd2\xbf\x36\x17\x79\x18\x83\x92\x59\xbf\x7b\x84\x76\x52\x7d\xbc\x7f\x7d\x3f\x4e\x52\x91\x0b\x2d\x34\x89\x42\xa9\x79\x2e\xe9\xa0\xa4\x13\x32\xa5\x7b\xf2\xc9\x08\x47\x48\xce\xe1\x0d\xf0\x9c\xe9\x05\x76\x92\x8d\x9a\x22\xcc\x03\x25\xe0\xec\xf2\xdc\x68\xdd\x3f\xeb\xea\xd1\x71\xe6\xed\x27\x86\xff\xd1\xc9\x71\x92\x5a\xb1\x0c\x3d\xaa\xd6\xdd\x96\x3f\x1f\x54\x8b\x0a\x62\xab\xd1\x63\xd4\x4c\x18\xee\x48\x29\x8e\xa5\x77\x23\xb3\x44\xbb\x94\xb8\x1a\xad\x8c\x7d\x92\x7a\x31\x24\x47\x1c\x26\x4f\x70\xa3\x58\xdc\x8e\xbe\x8b\x7f\xbe\x4a\x8b\x58\xae\x9e\xa6\x4a\x5c\xfa\x7b\xe8\x43\x7c\xdc\xe8\x6c\x75\xea\x62\xe8\xd4\x53\xe9\x72\x5a\x1f\x8e\x7b\x3b\x29\x24\x56\xb9\xe4\x79\x5d\xd9\x6e\xb2\x67\x67\x8c\x14\x4c\xa4\x94\xcb\xf4\xfa\x9b\xbb\x2d\x01\x19\x2c\xc9\xb3\x1e\x56\xf7\xaa\x21\xd3\x82\xfe\x77\xd2\x79\x9a\x3f\x1b\xb9\x20\x4f\x08\xd2\xf7\x93\xd7\xbf\x8f\x33\x07\x79\x76\x44\x76\x56\x71\x34\x4a\x66\x59\x81\x1e\x6d\xab\x80\x61\x42\xc4\x9b\x2b\x53\x0f\x07\x8a\x9c\xaf\xe2\xa9\x98\x7e\xf3\x8c\x3c\xf8\x63\x85\xdc\xe5\x63\x3c\x0c\x99\x45\xf0\x2b\x43\xe9\x9f\x4a\x38\xda\x0f\x58\x13\x00\xce\x34\x95\xd7\xcd\x09\x38\x06\x78\x79\xd5\x12\x54\x6a\x21\x2d\x72\xaf\xd6\xe0\x73\x6b\xc2\x22\xaf\x0a\xf2\x78\x74\x00\x37\xd6\xa2\x2b\x8d\x16\x74\xa8\x34\xa8\xd4\xe9\x7d\xbb\xa6\xcd\x1e\x1a\xcc\x5a\x5c\x0a\x56\x02\xdc\x5c\x41\x8b\x97\x43\x1f\x6f\x26\x95\x83\xec\xd2\xdb\xc6\x23\x3e\x51\x36\xe9\x2e\xec\xe0\x9f\xb9\x54\xd8\x68\x03\x2f\x5e\x5e\xd5\x9a\x3b\xc8\x59\x59\xa2\x76\x74\xd4\xdb\x35\x78\x59\x20\x30\x08\x0e\x6d\x75\x80\xb5\xe5\x65\x1b\x55\xaf\x81\x6d\xc4\x7e\x71\x73\xb5\x01\x34\x01\x1e\x03\xdd\xd1\x15\x49\x34\x17\x6a\x27\x7d\x48\x7d\x8c\x16\xe5\x55\x8e\x7a\xcb\xbb\x40\x18\x74\xfa\xf2\xd2\x57\xa2\x00\x66\x8b\x8c\xd8\xa3\x95\x46\x48\x0e\x33\xc6\x9f\x42\x19\xeb\xaf\xde\x52\x86\xa2\xc3\x4a\x51\xdf\xf0\xf0\x59\xba\x08\x6a\xb5\x77\x2e\x15\x66\xf0\xaa\xf1\x5b\xb5\xae\x6a\x33\x13\x51\xb1\xc6\x14\x6d\x50\x8d\x25\x07\xe2\xa8\x62\x31\x41\xc7\xec\x86\x49\xca\x23\x84\x87\x0d\x5a\x47\xc7\x50\x4c\xbb\xbd\x33\xbf\x45\xf3\xce\x78\x1c\xc3\x8e\x55\x2b\xe3\xd5\xb7\xa1\x08\x68\x2c\xbb\x88\x63\x8f\xef\xb5\x31\x8d\x95\xde\x64\x0a\xb7\xef\xdf\xbd\x7b\x73\xf7\xf8\xf6\xc7\x2a\x0a\xe8\x52\x79\x1f\xaf\x34\x5b\x4d\x8e\xad\xa6\x12\xbc\x98\xdc\x5e\x11\xb4\xc2\xe8\xb6\x5f\xc5\xc2\x2d\xe1\x59\x49\x7b\xbd\x5d\x09\xad\xa4\x52\x14\x5f\x5c\x21\xb3\xc4\xe9\x0d\xe3\xf9\x5e\x0c\xb6\x68\xe6\x8c\x02\x35\x68\xf9\x39\x20\x50\x62\x74\xa6\xbe\x0b\x44\xbf\x21\xd5\x23\x89\x19\x25\xcb\xe1\xc6\xd5\xa4\x4f\x0c\xa9\x00\xec\xf0\x56\x8d\x2b\x22\xb7\x9f\xfd\x0e\x5d\xc3\xca\x2a\x9e\xba\x12\xf8\xaf\x4e\xfa\xa4\x5a\xec\x45\x10\xa5\x64\x9f\xdc\x28\xe1\x6a\xd5\x27\xaf\xab\xf6\xca\x35\x48\xcd\x55\x10\x5d\x8c\x68\xbc\x7f\x3f\x79\xed\x32\x80\xbf\x20\x67\xc1\x51\xd9\x4b\xc6\xba\xf4\x70\x7f\xf7\xf6\x47\x8a\xe1\xb4\xa2\xb2\x0c\xb1\xd4\xc0\x94\x4c\x4d\xa0\xa4\x40\xdc\xdd\x47\xbf\x92\x90\xb3\x92\x7c\xd6\xc5\x06\x91\xf6\xd1\xfd\x72\x54\xa5\x83\x82\x6e\x2e\x2e\xd8\x4a\x0b\x62\x16\xdf\xc6\x33\xa7\x93\xa4\x30\xb1\x7c\x5e\xa0\x27\x57\x9b\xab\xd8\xdc\xf8\x4d\x8e\xa5\xee\x9e\x43\x6a\x0e\x1e\xed\x3a\xc4\x65\x3b\x7d\x07\x33\xab\xd2\x56\xab\xf1\x70\x42\xdf\x81\x2d\x16\x16\x17\xa4\xdb\xb4\x25\x40\x4b\x88\x57\x7b\x8b\xc9\x50\xf5\x91\x5e\x5d\x51\x9a\xf0\x74\xb5\xa0\x56\x2e\x3b\x72\x5d\xd3\x75\x8a\x61\x92\x16\x9f\xd3\x7c\xe0\x5e\x2e\xf1\xe1\x6b\x7d\xbf\x0d\x76\xa7\xbe\x4d\x7e\xab\xd4\x6d\xdc\x2b\x06\x79\x8d\x76\xac\x88\x8c\x52\x26\x9c\xdb\xcd\x38\x58\x42\x74\xdb\xe3\xb4\xb2\x65\x47\x8b\x8b\x87\x86\x1a\x6c\x37\x6a\x63\x03\x23\x4d\xc7\x1a\x21\x5a\xe2\x27\x0d\x8f\x39\xba\xae\xcb\x20\x95\x29\xa9\xbf\x11\x55\x4f\xf1\xe3\x2d\xd3\x2e\x4a\xe4\x68\x6f\xf7\xf8\x02\x77\x74\x66\x77\xd0\xac\x13\x3c\x7c\xe9\xd9\xba\x21\x71\x78\x2c\x3b\x88\xc7\x3d\x6f\xac\x35\x36\x3e\xfd\x69\x18\xc7\x9f\xe3\xf4\x03\xa6\xe3\x68\x9b\xf6\xbf\xfa\x78\x77\xd0\xfe\xe9\xb0\x58\xcb\x1e\xb1\xbf\x4f\x32\x0c\xeb\xbf\xc3\xef\x4f\xa7\xdd\xde\xdb\xc3\xe4\x04\xb9\x97\xbd\x72\x7f\x81\x1f\x98\x67\x0a\x30\xe2\xd6\x90\x8e\xff\xdc\x9a\xa2\x54\xe8\xbb\x9c\x83\xe8\x7e\x69\x37\x51\x0e\xc5\x30\x80\x62\xce\xbf\x4f\x6d\xe3\xcd\x4f\x3d\x9d\xc9\x78\x6e\x6c\xc1\xfc\x18\x68\xed\x90\x2a\xbf\xce\x55\x3a\x28\xc5\x66\x0a\xc7\xe0\x6d\xe8\x5e\x72\x30\x2d\x00\x14\xe8\x1c\x5b\x74\x26\x94\xa3\x7b\xfb\x9a\x02\x47\x37\x96\x39\x73\xdd\x00\x01\x48\x8f\x45\xcf\xab\xbd\x30\x7f\x20\x2a\xa7\x84\x39\xad\xeb\x21\x78\xd8\x5c\x69\x1c\x84\xe8\x24\x7d\xd3\xe8\x87\xeb\x0c\x22\xfd\x89\xbc\x1e\xbf\x7d\x42\x3f\x53\x40\x2c\x0f\xca\x77\xd0\xc0\x1d\x2a\x4c\x3d\x96\x27\x58\x99\xf8\x1e\x24\x7a\x8a\xa9\xd3\x38\xc1\xe0\x69\x9c\x04\x48\x1a\xc7\x8c\x7f\x36\xc1\xe3\x8e\x90\xc6\xf9\xee\x70\x94\x24\x9c\xea\x30\x67\x2a\xd5\x5b\x1d\x74\x2d\x63\xd6\xb2\xee\x76\xf3\x09\x84\x0e\x93\x38\x04\xed\xb7\x88\xae\x23\x00\xf5\xdc\x5b\xbe\xd5\xcd\xe5\xd7\xdd\x5d\x7a\x48\xee\xde\x68\xce\xbd\xbd\xf4\xc9\xb9\x73\xa7\x39\xf9\xfe\x72\x04\xf0\x03\xce\xb3\x03\xb8\x53\x92\x63\xf5\x33\xd1\x0c\x01\x75\xec\x0b\xc5\xfe\xd6\x2c\x78\x02\x8d\xa7\x9f\x8a\x09\xb0\xb4\x78\x96\x00\x6d\xf7\x41\xac\x40\x4b\xae\xe2\xf0\x73\x48\x8d\x57\x0d\x6b\x56\xa8\xf8\xeb\x8a\xd1\x4e\x8a\x78\x99\x76\x72\xa1\xe5\x5c\x72\xa6\x3d\xac\x62\xdf\x28\xb2\x93\xfe\xb2\x7d\xa3\xd3\x66\x5f\xfa\xd3\x2f\x67\x7b\x53\x9b\xef\x55\xaa\x4f\x63\x9a\xa9\x18\x24\xc3\xea\x23\x95\xcd\x5b\x80\x74\x41\xdb\x2a\x4b\x9c\x37\x96\x72\x6a\x9a\xd9\x44\x18\xe3\x1c\x4b\x8f\xe2\x6e\xff\xa3\x95\x8b\x54\x5b\xd5\xdf\xa4\xc4\x47\x6e\x74\xba\x0c\xb8\x31\x7c\xfc\x34\x48\x54\x51\x7c\xa8\x85\xa1\xc9\xff\x06\x00\x00\xff\xff\xb9\xf2\xc3\xaa\xe5\x23\x00\x00")

func configCrdsKudoDev_instancesYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_instancesYaml,
		"config/crds/kudo.dev_instances.yaml",
	)
}

func configCrdsKudoDev_instancesYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_instancesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_instances.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configCrdsKudoDev_operatorsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4d\x8f\x1b\x37\x0c\xbd\xcf\xaf\x20\xd2\x43\x2e\xf5\xb8\x69\x81\xa2\xf0\x2d\xd8\xe6\xb0\x68\x93\x06\xd9\x60\x2f\x45\x0f\x9c\x11\x6d\xb3\x3b\x23\xa9\x24\x65\x74\xf3\xeb\x0b\x49\xe3\x8f\xf1\xda\xbb\x6d\x81\xcc\x4d\x4f\x22\x1f\xf5\xf8\x31\x6a\x16\x8b\x45\x83\x91\xef\x49\x94\x83\x5f\x01\x46\xa6\xbf\x8d\x7c\x5e\x69\xfb\xf0\x93\xb6\x1c\x96\xbb\x37\x1d\x19\xbe\x69\x1e\xd8\xbb\x15\xdc\x24\xb5\x30\x7e\x22\x0d\x49\x7a\xfa\x99\xd6\xec\xd9\x38\xf8\x66\x24\x43\x87\x86\xab\x06\x00\xbd\x0f\x86\x19\xd6\xbc\x04\xe8\x83\x37\x09\xc3\x40\xb2\xd8\x90\x6f\x1f\x52\x47\x5d\xe2\xc1\x91\x14\x86\x3d\xff\xee\xbb\xf6\xfb\xf6\xc7\x06\xa0\x17\x2a\xe6\x9f\x79\x24\x35\x1c\xe3\x0a\x7c\x1a\x86\x06\xc0\xe3\x48\x2b\x08\x91\x04\x2d\x88\xb6\x0f\xc9\x85\xd6\xd1\xae\xd1\x48\x7d\x26\xdb\x48\x48\x71\x05\x07\xbc\x9a\x4c\x71\xd4\x3b\xfc\x36\x59\x17\x68\x60\xb5\x5f\x66\xf0\xaf\xac\x56\xb6\xe2\x90\x04\x87\x13\xb6\x82\x2a\xfb\x4d\x1a\x50\x8e\x78\x03\xa0\x7d\x88\xb4\x82\x0f\x99\x2a\x62\x4f\xae\x01\xd8\xe1\xc0\xae\x5c\xa3\x92\x87\x48\xfe\xed\xc7\xdb\xfb\x1f\xee\xfa\x2d\x8d\x58\x41\x00\x47\xda\x0b\xc7\x72\xee\x10\x03\xb0\x82\x6d\x09\xea\x51\x58\x07\x29\xcb\x3d\x23\xbc\xfd\x78\x3b\x99\x47\xc9\xa0\xf1\xfe\x8a\xf9\x3b\xc9\xe9\x01\x3b\x23\x7a\x9d\x23\xa9\x67\xc0\xe5\x2c\x52\x25\x9c\x72\x41\x0e\xb4\x52\x87\x35\xd8\x96\x15\x84\xa2\x90\x92\xaf\x79\x3d\x71\x0b\xf9\x08\x7a\x08\xdd\x9f\xd4\x5b\x0b\x77\x24\xd9\x09\xe8\x36\xa4\xc1\xe5\xd4\xef\x48\x0c\x84\xfa\xb0\xf1\xfc\xe5\xe0\x59\xc1\x42\xa1\x1c\xd0\x68\x52\x7c\xff\xb1\x37\x12\x8f\x43\xd6\x30\xd1\xb7\x80\xde\xc1\x88\x8f\x20\x94\x39\x20\xf9\x13\x6f\xe5\x88\xb6\xf0\x3e\x08\x01\xfb\x75\x58\xc1\xd6\x2c\xea\x6a\xb9\xdc\xb0\xed\xab\xb8\x0f\xe3\x98\x3c\xdb\xe3\xb2\xd4\x22\x77\x29\x27\x74\xe9\x68\x47\xc3\x52\x79\xb3\x40\xe9\xb7\x6c\xd4\x5b\x12\x5a\x62\xe4\x45\x09\xdc\x97\x22\x6e\x47\xf7\x8d\x4c\x25\xaf\xaf\x4f\x22\xb5\xc7\x9c\x75\x35\x61\xbf\x39\xc0\xa5\xc8\xae\xea\x9e\x6b\x2d\x67\x17\x27\xb3\x1a\xff\x51\xde\x0c\x65\x55\x3e\xbd\xbb\xfb\x0c\x7b\xd2\x92\x82\xb9\xe6\x45\xed\xa3\x99\x1e\x85\xcf\x42\xb1\x5f\x93\xd4\xc4\xad\x25\x8c\xc5\x23\x79\x17\x03\x7b\x2b\x8b\x7e\x60\xf2\x73\xd1\x35\x75\x23\x5b\xce\xf4\x5f\x89\xd4\x72\x7e\x5a\xb8\x29\xbd\x0c\x1d\x41\x8a\x0e\x8d\x5c\x0b\xb7\x1e\x6e\x70\xa4\xe1\x06\x95\xbe\xba\xec\x59\x61\x5d\x64\x49\x5f\x16\xfe\x74\x04\xcd\x0f\x56\xb5\x0e\xf0\x7e\x56\x5c\xcc\xd0\xbe\x05\xef\x22\xf5\xb3\xd6\x70\xa4\x2c\xb9\x7c\x0d\x8d\x72\xd1\xcf\xe6\xc8\xf5\x6e\x3c\x67\x98\x6d\x5c\xb9\x4a\xa9\xa3\xd4\x91\x78\x32\xd2\x0b\xcd\xfc\x82\xa5\x0b\xff\xd5\x66\x44\xf6\x86\xec\x49\xf4\xdc\x86\x8d\xc6\x27\xe0\x99\x6a\xef\x0f\xe6\x13\xde\x91\xe6\xa9\x70\x18\x68\x47\xff\xed\x13\x4f\xd7\x54\xab\x1f\x8d\xc8\xc3\xa5\x8d\xb3\x10\xde\xe5\x73\xa5\xb5\x3c\x84\x82\xe1\x50\x8d\x01\x9d\x13\xd2\x32\x71\x72\x1d\x62\x5f\x9a\xe0\xa2\xcb\xfa\xbf\x70\xcf\xc6\xfb\xac\x90\x47\x27\xff\x22\xe6\xfc\xc3\xa8\xd3\x20\x29\x49\xb1\x82\x20\x10\x64\x83\x9e\xbf\x94\x51\x5b\xc0\xff\x11\xc3\xc5\xca\x3f\xdd\x42\x11\x7c\x9c\xed\x24\x79\xa2\xf3\x15\x8a\xcb\x6d\x65\x68\x49\x5f\x6e\xac\x72\x6c\xd6\x5a\xa1\xd3\x3c\xbc\x9e\xef\xad\x0b\x9c\x67\xd0\xf1\x21\x31\xbd\x59\x0e\x50\x89\x6a\x31\xbd\x1e\x8e\xbb\x00\x95\x77\x05\x26\xa9\xd6\x83\x5a\x10\xdc\xd0\x84\x1c\xaf\x84\x7d\x4f\xd1\xc8\x7d\x38\x7f\x4d\xbc\x7a\x35\x7b\x2c\x94\x65\x1f\xbc\xe3\xfa\xfe\x81\xdf\xff\x68\xaa\x57\x72\xf7\xfb\x60\x32\xf8\x4f\x00\x00\x00\xff\xff\x7d\x5f\xed\xd8\x7e\x09\x00\x00")

func configCrdsKudoDev_operatorsYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_operatorsYaml,
		"config/crds/kudo.dev_operators.yaml",
	)
}

func configCrdsKudoDev_operatorsYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_operatorsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_operators.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configCrdsKudoDev_operatorversionsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\xdd\x6f\xe3\x36\x12\x7f\xcf\x5f\x31\x48\x1f\xf6\x0a\xc4\xf6\x75\x7b\x38\x1c\xfc\xb6\xd8\x6d\x0f\xb9\xb6\x49\xb0\xc9\xf6\x70\x68\x7b\x08\x2d\x8d\x2d\x36\x14\xa9\x92\x94\xbd\xbe\x45\xff\xf7\xc3\xf0\x43\x92\x6d\x51\x56\xbc\x9b\xde\x1e\x10\xbd\x24\x96\x28\xce\x07\xe7\xe3\x37\x43\xea\x6c\x32\x99\x9c\xb1\x8a\xff\x88\xda\x70\x25\xe7\xc0\x2a\x8e\xef\x2d\x4a\xfa\x65\xa6\x0f\x7f\x33\x53\xae\x66\xeb\xaf\x16\x68\xd9\x57\x67\x0f\x5c\xe6\x73\x78\x5d\x1b\xab\xca\xb7\x68\x54\xad\x33\x7c\x83\x4b\x2e\xb9\xe5\x4a\x9e\x95\x68\x59\xce\x2c\x9b\x9f\x01\x30\x29\x95\x65\x74\xdb\xd0\x4f\x80\x4c\x49\xab\x95\x10\xa8\x27\x2b\x94\xd3\x87\x7a\x81\x8b\x9a\x8b\x1c\xb5\xa3\x10\xe9\xaf\xff\x3c\x7d\x39\xfd\xeb\x19\x40\xa6\xd1\xbd\x7e\xc7\x4b\x34\x96\x95\xd5\x1c\x64\x2d\xc4\x19\x80\x64\x25\xce\x41\x55\xa8\x99\x55\x3a\xbc\x69\xa6\x0f\x75\xae\xa6\x39\xae\xcf\x4c\x85\x19\xd1\x5c\x69\x55\x57\x73\x68\xee\xfb\x37\x03\x3b\x5e\x94\xeb\x30\x49\x10\xdf\x3d\x11\xdc\xd8\xef\xfa\x9e\x7e\xcf\x8d\x75\x23\x2a\x51\x6b\x26\x0e\x59\x70\x0f\x0d\x97\xab\x5a\x30\x7d\xf0\xf8\x0c\xc0\x64\xaa\xc2\x39\x5c\x11\x1b\x15\xcb\x30\x3f\x03\x58\x33\xc1\x73\x27\xa9\x67\x4c\x55\x28\x5f\xdd\x5c\xfe\xf8\xf5\x6d\x56\x60\xc9\xfc\x4d\x80\x1c\x4d\xa6\x79\xe5\xc6\xed\x33\x06\xdc\x80\x2d\x10\xfc\x1b\xb0\x54\xda\xfd\xdc\x67\x0f\x5e\xdd\x5c\x4e\xc3\x74\x95\xa6\xa7\x96\x47\x75\xd0\xd5\x31\x83\xe6\xde\x1e\xe1\x17\xc4\x59\x20\x9a\xd3\xc2\xa3\xa7\x1c\x48\x60\x0e\xc6\xf3\xa0\x96\x60\x0b\x6e\x40\x63\xa5\xd1\xa0\xf4\xa6\xd0\x99\x16\x68\x08\x93\xa0\x16\xbf\x62\x66\xa7\x70\x8b\x8e\x4f\x30\x85\xaa\x45\x4e\xd6\xb2\x46\x6d\x41\x63\xa6\x56\x92\xff\xa7\x99\xd9\x80\x55\x8e\xa4\x60\x16\xc3\x7a\xc4\x8b\x4b\x8b\x5a\x32\x41\x3a\xad\xf1\x02\x98\xcc\xa1\x64\x5b\xd0\x48\x34\xa0\x96\x9d\xd9\xdc\x10\x33\x85\x1f\x94\x46\xe0\x72\xa9\xe6\x50\x58\x5b\x99\xf9\x6c\xb6\xe2\x36\x1a\x7e\xa6\xca\xb2\x96\xdc\x6e\x67\xce\x7c\xf9\xa2\xb6\x4a\x9b\x59\x8e\x6b\x14\x33\xc3\x57\x13\xa6\xb3\x82\x5b\xcc\x6c\xad\x71\xc6\x2a\x3e\x71\x8c\x4b\x67\xf7\xd3\x32\xff\x42\x07\x2f\x31\x2f\x3a\x9c\xda\x2d\x59\x81\xb1\x9a\xcb\x55\x73\xdb\x19\x64\x52\xef\x64\x90\xb4\xcc\x2c\xbc\xe6\xf9\x6f\xd5\x4b\xb7\x48\x2b\x6f\xbf\xb9\xbd\x83\x48\xd4\x2d\xc1\xae\xce\x9d\xb6\xdb\xd7\x4c\xab\x78\x52\x14\x97\x4b\xd4\x7e\xe1\x96\x5a\x95\x6e\x46\x94\x79\xa5\xb8\xb4\xee\x47\x26\x38\xca\x5d\xa5\x9b\x7a\x51\x72\x4b\x2b\xfd\x5b\x8d\xc6\xd2\xfa\x4c\xe1\xb5\x73\x7f\x58\x20\xd4\x55\xce\x2c\xe6\x53\xb8\x94\xf0\x9a\x95\x28\x5e\x33\x83\x4f\xae\x76\xd2\xb0\x99\x90\x4a\x8f\x2b\xbe\x1b\xb5\x76\x07\x7a\x6d\x35\xb7\x63\x5c\xe9\x5d\xa1\x3d\x97\xbc\xad\x30\xdb\xf1\x90\x1c\x0d\xd7\x64\xc5\x96\x59\x24\xdb\xdf\x7b\x61\xda\x99\xb8\xcf\x39\xbd\x83\x56\x3d\x0e\x9a\x14\x0c\x7c\xd4\x95\x98\x11\x8b\xb7\xee\xe1\xfe\x8b\x3b\x32\xbc\xde\x1b\xdc\x08\xc0\xc0\x62\x59\x91\xc7\xe5\xd1\xfe\x6c\xc1\x2c\x64\x4c\xc2\x02\xf7\xa6\x04\xa8\x0d\xe6\xe4\xa6\x81\x38\xfd\xcb\x24\x70\x69\x2c\x93\x19\xfa\xd8\x80\x8d\x02\xa6\x63\x65\xc9\xb1\x42\x99\xa3\xcc\x0e\x14\xb3\x27\xc7\x9b\xce\x40\x60\x2e\xa0\xbb\x68\x23\xc4\xce\x1c\x91\x11\x95\x60\x84\x5b\x2c\x0f\x08\x25\x96\xbd\x21\x49\xd1\x66\x89\x1a\x65\xe6\x68\x7b\x0d\xe6\x49\x1a\xe9\xc5\x8e\x4b\xde\x17\x93\x13\xcc\xbc\xba\xb9\x8c\x91\x38\xca\x16\x98\xb1\x87\x74\x07\x55\xed\xaf\x25\x47\x91\xdf\x30\x5b\x8c\xa0\xfd\xe2\x72\xe9\x89\x79\xeb\x50\xc0\xa0\xe2\xe8\x57\xbb\x09\xf3\xce\x06\x90\xe5\xa0\x96\xbd\x33\x12\x6c\x00\x72\x63\x8d\xe1\x8d\x0b\x1f\x8d\x82\xd1\xb5\xc9\xc1\x32\x2e\x81\xf9\xe4\x09\xff\xb8\xbd\xbe\x9a\xfd\x5d\x25\xa6\x74\x52\x00\xcb\x32\x34\xc6\xbb\x5f\x89\xd2\x5e\x80\xa9\xb3\x02\x98\x89\x9e\x79\x4b\x4f\xa6\x25\x93\x7c\x89\xc6\x4e\x03\x0d\xd4\xe6\xa7\x97\xbf\xf4\x6b\x0f\xe0\x5b\xa5\x01\xdf\xb3\xb2\x12\x78\x01\x3c\x58\x53\x0c\xb1\xc1\x0a\x5c\x72\x26\x75\x34\x33\xc2\x86\xdb\x82\xcb\x94\x06\xa0\x52\x79\x10\x7b\xe3\xc4\xb5\xec\x01\x41\x05\x71\x6b\x04\xc1\x1f\x70\x0e\xe7\x14\x8e\x3a\x6c\x7e\x20\x70\xf3\xfb\x79\x62\xd6\x3f\x6d\x0a\xd4\x08\xe7\x34\xe8\xdc\x33\xd7\x64\x52\xba\x17\xed\xa5\x65\xd2\x39\xb8\xd5\x7c\xb5\x42\xed\x80\x4a\xdf\xe5\x12\x04\x05\xde\x2f\x41\x69\xd2\x80\x54\x9d\x29\xdc\xc4\xb4\x7a\x15\x66\x7c\xc9\x31\x3f\x60\xfa\xa7\x97\xbf\x24\x39\xde\xd5\x17\x70\x99\xe3\x7b\x78\x09\x5c\x7a\xdd\x54\x2a\xff\x72\x0a\x77\xce\x3a\xb6\xd2\xb2\xf7\x44\x29\x2b\x94\xc1\x94\x66\x95\x14\x5b\x92\xb9\x60\x6b\x04\xa3\x4a\x84\x0d\x0a\x31\x89\x2e\xba\x61\x5b\xd2\x42\x5c\x38\xb2\x37\x06\x15\xd3\x76\xd0\x5a\x23\x7e\xb9\xbb\x7e\x73\x3d\xf7\x9c\x91\x41\xad\x1c\x28\xa3\x1c\xb8\xe4\x84\x48\x08\x8a\xf8\xbc\xea\xac\x71\x2f\x2d\xb7\x97\xa9\xbd\xf9\x50\xf4\x2c\x98\x5c\xa1\x97\x17\x61\x59\x53\xae\x9b\xbe\x38\xc5\x8f\xf7\xa1\x45\x7b\xf5\x80\x8c\xfd\xc0\xf1\x3f\x4a\xd5\xa3\x85\x73\xd5\xc0\x08\xe1\xae\x3a\x56\x3e\x28\x1c\x55\x26\x5a\xa2\x45\x27\x5f\xae\x32\x43\xa2\x65\x58\x59\x33\x53\x6b\xd4\x6b\x8e\x9b\xd9\x46\xe9\x07\x2e\x57\x13\x32\xcd\x89\xb7\x01\x33\x73\xe5\xc5\xec\x0b\xf7\xe7\x64\x59\x5c\x61\x30\x56\x20\x37\xf8\x8f\x90\x8a\xe8\x98\xd9\x49\x42\x35\x91\xf0\x6a\xdc\x4a\xb9\x85\x8a\x21\xc3\x1c\x44\xa8\x26\x83\x6f\xa7\xf0\x36\x4e\x9d\x8a\x4f\x6b\xee\x52\x30\xab\x85\x35\x14\x76\x96\x7c\x75\x52\x26\x8c\x88\x7a\x7c\x2e\x7e\x71\xeb\x25\xc8\xf6\xdf\x25\xd7\xde\x14\x3c\x2b\x62\x81\x14\x44\x48\x48\xc0\x09\x99\xe7\x3e\xbd\x30\xb9\x7d\x72\x77\x24\xa3\xa8\x35\x71\xb4\x9d\x84\x92\x7d\xc2\x64\x4e\xff\x1b\x6e\x2c\xdd\x3f\xc9\x0a\x6a\x3e\x2a\x04\xbd\xbb\x7c\xf3\xc7\x38\x69\xcd\x4f\x8c\x37\xeb\xd1\x26\x70\x1e\x57\x3c\x63\x15\xa9\xdb\x04\xb1\x7e\xab\xb9\x76\x18\xc4\xb8\x52\x7d\x43\x89\x76\xa7\x85\x70\x78\x05\x95\xb0\x85\x5a\x37\x08\x83\x69\x24\x48\xab\x36\x54\x5d\xfd\x2c\xe1\x1b\x0f\x44\xe6\xf0\xef\xaf\xa7\x5f\x4d\xff\xd2\x9f\x57\x07\x85\x0b\xac\xf5\x2c\xd4\x64\xd7\x8b\x7b\x9e\xaf\x3b\x1d\x94\x43\x82\x7b\x95\x54\xf7\x11\xd3\x9a\x6d\x77\x8b\xd4\x80\x95\x07\x31\xfe\xb5\x9b\xb1\x09\x00\x11\x75\x18\x40\xa9\xea\x55\xe1\xcc\x45\x97\xae\xeb\x40\x1e\x27\xd0\xc2\x56\xd5\x07\xec\x71\x49\x91\xc6\x12\x7a\x29\x55\xce\x97\xdb\xd6\xf4\xa8\x5a\x0b\xd9\x7d\xef\xb5\x21\xc8\x3e\x0c\xd8\x3f\x06\xae\x0f\x2e\xdd\x20\x54\xff\x28\xa0\x0e\xac\x1f\x4f\x9d\x0c\xd3\x3d\xaf\xbd\x73\x3e\x01\x48\xff\xf4\x10\xfd\x29\x00\xfa\x93\xc0\xf3\x27\x03\xe7\x1f\x01\xcd\x1d\x08\xef\xe7\xf6\x24\x60\xde\x81\xe0\xbd\xb3\x3e\x16\x96\x1f\x02\xf0\xde\x69\x8f\x83\xf2\x41\x6f\x4d\x01\xf2\xcf\x1f\x8e\x0f\x8a\x95\x82\xe2\x9f\x1d\x10\x3f\x2a\x45\x12\x84\x7f\x96\x10\xfc\x48\x52\x3f\x0a\x5d\x3f\x0e\xb8\xa6\x8a\xd9\xff\x07\xd8\x3a\xa8\xb9\x04\x64\xfd\xcc\x00\xeb\x80\x08\x49\xec\x55\x31\xcd\x4a\xb4\xa8\x0f\x00\xcc\x98\x9e\xe7\x4d\x7c\x7b\x17\xd8\xae\x99\xe6\x6c\xc1\x05\xb7\xdb\x10\x98\xfb\x76\xd7\x76\xaf\x05\x52\x34\xf7\x9d\x61\xcb\x5d\x7f\x99\x00\x43\xdb\x2c\x7e\x6c\xbf\x34\x14\x7b\x23\xd0\xf9\x1b\x3f\xd2\x6f\xaa\x84\xd7\x42\xfe\xf6\xa9\xb2\x51\x12\x0d\xa9\xb4\x5a\xf3\x3c\x59\x67\x2e\x3c\x6e\x4c\x73\x0d\xc7\x0b\x8b\x2e\x77\x63\xd8\x6f\x7e\xb4\xcb\xc0\x40\x28\xb9\x42\xdd\x1d\x4a\x6b\x51\xa8\xcd\x40\x03\xaf\x15\x74\xc3\x85\x70\x9b\x36\x06\xf3\xd3\x64\xe0\xa6\x12\x6c\x3b\xb2\xd2\x7f\xd3\x8e\x0e\x5b\x09\x7e\xeb\x60\xb1\x85\x77\x97\xe6\x24\x06\x46\x76\x83\xce\xaf\x02\xfa\x21\xf9\xbb\x3b\x1a\x01\xba\x46\x4e\x42\x9e\x8f\xbb\x1f\xc9\x0e\xb3\x40\x57\xca\x75\x81\xe6\xbd\xdf\xa6\x7e\x7d\xfd\xee\xea\xee\x9e\x66\x91\x50\x9b\xb8\x4d\xe7\x7d\x45\x90\xc5\x24\xdb\xc0\x84\xc6\x02\x94\xfc\x59\xfa\xcd\x27\x17\xce\x2b\xc1\x33\x66\xe6\x00\x1f\x3e\xc0\xd4\xf9\xa2\x99\x3a\x2a\xf0\x7b\x02\x5d\x1e\x6d\x6e\xa4\xca\xbe\x03\xbd\xbd\x0d\x43\x3b\xfd\x99\x80\xa9\x77\xbc\x25\xce\x08\x36\xd5\x94\x5f\x60\xe3\x52\xb4\xdc\x4c\x88\xc6\x79\xcc\x05\xb8\xaa\x18\x6d\x81\xba\xe3\x9b\x64\x21\xa6\x5e\x2e\xf9\xb0\x7f\x2d\x94\x12\xd8\x5b\xb3\x04\xb4\x3c\x42\xcc\x3b\x3f\x12\x78\x4e\x29\xa6\x69\x43\x55\x82\x49\x6f\x26\x2b\xb4\x06\xf0\x3d\x66\x35\x85\xac\x4d\x91\xec\x39\x7b\x3c\xdc\x06\x4c\x07\x29\x4d\xb4\xab\xcb\x66\x4b\x2c\x74\x91\x3b\x41\xe9\xde\xef\x9c\xde\xa7\xfa\x42\x4b\x02\xc1\xc4\x90\x83\xe0\x8e\x2b\x07\xe9\xf1\x3d\x37\x96\x74\x48\xea\xdb\x70\x83\xc0\xed\x0b\x03\xf7\x39\x56\x42\x6d\xef\x4f\xf2\x2a\x17\x16\x27\x6e\xd0\x08\xe5\x6d\xab\xfd\xfe\x9d\x0f\xab\xf4\x7e\x23\xa2\x2b\x6f\xee\x3d\xc5\x53\x98\x3a\xa1\xaf\x40\xda\x3a\x48\x1a\x2c\xcf\xdd\xe1\x15\x26\x6e\x06\x12\xcb\x6e\xfe\x23\xad\xb7\x02\x32\x30\xa8\xc3\x76\xe2\x4d\xc1\x8c\x93\x99\x56\x03\xfd\x2e\xe8\x82\xca\x36\x0a\x0b\xb6\x2f\xa8\x0e\xa7\xb3\xca\xcd\x37\x42\xe9\x81\x70\xc9\x2a\x62\xc8\xbd\xe6\xcd\xc1\xd5\xb5\xee\xe9\x60\x9d\x94\xc8\xfb\x29\x4a\x3b\xe2\xc7\xfd\x55\x63\xb1\x0a\xb2\xc7\xd2\xff\xbb\x06\xf5\x24\xa6\x8e\x47\x12\x12\xd1\xfe\x98\x7e\xfc\x95\x0e\xfa\xfe\x3a\x62\xdd\xfe\x72\xdc\x0f\xcd\xb2\xa3\x85\x5b\x27\x6b\x50\x37\xbd\xda\xd1\x76\xd4\x47\xb3\x67\x3e\x30\x29\x74\x54\x14\x55\x01\xc6\x2a\x0a\x9e\xac\x3d\xee\x91\xd2\x0e\x1c\x5b\xba\x04\xeb\x9d\x9d\x7d\x13\xe1\xbe\x41\xc7\xb5\xef\xbe\x0d\xf4\x22\xe3\xe5\x16\x5a\x65\x59\xdd\xb3\xa9\xdd\xbd\xc6\xac\xa0\xbf\x8e\xad\x63\xa0\x3b\x66\x35\xc3\x50\x66\x1e\x8e\x52\x1d\xa5\xc1\x47\x93\x4e\x87\xa1\xfe\x71\xbd\x91\xec\xb1\xd3\x19\xab\x99\xc5\xd5\x76\xb4\x19\x5f\xeb\x1c\x7d\xcb\xae\xf1\xe7\x42\x6d\x3c\x2a\xaa\x17\x4e\x2f\xb1\xab\x33\xbc\xc6\x82\xc9\x99\x8f\x3a\x2d\x82\x72\xa7\xfd\x72\x50\x75\x22\xe6\x74\xe5\x1a\xd4\xe9\x51\x0d\xc9\x5a\x08\x82\x53\x73\xb0\xba\xee\x47\x69\xc3\xea\x1b\x56\xdc\xa9\x2a\xeb\xa8\x25\x21\xd9\x78\x65\x9d\x9a\x0c\x0f\x32\x57\x9b\x24\x28\x8d\xb5\x51\x8b\x7e\xee\x93\x1e\xd4\x6b\x92\x68\xaf\xd7\xed\xf0\xf1\x7d\xe7\x44\x8e\x1b\x0d\x6c\xcd\xb8\x08\x88\xd8\xeb\x6e\xe0\x7c\x14\x8c\x2c\x54\xef\x98\x79\xf0\xf5\xdd\x4a\xa8\x05\x13\x17\x50\x29\xb1\x2d\x95\xae\x0a\x9e\x01\xa7\x9c\x5c\xc6\xa3\x89\x91\x9d\xaa\x5e\x08\x9e\xf5\xf6\x28\x5b\x1e\x1d\xcf\x8f\x4c\xe5\xe9\x4d\xf8\x93\x4b\x9a\x23\x2f\xee\x9f\x57\x1b\xd0\x92\x3b\xae\x86\xe5\x02\x73\xe3\xb5\xa0\x8c\xe1\x51\x52\x37\x91\x09\x0d\x5d\xb7\xe3\x94\x0a\x06\xb5\xef\xa3\xaf\x15\xcf\x61\xa3\xb9\x3b\x95\x98\xb9\xd3\xc2\x50\xcb\x59\xc9\xb4\x29\x98\x10\xae\xb7\x4d\xc9\xc3\x77\xcf\xdd\xb1\x8c\x8a\xe9\xa4\x93\x64\xa8\x1d\x98\x70\x3d\x5a\x13\x36\x80\x69\x6a\x15\xaa\x33\xe2\xf1\x3b\x2e\x73\x62\x11\x21\x57\x1b\x69\x78\x8e\xf1\x0c\x6a\xaa\xc0\xaa\x2a\xad\x58\x56\x00\x37\x17\x9e\x1d\x27\x3f\x15\x24\xae\x07\xea\xea\x0d\xa9\xac\xef\x4a\x07\xda\x01\x6b\x27\xdd\x99\xbc\xe9\x57\xa3\xbc\x5f\x19\xca\xe0\x3c\x8a\xb9\xc0\x4c\x95\x08\xac\x5c\xf0\x55\xad\x6a\xd3\x1c\xd3\x0d\xf5\x4d\x0a\xff\x90\x62\xf4\x14\xfe\x89\x50\xf2\x55\x61\x41\xe3\x9a\x1b\x6e\xbd\x93\xb4\x42\x74\x1b\xd2\x21\xac\x0c\x95\x24\x91\x1b\x09\xdc\x98\x3a\x51\x50\x1d\xcf\xdc\xb9\x92\x03\x19\xfb\x58\x41\x46\xd7\x92\x59\x26\x3e\x6e\x8a\xa6\xbc\x3a\x36\xcd\x60\x92\xa9\x78\xaa\xba\x81\x31\x10\x61\x37\xce\xf2\x0a\xc3\x09\x50\xba\xbb\x08\x39\x82\xf9\x4e\xc1\x0a\x25\x85\x36\x57\xf9\x0e\x66\x55\xe6\x26\x8a\x51\x2c\xa2\x42\x99\xb7\xbd\xd7\x21\x68\x39\x16\x77\x11\x4f\xc7\xc0\xcf\x68\xe0\xf3\x80\x83\xe0\xe3\x71\x73\x25\x23\xe6\xa3\x27\x1b\x85\xb1\x46\xa0\x08\x18\x05\xc4\x2a\x35\xc0\xf7\x08\x8e\x9b\x43\xe3\x1f\x61\x8f\xa3\x14\xf3\xc9\x24\xde\x30\x69\xbf\xd1\x47\x1d\x70\xc8\x8f\x07\x97\xe8\x84\x6a\x3f\x56\x5f\x27\x56\xfc\x03\xfa\xdb\xcd\x9e\x91\x8c\x07\x1a\xb1\xf6\xeb\x1c\x02\xb6\x0a\xfe\xf5\xea\x87\xef\x5b\x86\x40\xa8\xac\xb7\x2c\xdc\xeb\x36\x52\x8a\x10\x39\x6a\xe7\xf2\x74\x43\x77\x1c\x3f\x1c\xd0\x27\x20\xd2\x7f\x80\xba\x47\x59\x75\xb5\xd2\x2c\xa7\x05\xff\x56\xab\x72\x10\xa1\xbd\xdb\x19\xea\xc4\xf2\xc8\x60\x0f\x96\x99\xf6\x20\xb8\x9f\xfd\xd0\x8a\xdc\x36\xf6\xa7\x01\x74\x9f\xe8\xd0\xc7\x89\xc7\x3e\x9e\xcf\x6a\xef\xc9\xfb\x7c\x56\xfb\xf9\xac\xb6\xe7\xf8\xf9\xac\xf6\xf3\x59\xed\x11\xc2\x3d\x9f\xd5\xfe\xec\xcf\x6a\x3f\x9f\x73\x7e\x3e\xe7\x7c\x1a\xe0\x4e\x1c\x8c\x4e\x90\xe9\xff\xc2\xd2\x32\x5b\x9b\xd1\xdf\x58\xba\xd1\x3b\x5f\x59\xaa\x85\x41\xbd\x1e\xf9\x99\x65\x0f\x0b\x7b\xb7\xda\xaf\xd2\xc3\x07\xf0\xcd\x2d\xc7\xe4\x24\x7c\x8a\xde\x3e\x05\xf0\xf4\x3b\x05\x15\x95\xed\x6c\x15\x4b\xac\x56\x42\x02\x39\x95\xc5\xfc\x6a\xff\x9b\xf4\x73\x9f\x65\xe3\x47\xe6\xee\x67\xa6\xa4\x2f\x5a\xcc\x1c\x7e\xfa\xe5\x0c\x42\x33\x20\x82\x70\x77\xf3\xbf\x01\x00\x00\xff\xff\xf7\x67\x74\x00\xcb\x3f\x00\x00")

func configCrdsKudoDev_operatorversionsYamlBytes() ([]byte, error) {
	return bindataRead(
		_configCrdsKudoDev_operatorversionsYaml,
		"config/crds/kudo.dev_operatorversions.yaml",
	)
}

func configCrdsKudoDev_operatorversionsYaml() (*asset, error) {
	bytes, err := configCrdsKudoDev_operatorversionsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/crds/kudo.dev_operatorversions.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"config/crds/kudo.dev_instances.yaml":        configCrdsKudoDev_instancesYaml,
	"config/crds/kudo.dev_operators.yaml":        configCrdsKudoDev_operatorsYaml,
	"config/crds/kudo.dev_operatorversions.yaml": configCrdsKudoDev_operatorversionsYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"config": &bintree{nil, map[string]*bintree{
		"crds": &bintree{nil, map[string]*bintree{
			"kudo.dev_instances.yaml":        &bintree{configCrdsKudoDev_instancesYaml, map[string]*bintree{}},
			"kudo.dev_operators.yaml":        &bintree{configCrdsKudoDev_operatorsYaml, map[string]*bintree{}},
			"kudo.dev_operatorversions.yaml": &bintree{configCrdsKudoDev_operatorversionsYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
