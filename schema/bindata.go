// Code generated by go-bindata.
// sources:
// gen.sh
// rego-abstract.in
// rego-abstract.json
// rego-attribute.in
// rego-info.json
// rego-spec.in
// rego-spec.json
// rego-type-mapping.json
// DO NOT EDIT!

package schema

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

var _genSh = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xcf\x31\x0a\x02\x31\x10\x85\xe1\x3e\xa7\x18\x57\x21\x5a\x24\x39\x80\x20\x28\x58\xd8\xca\x5a\x87\x49\x8c\xd9\xa8\x9b\x84\xcc\x78\x7f\x71\xd9\x42\x6c\x2d\xdf\x5f\x7c\xf0\x96\x0b\xe3\x52\x36\x0e\x69\x10\xa2\x86\xf6\x04\x55\x03\x48\x32\xd6\xee\xfb\xfe\x7c\x3a\x5c\xfa\xa3\xb5\x46\x76\xab\xb5\x47\x86\x16\x62\x51\xc8\xdc\x92\x7b\x71\xd0\x29\x6f\x3a\x69\xa2\x9c\xbb\x23\x6e\xe8\x59\xa7\x0c\xbb\x9f\x74\xa7\x92\xff\xf5\xa9\x06\xff\x65\x4f\x73\x72\x45\x1b\x41\xdd\x40\x7f\xae\x5c\x91\x51\xc7\xb2\x15\xb1\xa8\x79\x82\xaa\x8f\x08\xe4\x87\x30\x22\x68\xf1\x0e\x00\x00\xff\xff\x19\x53\xac\x2b\xf3\x00\x00\x00")

func genShBytes() ([]byte, error) {
	return bindataRead(
		_genSh,
		"gen.sh",
	)
}

func genSh() (*asset, error) {
	bytes, err := genShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gen.sh", size: 243, mode: os.FileMode(493), modTime: time.Unix(1521698190, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoAbstractIn = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x51\xbd\x6e\xe3\x30\x0c\xde\xfd\x14\x84\x73\x63\x72\xbe\xe9\x86\x6c\x39\xe0\x86\x02\x1d\x8a\x22\x9d\x0d\x5a\xa2\x6a\x05\x8a\x24\x90\x0c\x8a\xa0\xc8\xbb\x17\xb6\x52\x35\x4e\xd1\x4d\x20\xbf\x5f\xf1\xbd\x01\x00\x68\x7f\x89\x19\xe9\x88\xed\x16\xda\x51\x35\x6f\xbb\xee\x20\x29\x6e\xca\xf4\x77\xe2\xd7\xce\x32\x3a\xdd\xfc\xf9\xdb\x95\xd9\xaa\x5d\x17\xa6\x7a\x0d\x34\xf1\x76\x83\x28\xa3\xd1\xcf\x85\x25\x31\xec\xb3\xfa\x14\xe7\x75\x04\xbc\x22\x00\x43\x48\x6f\x02\x9a\xc0\x92\x09\xc8\x04\x08\x42\x0a\xc9\xc1\x80\x42\x80\xaa\xec\x87\x93\x92\x00\x46\xdb\x25\x06\xa6\x80\x93\x92\x80\x8e\xa8\x60\x30\xc2\x40\xc0\x74\x12\xb2\xe0\xa3\x26\x48\x3a\x12\x83\x64\x32\xde\x79\x53\xc0\x35\xe3\x39\xcf\x11\xd3\x70\xa0\xaf\x80\x68\xad\x9f\x60\x18\x9e\x38\x65\x62\xf5\x24\xed\x16\x1c\x06\xa1\x75\x53\x30\xf9\x76\x53\xfe\xaa\x70\x6b\xc2\xc5\x7c\xe1\x86\xcc\x78\xbe\x9a\xd5\xe5\xdd\xaf\x3c\x7a\x99\x6b\xdf\x34\x4e\x0e\x74\x9c\xba\x65\x26\xa1\xa8\x64\x61\x91\xbb\x4a\x79\xa5\xe3\x77\xfb\x72\x4e\x26\x37\xc9\xaf\x3a\x4b\xce\xc7\xb9\xa6\x74\xd5\xa4\x5d\x30\x2e\xcd\xf2\x75\xa9\x07\xac\xd4\x1f\xca\xcf\xe3\xbe\xdf\xed\xf7\xcf\x0f\xff\x5e\xf6\xff\xfb\xfe\x5e\xaa\xb9\x34\x1f\x01\x00\x00\xff\xff\x4d\x1e\x3d\x56\x62\x02\x00\x00")

func regoAbstractInBytes() ([]byte, error) {
	return bindataRead(
		_regoAbstractIn,
		"rego-abstract.in",
	)
}

