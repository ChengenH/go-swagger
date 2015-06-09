// Code generated by go-bindata.
// sources:
// ../schemas/jsonschema-draft-04.json
// ../schemas/v2/schema.json
// DO NOT EDIT!

package assets

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _jsonschemaDraft04JSON = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x57\x3b\x6f\xdb\x30\x10\xde\xfd\x2b\x04\xa5\x63\x52\xb9\x40\xa7\x6c\x45\xbb\x18\x68\xd1\x0c\xdd\x0c\x0f\xb4\x75\xb2\x19\x50\xa4\x42\x51\x85\x0d\x43\xff\xbd\xa4\xa8\x07\x29\x91\x92\x2d\xbb\x48\xb4\xc4\xe1\xbd\xbe\x3b\xde\x8b\xe7\x45\x20\xbf\x10\xc7\xe1\x73\x10\x1e\x84\xc8\x9e\xa3\xe8\x35\x67\xf4\x29\xdf\x1d\x20\x45\x9f\x19\xdf\x47\x31\x47\x89\x78\x5a\x7e\x8d\xf4\xd9\x43\xf8\xa8\x85\x3e\xe9\xff\x67\x48\xc6\x90\xef\x38\xce\x04\x66\x54\x49\x7f\x67\x1c\x02\xcd\x12\xa4\x20\x50\xad\xa2\xe3\x4e\x30\xc5\x8a\x39\x97\xdc\x1a\x71\x45\xd0\x6c\xdf\x38\x47\x27\x8b\x50\x11\xc5\x29\x03\xa5\x1c\x55\xe4\x47\x9b\x98\x62\xba\x12\x90\x2a\x7d\x5f\x7a\x24\x5c\x9f\x9f\xa5\x83\x1c\x12\xa5\xe2\x21\x0c\xca\x96\xa9\xec\xf8\xc3\x8c\xe5\x12\xd7\x5f\x58\x51\x01\x7b\xe0\x7e\x10\xb8\x66\x18\xc2\xc0\x69\x91\x4a\x8e\xe5\x25\xfa\x7f\x40\x82\x0a\x22\x96\x43\x3b\x88\x90\xdf\x0a\xea\xda\x82\x1d\x19\x91\x8b\xfa\x58\xa5\x21\xc5\x1c\x6b\x9d\x0a\x42\x50\x06\x1b\x27\x8c\x1c\xa7\x19\x81\x3f\xd2\x97\x7c\x68\x1a\x68\xe5\xc0\xba\x8d\x74\x10\x6e\x19\x23\x80\xa8\xfa\xd9\x3a\x1e\x84\xb4\x20\x44\xff\x4d\xb7\xfa\x84\x6d\x5f\x61\x27\xd4\xaf\x5c\x70\x4c\xf7\xa1\xcf\x7e\x45\x9d\x73\xcf\xc6\x65\x36\x7c\x8d\xa9\xf2\xf2\x94\x28\x28\x7e\x2b\xa0\xa1\x0a\x5e\x40\x07\x73\x61\x80\x6d\x6d\x34\x8e\xe9\xd3\x8c\xb3\x0c\xb8\xc0\xbd\xe8\xe9\xa2\xf3\x78\x53\xa3\xec\x01\x49\x18\x4f\x91\xba\xab\xb0\xe0\x38\x74\xc6\xaa\x2b\xca\x7b\x6b\x16\x58\x10\x98\xd4\xeb\x14\xb5\xeb\x7d\x96\x82\x26\x4b\xcf\xe6\x71\x2a\xcf\xb0\x4c\xcd\x2a\xf7\x3d\x6a\x9b\x74\xf3\x56\x5e\x8f\x02\xc7\x1d\x29\x72\x59\x28\xbf\x5a\x16\xfb\xc6\x4d\xfb\xe8\x58\xb3\x8c\x1b\x77\x0a\x77\x86\xa6\xb4\xb4\xf5\x64\x93\xbb\xa0\x24\x88\xe4\x1e\x84\xad\x13\x37\x21\x9c\xd2\x72\x0b\x42\x74\xfc\x09\x74\x2f\x0e\xbd\x9e\x3b\xd5\xbc\x2c\x1f\xaf\xd6\xd0\xb6\x52\xbb\xdf\x22\x21\x80\x4f\xe7\xa8\xb7\x78\xb8\xd4\x7d\x74\x07\x13\xc5\x71\x05\x05\x91\xa6\x91\xf4\x7b\x38\x3d\xe9\x1e\x6e\x1d\xab\xef\x3c\x0c\x74\xbf\x7d\xd5\x6c\xce\x89\xa5\xbe\x8d\xf7\x66\xce\xee\xd1\x86\x67\x80\x34\xad\x8f\xc3\xb3\xae\xc6\x1c\xe3\xb7\xc2\x96\xd9\xb4\x72\x0c\xf0\xab\x92\xe9\x5a\x05\xee\x5c\xb2\x87\xc6\x7f\xa9\x9b\x17\x6b\xb0\xcc\x75\x77\x96\x16\xb7\xcf\x1c\xde\x0a\xcc\x21\x1e\x53\x64\x0e\x73\x4f\x81\xbc\xb8\x07\xa6\xe6\xfa\x50\x55\xe2\x5b\x4d\xad\x4b\xb6\xb6\x81\x49\x77\xc7\xca\x68\x1a\x90\x67\xd7\x78\x3f\x3c\xba\xa3\x8e\xdd\xe8\x7b\xc0\x8a\x21\x03\x1a\x03\xdd\xdd\x11\xd1\x20\xd3\x46\x72\x55\x7d\x93\x0d\xb3\xcf\x34\x52\x46\x03\xd9\x8d\x75\xe2\x0e\x42\xbd\xb9\xdf\xe9\xdd\x34\xb6\x24\x9b\x5b\xa4\x56\x3f\x6b\xac\xd8\x01\x30\x1e\x25\xce\x3a\x77\xc6\x73\xd4\xbd\x96\xc9\xf5\x06\xbc\xca\xf8\x44\xb0\x2e\x09\x5a\xf3\xf5\x3a\x94\x7b\xb7\xa8\x9f\x7f\x17\x8e\x58\x53\xb2\x0e\xfc\xf5\x92\x8c\xc2\x4c\x49\xca\x84\xe7\x7d\x5d\xb6\x2f\x7e\x4f\x79\xba\x96\xe6\x75\xb7\x87\x9b\x0d\xdc\xb5\xbd\xae\xbb\x85\xb8\x8e\x64\x67\xd1\xe8\x18\xe5\xe2\x5f\x00\x00\x00\xff\xff\x4e\x9b\x8d\xdf\x17\x11\x00\x00")

