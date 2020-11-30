// Code generated for package internal by go-bindata DO NOT EDIT. (@generated)
// sources:
// _template/header.tmpl
// _template/main.tmpl
// _template/registry.tmpl
package internal

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

var __templateHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8e\x41\x6a\xc3\x30\x10\x45\xf7\x73\x8a\x8f\xe8\x22\x86\x56\x3e\x41\x37\x69\xba\xe8\x26\x29\x34\xfb\x22\x5b\x23\x45\xb8\x1e\x19\x75\x5c\x5a\x84\xee\x5e\x4c\x42\x96\x33\x1f\xde\x7b\xb5\xc2\x73\x48\xc2\x30\x17\x76\x9e\x8b\x41\x6b\xd4\xf7\x78\xc9\x9e\x11\x59\xb8\x38\x65\x8f\xe1\x0f\x31\xab\x8f\x2c\x8f\x38\x9c\x70\x3c\x9d\xf1\x7a\x78\x3b\x5b\xa2\xc5\x8d\x93\x8b\x8c\x5a\xf1\x60\xdf\x6f\x47\x6b\x44\x69\x5e\x72\x51\xec\xc8\x8c\x59\x94\x7f\xd5\x90\x09\xb3\x1a\x22\x13\x93\x5e\xd6\xc1\x8e\x79\xee\x37\x6a\xaf\xbe\x1f\x92\x18\xea\x68\x53\x1f\xf3\x53\x5e\xae\x59\x49\x53\x16\x84\x5c\x30\x31\x2f\x49\x22\xae\xd4\x6f\x4b\x3f\xae\xe0\x13\xcf\x18\x92\xd8\xfd\x1a\x02\x97\xda\xee\xcf\x9b\xd1\xee\xdd\x38\xc5\x92\x57\xf1\xbb\xee\x3e\x86\x59\xed\x87\x96\x24\x91\xcb\x4e\xd2\x57\x47\x54\x2b\x58\xfc\x96\xfd\x1f\x00\x00\xff\xff\xd8\x6f\x86\x9d\x11\x01\x00\x00")

func _templateHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		__templateHeaderTmpl,
		"_template/header.tmpl",
	)
}