func regoAbstractIn() (*asset, error) {
	bytes, err := regoAbstractInBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rego-abstract.in", size: 610, mode: os.FileMode(420), modTime: time.Unix(1521698042, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoAbstractJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x51\x6f\xdb\x36\x10\x7e\xf7\xaf\x38\xb8\x05\x0a\x0c\xf1\xdc\x01\x43\x81\xe6\x2d\x0f\x7d\x28\xb0\x87\x62\xd8\x53\x83\x2c\x38\x89\x27\x8b\x0d\x45\xaa\x47\x2a\xb1\x37\xe4\xbf\x0f\x24\x25\xd9\x4a\x64\x59\xb2\x55\x74\x79\x89\x4c\xde\x1d\xbf\xfb\x78\x77\x3c\x4a\xff\x2e\x00\x00\x96\x6f\x6d\x9a\x53\x81\xcb\x6b\x58\xe6\xce\x95\xd7\xeb\xf5\x37\x6b\xf4\x2a\x8e\xfe\x6a\x78\xb3\x16\x8c\x99\x5b\xbd\xff\xb0\x8e\x63\x6f\x96\x57\x51\xd3\x49\xa7\xc8\xeb\xdd\x24\xd6\x31\xa6\xae\x99\x10\x64\x53\x96\xa5\x93\x46\x87\x69\x0d\x58\x4b\x00\x2a\x65\x9e\x2c\x38\x03\x82\x52\x85\x4c\x80\x60\xc9\x81\xc9\x20\x41\x4b\x80\xce\xb1\x4c\x2a\x47\x16\x50\x8b\xb5\x61\x60\x52\xe8\x2d\x59\x70\x39\x3a\x48\x51\x43\x42\xc0\x54\x59\x12\x20\xb5\x33\x60\x5c\x4e\x0c\xb6\xa4\x54\x66\x32\x8d\xc2\x2d\xc6\x5d\x19\x20\x9a\xe4\x1b\xed\x01\xa2\x10\xd2\x8b\xa1\xfa\xc2\xa6\x24\x76\x92\xec\xf2\x1a\x32\x54\x96\xae\x16\x51\xa6\x3c\x9c\x89\x5c\x45\xdd\x16\x61\x67\xbc\xb3\x1a\x32\xe3\xae\x5e\xac\x9d\x7c\xc1\xca\x1f\xd2\x06\xb7\x0f\x3c\x36\x19\xb8\xdc\xfb\x56\x32\x59\xd2\x8e\x04\x74\x70\xb7\xa6\xa4\xa3\xe2\xf5\xf2\x71\x3b\x99\x32\x6f\xfe\xcd\x5a\x50\x26\x75\x70\xd3\xae\xdb\x45\x96\x1d\x8d\xe7\x45\xf7\xe9\xb9\xdd\xc0\x56\xf5\x88\xf3\x3d\xbe\xb7\xd1\xd0\xca\x0c\x13\xe0\xc3\xa2\x11\x0d\x44\x74\xb7\xf0\xa5\x76\xef\x4e\xee\x81\x0d\xee\x68\x47\x94\xe9\x7b\x25\x99\xc4\xf2\x1a\x6e\x5f\xf3\xa7\xb1\x78\x09\x7c\xbf\x7c\xcf\xf8\xa1\x53\x9d\xd9\xbb\x3a\x90\x5a\xc9\x23\x01\xb5\x77\xc1\xa7\x06\x89\xfb\x34\x47\xee\x17\xe9\x63\xf1\x4f\xda\xd0\xb6\x8c\xb9\x81\x60\x1d\x4b\xbd\x39\xe0\xb5\xa8\xac\x83\xdc\x68\xc3\x3e\xe7\x12\x82\x47\x54\x52\xf4\x38\xd2\xe1\x38\x9a\x59\xbe\x12\x7a\xee\x21\x60\x0f\xdb\xc8\xf4\x88\x6f\x7d\xc0\x6f\xa2\x9e\x07\x54\x91\x85\xcc\x30\xe0\x8b\x88\xf0\x70\xe0\x93\xae\x8a\x53\x78\xfb\xf2\xad\x15\x3a\x9e\x2c\xc7\xfc\xee\x37\x14\x24\x4b\x74\x8e\x38\x38\xf0\xf7\xed\xcd\xea\xeb\x6f\xab\x8f\x77\xb7\xb8\xfa\xe7\x66\xf5\xf5\xfd\xea\xe3\xdd\x2f\x6f\x5f\x93\x06\x9d\x44\x6b\x47\xfa\xa8\xac\x9c\xd9\x90\x26\x46\x17\xc2\x73\x1c\x91\x7f\xe5\x07\x65\x13\xa4\x85\x8e\x19\x48\x76\xa1\xa8\x24\x98\x3e\x90\x3e\xb9\xf3\x89\x31\x8a\x50\x8f\xdb\xfa\x94\x29\xa4\xea\xbd\xd1\x6a\x77\x26\x5e\x5f\xcf\xbd\xba\x8f\x4d\x7f\x10\x88\x2a\x44\x70\x30\x4d\xe0\x33\xa6\xaf\x1a\x5c\x86\x5b\x50\x86\x95\x72\xf7\x86\x05\xf1\x24\xdc\x21\x58\x9b\x3a\xdd\x21\x3d\x1c\x47\x68\xa1\xb6\x0d\xc1\xb6\xf7\xe4\x81\x8e\x46\xe6\x45\xe0\x03\x94\x49\xe0\x1b\x68\xfd\x4e\x9c\x00\x79\xdb\x24\xd9\x1e\xee\x15\x2c\xa5\x76\xb4\x21\xf6\x8f\xba\x52\x2a\xfe\x2f\x92\x38\xd2\x54\xea\x36\xaf\xee\x46\x3a\x58\x32\xa5\x97\xa5\xc0\x81\x8d\x79\xb9\x3f\x5c\x76\x0a\xf3\xed\xd0\x54\xde\x27\xd5\x62\xda\x62\x51\x2a\x9a\x18\x1b\x9f\xa2\x56\x1d\x17\xbe\x10\xff\x3f\x03\x83\xb6\xa5\xb1\x97\x44\x45\x59\x25\x4a\xa6\x6a\x07\x8d\xa5\x91\xb1\x71\x44\xac\xce\x27\x2f\xe9\xb8\xa2\x71\x5b\x94\x49\xe5\x88\x31\x51\xd3\x72\xb7\x5b\x30\x13\x8a\xf5\xc6\xef\x55\x34\xe8\x2b\x4d\x59\xb1\x77\xcb\xce\x1a\xf2\x99\x61\x92\x1b\x7d\xef\xcb\xd8\xf9\x27\x12\xd4\x66\x66\xaf\x86\x99\xe1\x02\xdd\x24\x64\x51\x65\x74\x1e\x92\xef\x3f\xfa\x5a\xc5\x3d\x06\xa6\x63\xda\xd1\x42\x81\x52\x0d\x36\x15\xb9\xd1\x83\x16\x3e\x7f\x79\xfc\xfd\xc4\xfc\x87\xa1\xf9\x54\x0a\x1e\x9a\x17\xf8\xf2\x66\xd0\xfc\x8d\xcb\xcc\x0d\xf9\xae\xe8\xcc\xf8\x78\x92\x4a\x41\x8e\x8f\xfe\x2a\xb8\x6f\x5b\x6a\x93\x73\xc6\x8a\x14\xa4\x9d\xcc\xe4\xd9\x48\x0f\x8e\x79\x84\x02\xa5\x86\x03\x93\x97\xd5\x92\x4a\xcb\xef\x15\x7d\xae\x5b\x55\x5f\x4e\xc6\xb9\xa4\x05\x6d\x67\xf0\x46\x43\xb4\x34\x27\xdd\x05\x6e\xef\x15\xe9\x8d\xcb\x27\x01\x2c\x70\x2b\x8b\xaa\x80\xa8\xfa\xba\xd7\x6a\xee\x05\x83\xbd\x7a\x8b\xb8\x39\x87\x46\x23\x9e\xde\x56\x35\x80\x8f\xf4\x86\x0d\xde\xf6\x44\x9c\x13\xb0\xd4\x67\x51\x2c\xf5\x4f\xa3\x58\xea\x73\x28\xae\x01\xff\x04\x8a\xc3\xfb\x80\x29\x58\xbd\xc2\x0f\x6d\xf2\xc2\xcd\x62\xc6\x06\xa2\xbd\xa9\xfc\x90\xfe\xa1\x64\x59\x20\xef\x2e\xeb\x1f\x62\x9d\x82\xda\xd6\xec\x4d\x04\x13\x8a\x4b\xee\xb0\xd2\x82\x37\x11\xee\xb1\x33\x03\x6b\x5f\x55\x9d\x8d\xab\xb6\x30\x27\x2c\x4b\x29\xd3\xb4\xa6\xab\x03\x2a\xea\x03\x6a\x01\x36\x37\x95\x12\x3e\x1e\x05\xa1\x72\xf0\x24\x5d\x0e\x29\xf2\xc9\x8c\x99\x08\x78\xf6\xfe\xc4\xce\xdf\x9f\x58\x67\x2e\xda\xeb\x5a\x7f\xae\x4b\xcd\xe8\x26\xc4\x56\x49\x6d\x7c\x7a\x95\x2c\xb0\x2c\x7d\xed\x71\x06\x4a\x99\x3e\x84\x82\xf4\x94\x93\x0e\xe5\x33\x16\x75\x1b\xde\x0b\x39\x03\xef\x68\xeb\x88\x35\xaa\x77\x73\x96\x53\xc7\xa8\xad\x24\x7d\x41\x3c\xef\x4d\xcc\x19\x0e\x93\x39\x0d\x74\x1d\x7e\x51\x18\x49\xd3\xd9\x17\x9f\xd3\xaf\x4b\x87\x4f\xe4\x20\x92\x29\x83\xc7\x88\x0b\x02\xc3\x01\xbb\x87\x3a\x30\xaf\xa4\x1d\x5c\xa2\xf7\x9b\x42\x47\xc2\xc9\xde\xcf\x02\x7b\x04\x75\x68\x8e\xbe\x45\x2d\xfa\x7f\xd5\xdf\x61\x16\xcf\x8b\xff\x02\x00\x00\xff\xff\xa8\x66\x67\x8a\x9f\x1b\x00\x00")

