// Code generated by go-bindata.
// sources:
// ../../resource/test_certs/ca.crt
// ../../resource/test_certs/ca.key
// ../../resource/test_certs/node.crt
// ../../resource/test_certs/node.key
// DO NOT EDIT!

package securitytest

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
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

var _test_certsCaCrt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x92\x4b\x73\x9b\x3e\x14\x47\xf7\x7c\x8a\xff\xde\xf3\x9f\xa8\xd8\xad\x61\xd1\xc5\xd5\x03\x90\x89\x34\x11\x60\x13\x65\x07\x76\x10\x0f\x3f\x62\x70\x2c\xe2\x4f\x5f\xec\x45\xa7\x9d\xde\xe5\x59\xfc\xe6\xcc\x99\xfb\xff\xfd\x30\x0b\xb9\xfc\x8f\xb0\x24\xe3\x01\x27\x90\xb1\x07\x75\x04\xe7\xf8\x9d\x12\x02\x69\x6a\xc0\x72\x0c\x86\x2b\xd9\x55\x9b\xb1\x63\x55\x1d\x3d\xbf\x17\x68\xb6\x2a\xab\x65\x74\xa4\xf0\x8c\x4d\x77\xae\xbb\x26\xf4\x2d\xc2\xa0\x06\x1b\xdf\x58\xea\x08\x0c\x21\x7c\x5b\x33\x52\x8b\x95\x72\xfd\xb6\x98\xaf\xae\x3a\x97\x27\x91\xac\x2d\xb3\x9a\x6e\x94\xa2\xcc\x8e\xb4\x74\x65\xbf\x3d\xf8\xb5\x76\x8d\x51\x88\xd9\xa8\xde\x4a\x91\xad\xad\x23\x5b\x58\x88\xcc\x20\x49\x95\xcd\x1f\x50\xdb\x89\xcd\x7f\xb3\x16\x7a\x91\x70\xcb\xe0\x31\x16\x33\xbb\xff\x73\x6c\x74\x82\x0c\x32\x6c\xe4\x06\x83\xc8\x28\x93\x57\xed\x5e\xbe\x4a\x37\x68\x0b\x82\xa9\xca\xf0\x56\x00\x0a\x49\x7a\x0e\x53\x5e\xce\xa9\x62\x93\xfa\x1a\x00\x0d\x20\x98\x21\x4a\x11\xcf\x39\xcf\xca\xc1\x3d\x86\x2f\x33\x8d\x9b\x03\x2a\xfc\xe8\x74\x19\xcf\xd7\x83\xcf\xf5\x17\x1d\xcb\xdd\x65\x6f\xb1\xee\xcb\xb7\xbd\xa1\x88\x45\xfa\x13\x1d\xfa\x9d\xa9\x8a\x62\xd9\xd0\x3d\x39\xc4\x59\xe8\x8c\x30\xcb\xdf\xc7\xe3\x76\xa4\xed\x79\x28\x87\x8f\x20\x78\xab\xc0\x4c\x65\x20\x6c\x57\x2d\x74\x02\x16\xf7\x46\x3b\x6a\x19\x7e\xb2\x8a\x4d\xa1\xe1\x3b\x85\xf4\x6e\x1d\x25\x02\x3b\x50\x79\x8c\xd0\x29\xa4\x8a\x9e\xc0\x30\x22\x60\xf8\x4b\x99\x58\x39\x8d\xc9\xe4\xf3\xe8\x56\x97\xa5\x3e\xed\x4a\x76\xf3\xf3\xb1\x7b\xdd\x91\x38\x7c\xc9\x13\xa7\xf6\x3a\xf7\x99\xef\xf1\x38\x17\xd7\x79\x93\xe2\x45\xa4\x92\xfe\xd8\x10\xab\xed\x2d\x8e\x84\xd7\x7a\x3f\xf2\x3e\x8a\x25\x9a\xe5\x1f\x4b\xaf\x06\xf4\xba\xb2\xd5\xba\xd9\x5c\x82\xc5\x4f\xe7\xf1\x09\x4c\xd2\x7f\xbf\xe3\x57\x00\x00\x00\xff\xff\xcb\x20\xf7\xeb\x3a\x02\x00\x00")