func _templateHeaderTmpl() (*asset, error) {
	bytes, err := _templateHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_template/header.tmpl", size: 273, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __templateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x5a\x5b\x6f\xd3\xca\x16\x7e\xcf\xaf\x58\x54\x85\x63\x1f\x52\x87\xe7\x96\x3e\x9c\xd2\x72\x88\x84\x10\xa2\xb0\x5f\x10\xaa\x26\xf6\x72\x3a\xaa\x33\x8e\xc6\xce\x85\x6d\xfc\xdf\xb7\xe6\x66\x8f\xed\x71\xec\x46\xdd\x6c\xd8\x79\x4a\xe6\xb2\x66\x5d\xbe\x75\x9b\x49\x51\x40\x84\x31\x65\x08\x27\x2b\x42\xd9\x09\x94\xe5\xa4\x28\xe0\x74\xfd\xb0\x84\xf3\x4b\x38\x0d\x3e\x92\xf0\x81\x2c\x51\x8f\xe7\xb8\x5a\x27\x24\x47\x38\xb9\x47\x12\x21\x3f\x81\x53\x31\x23\xa6\x38\x61\x4b\x84\xd3\x4c\xee\xba\xcd\xf9\x26\xcc\x33\x31\x37\x9b\x81\x20\x98\x05\x6f\xd2\xd5\x0a\x59\x2e\xc6\xf2\xef\x6b\xd4\xa3\x1f\xc8\x4a\x10\x87\x4c\xee\x80\x62\x52\x14\x67\x86\x56\x2c\x59\xc8\x82\xb7\x14\x93\x48\x12\x03\x00\xd0\x04\x63\x9b\xa0\xd8\x44\x63\x20\x2c\x02\x4f\xce\xb0\x88\xe6\x34\x65\x24\xf1\xc1\x63\x69\x0e\xcd\xc1\xab\x34\x4d\x7c\xc5\x9c\xe0\xef\x4b\x86\x70\x8b\xb9\xa2\x6a\x18\x12\xb4\xfe\xdf\x1e\xbc\xc7\x64\x8d\x3c\x0b\xe4\x81\xc8\x22\xc3\x93\x3e\xff\x34\x0e\xae\xd3\xcd\x22\xc1\xdb\x84\x86\x58\x4f\xda\x24\xbe\x7e\xfb\xfa\x4d\x8d\x7c\x16\x5a\xb0\x08\x60\x92\xa1\xa6\x72\x70\xff\x81\xdd\xce\x1d\x3d\xeb\x15\xf7\xd6\xd7\x72\x52\x1b\x4b\xef\x15\x9b\xe6\xd7\x40\x33\xf8\xfc\x1e\xa4\xd5\x68\x04\x69\xdc\x5c\x14\x4c\xc2\x94\x65\xb9\x73\xe7\x25\xbc\xda\xab\xf1\x77\xb8\x9f\x5f\x43\x59\x5e\xc8\x53\x6e\x58\x98\x46\x08\x74\xb5\x4e\x50\x18\x31\x83\x05\x65\x81\x1a\xe5\xc1\x24\xde\xb0\x10\x3c\xb5\xf1\x13\x86\x48\xb7\xc8\x85\x28\xff\x6d\x9c\xe1\x6b\x32\x7a\xe1\xd5\x26\xfe\x1f\x5f\xca\x65\x82\xd8\xd5\x26\x8e\x91\xfb\x80\x9c\xa7\x1c\x0a\x29\x38\x35\xac\xdb\x44\x2f\x2f\x81\xd1\x44\xaf\x10\x1f\x8e\xf9\x86\x33\x88\x57\x79\x70\x23\x36\xc7\xde\x49\x48\xd8\x7f\x72\x40\xc5\xb5\x26\x41\x76\x5a\xa7\x40\x32\x41\xe1\xc4\x97\x14\x94\x03\x59\xec\x04\x1f\x37\xf9\xfc\xda\x73\xa8\xc7\x1f\x46\xbb\x46\x96\x1b\xc3\x5d\xf8\x59\x0b\xcc\xa4\x4b\xe6\xc0\x78\x50\xb5\x5a\x9e\x2a\x66\xde\x91\xcc\xeb\xcc\xce\x59\x84\x7b\xa9\xf1\xa2\x0d\xa0\xe6\xf9\x7f\x60\x98\xa7\xdc\xc6\x61\x4b\x11\x6a\xc1\x3b\x19\x3d\xbc\x04\x59\xd7\xc8\x41\x03\xbd\xbe\x52\x6a\x2c\x4c\x58\xe8\x43\x34\x4c\xa0\x2c\x69\x24\xd0\xa5\xa1\x7f\x27\xbe\x4a\xa6\xa6\x1d\x87\xac\xf8\xe2\xe9\xae\x28\xc4\x86\xb2\xdc\x16\x05\xb2\xa8\x2c\x85\xde\x95\x0d\x06\x78\xb1\x84\x77\xd2\xae\x00\x34\x46\x6e\x9e\xee\xb4\x6c\x46\xbe\xbb\x29\x6c\x6b\x5e\x78\xba\x1b\x52\x76\xad\x07\xeb\x64\x33\x39\x67\x39\xf2\x98\xd4\x61\xc4\x7c\x68\x0c\xdb\x2e\xe6\x7b\x70\xbf\x61\x64\x91\x20\xe4\x69\x1f\xf6\xcf\x21\x96\xc8\x51\x8a\xfa\x44\x76\x46\x57\xa8\x1c\x1b\x76\x34\xbf\x07\x2a\xf1\xf3\x3c\x12\x91\x44\x78\xca\x14\x68\xb4\xf7\x1b\x87\x37\x65\x90\x86\x99\x58\x2c\x23\xe7\x42\x37\xdb\xc0\xed\xf2\xfe\x85\x5c\xf1\xec\x9f\x11\xeb\x1c\x9e\xef\x94\x4c\x53\xc1\x46\x2d\x98\x3b\x3c\xf7\x40\x44\x9d\xf5\x56\x84\xbe\xb2\xf4\xb6\xfe\x61\xe3\xf7\x61\xaf\x13\xe1\xfb\xf9\xe8\x45\xd2\x01\x14\x1d\x88\x26\x46\x45\x63\xe2\xe9\xf1\x06\xa0\xcd\x58\xdb\x05\x4c\x0d\x96\x01\x46\x1f\x09\xa5\xa7\x94\x42\x03\xa6\xc2\x4a\x7f\x1a\x3f\x88\x91\xa1\xd8\xd9\x87\x20\x27\xa0\x1c\xa9\xa3\x77\x47\xb7\x78\xb0\x34\xc4\x68\x32\x91\x35\xe1\x81\xd4\x56\xc5\xf2\xd6\xa9\xb3\x59\xb7\x12\xcb\x30\xcf\x60\x4b\x92\x0d\x9a\xd2\xa3\x9e\x0b\xad\xed\x52\xd5\xe3\x6b\x87\xf6\x39\x9e\x3a\xa2\xe2\xac\x59\xc9\xe9\xb2\xcd\x51\xa3\xa9\x71\xa9\x85\x66\xa9\xe5\xeb\x6a\xb6\x23\xa8\x9d\xba\x45\x34\x56\xe7\xb6\x42\xc3\xb8\x6c\x7d\x8b\xf9\x81\x6c\xad\x8c\xa8\x98\x3e\x8e\xfe\x17\x96\x8d\x38\x61\xe2\x46\xee\x93\x89\x30\x18\x71\x94\x0e\x5b\xc5\xec\x60\xe9\x34\x9b\x75\x2b\x7c\x85\xe1\xf1\x78\x13\x7d\x82\x20\xb4\x48\xd3\x04\x09\x83\xdd\x3d\x0d\xef\x45\x90\xca\xf9\x46\x22\x45\xad\xda\x91\x4c\xe0\x78\x3c\x3a\xdb\x8c\x79\x3e\x3c\x01\x40\x1b\xf8\x9c\x42\xfa\x20\xf9\xf6\xeb\xea\xf8\xd9\x93\x97\x8a\x56\x60\x90\xfc\x4f\x21\x26\x49\x86\x56\x7c\xd1\xb3\x03\x46\x9e\x4a\x8d\x1a\xbb\x6a\x2b\x4f\xdc\xc1\x48\x18\xe4\x1a\x5d\x1d\x86\x1a\x7d\x44\x87\xa1\x36\xfc\xb4\x0e\x23\x42\x77\x87\x91\xa7\x9d\xac\xd7\xce\x73\x75\x9e\x78\x93\xb2\x6c\xb3\xc2\x9e\x9e\xe3\x98\xf4\xd6\xc3\x56\x37\x8d\x8d\xed\x67\xfa\x03\xe2\x38\x7f\x7f\x3a\x98\x1e\xa8\xd0\x1c\x9c\x76\x6b\xe9\xbf\xc5\x61\x0c\x3f\xad\xe3\x8a\x76\x8d\xec\xe6\xb7\x5b\x14\x8a\x8f\xba\xb1\x79\x8f\x6c\xda\x87\x9b\x46\x8f\xe2\xb7\x05\xed\xad\xaf\x8f\x47\xcf\xb8\xe2\xc8\x7c\x9a\xe2\x88\x8e\x89\x46\x7b\x21\xc7\xab\x0b\xf9\xed\x75\x2d\xa2\x1c\x78\xf9\x12\x8a\x49\x53\x7f\x83\xd5\x73\x25\x2f\x63\x47\xeb\x6a\x8c\xbe\x7e\x96\xce\xba\x7a\x13\x9f\x2d\xe1\xb2\xbf\x74\xdd\x25\x75\x74\x2c\x34\xa1\x30\xaa\x55\x5d\x0f\xbc\xae\xf4\x74\x51\x7d\x93\x5a\x6f\x2b\xbd\x05\x55\x18\xd7\xa9\x2a\x56\x65\xce\xd0\x56\x50\xd1\x58\x31\x5d\x6d\x6b\x56\xc3\x75\x1d\xff\x1b\x19\xa5\x75\xfd\xe7\x68\xec\x6d\xd3\x55\x75\xc0\x01\xd3\x59\x2d\xb3\x58\x1d\xb8\xf3\xd8\xa1\xb6\xf9\xd7\xd0\xc8\x20\x24\xda\x8e\xd9\xec\x90\x7e\x3b\x10\x74\x1d\x05\x1e\x15\xb8\x24\xfb\xe9\x0e\x2e\x81\xac\xd7\xc8\x22\x8f\xa7\xbb\xa9\x52\xd9\x98\xc8\x30\x9c\x79\x35\xd9\xc1\x6a\x8d\xa7\xbb\xee\x81\xae\x3c\xfb\xf4\x47\xf7\x48\xdb\xa3\xdf\xb2\x93\x75\x1f\x79\xa1\xf6\x54\x21\xea\xd7\xcb\xb1\x63\xfb\xae\x8e\x02\x0f\x86\xb1\xf1\x17\x34\xc7\x05\xad\x9f\xab\xa7\x3e\x48\x1f\x1f\xa5\xfe\x25\x38\xe8\x56\xa7\xee\x0b\xc9\x51\x57\x49\x12\x51\x59\xd3\x07\x67\x33\x90\x0f\x4f\xf2\xe1\xd0\xea\xf3\xaa\xc1\x94\xd7\x6f\x55\xf6\xce\xfe\xee\xaf\xd5\xfc\x55\x94\x3c\xdf\x41\x06\x0a\xc3\xe6\x8b\x2e\x25\x79\xc5\x65\xb5\xa3\x37\x2c\xdb\x70\xca\x96\x40\x0d\x85\x0c\x28\x83\x30\x5d\xad\x69\x82\x67\x39\x5d\xa1\x79\xe4\xb0\x5f\xd6\x44\xba\xf7\x26\x77\xf6\xeb\x18\x5c\x9a\xf3\xf4\xaa\xa2\xd4\x0b\x74\x73\xeb\x5a\xe0\xd6\xe0\x9d\x4b\x2a\xf7\x6e\x2d\x8a\x6f\x8b\x55\xbf\xfa\xc6\xf2\xd5\x77\x5e\x8b\x56\x3f\xfc\xda\x57\x2b\x6b\x8e\x99\x34\x51\x05\x43\xd3\xdf\x2e\x91\x21\xa7\xa1\x7c\x63\x0c\xf4\xb3\xec\xcd\x9e\x08\xa3\x9e\x8b\xef\xb0\x74\x06\x58\xe3\x40\x8b\x4d\xec\xcb\x65\x6d\xdf\x11\x63\xe2\xb3\x26\x8c\x86\x9e\x84\xb9\x18\x92\xdc\x41\xb6\xa3\x79\x78\xaf\x1e\x5d\x96\x81\x27\xce\x56\x77\x76\x46\xac\x50\xb6\xb0\xb2\x5b\x33\x88\xca\xe0\x4c\xc9\x06\x21\xc9\x50\x5d\x17\x84\x41\xe5\x4d\x5a\xe8\xd0\x92\xcd\x68\xcc\xec\x8b\x30\x26\x9b\x24\x3f\xd7\x3c\x6d\x0d\x47\xd5\xa3\x78\xad\xb1\x0a\x2b\x3a\x0a\x58\x28\xa8\x7e\x6b\xa3\xcb\xdf\x6d\xbc\x56\x94\x26\xf6\x9d\x48\x43\x75\xad\x0b\x12\xc2\xbf\x43\x84\x67\x19\x72\x4a\x12\xfa\x27\x11\x2d\x6a\x85\xcb\x3a\x30\x2b\x07\x72\x91\x13\xa6\x68\xde\x8e\x78\xad\xec\x2c\x6f\x4b\xaa\x2b\xa7\xa8\xb2\xeb\x62\x13\x07\x1f\x11\x1f\xe6\xd7\x3a\x1e\xf6\xc5\xc1\x3a\x36\xc8\xbd\x56\x3c\xd1\x06\xa5\x91\xf5\x72\x76\xc8\x94\x3a\xd8\x48\x4b\x36\x0c\xa9\xae\x4a\xce\xab\x23\x8d\xee\x84\x03\xb7\xed\x1b\x54\xab\xb6\x26\xd2\x87\x96\xef\x74\xb3\xde\xd6\xe4\x35\x81\xda\xe1\x77\x2b\x29\xe7\x60\xb0\x8f\x83\x03\x11\xbd\x6c\xeb\xee\xc5\x76\x2a\x63\xab\x23\x32\x1b\x78\x3a\xd5\xfd\x68\x36\x04\x10\x3e\xe0\xee\x0b\xc3\xfd\x1a\xc3\x1c\xa3\xf9\xb5\x47\x23\xdf\x5c\x17\xd5\x7f\x39\xa8\x11\xb4\x48\xf7\x98\x41\x7e\xdf\x76\x86\x35\x4f\xb7\x54\x5a\x80\xe8\xff\x61\x04\xb6\xcf\xe8\xed\x57\xe9\xbe\xfe\x2f\x09\x54\xff\x85\x30\xc4\xfb\xbd\xa2\xe7\xa6\xd0\x02\x7f\x7d\x82\x49\x20\x0b\xe5\xff\xf6\x54\x75\x63\xd8\xf1\x83\xd6\x2d\xe1\xe2\x91\xef\x64\x0d\x25\x5b\xc2\x3a\x6e\x06\xb7\xc3\xb1\x72\x84\x7f\x1d\x64\x43\x18\x29\x52\x59\xdf\xfd\x96\xb5\x68\x96\x38\xa2\x44\x70\x24\xf6\x81\x7f\x82\x3c\x5e\xf9\xfa\x49\x6f\xb4\xf2\x7f\xfc\xe8\x72\x7a\xec\xc3\x9f\xf5\xdf\xa1\xee\x1b\xa5\x26\xd2\x3a\x2c\xa8\xf9\xf5\x27\x65\x2b\xb7\xea\xaf\x7f\x05\x00\x00\xff\xff\xac\xc9\xa7\xcc\xa3\x25\x00\x00")

