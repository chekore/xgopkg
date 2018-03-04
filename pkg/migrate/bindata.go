// Code generated by go-bindata.
// sources:
// README.MD
// README_zh-CN.MD
// bindata.go
// migrations/201803041120_add_package.sql
// migrations/bindata.go
// DO NOT EDIT!

package migrate

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

var _readmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x70\x49\x2c\x49\x4c\x4a\x2c\x4e\x55\x48\xc9\x4f\x2e\xcd\x4d\xcd\x2b\x01\x04\x00\x00\xff\xff\x57\x6d\xe5\x94\x13\x00\x00\x00")

func readmeMdBytes() ([]byte, error) {
	return bindataRead(
		_readmeMd,
		"README.MD",
	)
}

func readmeMd() (*asset, error) {
	bytes, err := readmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "README.MD", size: 19, mode: os.FileMode(420), modTime: time.Unix(1520139157, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _readme_zhCnMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x78\x36\x75\xc3\xb3\xde\x75\x4f\x77\x4d\x7e\x36\xad\xfd\xd9\xc2\xc5\x5c\x5c\xca\x48\x62\x2f\xf6\x37\x3e\x5f\xbe\x1b\x10\x00\x00\xff\xff\xc8\xfa\xce\xd5\x25\x00\x00\x00")

func readme_zhCnMdBytes() ([]byte, error) {
	return bindataRead(
		_readme_zhCnMd,
		"README_zh-CN.MD",
	)
}

func readme_zhCnMd() (*asset, error) {
	bytes, err := readme_zhCnMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "README_zh-CN.MD", size: 37, mode: os.FileMode(420), modTime: time.Unix(1520139141, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1520139360, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations201803041120_add_packageSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xcc\x31\x8b\xc2\x30\x18\x87\xf1\xbd\x9f\xe2\x3f\x26\x70\x43\x39\x38\xb8\xa5\x43\xda\xbe\x6a\x30\x4d\x25\xa6\x43\xa7\x26\xd6\x1a\x4b\x31\x16\xa9\xe0\xc7\x97\x22\xb8\x3f\xbf\xa7\x30\x24\x2c\xc1\x8a\x5c\x11\xdc\x2b\xcc\xdd\xec\xfb\xc9\x87\xc1\x81\x25\x80\x1b\xcf\x0e\xa7\x31\x8c\x71\x61\xbf\x29\x87\xae\x2d\x74\xa3\x14\x44\x63\xeb\x4e\xea\xc2\x50\x45\xda\xfe\xac\xe9\x3c\x85\x2e\xfa\xdb\xe0\xd0\x5f\xfd\x83\xfd\xa5\x1c\x25\x6d\x44\xa3\x3e\x64\x6d\x0e\x46\x56\xc2\xb4\xd8\x53\x0b\xb6\xbe\x79\xc2\x41\x7a\x2b\x35\x65\x32\xc6\x7b\x99\x7f\x49\xb1\x13\xe6\x48\x36\x7b\x2e\x97\xff\x77\x00\x00\x00\xff\xff\xc9\x8c\x4f\xb2\xa5\x00\x00\x00")

func migrations201803041120_add_packageSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations201803041120_add_packageSql,
		"migrations/201803041120_add_package.sql",
	)
}

func migrations201803041120_add_packageSql() (*asset, error) {
	bytes, err := migrations201803041120_add_packageSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/201803041120_add_package.sql", size: 165, mode: os.FileMode(420), modTime: time.Unix(1520135929, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsBindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x5b\x6f\xdb\x38\x16\x7e\x96\x7e\xc5\x19\x03\xed\x4a\x33\x1a\xd9\xf2\xdd\x5e\xe4\x61\xa6\x17\xa0\x8b\x6d\x67\xd1\x76\xb1\x0f\x65\x11\x50\x12\x69\x13\x95\x45\x97\x92\x13\x26\x81\xff\xfb\xe2\x90\x94\x2c\x27\x8e\x93\x29\x16\x0b\x4c\x00\xc7\x16\xc9\xf3\x9d\x2b\xbf\x43\xb1\xdf\x87\x57\x32\x67\xb0\x62\x25\x53\xb4\x66\x39\xa4\x37\xb0\x92\xbf\xa6\xa2\xcc\x69\x4d\x63\xbf\xdf\x87\x4a\xee\x54\xc6\xaa\x25\xfe\x1e\x0e\x92\xf9\x60\x34\x18\x27\xc9\x70\x70\x49\xf3\xfc\x72\x4b\xb3\x6f\x74\xc5\xe2\xea\x7b\x81\xf3\xaf\xff\x80\x0f\x7f\x7c\x86\x37\xaf\xdf\x7d\xfe\xc9\xf7\xdd\x24\x6c\xc4\x0a\xc1\x7d\x5f\x6c\xb6\x52\xd5\x10\xf8\x5e\x2f\xbd\xa9\x59\xd5\xf3\xbd\x5e\x26\x37\x5b\xc5\xaa\xaa\xbf\xba\x15\x5b\x1c\xe0\x9b\x1a\xbf\x84\xb4\xff\xfb\x42\xee\x6a\x51\xe0\x83\x34\x02\x5b\x5a\xaf\xfb\x5c\x14\x0c\x7f\xe0\x40\x55\x2b\x51\xae\xcc\x5c\x2d\x36\xac\xe7\x87\xbe\xcf\x77\x65\x06\xce\x8b\x8f\x8c\xe6\x01\xfe\x80\x2f\x5f\x51\x6d\x04\x25\xdd\x30\xb0\x62\x21\x04\xcd\x28\x53\x4a\xaa\x10\xee\x7c\x6f\x75\x6b\x9e\x60\x79\x01\x68\x55\xfc\x81\x5d\x23\x08\x53\x81\x31\x1b\x9f\x7f\xdf\x71\xce\x94\x81\x0d\x43\xdf\x13\xdc\x08\xfc\x74\x01\xa5\x28\x10\xc2\x53\xac\xde\xa9\x12\x1f\x23\xe0\x9b\x3a\x7e\x83\xe8\x3c\xe8\x21\x10\xbc\xf8\xbe\x84\x17\x57\x3d\x6b\x89\xd1\x15\xfa\xde\xde\xf7\xbd\x2b\xaa\x20\xdd\x71\xb0\x7a\xac\x12\xdf\xbb\xb4\xe6\x5c\x80\x90\xf1\x2b\xb9\xbd\x09\x5e\xa6\x3b\x1e\xc1\xea\x36\xf4\xbd\xac\x78\xd3\x58\x1a\xbf\x2a\x64\xc5\x82\xd0\xff\x5f\xd9\x83\x30\x16\xff\x11\x20\xa6\x94\xb5\xdb\x0d\xa6\x3b\x1e\xff\x8e\xa6\x07\x61\x84\x2b\xfc\xbd\xef\xd7\x37\x5b\x06\xb4\xaa\x58\x8d\x21\xdf\x65\x35\xa2\x18\xff\x5c\x3e\x7c\x4f\x94\x5c\x02\xc8\x2a\x7e\x2b\x0a\xf6\xae\xe4\xb2\x95\x73\x29\x6c\xc6\x3b\x08\x26\x87\x00\x2e\x8d\xbe\x57\x89\x5b\xf3\x2c\xca\x7a\x3a\xf6\xbd\x0d\x96\x35\xb4\xa0\xef\x65\xce\xcc\xe0\x67\xb1\x61\x80\x65\x12\xe3\x2f\xd4\x63\x4a\x25\xe0\xe2\xbe\xae\x10\x3e\xd0\x0d\x0b\x42\xa7\x01\x75\x3a\x2f\xb9\x88\x51\xbb\xbf\x3f\x23\xfb\x49\xdc\xa2\xac\xb1\xe6\x58\x14\x0d\x3d\x2b\x8a\xb6\x06\x61\xd7\xf2\x63\x00\x74\xed\x29\x00\x74\x2e\x08\x0f\x8e\x3e\x40\x70\xde\x3f\x0e\xf2\xae\x7a\x2d\x54\x10\x42\x2a\x65\xd1\x95\xa6\x45\xf5\x84\xe7\x37\x95\x75\x9c\x29\x4e\x33\x76\xb7\xef\x48\xbb\x92\xc0\x2a\xbf\xbc\x7c\x8c\x4c\x3e\x7d\x2f\xe0\xc2\xd5\x46\xd0\x23\x3a\xe1\x44\xcf\x53\xa2\x07\x73\xa2\x07\x83\xd3\x1f\xce\x89\x1e\x65\x44\x67\x19\xd1\xa3\xc4\xae\xcf\x86\x44\x8f\x06\x44\x27\x73\xa2\xe7\x33\xa2\x79\x42\x74\x9a\x13\xbd\xe0\x44\x33\x9c\xe3\x44\x0f\xa7\x44\xcf\x06\x44\x8f\x47\x44\x8f\x16\x44\x8f\xe6\x44\xa7\x73\xa2\xe9\xc4\x8e\xe5\x94\xe8\x94\x11\x3d\xa5\x16\x6b\x9c\x13\x3d\x9c\x10\x4d\xa7\x76\x9e\xce\x2c\x46\x3e\x25\x3a\xa1\x44\x8f\x53\xab\x3f\x99\x12\x4d\x17\x44\xb3\x01\xd1\xd9\x8c\xe8\x05\xae\x1b\x5a\x6c\xd4\x9b\x72\x2b\x8b\x98\xc3\x31\xd1\x43\xb4\x1d\xed\xa6\x44\x4f\x32\xa2\x93\x84\xe8\x3c\x23\x7a\x98\x5a\x9f\xf2\x9c\x68\x96\x11\xcd\xf1\x79\x61\xfd\x31\xeb\x13\x6b\xcf\x1c\xfd\xc4\x39\x4e\xf4\x80\x39\x6c\x9c\xcf\x88\x9e\x25\x44\x4f\x13\xab\x73\xe8\x64\x29\x23\x7a\x98\x13\x3d\x1b\x13\x4d\x47\x44\x27\x63\xa2\xc7\x63\xa2\xa7\x23\xa2\x59\x4a\xf4\x98\x11\xcd\xa8\x8d\xe1\x04\xfd\x9e\xd8\x58\x70\xc4\xce\x88\x66\x0b\x1b\xef\xf9\x84\xe8\x21\x23\x9a\x53\xa2\xf3\xd4\xfa\x9b\x0f\x88\x9e\x70\xa2\x79\x4e\xf4\x7c\x64\xbf\x31\x9e\x49\x66\x6d\x9d\xe6\x56\x17\xea\x1d\x61\x6c\xc7\x76\x0c\xed\x1e\x4f\x89\x9e\x4c\xad\xde\x74\x4c\x74\x3e\x27\x7a\x32\x22\x7a\x90\x12\x9d\x4e\x6d\x2e\x66\x0b\x3b\x3f\x4e\x88\x9e\x51\x1b\xa3\x11\xe2\x4e\x88\x1e\x0d\x89\xce\x30\xa7\x29\xd1\x8b\x05\xd1\x33\x4e\xf4\x78\x41\x74\x8a\x39\x41\xdf\x30\x6f\x98\x03\xb7\x06\x6d\xc7\xdc\x60\xfd\xcc\x66\x0f\x6b\x0a\x3f\x26\xde\x19\xd1\x63\xcc\xdb\xd0\xfa\x72\x58\xd7\x6b\xba\xcd\xb9\x7a\x76\x9c\x78\xaa\xd7\x34\xcc\xd9\xe9\x55\xbe\xe7\x9d\xdd\x1d\x91\xef\x79\xbd\x73\xad\xb8\x17\xf9\x5e\xd8\x72\xdb\x39\x28\xb4\xe9\x67\x43\xce\x5d\x9b\x0c\x3b\xb7\x2d\xf0\x19\x8e\x3d\xd5\x70\xda\x3e\x61\x98\x7e\x79\x71\x9f\x35\xee\x90\x4f\x97\xf0\x84\x53\x80\xd4\xb9\x84\x64\x3a\x89\x00\x49\x70\xd9\xe5\xc8\x60\x3c\x1c\x84\x66\x1c\xa9\x6d\x69\xa9\xef\xdf\xa5\xd0\x41\x32\x19\x0e\x92\xd1\x64\x31\x5c\x44\x30\x08\xf7\xbe\x47\xd1\x82\x97\xc6\xeb\x3b\xe3\xea\x12\x9c\xc7\x68\xde\xd2\xfc\xdf\xb7\x99\xa1\x6d\x27\xeb\xf7\xe1\x37\xd3\xc7\x0a\x49\xf3\x0a\x68\x99\x83\x5d\x53\x41\xbd\x6e\x7a\x1c\x97\xca\x3c\xad\xc4\x15\x2b\x4d\x3f\x35\xa7\xa8\x77\x75\xbb\x96\x96\x36\xd4\x20\x78\x47\x2e\x93\xbb\x22\x87\x52\xd6\x90\x32\xe0\x72\x57\xe6\x20\x15\x4a\x1e\x4d\xa0\x66\x96\xc7\x36\xb1\xc6\x98\xe0\xa9\xc3\x4c\x46\xcb\x52\x96\x22\xa3\x05\xf6\x33\x74\xdd\x1d\x98\xe2\x8f\x6c\x5b\xd0\x8c\x05\xb6\xe9\xf7\x08\xe9\x45\xd0\xeb\xf7\x22\xf8\x35\xb1\x19\xe5\x11\xc8\x6f\xa6\x04\x5c\xbe\xbe\x1c\x83\x7d\xfd\x3b\xce\x63\xb2\x69\x5b\x2d\x1c\xab\xe1\x44\x39\x3c\x7a\x00\xb1\x21\x7d\x51\x41\x46\xcb\xbf\x61\x94\xa8\x39\x85\x1a\x07\x4e\x9c\x4a\xf0\x58\xd2\xa6\x26\x76\x79\xc3\x04\x79\x7b\xff\x69\x1d\x18\x46\x13\x5c\x87\x1a\xba\xbc\xbe\xdf\x55\xb5\x5d\x24\x2a\x28\xc4\x37\xe6\x32\x9d\xee\x6a\xd8\xd2\x52\x64\x15\x5c\xaf\x59\xe9\x46\xaf\x4d\x4a\x1a\x23\x5c\x3a\x9b\x34\x57\x62\xb3\x2d\x04\x17\xac\x82\x8a\x72\x06\xa2\x14\xb5\xa0\x85\xb8\xa5\xb5\x90\x25\x48\x0e\xab\x42\xa6\xb4\x80\x2b\xaa\x04\x4d\x0b\x56\xb9\x74\xb6\x36\x1c\xa7\xd4\x66\x14\x63\x78\x08\xf2\x61\xd9\xc9\xad\x67\x2c\x0e\x7a\xa6\xb0\x96\x6e\x71\x0f\x7e\xb1\xe7\xde\x5f\xa0\x17\x2e\x01\x1f\x99\x52\x36\x42\x41\x18\x1e\x1d\xe3\x68\xb7\xda\xcd\xb9\xeb\x5c\xc5\x9b\x3d\xfd\xff\x2e\x7b\xb4\xea\x5e\xe9\x77\xce\x8f\x7f\xb9\xfa\x37\x41\xfe\xc1\x3d\x20\x8c\xc3\xcf\xda\x02\x8d\x9a\xc7\xb6\x81\x59\x84\x9e\x55\x47\x79\x2e\xcd\x88\xec\xe4\xad\xea\xe6\xc2\x08\x04\x58\xa9\x87\xa3\xb2\x15\x59\x5e\xc0\x86\x7e\x63\x41\x33\x13\xc1\x20\x82\x82\x95\x41\x13\x50\x2c\x3c\x2c\x9d\xd2\xa5\x46\xd1\x72\xc5\xda\x70\x9b\xc0\x59\xa4\x0b\xa0\xdb\x2d\x2b\x73\x93\xac\xaa\xb1\xba\xeb\x30\x8e\x3b\x37\x5a\x79\x51\x01\x85\x1a\x37\x59\x04\x6b\x59\xe4\x68\x1c\xa3\xd9\xda\x15\x9f\x7b\xed\x95\x2a\x82\x0d\xc2\xe7\x50\x4b\x10\x75\xe5\xea\xd7\x1c\x53\x1b\x28\xf4\x64\xfb\xc5\xba\xf1\x15\x7d\x7f\xd8\x3a\xef\xfc\x27\xda\xf2\xf2\x6c\x2f\x8d\xba\x49\x78\x2d\xd4\x51\x0a\xf0\x85\xd7\xe5\x21\x65\x85\xbc\x06\x0a\x19\x53\x35\x15\x25\x8a\xe4\x42\xb1\xac\x96\xea\x06\xd8\x26\x65\x79\xce\x72\x10\xe5\x41\xee\xe1\x4b\xfd\x5b\xa9\x80\x69\xba\xd9\x16\x0c\xf7\xe3\x8d\xdc\x81\xda\x95\x9d\x55\x20\x4b\xc0\xef\x7e\x1c\xc7\x66\xdf\x9b\xc1\x4c\x96\xa8\xd2\x98\x84\x30\x5c\x16\x85\xbc\xc6\xa8\xae\x05\x53\x54\x65\xeb\x1b\x73\x4f\x80\x7f\x46\xba\x79\x00\xe0\x52\xc6\xb5\xae\x0f\x03\x62\xb3\xea\x4c\x03\xd0\x78\x5b\xae\xba\x03\x69\x33\x50\xb7\xcc\x8b\xaf\x24\x3d\x04\xee\x85\xc7\x24\xdc\x14\xd8\x5d\xcf\xe9\xc1\xad\x2c\x36\xab\xde\xbe\x1b\x51\x2b\xdb\xc7\xf1\x47\xe5\x8d\x19\x28\x6d\xd4\xdf\x97\x6f\xd0\x43\x13\x93\xc3\x78\x29\x6b\xa6\x45\x55\xdf\xc7\x6d\x48\xef\x18\x05\x57\x89\xa2\x78\xa8\xdc\xb8\xb6\xef\xee\x2d\x5c\x7f\xbf\xc3\x37\x7b\xe9\xc0\x71\x25\xbe\x24\x3a\x9e\xaa\x15\x63\x86\xbd\x70\x9f\x99\x6d\x82\x34\x34\x30\x7b\xe9\x47\xb9\xd0\xdb\xd2\x7a\xfd\x4f\x51\xd5\x5d\x91\x4f\xdb\x42\xd4\xc1\x31\xa4\x11\x42\x01\xdc\xd4\x97\x11\x6c\x0f\x9b\xba\x85\x30\x6c\x68\x2c\xbe\x00\xfc\x8a\x5f\xad\x45\x91\x2b\x56\x7e\xd9\x7e\xc5\x29\xc1\xc1\xce\x76\xc8\xf3\x47\x3a\xbb\x67\x79\x72\xdf\xdc\x64\x18\x5d\x6f\x31\xb2\xcf\xb8\x16\x39\x87\x8b\xa4\x73\xf5\x38\xb7\x1d\xf9\xd4\x10\x5c\x86\xcf\x1f\x8e\x58\xee\x68\x9d\x35\xe6\xea\xc0\x73\xea\x2a\x3a\x08\x1d\x31\x1d\xce\x74\xef\x56\x5c\xce\x3b\x77\x23\xc6\x49\xb3\xe9\x4e\xf2\x94\xef\xb5\x5a\x3b\x9c\xf6\x73\x53\x3b\xfb\x96\xf8\x0c\xec\x05\xbc\x74\x3f\xef\x4c\x94\x4e\x88\x3c\x83\xf8\x5a\x8c\xb3\x0c\x78\x12\x7c\xbf\x8f\xfc\xbd\xa5\xc6\x8f\xac\xaa\xa5\x62\x36\x39\xca\x3e\x98\x93\x85\xe5\xf3\x5d\x99\xb3\xee\x31\xa4\xa5\x45\xbb\xa1\xba\xd2\x41\x2e\xd4\xbd\x8b\x40\x7b\x3a\xb9\xf3\x3d\xdc\x85\xcf\x3d\x6a\xb9\x9c\xb8\x17\x1c\x4f\x34\x87\x8f\x56\xb4\x3d\xa8\x3c\x47\xdc\xde\xf2\xc9\x2a\x7e\xff\x2d\x17\xea\xb7\xa2\x08\x2e\x91\xc0\xff\x45\xeb\xb5\x35\xb8\xb9\xf7\x8c\x1b\x5e\x08\xc3\xe8\xe8\x2d\x68\x30\x9b\x4c\xce\x5e\x48\x1e\xab\xb2\x97\xab\xf1\x7f\x94\xa8\x19\x62\xdc\xd7\x67\x54\x44\x60\x23\x82\xce\xc5\xf6\x76\xea\x4f\x39\xf3\x6a\x8d\x2f\x63\xd5\x23\xd8\x0d\xaa\xbd\xb2\xba\xff\xfc\x1c\x45\xc7\x97\x4b\xf7\xca\xa4\xfa\x33\x75\x02\x8a\x65\x3b\x55\x89\x2b\x56\x9c\xa8\x99\xea\x5c\xd1\x64\x6e\x4b\x1d\x67\xbf\x4d\x93\xef\x61\xd7\x15\x05\x3b\xe3\xcf\xe9\xfa\xb4\x7b\xbf\xdf\x87\xd7\x42\xf9\x0d\xaf\x1a\x6d\x07\x2a\xc9\xba\x2c\x62\xe3\x7e\xc2\xf0\xb6\x78\xfe\x21\x45\xe9\x98\xde\x48\x86\xe7\x8f\xac\x26\xce\x8e\x47\x8f\x63\x6d\xef\x17\x1e\xa6\xb5\x0d\xcf\xe1\x40\xf8\xa3\xad\xa7\xbd\xb5\xec\xda\xee\x38\xb2\x6d\x9c\xb9\x50\xfb\xe8\x19\x8d\x29\x8e\x63\xf3\x41\xdb\xff\x1b\x00\x00\xff\xff\x93\x30\x58\xfc\xfc\x18\x00\x00")

func migrationsBindataGoBytes() ([]byte, error) {
	return bindataRead(
		_migrationsBindataGo,
		"migrations/bindata.go",
	)
}

func migrationsBindataGo() (*asset, error) {
	bytes, err := migrationsBindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/bindata.go", size: 6396, mode: os.FileMode(420), modTime: time.Unix(1520135997, 0)}
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
	"README.MD":                               readmeMd,
	"README_zh-CN.MD":                         readme_zhCnMd,
	"bindata.go":                              bindataGo,
	"migrations/201803041120_add_package.sql": migrations201803041120_add_packageSql,
	"migrations/bindata.go":                   migrationsBindataGo,
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
	"README.MD":       &bintree{readmeMd, map[string]*bintree{}},
	"README_zh-CN.MD": &bintree{readme_zhCnMd, map[string]*bintree{}},
	"bindata.go":      &bintree{bindataGo, map[string]*bintree{}},
	"migrations": &bintree{nil, map[string]*bintree{
		"201803041120_add_package.sql": &bintree{migrations201803041120_add_packageSql, map[string]*bintree{}},
		"bindata.go":                   &bintree{migrationsBindataGo, map[string]*bintree{}},
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