func regoAbstractJsonBytes() ([]byte, error) {
	return bindataRead(
		_regoAbstractJson,
		"rego-abstract.json",
	)
}

func regoAbstractJson() (*asset, error) {
	bytes, err := regoAbstractJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rego-abstract.json", size: 7071, mode: os.FileMode(420), modTime: time.Unix(1522113745, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoAttributeIn = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\x4d\x6f\xe3\x36\x10\xbd\xef\xaf\x18\x18\x05\x16\x28\x62\xa0\x05\x8a\x02\x9b\x5b\x0e\x7b\xd8\xdb\xa2\xe8\x69\xd3\xd4\x18\x89\x23\x6b\x1a\x8a\xd4\x92\xc3\xc4\x6e\xd1\xff\x5e\x90\x92\x3f\xe4\xc8\xb6\x64\x69\xb1\xf5\x25\xb1\x38\x7c\x7a\xf3\xf4\x38\x1c\xca\x70\xf4\x59\x08\x8b\xa6\xc5\x3d\x2c\x1e\x44\x1c\x67\x41\x68\x71\xf7\xae\x13\xa1\xc8\xe7\x8e\x6b\x61\x6b\x52\x9c\x01\xdc\x85\x82\x2d\x00\xc1\xd7\x94\x73\xc1\x39\xa6\x90\x93\xd9\xb2\xad\x13\xbc\xcd\xfe\xa2\x5c\x4e\x47\x51\x29\x8e\xb3\x50\x7f\x76\xb6\x26\x27\x4c\x7e\x71\x0f\x05\x6a\x4f\x27\xa1\x8e\xbe\x06\x76\xa4\x16\xf7\xf0\xd8\x19\x49\xa3\x06\xab\x53\xe2\x87\xdb\xf7\x5c\x3f\x4e\xaa\x33\xfa\x74\xf7\xae\x7b\xdf\xfa\x98\xd8\x3f\x6f\x91\x50\x6b\xfb\x4a\x6a\x95\x97\xe8\xfa\x43\xfa\x54\xfc\x8d\xd6\xb4\xa9\x41\x4a\x94\xa8\xa0\x38\x36\xeb\x23\x5d\xab\xe0\x05\x4a\x6b\xac\x03\xb1\x90\x11\xbc\xa0\x66\xd5\x93\x48\x47\xe3\x06\x66\xf1\x26\xe8\xdf\x1e\x01\x0e\xb4\x2d\xe7\x67\x72\xeb\x23\xfe\xd0\xcc\x8b\x84\x02\x79\x28\xac\x03\x3c\x71\x44\xa4\x03\x1f\x4d\xa8\xae\xf1\x45\xe7\x70\x7b\x2e\x88\x85\xaa\xf3\xb4\xfa\xf2\xee\x07\x4a\x91\x35\x8a\x90\x4b\x09\xfc\xf9\xf8\xb0\xfc\xf2\xf3\xf2\xc3\xd3\x23\x2e\xff\x7e\x58\x7e\xf9\x69\xf9\xe1\xe9\xc7\x3f\x7e\x78\xab\x5a\x52\x6e\x98\x96\x41\xec\x9a\x0c\x39\x94\xe4\xcf\x61\x4a\xfe\x5e\xd2\x91\x6e\xec\xa1\x03\x03\xd9\x16\xa4\x24\xc8\x30\x7f\x26\x73\xf5\xd1\x67\xd6\x6a\x42\x33\xec\xd9\xe7\x8e\xd2\x5a\x5d\x59\xa3\xb7\x37\xf2\xcd\xd1\x40\x9c\x1e\xcd\xe9\x49\x40\x85\x64\xe1\x04\x4d\x10\x97\x4c\x5f\x39\x98\xc6\x5b\x51\x81\x41\xcb\xca\x3a\x45\x6e\x14\xef\xe4\xd6\xe4\xcd\x53\xd1\x83\x27\x05\xe8\xa1\xc5\x86\x84\x1d\x33\x79\xa6\xb3\xd6\x9c\x44\x3e\x51\x19\x45\x7e\x47\xad\x3f\x89\x2b\x24\x1f\x77\xab\xec\x40\xf7\x0e\x16\x6c\x84\xd6\xe4\xe2\xbf\x26\x68\xdd\xfc\xad\xb2\xe6\xca\xae\x54\xef\x17\xd6\xd3\xc0\x04\x6b\x47\xf9\xb4\x25\x70\x84\x31\xaf\xf6\xc7\xb7\x1d\xa3\xfc\xfe\xd2\x58\xdd\x47\x15\x63\xda\x60\x55\x6b\x1a\xe9\x8d\x8f\xcd\xac\xd6\x17\xb1\x12\xff\x3f\x8d\x41\x9b\xda\xfa\x29\xae\xa8\x43\xa6\x39\xd7\x5b\xd8\x21\x0d\xf4\xc6\x99\xb0\x76\x3d\xc5\x48\x71\x81\x86\x3d\xa2\x82\xb5\x90\xc3\x4c\x8f\x5b\xbb\xdd\x82\x99\x51\x53\x6f\xe2\xb3\x6a\x00\x63\xa5\xa9\x83\x8b\x69\xf9\x59\x2d\x5f\x58\x47\xbc\x36\xab\x58\xc6\x6e\xdf\x91\xa0\x85\x99\xbd\x1a\x16\xd6\x55\x28\xa3\x98\x35\x53\x06\xaf\x43\x8a\x0d\x48\x5f\xaf\x78\xe0\xe0\xe8\xdc\xec\x06\xa1\x42\xd6\x17\xbb\x8a\xd2\x9a\x8b\x08\x9f\x3e\xbf\xfc\x72\x65\xfc\xd7\x4b\xe3\x39\x2b\x77\x69\x5c\xa1\xf4\xd8\x37\x7e\x86\xad\xcc\x35\xc5\xb6\xe8\x46\x7f\xbc\xb2\xd6\x50\xe2\x0b\x01\xc2\xa1\x6d\x69\x21\xe7\xf4\x0a\x2b\x32\xc2\x05\xdf\xcc\xf4\x68\x9b\x47\xa8\x90\x0d\x1c\x41\x4e\xab\x25\xc1\xf0\xd7\x40\x9f\xda\x5e\x35\x96\x93\x61\x29\x19\x45\x9b\x19\xb2\x31\xd0\x20\xcd\x29\x77\x85\x9b\x95\x26\xb3\x96\x72\x14\xc1\x0a\x37\x5c\x85\x0a\x9a\xa9\x6f\x7b\xad\xdd\xc1\xe0\x62\xb3\xbe\x67\xbc\xdb\x87\x06\x33\x1e\xdf\x56\xed\x08\x9f\xe9\x0d\x77\x7c\xf7\x3b\xe2\x9c\x84\xd9\xdc\x24\x31\x9b\xef\x26\x31\x9b\x5b\x24\x6e\x09\x7f\x07\x89\xd3\x0b\x81\x31\x5c\xe3\x84\x6f\xda\xe4\xa5\x93\xc5\x8c\x0d\xc4\xfe\xa4\xf2\x4d\xfa\x87\xda\x71\x85\x6e\x3b\xad\x7f\x68\xea\x14\xb4\x58\xb3\x37\x11\x8e\x50\x4d\x39\xc3\xb2\x87\x08\x91\xce\xb1\x33\x13\xdb\xbf\xab\xba\x99\x57\x8b\x30\x27\x2d\x4f\xb9\xa3\x71\x4d\x57\x87\x54\x33\x1f\xd0\x28\xf0\xa5\x0d\x5a\x45\x3f\x2a\x42\x2d\xf0\xca\x52\x42\x8e\xee\xea\x8a\x19\x49\x78\xf6\xfe\xc4\xcf\xdf\x9f\x78\xb1\x93\x9e\x75\x3b\x7f\xae\x43\xcd\xe0\x26\xc4\x87\xac\x05\x1f\x5f\x25\x2b\xac\xeb\x58\x7b\xc4\x42\xcd\xf9\x73\x2a\x48\xaf\x25\x99\x54\x3e\x9b\xa2\xee\xd3\x7b\x21\xb1\xf0\x9e\x36\x42\xce\xa0\x7e\x3f\x67\x39\x15\x87\xc6\x33\x99\x09\x7e\x3e\x40\xcc\x69\x87\xd1\x9a\x26\xb9\x6c\x31\x7a\xd7\xb9\xf9\xe0\x73\xfd\x7d\xe9\xe5\x1d\x39\x85\x14\xda\xe2\x39\xe1\x52\xc0\x65\xc3\x1e\xa8\x5e\x18\xd7\xec\x2f\xde\xa2\xf7\x47\x85\x4e\x84\x70\xef\xef\x02\x07\x06\xad\x35\x07\x9f\xa2\xde\x75\xbf\xfd\x17\x00\x00\xff\xff\x00\xae\xcf\x69\x4c\x19\x00\x00")