func _templateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		__templateMainTmpl,
		"_template/main.tmpl",
	)
}

func _templateMainTmpl() (*asset, error) {
	bytes, err := _templateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_template/main.tmpl", size: 9635, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __templateRegistryTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\x31\x4b\x04\x31\x14\x84\xfb\xfc\x8a\x21\x6c\xa1\xa0\x7b\xa2\xdd\x81\x9d\x85\x82\x82\x1c\xd7\x89\xc5\xe3\x76\x36\x86\xbb\xe4\x42\x92\xc5\x5b\x42\xfe\xbb\xc4\x5d\xac\x2c\xdf\x0c\xf3\xf1\xbe\x52\x30\x70\xb4\x9e\xd0\x91\xc6\xa6\x1c\x67\x8d\x5a\x55\x29\xe8\xc2\xd1\x60\xfb\x88\xae\x7f\x97\xc3\x51\x0c\xd7\x3c\xd3\x85\x93\x64\x42\x7f\x51\x06\x46\x8d\xae\x35\x6a\xb3\xc1\x7e\x0e\x4c\x88\xcc\x53\xf4\x09\x4e\x42\xb0\xde\x60\x8c\x67\x87\x3c\x07\xc2\x0e\x09\xf9\x8c\xfd\xeb\x72\x7a\x71\x4c\xbd\x1a\x27\x7f\x58\xa6\x6f\x12\xae\xae\xdb\xee\x63\xb2\x3e\x3f\xdc\x7f\xa6\x1c\x1b\xa1\x28\xac\xd4\x7f\xcb\x52\x6e\x11\xc5\x1b\xa2\xe3\x89\x6e\x79\x7a\xb7\xea\xb4\xdf\x00\xe0\xee\xd2\x9c\x5a\xdf\x3f\xf3\xf2\xf2\x84\x5a\xb7\xd0\x7f\xd9\x4e\xbe\x51\xab\xbe\xf9\x85\xd1\x0f\xcb\xac\xaa\xaa\x9a\xf3\x1a\xfc\x04\x00\x00\xff\xff\x97\x1b\x2c\x20\x2e\x01\x00\x00")

func _templateRegistryTmplBytes() ([]byte, error) {
	return bindataRead(
		__templateRegistryTmpl,
		"_template/registry.tmpl",
	)
}

func _templateRegistryTmpl() (*asset, error) {
	bytes, err := _templateRegistryTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "_template/registry.tmpl", size: 302, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"_template/header.tmpl":   _templateHeaderTmpl,
	"_template/main.tmpl":     _templateMainTmpl,
	"_template/registry.tmpl": _templateRegistryTmpl,
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
	"_template": &bintree{nil, map[string]*bintree{
		"header.tmpl":   &bintree{_templateHeaderTmpl, map[string]*bintree{}},
		"main.tmpl":     &bintree{_templateMainTmpl, map[string]*bintree{}},
		"registry.tmpl": &bintree{_templateRegistryTmpl, map[string]*bintree{}},
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