func test_certsCaCrtBytes() ([]byte, error) {
	return bindataRead(
		_test_certsCaCrt,
		"test_certs/ca.crt",
	)
}

func test_certsCaCrt() (*asset, error) {
	bytes, err := test_certsCaCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test_certs/ca.crt", size: 570, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certsCaKey = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xd1\xcb\x96\x63\x40\x00\x80\xe1\xbd\xa7\xe8\xbd\x33\xc7\x25\xa4\x59\xcc\xa2\x8a\x62\x4a\xa4\x94\x12\xd7\x1d\xe2\x1e\x89\x16\xa2\xe5\xe9\xe7\x4c\xaf\xe7\xdf\xfe\xcb\xef\xd7\xbf\x20\xb2\x31\xf9\x60\x01\xf8\xa0\x0c\x47\xe0\x82\x3e\x4e\x28\xfd\x39\xdc\x19\x63\xe8\xf9\x18\x02\xe0\x40\xe0\xee\xb3\xba\xbe\xf3\x32\x15\xd4\xc6\x3e\x15\x6c\x16\xab\xee\x60\x7f\xf1\x85\xd8\xf6\x98\x5a\x8b\x2c\x27\xc0\xea\x96\x65\x8c\x01\xf5\xfd\x6b\xf7\x19\xe4\xdc\x22\x43\x7d\x9a\x57\xec\x85\xa7\x74\x3a\x17\xc8\x54\xb3\x4f\xfb\xfa\x46\x9e\xf7\xf5\xbd\xee\x43\x18\xdd\x34\x03\x6c\x08\x00\xdf\x01\x9e\x62\x23\x5a\x51\xba\x26\x65\x26\x51\xa1\x89\x8b\xd1\xe3\xda\x24\xb2\xaa\x5d\xaf\x1f\xb0\x96\x78\x7b\x27\x5b\x3c\xa5\xcd\xac\xe6\xe2\x81\x54\xd1\xc1\x2c\x74\x26\x1b\x97\xab\x7f\xad\xe3\x71\x14\xcc\x6e\x4d\xf3\xe3\x30\x46\x42\x99\x14\xa5\x06\x5d\x3e\xe4\x1e\x3e\x6e\x81\x67\x91\x7c\x7e\x0e\xf5\x74\x91\xd9\x1e\xb7\xc0\xa9\x87\xed\xd1\xeb\xc5\xc4\x0c\xc2\x63\xd8\xbf\x4b\xdc\x98\x86\x9e\x86\xae\x03\x3a\x04\xa4\xd1\x74\x21\x5e\x27\xbf\xee\x77\x0e\x66\x2b\xb9\x16\x34\x4b\x14\x6d\xa9\x6c\xf1\xbb\x95\x54\x9a\x14\x19\x1c\x1a\x30\xf2\x7c\x69\x60\x43\x1f\x9e\x47\x73\x0d\x6e\x53\x52\xc3\xd7\xc0\xac\xf8\x95\xc8\xc1\x6d\xdd\xb4\x46\x5d\x98\xca\x69\x0a\x1c\xff\xc8\xa8\x8d\xc2\x8c\x81\x0e\xc6\x0e\x1d\xe8\xac\x2c\x93\xdb\x6b\xf9\x5b\xc3\x24\x09\xdc\x29\x18\x8d\x33\x59\xf6\x9e\xe8\xab\x77\x2f\x2c\x6b\xa4\x0d\x8b\x37\xdc\x54\x34\x45\x9c\x2f\x3c\x7d\x69\x31\x98\x78\xbc\x4b\x0f\xa5\x57\x24\x45\x15\x9d\x53\xe8\x93\x78\x3f\x6c\x4c\x7c\x1d\x85\x39\x38\xde\xc5\xdf\xdc\x0f\x1a\x22\xe6\xff\x31\xff\x06\x00\x00\xff\xff\x6b\x78\xcb\x8b\xed\x01\x00\x00")

func test_certsCaKeyBytes() ([]byte, error) {
	return bindataRead(
		_test_certsCaKey,
		"test_certs/ca.key",
	)
}