func regoAttributeInBytes() ([]byte, error) {
	return bindataRead(
		_regoAttributeIn,
		"rego-attribute.in",
	)
}

func regoAttributeIn() (*asset, error) {
	bytes, err := regoAttributeInBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rego-attribute.in", size: 6476, mode: os.FileMode(420), modTime: time.Unix(1522092547, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoInfoJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\xbd\x4e\xc3\x40\x0c\xc7\xf7\x3c\x85\x75\x30\xb6\x0d\x13\x43\x36\xc6\x0e\xa0\x4e\x2c\x88\xc1\x4d\x7c\xad\xab\x34\x77\xd8\x0e\x02\x55\x7d\x77\x74\xb9\xa6\x4a\x3f\x36\x96\x44\xf9\xd9\xff\x8f\xf8\x50\x00\x00\xb8\x47\xad\xb7\xb4\x47\x57\x81\xdb\x9a\xc5\xaa\x2c\x77\x1a\xba\x79\xa6\x8b\x20\x9b\xb2\x11\xf4\x36\x7f\x7a\x2e\x33\x7b\x70\xb3\xac\x34\xb6\x96\x92\x6e\xd9\xf9\x30\xc2\x86\xb4\x16\x8e\xc6\xa1\x4b\xa3\x57\x32\x04\xee\x7c\x90\x3d\x26\x06\xb8\x0e\xbd\x01\x82\x92\x41\xf0\xa0\x91\x6a\xf6\x5c\x0f\x43\x5d\x9c\xad\x7f\xe3\xe0\x1c\xd6\x3b\xaa\x6d\xa4\xd8\x34\x9c\xf6\xb0\x5d\x49\x88\x24\xc6\xa4\xae\x02\x8f\xad\xd2\x69\x45\xe8\xab\x67\xa1\xc6\x55\xf0\x31\x90\x81\x46\x21\xcf\x3f\x27\x9b\xbc\x17\x82\x4d\xbf\xbf\x49\x34\x55\x1e\xc8\xe7\xac\x28\x4e\xba\x49\xcc\xe1\xc6\x6f\xca\x2e\x6a\xab\x09\x77\x9b\x89\xff\xbd\xd3\xac\x06\x13\xb0\x00\xbd\x52\x7a\xd5\xd8\xb6\x80\x91\xd5\x9d\x75\xc7\xeb\xca\xff\x8b\x7c\xc3\x3d\xa5\xab\xdb\x96\x20\xd9\x5d\x9e\xff\x56\xec\xb1\x6f\x53\x66\xce\xbe\xdb\x6a\x3c\xdc\x4d\xb1\xab\xe8\xf7\xbc\x37\xa6\xbf\xac\x96\xd7\x71\xe3\x9f\x70\x67\xb4\x21\x99\xc4\x15\xf9\x79\x2c\xfe\x02\x00\x00\xff\xff\xb8\x06\xde\x01\xb3\x02\x00\x00")

func regoInfoJsonBytes() ([]byte, error) {
	return bindataRead(
		_regoInfoJson,
		"rego-info.json",
	)
}