func jsonschemaDraft04JSONBytes() ([]byte, error) {
	return bindataRead(
		_jsonschemaDraft04JSON,
		"jsonschema-draft-04.json",
	)
}

func jsonschemaDraft04JSON() (*asset, error) {
	bytes, err := jsonschemaDraft04JSONBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "jsonschema-draft-04.json", size: 4375, mode: os.FileMode(420), modTime: time.Unix(1441639065, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _v2SchemaJSON = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5d\xdf\x73\xdc\xb6\xf1\x7f\xcf\x5f\x81\xb9\x78\x46\xf6\x24\xd6\x39\xfe\x7e\x5f\xea\x97\x8c\x1a\x39\x89\x5a\xbb\xd2\xf8\x9c\xf6\xc1\x95\x67\x70\x24\x4e\x87\x84\x3f\x2e\x04\x29\xe9\xea\xea\x7f\xef\x02\xfc\x71\x04\x01\x90\x20\x89\x3b\x9d\x6d\x7a\xa6\x8d\x8e\x04\x16\x8b\xc5\x62\xf7\xb3\x0b\x10\xf8\xf4\x0d\x42\xb3\x94\xa6\x01\x99\xbd\x42\xb3\x33\xf4\xb7\xc5\xe5\x3f\xd0\xc2\x5b\x93\x10\xa3\x55\x9c\xa0\xc5\x1d\xbe\xb9\x21\x09\x7a\x79\xfa\x02\x9d\x5d\x5d\x9c\xce\xbe\xe7\x15\xa8\xcf\x4b\xaf\xd3\x74\xf3\x6a\x3e\x67\x79\x91\x53\x1a\xcf\x6f\x5f\xce\x99\xa8\x7b\xfa\x3b\x8b\xa3\x6f\xf3\xc2\x4f\xf2\x47\xb5\x1a\xfc\xe5\xf3\xa2\x60\x9c\xdc\xcc\xfd\x04\xaf\xd2\xe7\x2f\xfe\xbf\xa8\x5c\xd4\x4b\xb7\x1b\xc1\x54\xbc\xfc\x9d\x78\x69\xfe\x2c\x21\x7f\x66\x34\x21\xbc\xf9\x0f\xf0\x1b\x9e\x14\xad\x8b\xd7\x9c\xb3\x68\x15\x97\x7f\x6f\x70\xba\x66\x33\xf8\xfb\x5a\xd4\xc5\xbe\x4f\x53\x1a\x47\x38\xb8\x4a\xe2\x0d\x49\x52\x4a\x18\xd0\x59\xe1\x80\x11\x51\x00\xca\xa7\x24\x89\xa4\xb7\x9f\x72\x52\x1f\xef\x9f\x57\x3f\x78\x97\x12\xb2\xe2\xac\x7d\x3b\xf7\xc9\x8a\x46\x82\x2c\x9b\xdf\x92\xc8\x8f\x93\xd7\xf7\x29\x89\x18\x3c\x98\x89\xd2\x0f\xf0\xff\x0f\x39\x79\x0d\xdd\x92\xfb\x1a\xed\xb2\xdb\x2c\x4d\x68\x74\x53\xf4\x05\x9e\x93\x28\x0b\xab\x6e\x8b\x27\x30\x26\xb3\xe2\xd7\x75\x55\xcc\x27\xcc\x4b\xe8\x86\x73\xc4\xa9\xbc\x5f\x93\x6a\x0c\x6f\x49\xc2\xf9\x42\xf1\x0a\xa5\x6b\xca\x90\x1f\x7b\x59\x48\xa2\xf4\xb4\xe0\xb4\x2e\xc2\xce\xce\x8a\x52\x52\xbd\x75\xcc\x52\x9b\x8e\x14\x62\xe6\xaf\x3e\x7e\xf8\xf8\xe9\x61\x8e\x5e\xfd\x1b\xfe\x5d\x7f\xf7\xf4\xc7\x57\xf0\x97\xff\xdd\xb3\x1f\x9f\xcc\xda\xfa\xc3\x1b\x42\x4f\x23\x1c\x12\x04\x1a\x4a\x37\xcf\xf2\x1e\x11\xa1\xa0\xe8\xf5\x3d\x0e\x37\x01\x79\x85\x4e\x76\x8a\x79\x22\x73\xba\xc4\x8c\x5c\x81\x72\xf4\xe5\x76\xde\xca\x16\xa7\x8a\xb8\xce\xa1\x34\xd6\xb1\x33\xc7\x1b\x7a\xd2\x90\xb5\x50\xf8\x9a\x42\x18\xc5\x5d\x14\x7c\x43\x41\xc6\x12\x05\x0f\xde\x66\x0d\x12\x0d\xe6\xce\x50\x00\xd5\xb8\x90\xde\x5e\xbc\x7d\x8d\x78\x4f\x19\xc2\x9e\x47\x36\x29\xf1\xd1\x72\x5b\x31\xbb\xeb\x9e\x9e\x89\x90\xf8\x14\xbf\x87\xea\x2a\x1b\xa0\xdc\x7e\xe6\xf5\x67\xa3\x68\x1a\x79\x38\x42\x05\x8d\x51\x6c\x88\x29\xdf\x29\xcd\xca\x32\xec\x6a\xd6\x5e\x77\xd7\xaf\x17\x6e\xb4\x9f\x80\x5a\x82\xc2\x58\x31\x51\x94\x3d\x37\x51\x4b\x08\xdb\xc0\x43\x1b\xfd\x28\x8b\x1a\x69\x31\xe2\x65\x09\x4d\xb7\x16\xaa\x56\x96\xd4\xd6\x3f\xef\x23\x27\x5d\x25\x89\x6a\x8a\x6f\x98\x6e\x16\xe2\x24\xc1\xdb\x9d\x1e\xd0\x94\x84\xf5\x72\xc6\x06\x81\x5e\x69\x12\x1f\xaa\xda\x59\x44\xff\xcc\xc8\x45\x41\x23\x4d\x32\x22\xf1\x40\xee\xf9\x04\xc7\xc1\x79\xec\x59\x74\x49\x2a\xdd\xb0\xf0\x3a\x1d\x52\xcc\xa9\xc6\xad\xe9\x66\xcb\x2f\x24\x22\x09\x0e\x10\xaf\x9e\x84\x98\x3f\x46\x78\x19\x67\xa9\x66\xb6\x2a\x5e\x51\x3c\x2d\xcc\x7d\x55\xac\x72\xf4\x8a\xcf\xe8\xf2\x8c\xe5\xd4\x32\x78\x47\xf1\x5a\xf6\x90\x2d\x02\xd4\x7a\xc9\x52\x8e\xf2\xc0\x69\x3c\x66\xad\x1b\x8d\xd6\x0c\x06\x5c\x27\xdb\x33\x94\xab\x04\xc2\x91\x0f\x56\x87\x78\x14\x2c\xb7\x20\x5a\xf7\x24\x35\xce\xbe\x57\xa5\x3a\xa6\x75\x06\x20\x27\x4a\xa9\x57\x79\x64\x70\xed\x4b\x70\xd0\x9d\x8d\xcb\x94\x86\x33\x10\xc4\x11\x07\x04\xb5\xe7\x92\x0b\x5d\xac\xe3\x2c\x00\xcf\x40\x90\x4f\x57\x2b\x92\x00\x46\x40\xab\x24\x0e\x45\x09\x21\xa7\x53\x84\x7e\xa1\xe9\xaf\xd9\x12\xfd\x1c\xe0\xdb\x18\x74\x0f\xbd\xc5\xc9\x1f\x7e\x7c\x17\x21\x40\x16\x38\x08\xe2\x3b\xe2\x1b\x7a\x01\x6a\x14\xb2\xcb\xd5\x82\x24\xb7\xd4\x1b\x33\x8e\xdc\xeb\x0a\x62\x9c\x7b\x96\x93\x13\xa8\xb5\x5d\x8a\xe0\x32\x53\xec\xa5\x76\xea\x5a\x16\xd6\x52\x0a\xa0\x41\x30\xba\x76\x94\xca\xc2\xaa\xc2\x37\x1d\x7a\x83\x3b\x5b\x93\xf1\x53\x5e\x53\x32\x19\xa5\x34\x60\x60\x40\xd7\x24\x0d\xeb\x39\xfd\x0d\x73\x91\xc3\xb0\x91\x43\x48\x7d\x50\x30\xba\xda\x42\x59\x94\xa3\xba\x9c\xcb\x42\x12\x08\xda\x85\x80\x61\x0e\x91\x02\x8e\xe8\x7f\x44\xbf\x0c\x23\x9b\x25\xc1\x48\x5e\x7e\x7b\xf7\x06\x6d\x62\x0a\xfc\x00\x33\x05\x8e\xf3\x54\xb9\x9e\xca\x84\xf2\xe7\x9c\x06\xb8\x3b\x3d\x6b\x30\xe5\xe9\x58\xe6\x04\x0d\x04\xc3\x05\xde\x9e\x59\x49\xc9\xc0\x65\xce\x4c\x9b\xe5\x3d\x98\xb1\x97\x74\x5f\x9d\x4f\x46\xdd\xd7\xfb\x3c\xa1\x8d\x03\xfd\xdb\xfe\x14\xbc\xae\xd4\x45\x17\x05\xfc\x3d\x45\x17\xe9\x09\x43\x24\xf2\xe2\x2c\xc1\x37\x60\x44\x41\xe3\x32\xc6\xfd\x12\xba\x5c\x00\x28\x8e\x43\x18\x08\xba\x0c\xaa\x6a\x07\xd5\xfb\xaa\x4d\x2b\x5d\x3f\x16\x1d\x52\x42\x00\x4b\xeb\xf9\x8e\x04\x20\xeb\xdb\x3c\x84\x63\xa5\x0c\x68\xe4\xd3\x5b\xea\x67\x80\xc4\x80\x0d\x21\x21\x76\x8a\x40\x62\x5b\x14\x66\x10\xcd\x80\x8f\x4c\xca\x8a\x45\x95\x93\x32\xbc\x3c\x39\x55\xc2\xc8\x3d\x0a\xa3\xa6\x0e\x10\xa8\x5a\x11\xe3\x3d\xe5\xb0\xb8\x6d\x14\xdb\xe6\x8e\x4d\x00\x65\x92\xbe\x81\x6e\x27\xc2\x2f\x92\x49\x0a\x9f\x8d\xd1\xbc\x8c\x44\x72\x20\x04\x68\x92\xe7\xb4\xf2\xf6\x59\x81\x79\x96\x42\xcd\x61\xb0\x72\x72\x0c\xc6\x91\x3f\x29\x82\x69\xbf\x00\x86\x22\x1c\x95\x23\xe4\x86\xaa\x69\x22\xb8\x3d\xf6\xbd\x6a\xaf\x7f\xf7\x13\x02\x38\x97\x81\x9f\x15\x8e\x81\x09\x5c\x50\x0b\x56\xa5\x6e\xe9\x62\xc9\x3d\xf6\xaa\x6c\x6e\xbf\x9d\x32\x45\x79\x3d\x7b\x23\xfb\x8c\x06\x83\x6a\xac\x56\xb6\x5a\xe5\xda\xc4\xcb\x2e\x2f\xc6\xcd\xb9\xe2\xc4\x4c\xfe\xc9\x3e\x26\x70\xe1\x3a\x8e\xdd\xfa\x93\x3c\xdd\x36\x66\x88\x95\x04\x41\x48\x43\xf2\x3e\xa7\xd1\x99\x2e\xd4\xb8\xd6\x2a\xdb\x55\x42\x80\x5f\xdf\xbf\xbf\x42\x21\x40\x38\x70\xf9\x0d\x8b\xc2\xd9\xc0\x8d\xa1\xec\x09\x81\x76\x49\xa3\x81\x38\xe8\x88\xe2\x7c\x39\x3b\x24\x09\x43\xce\x10\x89\x57\x6a\x96\x48\x37\x54\xb5\x97\x0f\x52\x75\x43\x9a\xa8\x51\x70\x06\x0e\x22\xc4\xc9\x76\x54\xfc\xbd\x4c\x28\x81\x88\x35\xa7\x54\xaa\x45\x35\xf6\x8f\x16\xfc\x57\x1c\x7c\x3f\x22\xba\x37\x18\x5a\xf1\xce\x36\xa5\xd6\xa4\x59\x31\x76\xe1\xbb\x48\xfb\x14\x01\x27\xdd\xa5\x5c\xba\x64\xaf\x49\x6f\x1b\x84\xdb\x33\xc5\xdd\x22\x16\x4d\x9a\xbb\xc9\x96\x26\xf9\x3f\x88\xad\x82\x8e\x2b\xb6\xb4\x59\xf0\x16\x92\xbb\xf2\x66\x9a\xba\x5c\x78\x0b\x49\xc5\x0a\x36\x26\xb1\xb2\xee\xd2\x42\x4b\x59\x7b\x69\x52\xf3\x39\x0e\xf1\x70\x4a\x8c\xda\xb9\x8c\xe3\x80\xe0\xa8\xa9\x9e\x2b\x9c\x05\xa9\x84\xa6\x15\x46\xd5\xb4\x7d\x1b\xa7\x52\xea\x5e\xd0\x32\xc6\x48\x02\xf8\xbb\x02\x42\x47\xe4\x34\x0a\xc2\xbd\x81\xd0\x0d\xb1\xcc\x08\xee\x7c\xb4\x5e\xf9\x33\x47\x74\xe4\xe5\xd4\xe1\x84\x7c\x12\xc0\xdc\x72\x42\x2a\xde\x34\xa3\x81\xe1\xb4\xd6\x04\x2b\xd3\x65\x98\xa0\x70\xea\xad\x1d\x51\x72\x64\xb7\xb4\x93\x4e\xbb\x9a\x67\x9d\x9c\xc8\xeb\x56\x61\x2c\x4f\x29\x31\x61\xbb\x09\x05\x4b\x9e\xf0\x44\x04\x8e\xb6\xe8\x16\x07\xd4\xcf\x11\x26\x83\x60\x23\x83\x32\xb1\x2f\xc2\xa6\x93\xc2\xdc\xd4\xb3\x12\x21\x95\xa7\xec\x0f\x6e\x67\xfd\xd3\x0f\x2f\x9e\xff\xe5\xfa\xd3\xff\x3d\x3c\x7b\xf2\xdf\x8f\x4f\x8b\xf6\x9f\x3d\xe9\x67\xc1\xff\x89\x83\x8c\x18\xf2\x1c\x7b\x30\x2b\x51\x9c\x36\x40\xa8\x7e\x84\x2c\x65\xd4\x29\x25\x6d\x37\xfa\x77\x64\xd7\x95\x2e\xf5\xcb\xe5\x59\x53\xc1\x38\x22\x97\x2b\x29\x86\xe8\x31\x3a\xda\x81\xb1\xa8\xcf\xb7\x00\xbd\x23\x62\x6d\xc9\xd3\x2c\x89\x5c\x6b\x59\x1f\x1e\x14\xd5\xa7\x53\xd9\xc4\xbe\x23\xeb\x6a\xdb\x93\x54\x53\x95\x76\x53\x62\x2d\x52\x93\x93\x5f\x4a\x93\x3d\x28\xad\x68\x40\x16\x3a\x6a\xb5\x5f\xd7\x46\xbb\x6d\x6d\x21\xcb\xc2\x86\x48\x41\x89\xd5\x5b\x48\x55\xa5\x5b\x26\xef\x91\x61\x15\x49\x89\x55\xb9\x39\xcf\xa4\xe5\x4d\xcc\x5a\x9a\x77\x06\xf8\xf4\xd3\x4c\x90\xb4\x9e\x5f\xa9\x9c\x53\x91\x98\xd2\x85\x73\xca\x0e\x38\xf1\x54\x53\x92\x9b\x71\xb1\xa2\xde\x7c\x4a\xa3\x94\xdc\xa8\x8f\x75\xe8\x1c\x95\x29\x86\xce\x09\x51\xa5\xc4\x7a\x5b\x08\x5d\xc2\xc2\x04\x35\x12\x1a\x52\xbe\xca\xc0\xf2\x04\x85\x96\x9e\x17\x07\x01\x0c\x25\x54\xf8\x59\xcb\x93\x69\x85\xbb\x51\xcb\x80\x22\xcb\x60\xc5\x82\x64\x59\x58\x4b\x29\xc4\xf7\x34\xcc\x42\x3b\x4a\x65\x61\x83\x01\xf1\x82\x8c\x81\x50\xde\xf6\x21\xa9\xd4\xd2\x73\x09\xe5\xed\xb9\x2c\x0a\x77\x70\xd9\x87\xa4\x52\xcb\x24\xcb\x37\x24\xba\x49\x2d\xf1\xef\xae\xb8\xa9\xcf\xbd\xa8\x55\xc5\x4d\xb8\xbc\xd8\x39\x69\xb7\x14\x25\x0a\x9b\x7a\x79\x61\x3f\x55\xaa\xd2\xa6\x3e\xf6\xa1\x55\x96\xd6\xd2\x92\x33\x86\x16\xe4\xea\x15\xf4\xba\x12\x59\xeb\x47\x64\xd4\x09\x98\x79\x14\x3c\xe5\xa5\x12\x06\x1b\xfa\xb8\x2b\x6f\x98\xf9\xfd\x61\x90\xe2\x99\x1f\xc9\xe9\x36\x8b\xb7\xec\x4e\x85\xe0\xa9\x70\x54\x5b\x1e\x3a\x25\x62\x25\xfc\x0e\x82\x2b\x74\xff\x9c\x67\x3d\x45\x64\xd5\xbd\x6b\x86\xe7\x8d\x35\x65\x8c\xbb\x0f\x97\xb1\xbf\xbd\xaa\xd6\xf5\xc6\x6d\x7c\xa8\xbb\x16\x69\xdf\x9f\x8c\x1b\xaf\x8f\x31\x6d\xe3\x2a\xbb\x9d\xa7\xd6\x35\xc9\xed\x2a\x58\xe7\xcb\xf7\x94\xc7\xc5\x7c\x8f\x9b\xd8\x3d\x43\x21\x8a\x2e\xd0\x25\x2f\x9d\xb1\x71\xfb\xdb\x1c\xef\x18\xd9\x31\x6e\x40\x11\x63\x04\x76\xce\x09\x83\x95\x2b\x12\xc2\x41\xec\x61\xbd\xd0\x6c\xa0\x18\xd7\xe5\x6e\xbc\x54\x53\xe0\x3e\xb9\x52\x13\xdb\x77\x6b\x22\x12\x20\x71\x82\x20\x76\xcf\xbf\x6c\xa8\xd8\xe6\x83\x55\xb6\xc7\x4b\xe4\x09\x2c\x1c\x9c\x0e\xc8\xc4\x6a\xc3\x39\xbb\x38\x6d\xc4\xb6\x8a\x1c\xb7\x57\x16\x62\x91\x2d\x17\x4d\x46\x8e\x2d\xec\xe9\x9c\xeb\x9f\xa9\x06\x1c\xcf\x44\x93\x03\x3d\xfe\x4f\x3f\xd5\x26\xa3\x3a\xd4\xa8\x1e\x3e\x36\x35\x04\xa1\x86\x90\x75\x8a\x4d\xa7\xd8\x74\x8a\x4d\x5b\x7b\x3d\xc5\xa6\x5f\x68\x6c\xfa\x4d\xfd\xbf\x25\x4e\x02\xde\x93\xed\x04\x93\x26\x98\x54\x7b\x2a\x74\x62\x42\x49\xfb\x43\x49\x82\x99\xd7\xe1\x26\xdd\x36\x57\x15\xa5\x96\x6d\x76\xbf\xb4\xb1\x25\x9a\x61\x88\xc1\x94\xe2\x49\x19\x5c\x53\xdb\xe5\xb6\x60\x38\x0a\xb6\x5c\x6f\x45\xc2\x86\xaf\x8a\x73\xa6\x78\xce\x26\x33\x7d\x33\x31\x21\xbc\xa3\x44\x78\xff\x82\x01\x7c\xcb\xad\xfe\x04\xf5\xd0\x04\xf5\x26\xa8\x37\x41\x3d\xd4\x84\x7a\xdc\xe4\x9d\xe3\x14\x4f\x68\x6f\x42\x7b\xb5\xa7\xa5\x5a\x4c\x80\x6f\x02\x7c\x3a\xde\x3f\x0f\xc0\xd7\x78\xc8\xf7\x69\x4d\x20\x10\x4d\x20\x70\x02\x81\x5d\xbd\x9e\x40\xe0\xd7\x04\x02\xf9\x27\x2c\x9f\x27\x00\x34\x7d\xb6\x59\x3c\x2d\x1e\x75\x6f\x9f\x1c\x04\x18\xb5\x4e\x4d\xfa\xd6\xb1\xd6\xb4\xa8\xe1\x1c\x62\x1e\x39\x8c\xe4\x8a\x35\x41\xc8\x69\x65\xb5\xfa\xf7\x75\x40\xae\x09\x69\xa1\x09\x69\x4d\x48\x6b\x42\x5a\xa8\x89\xb4\xa2\x38\xfa\xeb\x21\x36\xa9\xea\x3f\x1e\x19\xf4\x75\x9a\x71\xd3\x9c\x4e\x74\x16\xf4\x5a\x32\x8e\x03\x29\x9a\x96\xab\x07\x92\x33\xa0\x61\x65\x64\xaf\x1b\x18\x5a\x33\xa4\x83\x04\x2e\xef\x62\x1e\xd8\x09\x45\xd1\x3a\xd8\x57\x76\x64\xda\x7e\x4d\x7b\x06\x80\x2b\x47\x8c\x94\xd5\x8f\xe1\x04\x04\x83\xa5\x13\xd0\x73\x3a\x07\x3d\xc6\x4b\x09\x17\x5c\x7f\xe7\x3e\x1c\xb8\x68\x8f\x5d\x2d\xfb\x67\x79\xb4\xfb\x7c\xd7\x9d\xb9\x74\x1a\xad\x35\xbc\x1e\xd1\xa0\xe6\x2b\xd0\x5e\x70\x67\x50\x93\x6d\x98\xa8\xd3\x66\x0f\x68\xb1\xeb\x73\x8e\x0e\x20\x36\xa4\x45\x17\x68\x6d\x40\xbb\x4e\x20\xdd\x90\xfe\xba\xc0\x7d\xa3\xfa\x3b\x0a\x1c\xda\xb6\x2c\xf9\x97\x98\x89\x38\xe4\xa2\x88\x99\x86\x01\x49\x07\x2d\x9f\xe7\xf3\xe9\xc5\x20\xf0\x39\x40\xe6\xa3\x10\xea\x3e\x25\xbd\xef\x86\xdb\x05\x6d\x81\x81\x07\x08\xbb\x13\x28\x83\xc0\xcd\xc7\x28\x1c\x42\xea\x07\x69\xbd\x5d\xf4\xa6\xb4\xdf\x18\x06\x72\xaf\x7f\x26\x67\x3e\x6c\x23\x94\x21\x96\xcc\x18\xc6\x74\x7e\xfa\xce\xcb\x44\xdb\xc3\x1e\xa1\xd0\xcc\xa8\xca\x95\x6b\xbf\x9a\x99\xd3\x0a\x4d\x3c\xe8\x01\x95\x26\x15\x36\x06\x4f\xd5\x02\x28\x8b\x94\xd3\xa3\x89\x51\xb7\xd0\x29\x1b\xb5\x1f\x94\x97\xfa\xb3\xfd\x7a\x32\x28\x0f\xd6\xa8\xa1\xc3\x41\xa0\xa2\xb6\x96\xb3\x09\x4d\x9d\x33\x1d\x5a\x68\xdb\x29\x3d\x77\x86\x98\xc1\x1c\x21\xe5\x7d\xea\x9e\x7b\x7d\x38\xeb\x27\x50\x9f\x72\x6c\x0e\x62\xc2\x69\x9c\x0c\x89\x4e\x12\x88\xf9\x2f\xa3\xc0\x78\x30\xe3\xe0\x23\xd8\xee\x43\xe5\x84\x54\xbd\x0c\x78\x41\x03\x26\x74\x7f\x24\x62\xb1\x02\x50\x17\xe8\x2e\x66\xb4\xfe\xba\xb0\x76\x0c\xcc\x57\x1a\x46\x4f\xe1\xf0\xd1\x84\xc3\x8f\x82\x6c\xdc\x2c\x47\xd9\x6d\xd3\xd9\x9b\x89\x3a\x5e\x03\xd3\x5c\xc1\xb2\xb0\x32\xa3\x6e\x6b\x98\xce\x14\x9a\x56\x17\xbb\x28\x4d\xab\x8b\xd3\xea\xe2\xb4\xba\xf8\x78\xab\x8b\x8f\x00\x03\x25\x9f\xa4\xbb\x0a\x71\xec\xc5\x83\x25\xcd\x77\x39\x86\xe1\x57\x3d\xcc\x94\xfe\x76\x5c\x44\xa8\xa3\x31\xdc\x5f\xaa\x4e\x51\x89\x4b\xf5\xce\xc2\xea\xd6\x08\xd3\x59\xf9\xb2\xa4\xe5\xd0\x69\x3f\x1e\xdf\x62\xfb\x92\xb6\x0b\xe2\x04\xdd\x0d\xee\xbc\x80\x4e\x6f\x87\x40\x09\xe8\xfd\x90\x9a\xa0\xf2\x09\x5d\x66\xea\x81\xcc\xa3\x41\xe0\x5d\x82\x37\x1b\x57\x47\x90\x1f\xcb\x5c\xe5\x17\x7a\xba\xd2\xa0\x3e\x57\x86\xb9\xd6\xb6\x91\x67\xc9\x3a\x03\xf8\xc7\x32\xae\x1d\xd7\xc9\x0e\xb7\x75\xba\xf3\x75\xad\x32\x58\x4b\xcc\xa8\x77\x96\xa5\x6b\x7e\x37\x44\xbe\x81\x74\xa1\x1c\xa7\xdf\x48\x6b\x59\x11\xc6\x1b\xfa\x77\xb2\x75\x43\x2b\xc6\xc0\xe0\xcb\x0b\x08\xcc\xa8\x47\x53\x97\x34\xaf\x30\x63\x77\x71\xe2\xbb\xa4\x79\xb6\xe1\x7c\x3a\x14\x65\x41\xd6\xf3\x08\x63\x3f\xc5\x3e\xd1\x52\xad\xfe\xbe\xd6\x6a\x5e\xdb\x38\xef\xd7\xd2\x3c\xc6\xe9\xb8\xa2\xb7\x2e\xb7\x33\x1f\x9f\x29\x69\xcc\xaf\x03\x8c\x61\x03\x45\x34\xb6\xb4\x1d\x78\x84\xf3\xee\x77\x0f\xf1\x50\xcf\xd5\x6f\x67\x7e\xeb\x79\x65\x8d\x5c\x88\xf3\xe3\x39\x8e\x4f\x37\x0d\xf6\xfa\xb0\x3a\xba\x0a\xe2\x3b\xe9\xde\x02\xe0\x29\x4e\x8a\x3b\x62\x7f\xeb\x73\xd7\x9c\x1b\x8d\xcd\x85\x62\x91\x04\xe3\x7c\x8f\x6e\x8d\x16\xc2\xef\x6e\x8f\x79\xd0\x6b\xdb\xcb\x4d\x44\x1f\x16\x79\x0d\x2d\x35\x45\xca\x3d\x7a\x62\x71\x9f\xf0\xe7\x3f\x2b\x14\xc4\xf1\xb8\xb3\x22\x8d\xff\x20\x5f\xfe\x6c\xd8\x14\x42\x3f\xf4\x6c\xa8\xa4\x3b\xcd\x02\x79\x16\xe8\x30\xf2\x34\x11\x8a\x96\xf7\x38\x11\xf0\x4e\xee\xd3\x5c\x38\x96\xb9\xa0\x06\x76\x47\x86\x94\xbe\xbe\x69\x52\x0d\xc9\x17\x86\x9f\xa6\x49\x88\xf4\x93\x70\xd1\x1c\x45\x07\x0b\x0f\x72\x97\xe5\x56\xe5\xfb\x44\x1d\x2e\xc9\x54\xd7\x36\x2b\xf2\xed\x58\x87\x69\x5c\xeb\xd7\xcd\x92\xe6\x8b\xe8\x1d\x0d\x14\x11\xe2\x13\x1f\xa5\xb1\x38\xcf\x06\xe1\xe2\x8e\xbe\xfc\xee\xd5\x20\xd0\x5e\x29\x51\xf2\x26\x1b\x31\x4d\xd7\x07\xa7\x3b\x95\xbb\xe4\x25\x11\x59\x93\x31\xdc\xe3\xa6\x4d\xc3\x59\xaf\x83\xd5\xee\x5b\x1d\x24\xfc\x34\xc1\x11\x03\x9e\xf8\x85\x1e\x69\xec\xc5\x41\xf9\x6d\xba\xb8\xc2\xbf\x4d\x9c\xc6\xd9\xaf\x33\x8f\x62\x43\x92\x6c\x21\xf8\x13\x26\x3f\xba\x53\x7e\xd7\x0c\x5e\x6f\xd9\xb4\x6c\xc7\x30\x31\xaf\xb2\x3e\xf3\xd8\xad\x74\x3d\x88\xfc\x33\x95\x7f\x6e\xe8\x46\x77\x99\xf8\x6e\xe9\x48\x90\x6b\xe5\x72\x77\x9e\xcf\xa1\xd8\xad\x3d\x08\xeb\x07\x09\xd9\xf1\xdf\xdc\xda\xe7\x68\x5b\x5f\x49\x5e\xef\x24\x9c\x6e\xe5\xdb\x35\xd5\xdc\x62\xe3\x6c\xfb\x5e\x65\xc0\x75\x7b\x03\x5c\x7e\xc1\x56\x35\xa4\xec\xc8\x71\xf6\xd5\x5a\xd9\x44\xcb\xee\x1f\xf7\x5f\xaa\x55\xfd\x52\xf6\xf0\x38\xfb\x3a\x4d\xed\x97\xd3\xb6\xf4\x1b\x8a\x6a\xe3\xa5\x6c\xfd\x71\xff\x95\x4e\x4d\x8a\x7b\x6d\x4d\xfe\x2a\x67\x87\x15\x9a\x1b\x92\x9c\x7d\x6d\x56\x13\xa3\xb2\x77\x72\x9f\x52\xdc\x67\x63\x7a\x21\xea\xf7\x3c\x39\xfd\x92\xac\x9a\x08\x91\x3b\xe5\x8f\x9a\x0a\x2f\x83\x21\x0b\x04\xad\x8f\x4a\x05\x4f\x8a\xbf\x1a\xb5\x49\xa6\xd7\xc7\xe2\x32\x4c\xff\x86\xff\xef\xe1\x7f\x01\x00\x00\xff\xff\x48\xaa\xbd\x2f\x28\x9c\x00\x00")

func v2SchemaJSONBytes() ([]byte, error) {
	return bindataRead(
		_v2SchemaJSON,
		"v2/schema.json",
	)
}

func v2SchemaJSON() (*asset, error) {
	bytes, err := v2SchemaJSONBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v2/schema.json", size: 39976, mode: os.FileMode(420), modTime: time.Unix(1441639065, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"jsonschema-draft-04.json": jsonschemaDraft04JSON,
	"v2/schema.json": v2SchemaJSON,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"jsonschema-draft-04.json": &bintree{jsonschemaDraft04JSON, map[string]*bintree{
	}},
	"v2": &bintree{nil, map[string]*bintree{
		"schema.json": &bintree{v2SchemaJSON, map[string]*bintree{
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