func test_certsCaKey() (*asset, error) {
	bytes, err := test_certsCaKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test_certs/ca.key", size: 493, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certsNodeCrt = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xd6\xcf\x77\xa2\x3a\x14\x07\xf0\x7d\xfe\x8a\xb7\x9f\xf3\x8e\x21\x81\xa9\x2e\x13\x12\x15\x6d\xa2\xa0\xc8\xd0\x9d\xe2\x34\x88\xad\x7d\x6a\x35\xc2\x5f\xff\x90\x19\xfb\x43\x6e\x37\xf5\xdc\x93\x9b\xe4\xf3\x4d\x38\xf0\xef\xf5\x8f\xcb\x41\xa0\xff\xf1\x65\x34\x0f\xfa\x81\xcf\xe6\xb2\xa9\x22\x15\x04\xc3\xa2\xf0\x7d\x5e\xed\x0d\xb3\x01\x67\x26\x08\xc7\xc5\x61\x38\x1d\x0f\x37\x89\x3f\x48\x2e\x47\x8e\xbb\xf9\xfe\xf7\xaf\x39\x7b\xe4\x66\xbb\xcf\xb7\x9b\x41\xcf\x62\xce\xc2\xa3\x1d\x57\x72\x86\x14\x67\x03\xe6\xc4\xd2\xcf\xd5\x28\x24\xbd\x62\x49\x47\xe7\x34\xd1\x6f\x2a\x8a\xad\xb4\xa9\x58\x84\xa1\x90\xf6\x22\x56\x44\x1f\xb2\xd7\x5e\x9e\x12\x63\x42\x2c\xed\x30\xcf\xb4\x9a\xc7\x16\xe9\x82\x79\xf5\x0f\x5c\xff\xbf\x24\x4d\x31\xb5\xf5\x6f\xf7\xa3\x56\xb0\x77\x15\x05\x56\xb2\x66\xb2\xb1\xb4\x2f\x5f\x27\xbb\xa0\x7e\xc5\x16\xdc\xe8\x05\x67\x6a\x2e\xb6\xfa\x9c\x92\xf7\x72\x45\xfa\xc5\xd2\xe7\x93\x15\x89\x5e\x54\xdf\x5a\x11\xa6\xa3\xf1\xdb\x53\x90\x9f\x33\xcd\x42\xc9\x79\xc8\xc4\xcc\x32\x3b\x63\x23\x8e\xd8\x74\xe7\x0d\x76\x64\x51\x64\xf8\x4c\x9c\xfe\xaa\xb7\x7f\x1d\x74\xb1\x89\x92\x8d\x78\x9c\x3d\xe0\x87\x9d\x2f\x96\xef\xbb\xf3\xf6\xb4\xf8\xe1\x4f\xbb\x21\xb5\xf3\xd3\x38\x12\xdd\xa7\x9e\xee\x6c\x4f\x46\x1f\xd1\x7a\x77\x11\xa7\x27\x23\xf6\xe4\x47\x5a\x6c\xe2\xc0\x5d\x94\x97\xae\xbb\xf5\x59\xbd\x69\xb6\x9c\xf8\xbc\x30\xd6\x98\x14\x2b\xe6\x5e\xb3\x5a\x0b\x2b\x79\xc7\x86\xb2\x0e\x9c\xbd\x09\xb6\x46\xd7\xed\x0f\x67\xb1\xec\x17\x2c\xe6\xc6\x1c\xb8\x91\x7d\x1e\x66\x82\x85\x69\x30\xb6\x69\xbd\xdd\x78\x58\x0f\xb6\xe2\x1a\x68\x84\xe7\x2c\x1c\x76\x38\x8b\x6d\x73\x62\xf5\xe9\x21\xfe\xdc\x9c\xc2\x5a\x86\x33\x9f\xff\xce\xea\xd5\x7e\x15\xc6\xbc\x1c\xeb\xa0\xf2\xd5\xa0\x56\xd3\x99\x2f\x86\x7f\x4f\xc6\xf7\xfe\xd4\x83\x40\x7d\x84\xf8\xf8\x7a\xa9\x83\xeb\x1f\x8d\xb9\xfc\x09\x4f\x04\xa7\xd5\xa0\x57\xa4\x49\xf9\xd9\x58\xb6\x1a\xf1\xad\x0f\x7d\x36\xc6\xad\x46\xbd\x69\x35\xd2\xfb\x05\xd1\x52\x98\x56\xe3\x64\xf6\xd1\xa8\x6f\x5b\xfd\xb2\xfd\xcf\x5a\x3d\x0e\xb5\x8a\x1b\x60\x60\xd9\xae\xe9\xbf\x13\xa2\x6f\x45\x60\x65\x0d\x4c\xa8\xbf\x4c\x88\x6e\xc5\x09\xb0\x45\x80\x52\xde\x53\x50\x53\x84\x06\xb6\x57\x2e\x01\x4a\x89\x74\x7b\xe5\x12\xa0\x94\x00\xa5\xbc\x52\xee\x43\x2c\x01\x4a\x09\x50\x2a\x05\x84\x58\x01\x94\x0a\xa0\x54\x0a\x08\xb1\x02\x28\x15\x40\xa9\xee\x29\xa8\x29\xb6\xc3\xa9\x00\x4a\x05\x50\x30\x02\x6e\x18\x06\x28\x18\xa0\x60\x05\x84\x88\x01\x0a\x06\x28\xf8\x46\x41\xdf\x8a\x6d\x0a\x06\x28\xf8\x2b\xe5\x16\xa2\x03\x50\x1c\x80\xe2\xdc\x53\x50\x53\x6c\xaf\xec\x00\x14\x07\xa0\x38\x08\xb8\x61\x0e\x40\x71\x00\x8a\x33\x01\x1e\x67\x02\x50\x08\x40\x21\x0a\x08\x91\x00\x14\x02\x50\x88\x06\x42\x24\x00\x85\x00\x14\x72\x4f\x41\x4d\xb1\xbd\x45\x0a\x50\x28\x40\xa1\x08\xb8\x61\x14\xa0\x50\x80\x42\x35\x10\x22\x05\x28\x14\xa0\xd0\x09\xf0\x38\x53\x80\xe2\x02\x14\x57\x01\x21\xba\x00\xc5\x05\x28\xee\x3d\x05\x35\x45\x60\x65\x80\xe2\x02\x14\x17\x01\x37\xcc\x05\x28\x1e\x40\xf1\xa0\x17\x8b\x07\x50\x3c\x80\xe2\x41\x2f\x16\x0f\xa0\x78\x00\xc5\x83\x5e\x2c\x1e\x40\xf1\x3e\x29\x99\x7c\xb6\x8c\xb1\xf6\x97\x9a\x40\x61\xc8\xe9\x44\x9f\xf0\x60\x3a\x7e\x29\xe3\xff\x56\x71\x94\x1f\x8e\x87\xa5\xca\xd4\xa8\xf8\x69\xf6\x19\x5d\xe3\x79\xf2\x90\xd2\xd5\x2e\x79\xad\xb6\x9d\xfd\xb3\xed\xe0\xf9\xcf\xe0\x91\x4d\x5c\x92\x27\xdd\x29\xf2\x66\x51\x67\xd7\x39\x9f\x7b\x62\xbd\x4c\xb6\xf9\xe2\x8d\x3f\x8f\x23\xba\x54\xa8\xf9\x8c\x94\x5a\xb4\x3f\x2d\xff\x0f\x00\x00\xff\xff\xff\x98\x35\x04\x77\x0a\x00\x00")