func regoInfoJson() (*asset, error) {
	bytes, err := regoInfoJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rego-info.json", size: 691, mode: os.FileMode(420), modTime: time.Unix(1521233328, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoSpecIn = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x4b\x8f\xdb\x36\x10\xbe\xef\xaf\x18\x28\x3d\x66\xe3\x9e\x7a\xf0\x2d\x7d\x00\x2d\xd0\x02\xc5\xd6\x3d\x15\x81\x41\x93\x23\x9b\x09\x45\x2a\xc3\xd1\xba\x8b\xc0\xff\xbd\xd0\xc3\x94\x64\x3d\x56\x8a\x65\x07\xbd\x2c\xe0\xe1\x70\xe6\x9b\xc7\xa7\x19\xee\x97\x07\x00\x80\xe8\x3b\x2f\x0f\x98\x88\x68\x0d\xd1\x81\x39\x5d\xaf\x56\x1f\xbd\xb3\x8f\xa5\xf4\x9d\xa3\xfd\x4a\x91\x88\xf9\xf1\xfb\x1f\x56\xa5\xec\x4d\xf4\xb6\xbc\xc9\x9a\x0d\xe6\xf7\xfe\x4a\x51\xea\x58\x4b\xc1\xda\xd9\xf3\xa9\x42\x2f\x49\xa7\x85\x68\x0d\xd1\x13\xa6\x84\x1e\x2d\x7b\x10\x16\xdc\xee\x23\x4a\x86\xa3\xe6\x03\x08\x63\x40\xe7\x62\x66\xd2\xbb\x8c\x31\xd7\x50\x40\x68\x0a\x7b\x3e\xb8\x7b\x49\x0b\x6f\xe5\xdd\xb3\x54\x28\xa5\x73\x35\x61\xfe\x24\x97\x22\xb1\x46\x1f\xad\x21\x16\xc6\x63\xa5\x42\xf8\x39\xd3\x84\x2a\x5a\xc3\x3f\x85\xa4\x90\x26\x4e\xa1\x89\x8a\xdf\x1f\x2a\xc5\xb4\x69\xe1\x4b\xad\x5a\x23\x6b\xc9\x5b\xa8\x04\x91\x78\xa9\x40\x85\xc3\x8b\x1c\xfc\xae\x3d\x83\x8b\x9b\x91\xba\x18\xf8\x80\x40\xe7\xec\xa0\xaa\x72\xf3\xee\xd2\x96\x66\x4c\xba\xfe\xcb\x12\x12\xc6\xb9\xfd\x37\x2b\x85\xb1\xb6\x45\x3e\xfc\x2a\x78\x89\x5a\x37\x4e\xe1\xd7\xa9\xf6\x10\xd5\xd9\x5e\x24\xc2\x60\x0e\x76\xc8\x47\x44\x3b\x10\x65\x51\x69\xc7\x07\x24\x70\x16\xfd\xf5\x41\x9f\x1d\x4f\x89\xb9\xec\x80\x4e\xbc\xfd\x86\x1b\xed\x52\x5b\x3c\x85\x5e\x0f\x8a\xed\xce\x19\x70\x11\x98\xf3\x47\x71\x3e\x9e\xd3\x5f\x9d\x51\x1e\x76\xc2\x6b\x09\xda\xc6\x8e\x92\x22\xc2\x19\xad\xd3\xcb\x9c\x70\xfa\x3a\x83\x1a\x5d\xd2\xc3\xa4\x70\x9a\x0a\xf9\x49\xec\xf1\xc2\x7e\x75\xd1\xf3\xd6\x8a\x64\xe8\xd0\x65\x24\x71\x50\x01\x2d\x6b\x7e\x19\x3c\x6e\xe6\xab\x75\xfa\xe1\xed\x43\x1b\xff\x00\xc1\xeb\x54\x18\x2d\xfc\xc0\x61\xc7\x55\x93\xd0\xe5\x3d\x60\x57\x97\x03\x04\xf8\xe6\x77\xf1\xb2\x2a\xc1\xe6\x18\xbf\x82\xd2\x30\x11\x3a\x76\x3c\x93\xb6\xfb\xa8\x57\xf1\xd4\x91\x9e\x7a\x12\x2a\x09\x05\xe3\xe4\x2c\x6c\xfa\xa9\x2d\x85\x85\x1d\x42\x69\x4c\xbd\x1a\xfe\xce\x39\x83\xc2\x76\x71\xf7\x21\x54\x68\x70\x1e\xc2\x61\x80\xa5\xad\xc5\x01\x36\x01\x4c\xcf\x63\x43\x34\x9d\xe0\x1d\xa8\x43\x2d\xd0\x87\xb4\x49\xae\xa9\x48\x7f\x12\x09\x9a\x47\x29\x3c\x2a\x78\x46\xf2\x2d\xb4\x15\xd3\x97\x05\xf9\x2f\xa3\x55\xf3\x89\xb9\x13\x1e\xdb\x34\x04\x3e\x68\x7f\x21\xaa\xac\xc3\xff\x8a\xa4\x7b\xe4\x85\x18\x4a\xc8\xa4\xf1\x79\x69\x0a\x9c\xc7\xc1\x1c\x90\x79\xe3\x9c\x3b\xa9\xba\x7f\x27\x1a\xa4\xa4\x9f\x17\xf8\xe8\x69\x0f\x95\xa5\x65\x93\xd9\x9e\x91\x73\x40\xa6\x26\x23\x61\xee\x43\xd3\x7a\xcc\xcf\x41\x98\x59\xfd\x39\xab\x6a\x1f\x52\xaa\xed\xbe\x00\x7a\x83\x52\x93\x73\x57\x53\x47\xc8\xfc\xc5\xe2\xcb\x5c\x3a\xc7\x33\x67\xfd\xac\xda\x67\xa9\x5a\x6e\x1c\x97\xc6\xae\xa1\xfa\x84\xcd\xba\x7e\x7a\xe4\xa0\xb7\xdb\xf7\x9b\xcd\xd3\x6f\x3f\xfe\xbd\xf9\x65\xbb\x1d\x7d\x7d\x8c\x6c\xca\x4f\x67\x95\xf1\x65\xf9\xbd\x31\xee\x58\x2c\x62\x0a\xa5\x11\x84\x20\xc2\x63\xa4\x7e\x8b\x1c\x5d\xbb\x5e\xdd\x87\xc7\x9d\x56\xe6\xd1\xbd\xb8\x11\x9a\xbf\x72\xb1\xfd\xaa\x8d\xce\x88\xba\x7d\xfc\xc5\x3a\x07\x31\xb9\xa4\xe8\x7e\x99\x11\xe5\xeb\x6e\xb7\xe5\xbe\xf1\xc6\x37\x14\x41\xb5\xef\x7d\x83\x08\x52\x42\x99\x23\x9a\x1f\x45\xde\xbd\xda\x43\xc3\xc4\x8d\x96\xd5\x91\x15\xab\xff\x5d\x3c\xdc\xa4\x43\xde\xe6\x2f\x2e\xbd\x65\x0c\x5b\xcb\xdd\x0b\x39\x7f\xca\x85\x1b\x90\xe5\x2b\xb3\xf0\x20\xc0\x68\xfb\x29\xff\x4e\x71\x27\xc4\x25\x67\xdd\x57\xcd\x8e\xde\x7c\x57\x83\xe3\xb6\xd9\x9e\x30\x5b\x46\x9b\xb5\x3b\x2f\x9a\x2f\x2a\x0f\x29\x52\x3e\xba\x5f\x9f\x23\x3f\x3b\x99\x25\x68\xf9\xe2\x5f\x2d\xa6\x77\xbc\x2f\x37\x2d\x16\xfe\x8e\xb7\xa3\x88\x1d\x95\x85\x2b\x8c\x40\xee\x69\xda\xb6\x32\xa7\xe1\x66\x7e\xa7\xfb\x11\x96\x46\x6a\x84\x4b\x02\x9c\xf3\xf9\xe9\x47\xb7\x47\xbe\x0d\xb4\x99\x64\xed\x47\x57\x1a\x59\x00\xe0\x10\x17\x1f\xca\xbf\xa7\x87\xff\x02\x00\x00\xff\xff\x21\xd1\x31\x60\x41\x18\x00\x00")

func regoSpecInBytes() ([]byte, error) {
	return bindataRead(
		_regoSpecIn,
		"rego-spec.in",
	)
}

func regoSpecIn() (*asset, error) {
	bytes, err := regoSpecInBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rego-spec.in", size: 6209, mode: os.FileMode(420), modTime: time.Unix(1521702680, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoSpecJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x5a\xdd\x6f\xdb\x36\x10\x7f\xcf\x5f\x71\x70\x0b\x14\x18\xea\xa6\x03\x86\x02\xcd\x5b\xb0\x15\x58\x81\x0d\x28\xba\x3d\x35\xc8\x02\x4a\x3a\x59\x6c\x28\x52\x25\xa9\xc4\x5e\x91\xff\x7d\x20\xa9\x4f\x5b\x92\x45\x8b\x49\xba\x97\x36\xe6\xc7\xf1\x77\xc7\xfb\xa2\xee\xbe\x9f\x01\x00\xac\x5e\xaa\x38\xc3\x9c\xac\x2e\x60\x95\x69\x5d\x5c\x9c\x9f\x7f\x55\x82\xaf\xdd\xe8\x1b\x21\x37\xe7\x89\x24\xa9\x5e\xbf\x7d\x77\xee\xc6\x5e\xac\x5e\xbb\x9d\x9a\x6a\x86\x66\xdf\x5f\x05\xc6\x34\xa5\x31\xd1\x54\xf0\x7a\x36\x41\x15\x4b\x5a\xd8\xa1\x0b\x58\x7d\xc6\x42\xa2\x42\xae\x15\x10\x0e\x22\xfa\x8a\xb1\x86\x7b\xaa\x33\x20\x8c\x01\x35\xc3\x5a\x4b\x1a\x95\x1a\xcd\x8a\x04\x24\x32\x4b\x4f\x35\xc7\xed\x0a\x7b\x9a\xdb\x5b\x8f\x92\x24\xa1\x66\x19\x61\x9f\xa4\x28\x50\x6a\x8a\x6a\x75\x01\x29\x61\x0a\xab\x25\x12\xbf\x95\x54\x62\xb2\xba\x80\x2b\x3b\x62\x47\x73\x91\x20\x5b\xd9\xdf\xd7\xd5\xc2\xa2\x4b\xe1\x7b\xbb\xb4\x45\xd6\x1b\xef\xa1\x22\x52\x92\x5d\x05\xaa\x99\xdc\x93\xc1\x1f\x54\x69\x10\x69\x97\x53\x91\x82\xce\x10\x64\x2d\x1d\x4c\x2a\xd9\xbc\xd9\xa7\x45\x35\xe6\x87\xe7\xbb\x2b\x94\x98\x1a\xfa\x2f\xce\x13\x4c\x29\xb7\xf2\x50\xe7\xcd\x29\xab\xde\x8e\x87\xe6\xd7\x43\x7b\xc2\xaa\x95\x76\x10\x0e\x1b\x72\x10\xa1\xbe\x47\xe4\x23\x5c\xda\x9b\x16\x3a\x43\x09\x82\xa3\x5a\xce\x74\x7d\xf0\x1c\x9e\x9d\x06\x1c\xf0\x3b\x4c\xb8\xa3\x2e\x2d\xc5\x87\x46\xd7\x9b\x85\x7d\xcd\x19\x39\xa2\xb1\x9c\x3f\xed\xfc\xb4\x4c\x7f\x17\x2c\x51\x10\x11\x45\x63\xa0\x3c\x15\x32\xb7\x1c\x7a\xa8\xce\xa0\xe5\x34\xb3\xc7\x2d\xa8\xa3\x25\x03\x96\xd4\xcc\x16\x24\xbe\x25\x1b\xdc\xa3\x5f\x6d\x54\xfa\x86\x93\x7c\x6c\x52\x94\x32\xc6\xd1\x05\xc8\x35\xd5\xbb\xd1\xe9\xae\xbc\x7a\xb3\xd7\xaf\xcf\xfa\xf8\x47\x0c\xbc\x15\x05\xa3\x44\x8d\x4c\x1e\x1c\xd5\x35\x68\xb7\x0f\xb4\x68\xaf\x03\x08\xa8\xae\x5f\xdc\xbf\x95\x86\xe6\x94\x7d\x35\x8b\xc6\x0d\xe1\x80\x8e\xd2\x92\xf2\xcd\x6a\x70\xe1\xc3\xc1\xe8\xc3\x80\x40\x63\x89\x44\xe3\x6c\x29\xfc\x3d\x6c\xda\x31\xe1\x10\x21\x38\x62\xc9\x51\xf6\x23\x21\x18\x12\x7e\x88\x7b\x08\x61\x82\x0c\xfd\x10\x8e\x03\x74\xb4\x82\x03\xec\x02\x98\x2f\xc7\xce\xd0\x7c\x03\x3f\x80\x3a\xa6\x02\x43\x48\xbb\xc6\x35\x17\xe9\xaf\x24\x47\xb6\x8e\x89\xc2\x04\xee\x50\xaa\x1e\xda\xca\xd2\xc3\x82\xdc\x6a\xe4\x89\xbf\x61\x46\x44\x61\xdf\x0c\x41\x67\x54\xed\x0d\x55\xd4\xe1\x7f\x65\xa4\x1b\xd4\x81\x2c\x54\xa2\x96\x14\xef\x42\x9b\x40\x1d\x0e\x7c\x40\x1a\xc5\xa9\x35\xa9\xda\xff\x44\x66\x50\x48\x7a\x17\xc0\xe9\x51\x05\x15\xa5\xb0\xc2\xec\xc7\x48\x1f\x90\x05\x2b\x25\x61\x4f\x63\xa6\x6d\x98\xf7\x41\x58\x72\xfa\xad\xac\xee\xbe\x11\x29\xe5\x1b\x0b\xf4\x11\xae\x5a\x0a\xb1\xd8\x74\x48\x6c\x5e\x2c\xca\xc9\x52\x08\xed\x19\xeb\xbd\xee\xbe\x2c\x92\x70\xe1\xd8\x11\x5b\x62\xea\x33\x32\xeb\xf6\xe9\x31\x9e\xfa\x5e\x36\x6b\xa6\xd3\xdf\x4b\xde\xbe\x97\x6c\xae\xd5\x17\xf5\x33\xe5\xba\x63\xb9\xa8\x3d\xfe\x91\x73\x54\x26\xee\x31\xb9\x89\x33\x22\xe7\x07\xc4\xcf\xb8\xc1\x6d\x01\x3a\x23\x36\x31\xb5\xd6\xd2\x91\x6b\x5e\x2a\x0d\x99\xe0\x42\x9a\x1c\x36\x42\xb8\x23\x8c\x26\x21\x8d\xae\x85\x2d\x68\xec\x91\x62\x5f\xba\x7d\x06\x50\x89\x0a\x52\x21\x81\xec\x69\x84\x81\x03\x1f\x78\x99\x3f\x71\xf0\x1e\x26\x04\x2e\xf6\x69\x8d\xd2\x32\xf0\xcf\xd5\xe5\xfa\xcb\xcf\xeb\xf7\xd7\x57\x64\xfd\xef\xe5\xfa\xcb\xdb\xf5\xfb\xeb\x9f\x5e\x2e\x89\xfc\xa4\xd4\x62\x83\x1c\xa5\x31\x64\x2f\xb7\xd0\x8a\x8d\x2a\xe8\x91\x81\x68\x67\x7d\x59\x44\xe2\x5b\xe4\x47\x6f\xde\xcb\x7f\xd9\x17\x00\x15\xfc\x46\x70\xb6\x3b\x11\xaf\xf1\x5e\x66\xbb\xd1\x4d\x85\x1a\x92\xd2\x6a\xb0\x7b\x5c\x80\xb1\x98\x21\x6f\xb0\x0c\x77\x82\x29\x29\x99\xbe\x11\x32\x41\xe9\x85\xdb\x2a\x6b\x1d\x6a\x7b\x42\x2f\x4d\xca\x4c\x14\x54\xb4\xc1\xd2\x36\x9c\xdc\xe2\xa8\x66\x2e\x02\x6f\xa1\x78\x3e\x41\x1c\xb4\x61\x26\x8e\x80\xbc\xaa\x8d\xac\x85\xfb\x1a\x56\x94\x6b\xdc\xa0\x34\x7f\xf2\x92\x31\xf7\x7f\x1e\xb9\x91\xda\x53\x37\x76\x75\x3d\x93\xc1\x42\x62\xbc\xcc\x04\x3a\x34\x7e\xd0\xc7\xdf\x5c\xb9\x7b\xbe\xa6\x48\x5e\x30\xf4\xd4\x8d\x0f\x6e\x57\xa5\x17\xc6\x11\xff\x98\x8a\x81\xdb\x42\xa8\x25\x5a\x51\x94\x11\xa3\x31\xdb\x41\x4d\x69\xa6\x6e\x8c\x2c\xab\xec\xc9\xac\xd4\xb2\xc4\x79\x57\x94\x52\xa6\x51\x92\x88\xf9\xd9\x6e\xdf\x61\x9a\x74\xcf\xf8\x1b\x73\x57\x8e\xa0\xf1\x34\x45\x29\x0d\x5b\x2a\xa8\xca\xa7\x42\x22\xdd\xf0\x1b\xe3\xc6\x4e\x8f\x48\x50\x91\x09\xee\x0d\xdd\x57\x53\x2f\x64\x6e\xcb\x6c\x3b\x44\x93\x7f\x0c\xa5\x8a\x2d\x06\x89\x63\xbb\x1d\x85\x9c\xd0\xfd\x6f\xc1\xbd\x05\x45\x26\xf8\x24\x85\x8f\x9f\xee\x7e\x39\x32\xff\x6e\x6a\x3e\xa6\x89\x9c\x9a\xb7\xcf\x90\xc1\xe9\x79\x96\xb9\x41\x93\x15\x9d\xa8\x1f\xf7\x94\x31\xc8\xc8\x1d\x02\x81\x36\x6d\xa9\x48\x86\xd4\x15\x9a\x98\xe7\x67\x4a\x4f\x46\xda\x09\xf3\x04\x72\x42\x39\x74\x48\x2e\xf3\x25\xee\xa9\xfc\xb1\x4a\x55\x8d\x3b\x99\xc7\x12\x4f\x70\x1b\x80\x1b\x0e\x8e\x52\x48\x71\xe7\x64\x7b\xc3\x90\x6f\x74\xe6\x05\x30\x27\x5b\x9a\x97\x39\xb8\xad\x87\xb9\x56\xfd\x2e\x98\xcc\xd5\x1b\xc4\x75\x1c\x9a\x8d\xd8\x3f\xad\xaa\x01\x8f\xe4\x86\x35\xde\x26\x22\x86\x04\x4c\xf9\x49\x22\xa6\xfc\xd9\x44\x4c\xf9\x29\x22\xae\x00\x3f\x83\x88\xbd\x3f\x7a\x75\xbf\x74\x3e\x4a\x92\x67\x5f\x16\x01\x13\x88\xe6\xa5\xf2\x28\xf9\x43\x21\x69\x4e\xe4\x6e\x59\xfe\xe0\xfc\x14\x54\xb4\x82\x27\x11\x12\x49\xb2\xe4\x0d\x4b\x15\x18\x12\xf6\x1d\x1b\x18\x58\xf3\xa9\xea\x64\x5c\x15\x85\x90\xb0\x14\xc6\xd2\xb3\x48\xd1\x03\xe5\xf6\xdb\x16\x01\x95\x89\x92\x25\xae\x5a\x47\x58\xd5\x38\x12\x13\x79\xd4\x62\x3c\x01\x07\xcf\x4f\x54\xf8\xfc\x44\x69\xb1\xe8\xae\xab\xfd\xa1\x1e\x35\xb3\x93\x10\x55\x46\x15\x71\x7f\x2f\x99\x93\xa2\xb0\x15\x01\x01\x05\x8d\x6f\xad\x43\xba\xcf\xaa\xb6\x12\xe7\xd4\x95\xfd\x2e\xa4\x05\xbc\xc2\xad\x46\xc9\x09\x7b\x15\xd2\x9d\x6a\x49\xb8\xa2\xc8\x17\xe8\x73\x4b\x22\xa4\x3a\x78\xcb\xd4\x8a\xab\xdb\x85\x34\x53\x4c\x27\x3f\x7c\x8e\x7f\x2e\x9d\x8e\xc8\x76\x49\xca\x04\x19\x13\x9c\x5d\x30\xad\xb0\x2d\xd4\x89\x79\x46\xd5\xe4\x11\x83\x35\x85\xde\x0a\x4d\x07\xcb\x02\x2d\x82\x4a\x35\x67\xbf\xa2\xce\x86\x7f\x0d\xb5\x6e\x4d\xd4\x5a\x3e\xd7\x4b\x8e\x94\x5a\x18\x13\xf7\xb6\x8b\x25\xc1\x98\x11\x69\x5c\x59\x4d\xbd\x6d\xe4\xba\x17\xfd\x0a\xcc\x61\xd7\xd6\x13\xd5\x60\x26\x9b\x8a\x3a\xac\xa9\x85\x15\x97\x93\xda\x61\x18\x69\x6b\x6f\x6a\xaf\x17\x06\x52\x29\x72\xeb\xba\xe2\x52\x4a\xe4\x7a\xa0\x5e\xf7\xcc\xed\x32\x63\x1c\x54\xcd\x32\xcf\xc0\xc1\x49\x5f\x7b\x1b\xed\x7d\x92\x8f\xbd\x13\x55\xad\xe1\xa6\xc2\x71\x25\x1d\x3b\xcd\xbf\xeb\x63\xf0\x1a\x9b\x96\x8f\x27\xbf\x48\xff\x16\x81\x66\x47\xe7\xab\x0a\xa3\xfc\xd6\xf8\x29\x7d\xc0\x62\xc8\x98\x7f\x52\xe1\x7d\x50\xde\x55\xd5\xfd\x71\xa5\x3d\x23\x56\x4c\x2a\xeb\x61\xbc\xe8\x56\x24\x14\x14\x28\x81\xc4\x33\xe2\xc8\x6f\x22\x2e\x73\xe4\x7a\xaf\x4f\x95\x0d\xf6\x46\x84\x8b\x16\x81\xfd\x78\x9f\x8b\xba\xde\xb1\x5f\x71\x0c\xda\x99\xe2\xe9\xa7\x87\x11\x3a\x22\xf3\x6b\xa2\x3e\x00\x7d\xdc\xcf\x30\xba\x0d\xea\xc7\x81\xe6\x69\xac\xc3\xe8\x1c\x91\x00\x00\xc7\x6c\xf1\xcc\xfd\xfb\x70\xf6\x5f\x00\x00\x00\xff\xff\x3d\x12\x41\x84\x7e\x31\x00\x00")