func test_certsNodeCrtBytes() ([]byte, error) {
	return bindataRead(
		_test_certsNodeCrt,
		"test_certs/node.crt",
	)
}

func test_certsNodeCrt() (*asset, error) {
	bytes, err := test_certsNodeCrtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test_certs/node.crt", size: 2679, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _test_certsNodeKey = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\xd1\x41\xb7\x72\x40\x00\xc6\xf1\xbd\x4f\xd1\xde\x79\x8f\x92\x98\x96\x63\x4c\xe8\x46\x43\xa2\xb1\x93\xe9\x4e\x84\x42\x43\x7c\xfa\xf7\xdc\xbb\xbe\xcf\xf6\x39\xe7\xbf\xf9\xfd\xfb\x99\x89\x6d\xd7\x5f\x84\x27\xb8\x20\xa1\x1b\xc3\x08\x2f\xbe\x30\xfd\x7d\x24\xcf\x75\xcd\x23\x77\x4d\x08\xf7\x26\x24\xcd\xc6\x6e\xd4\xb8\xcc\x97\x83\xba\xda\x5d\xb7\x6d\x6d\x83\x25\x0f\x93\xc2\x3a\x9c\x8c\xa5\xd1\x20\x2b\x7b\x37\xc3\x43\xc4\x32\x22\x20\x58\x8f\x91\xf8\x92\x42\x0b\xa4\x5b\x5f\x79\x08\xee\xf7\xac\xf9\x58\x22\xe5\x56\xab\xca\xb4\x2c\xce\xae\x16\x4f\x1f\xa0\x3d\x10\x1c\x31\x84\xc1\x1e\x62\x9b\xec\xd4\x1b\x4c\xb3\x40\x1e\x8e\xfb\xea\x69\x32\x81\xa5\x88\x9b\x2c\xb8\x78\xd3\xe6\x73\xaf\xd9\x3c\x2b\x1b\x98\x88\x3d\xbb\x64\xe4\xae\x16\x63\xbc\x95\xd1\x49\x47\x69\x92\x65\xf5\x31\x41\xdf\xe2\x76\x78\x15\xa1\xac\x14\xfa\x81\xd7\x07\x79\x77\x75\x3a\x29\x0d\xdc\x3b\x24\x22\xd6\x70\x6e\xc6\x17\xad\x4e\x56\x5d\x6f\xdc\x4e\xa7\xb4\xa1\x5a\xee\xab\x37\x45\x83\x57\xbf\xf5\xeb\x0b\x3e\xb4\xcf\xd4\xbb\xc2\x02\x43\xa5\x1a\x00\x0a\x78\x72\xf5\xe4\x5c\xaa\x8d\xf5\xbd\x9b\x9a\xbb\x52\xc5\xcb\x9e\x4f\xad\x9c\x55\xc0\x4f\x2e\x03\xc5\xc1\x58\x17\x3d\x72\x77\x29\x9d\xce\xce\xea\xbb\x04\xa2\x27\xcf\x08\x9d\x01\xfa\x58\x0e\x43\xf3\x6d\xe5\xaa\xce\x5a\x62\x21\xdf\x8e\x4e\x5b\x9a\x6f\x7f\xfe\xa9\x1b\xf6\x51\xa7\x64\x53\x7b\xc7\xa1\x7a\x05\x63\xee\x45\x15\xcf\xe9\xa3\x04\x0c\xeb\xb6\xdf\x6d\xdf\xeb\xac\x13\x96\xe5\xa4\x4f\x80\x5c\x84\x14\x09\x24\xa5\x0c\xfa\x80\x32\x86\x0c\x85\x0e\xd6\xe6\x35\x11\x12\xab\x33\xd2\x23\x45\x27\xef\xb1\xf3\x76\x72\x44\xd7\x6f\xe9\x17\x0d\xfb\xd6\xdf\x98\xff\x03\x00\x00\xff\xff\x94\xf9\x2a\xd6\xed\x01\x00\x00")

func test_certsNodeKeyBytes() ([]byte, error) {
	return bindataRead(
		_test_certsNodeKey,
		"test_certs/node.key",
	)
}

func test_certsNodeKey() (*asset, error) {
	bytes, err := test_certsNodeKeyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test_certs/node.key", size: 493, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
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
	"test_certs/ca.crt":   test_certsCaCrt,
	"test_certs/ca.key":   test_certsCaKey,
	"test_certs/node.crt": test_certsNodeCrt,
	"test_certs/node.key": test_certsNodeKey,
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
	"test_certs": {nil, map[string]*bintree{
		"ca.crt":   {test_certsCaCrt, map[string]*bintree{}},
		"ca.key":   {test_certsCaKey, map[string]*bintree{}},
		"node.crt": {test_certsNodeCrt, map[string]*bintree{}},
		"node.key": {test_certsNodeKey, map[string]*bintree{}},
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
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
		err = RestoreAssets(dir, path.Join(name, child))
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