func regoSpecJsonBytes() ([]byte, error) {
	return bindataRead(
		_regoSpecJson,
		"rego-spec.json",
	)
}

func regoSpecJson() (*asset, error) {
	bytes, err := regoSpecJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rego-spec.json", size: 12670, mode: os.FileMode(420), modTime: time.Unix(1522113745, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _regoTypeMappingJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\x3f\x6f\xdb\x3c\x10\xc6\x77\x7f\x8a\x03\xf3\x02\x2f\x50\x24\x56\xa7\x0e\x5e\x8b\x02\x5d\x0a\x74\xe8\x56\x74\x38\x8b\x27\xf9\x12\x8a\x64\x8f\xe7\xa4\x6a\xe1\xef\x5e\x50\x8a\x98\xc8\xb2\xe3\x00\x45\x36\xe9\xfe\xe8\x39\x3d\xbf\x23\xff\xac\x00\x00\xcc\x7f\xa9\xde\x51\x87\x66\x03\x66\xa7\x1a\x37\x55\x75\x9b\x82\xbf\x19\xa3\xeb\x20\x6d\x65\x05\x1b\xbd\x79\xff\xa1\x1a\x63\x57\xe6\x7a\xec\x54\x56\x47\xb9\xef\x5b\x1f\x09\xbe\x60\x8c\xec\xdb\x34\x65\x2d\xa5\x5a\x38\x2a\x07\x3f\xd4\xec\x38\x41\xc3\x8e\x80\x13\xec\x13\x59\xd0\x00\x1d\x46\xa0\x5f\x4a\xe2\xd1\x81\xf6\x91\x52\x8e\x3a\xf4\xed\x1e\x5b\x82\x14\xa9\xe6\x86\xeb\xb4\x2e\x92\x7d\x1c\x14\xc3\xf6\x96\x6a\x9d\xa2\x68\x2d\x67\x1d\x74\x5f\x25\x44\x12\x65\x4a\x66\x03\x0d\xba\x44\x8f\x25\x11\x35\xcb\xcc\xf2\xa3\x01\x43\x7a\xfd\x6e\xf6\x3e\x1a\x23\xd4\x64\xad\xab\xca\x52\xc3\x7e\x50\x48\x55\x8d\x4a\x6d\x90\xde\x94\xe2\xc3\xf0\x74\x28\xbf\x5d\x6a\xe7\x12\xa5\x71\x21\x54\x7c\xfc\x38\x95\x5c\xcf\x0b\x16\x56\x12\x4c\x5f\x2b\x6e\x6e\x7b\x68\xc9\x93\xa0\x06\x49\xc0\x1e\x82\x58\x92\x6c\xa7\x0a\xfa\xe4\x50\xe9\xd8\x6a\xf6\x39\xbb\x23\x16\x08\x0f\x1e\xb8\x8b\x8e\x3a\xf2\x8a\x59\x68\x7d\x3c\xc4\x49\xeb\x4b\xf6\x32\x82\x52\xfa\x32\x8a\x97\x90\x5c\x40\xd3\x8d\x1b\x68\x16\x3d\x87\xd5\xe9\xb7\xc3\xd3\x64\x66\x6a\x3e\x4f\xe7\xf9\x96\x5f\x24\xc4\x09\xee\xa8\x7f\xbe\xeb\x13\x06\x0e\x1e\xd0\xcf\x59\xfc\x9f\x20\xed\xb7\xf9\x09\x1a\x09\x1d\x60\x59\xfd\x37\x66\x21\xf4\x73\xcf\x42\xd6\x6c\xe0\xfb\x12\xc1\x20\x33\x0b\xff\x38\x66\x79\x01\xe2\xe3\xa0\x67\x30\x9e\xd8\xeb\xb2\xac\x76\x30\xe6\xf8\xcf\x17\x0e\x24\x95\xd3\xc8\x97\x7d\x26\x2f\xca\xab\x87\xf9\x1c\x1e\x32\xb5\x61\xb9\xd0\xf1\x6f\xca\x27\xe5\x0d\x66\xea\x62\x90\xd7\x4f\xf5\xe9\x9e\xbc\xee\xd1\xc1\x13\x63\x18\x3f\x01\x41\x20\x62\x7d\x97\xef\x4d\xdd\xa1\x42\x94\x70\xcf\x36\xdf\xa8\xff\x3c\xf7\xb9\xe3\x33\x5e\x7c\xab\xc3\xea\x6f\x00\x00\x00\xff\xff\x72\x98\x35\xd6\x4b\x06\x00\x00")

func regoTypeMappingJsonBytes() ([]byte, error) {
	return bindataRead(
		_regoTypeMappingJson,
		"rego-type-mapping.json",
	)
}

func regoTypeMappingJson() (*asset, error) {
	bytes, err := regoTypeMappingJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rego-type-mapping.json", size: 1611, mode: os.FileMode(420), modTime: time.Unix(1521685528, 0)}
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
	"gen.sh": genSh,
	"rego-abstract.in": regoAbstractIn,
	"rego-abstract.json": regoAbstractJson,
	"rego-attribute.in": regoAttributeIn,
	"rego-info.json": regoInfoJson,
	"rego-spec.in": regoSpecIn,
	"rego-spec.json": regoSpecJson,
	"rego-type-mapping.json": regoTypeMappingJson,
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
	"gen.sh": &bintree{genSh, map[string]*bintree{}},
	"rego-abstract.in": &bintree{regoAbstractIn, map[string]*bintree{}},
	"rego-abstract.json": &bintree{regoAbstractJson, map[string]*bintree{}},
	"rego-attribute.in": &bintree{regoAttributeIn, map[string]*bintree{}},
	"rego-info.json": &bintree{regoInfoJson, map[string]*bintree{}},
	"rego-spec.in": &bintree{regoSpecIn, map[string]*bintree{}},
	"rego-spec.json": &bintree{regoSpecJson, map[string]*bintree{}},
	"rego-type-mapping.json": &bintree{regoTypeMappingJson, map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
