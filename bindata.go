// Code generated by go-bindata.
// sources:
// data/currencies.json
// DO NOT EDIT!

package main

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

var _dataCurrenciesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x5b\x4b\x6f\xe3\xc8\xb5\xde\xcf\xaf\x28\x18\x77\x79\x71\xef\x7e\x76\xa4\x24\x4b\xb6\x24\x4a\x4d\x52\x76\x5b\xbb\x43\xb1\x4c\x56\x8b\xac\xd2\x14\x49\x2b\x52\x30\x40\xe0\x2c\x83\x6c\x82\xec\xb2\x08\x82\x20\x41\xde\xc8\x63\x82\x4c\x90\x19\x04\xc8\x6b\xed\xd9\x65\x99\xf4\x1f\xc8\x5f\x08\x0e\x25\x0f\xba\xaa\xdc\x47\x36\x67\xd1\xed\x76\xdb\xf5\x7d\xf5\x3c\xe7\x3b\x0f\x7d\xfd\x03\xc6\xce\x34\xaf\x9a\xa2\xae\xce\x3e\x64\xf8\x2d\x63\x67\xde\x64\xf2\xe5\x37\x8c\x9d\xad\x1a\xad\xb9\x5c\xed\x02\x28\xf9\xd9\x87\xec\xcc\x2b\x12\x90\x02\x24\x2b\xf8\xfa\xec\x7f\xed\xdf\x8a\x76\x65\xa2\x0a\xfc\xbd\xc9\xbb\x3f\x16\x69\x3b\x74\x32\x39\x6b\xff\xe7\xe3\xc3\x0f\xce\x5e\xf7\xfa\x14\xd5\x00\xaa\x9a\xf5\x40\x8b\x24\xe1\x20\x59\xaa\x8a\x02\x34\xc5\xf9\x3f\x36\x23\x12\x18\x8c\x83\x45\x48\x32\x36\x5a\x6d\x90\x0b\xff\x41\x31\xbd\xbd\xff\x85\xcd\x85\xd0\x06\x97\xef\x93\xab\xf3\x41\x27\x90\x8a\x6e\x0b\x43\x6c\x93\x2c\x0e\x48\xb2\xbc\xa9\x41\xf2\x8a\x33\x99\x35\x45\xad\x9b\xd2\x41\x8c\x03\x0b\x31\xa0\xa7\xaf\x1b\xc9\x45\xa7\xb9\x07\xd6\xdc\x5f\x7b\xe7\x14\x53\x8f\xcb\x5a\x43\xc1\xbc\x5b\x2d\x56\x20\x59\xef\xdc\x63\xb7\x1a\xe4\xca\x39\x6d\xef\xdc\x04\xee\x2d\xe6\x24\x70\x93\x80\x64\x1b\x5e\x91\x47\xed\xcc\x1f\x51\x0d\x9a\x45\x44\xee\xd4\x42\x8a\x9a\xa7\x2c\xaa\xa1\xe6\x55\x97\x0d\x43\x7c\x83\xf0\x7c\x4c\xae\xeb\x1c\x8a\x75\x01\x32\x65\x17\x15\x7e\xa9\xd8\x46\x35\x32\xa5\x28\xff\xf2\x43\x9b\x13\x29\x0c\xce\xe1\x05\xc9\x39\x14\x89\x86\xa2\x06\xdd\x85\x0c\xb1\x0d\xb2\xd1\x82\xbc\x11\xa3\x46\x66\xa0\xf1\xe9\xdc\x2a\x2d\x64\x4d\xb1\x9d\xd7\x36\x1b\x82\x1b\x6c\x17\x21\x69\x14\x2e\xf4\xc1\xe0\x69\x01\x05\xc5\xf4\x9f\x3f\x7c\x66\x53\x21\xb2\x41\x75\x39\x25\xaf\xca\x25\x94\xd0\xde\xf1\xd3\xb7\xe4\xd2\xb9\x26\x88\x6d\x90\x79\x0b\x92\xcc\x6b\x2a\x7c\x57\x1d\x2d\x10\x82\x1b\x6c\x13\x6f\x4c\xb1\x4d\x40\xb1\xb5\xd8\xd0\x46\xf5\x97\x36\x0b\x82\x9a\x2c\x37\xe4\x9a\x26\x22\xd9\xe1\x7a\x84\x7c\x77\x39\x47\xac\x1b\x6b\xc6\xd3\x31\x89\x35\x85\x15\x4f\x55\x7b\xf4\x29\x97\xf4\xf6\x3c\x7c\xf2\xf0\xfb\x87\xcf\x6d\x46\x24\x30\x2d\xdd\x8c\xbc\xd7\xd7\xbc\xaa\x9f\x61\xe6\x66\xd6\xfd\x0d\x96\xe4\x3a\x02\xbe\x65\x4b\x0e\xad\x45\xe8\x70\xd0\x88\x6e\xd0\xcd\xa6\xe4\x73\x99\x95\x20\x45\x97\xc7\x82\xb8\x06\xd1\x7c\x48\xde\xa8\x39\x6c\x1a\x60\xb8\xba\x61\x23\x24\x7a\xed\xb5\x90\x60\xa3\x22\x88\x81\x1a\x5e\x93\x67\x10\x6e\x41\xa6\x68\x59\x9e\xda\x7a\x1c\x6b\x80\x5d\x47\x31\x05\x16\x41\xa9\x40\xb2\x1a\x0a\x67\x5e\x38\xd2\x9c\x17\xed\x45\x22\xae\x13\xe1\x5e\xed\x27\xae\xe2\x77\x1f\x3e\x7d\xf8\xfc\xff\x9c\xa9\xdb\x4e\x24\x1a\x90\xbb\x1b\x6d\x79\x2a\xaa\x9c\xad\xb5\x7a\x77\x53\x5d\xbe\xb5\xf3\xd0\x10\xd9\xa0\x8a\x97\x11\x45\x15\x83\xdc\x1f\x4c\x6c\x95\x8b\xa2\x10\x32\xa3\xf8\xe2\x28\xb7\x09\x11\xdf\xb4\x7c\xb4\x99\xf5\x74\xc9\x0f\xef\x5a\x83\x23\x84\x3c\xdb\x8c\xfa\xf4\xc1\xf8\x90\x43\xd9\x55\xc6\xd9\x87\xe2\x7b\x53\x92\x4b\x55\x52\x00\xc3\x97\x3c\xe2\x7a\xcf\x33\x75\x27\x24\xb0\xb5\x92\x77\x5c\xd7\x22\x11\x85\x04\x56\x82\x5e\x93\x27\x36\x9e\x3a\xf3\xf0\xa6\x96\x72\xba\x1a\x90\xca\x09\x36\x9c\x5d\x71\x9d\xe2\xab\xe3\xd5\xaa\x49\x95\xa3\x92\xae\x06\x16\x64\x70\x43\x42\xe6\xa2\xd5\xa7\x9a\xcb\x52\xc8\x44\x90\x0a\xe2\x47\x0e\x5b\x70\x63\xb1\x85\x3d\x92\x4d\x55\x35\xb0\xb0\x35\xb4\x2b\x55\x28\x49\xfb\xa5\x1f\x38\x7c\x61\xcf\xe2\x5b\x92\xaf\xa9\xb7\xe7\xab\x9c\xad\x95\x6e\xe8\xc7\x34\xfe\xdb\xb7\x1d\xaa\xa5\xf5\x9a\x06\x21\x29\xf5\x07\x5a\xd4\x1a\xcf\x45\xc2\xfa\xd6\x31\x3b\x38\xd8\xd4\x75\x03\x32\xdc\x1b\x72\xa5\xb3\x36\xdc\x03\x2d\x1c\xdd\x36\xb0\x02\xba\x51\x3c\x24\x75\x1b\x88\x1a\xb1\x32\xd5\xe8\x94\x3b\xba\x2c\x1e\x5a\xba\x2c\xa0\x75\x99\x6c\xa3\x27\xdd\x6c\x38\xa7\x8f\xef\x8f\x8e\x2e\x0b\x6c\x5d\x36\xa3\x75\x99\xd2\xe9\xc1\x42\x3d\xa9\x2c\x70\xb4\x01\x37\x0e\xaf\x49\xdb\xaa\x9a\x3a\x67\x63\xd5\x1e\xd3\xf6\xd4\xe5\xfb\xa9\x4d\x87\xe8\xa6\x28\xf2\x49\x6d\x3e\xe1\xc9\x21\xf6\x2b\x84\x26\x6f\x9f\xab\xcc\x11\xd9\xd4\x4c\xd7\xe4\x3d\x9f\x42\x01\x5b\xdc\xa7\xf5\x16\x56\xb9\x73\xfb\x70\xb4\x09\x17\xce\x68\xb8\x46\x8b\xfa\xb0\xf3\xaa\xc9\x1a\xb1\x73\x21\xc3\x99\x05\xb9\x24\x9f\xc7\x54\xed\xa1\x4c\xda\x87\x5f\xf2\x5a\xac\xde\x95\x29\x47\xc4\xa5\xf5\x44\xbc\x80\xbc\xd5\x01\xaf\x73\xae\x0f\x91\x96\x27\x6b\x51\x14\x78\xac\x59\x53\xa4\x9c\x3c\xd9\x7f\x7c\xc7\x71\x3d\x81\xf5\x02\xe6\x03\x72\x2d\x73\xae\x9b\x3b\xdc\x1b\xd9\xf0\x3b\xc5\x2a\x45\x4a\xae\xe8\xff\x1d\x51\x80\xf8\x06\xe1\x2b\x8f\x7c\x72\xaf\xa0\x06\x8d\xe2\x6e\xf7\x62\x75\x87\xc8\xa6\xfe\x88\x69\xbd\x03\x8a\xc5\xaa\xe4\xad\xb3\x9b\x6b\x21\x57\x62\xc3\x59\xaa\x12\xed\x5c\x01\x44\x32\xa1\xe9\x14\x56\x24\xb8\xd6\xc0\x26\x5c\xb5\xaa\xb1\xc0\xaf\x0e\xa6\x9d\xab\x8a\x66\xa4\x86\x89\x54\x09\x85\x78\x96\x80\x89\x1c\xae\x99\x25\x5f\xa2\x3e\x79\xe1\xa2\x26\x3d\xbc\x67\x2b\xd4\x3e\xa2\xf5\xad\x4b\x14\xdd\x90\xd6\x21\xda\xb5\x91\x74\x87\xb0\x1d\x81\xcd\x87\x32\xf3\x48\xd9\x25\x33\x55\x1c\x6c\x83\xdc\x3b\xa7\x88\x83\x4d\xb4\x6b\x72\x17\x3c\xdd\xa6\x6f\x6e\x0b\xa5\xc5\x4b\xdf\xd9\xb5\xb5\x45\xfe\xe8\x94\xc4\xd3\x20\xa4\x78\xda\xfa\xe3\x60\x13\x8d\x8e\xc7\x7c\x5e\x88\x3d\x7f\x86\x5c\xf4\x97\xae\x60\xb4\x83\x31\xff\x9a\x3c\x5c\x5f\xd5\xd5\x16\x24\xb0\x4d\x53\x90\xa6\x7f\xee\x30\x5d\x5b\x87\xeb\x5f\x90\x71\x93\xdf\xe8\x46\xa6\xe2\xe9\xb8\x09\xc7\x9a\x0e\x92\x0e\xe3\x7b\xb0\x2b\x41\x7e\x99\xbf\xea\xa0\xac\xc7\x76\xb0\xdf\x9b\xd1\xb9\x40\x55\xa8\xb2\x0d\xb0\x5e\x9e\x0f\x9c\x59\x3b\xd5\x1f\x93\x3e\xb2\x0f\xf2\x31\xb0\x22\x75\x8b\x1b\x58\x21\xb0\x29\xde\xe2\x57\xa4\x78\x6b\xa0\xe6\x25\x46\xfe\xec\xa3\x86\xd7\x7b\xda\x68\xbf\x72\xe4\x5d\xfc\xca\x92\x77\x01\x69\x57\x47\x4a\xa6\x8d\x6e\x2d\x6a\xb9\x39\x21\x35\x26\x8e\xfa\x0b\x2c\x83\x7b\xd1\x3f\xa5\xfe\x94\xe4\xd5\x51\x01\x0a\xc8\x29\xb6\x70\xe3\x28\xc0\xbe\xe5\x8e\x2e\x26\xa4\x7d\xbf\xa8\x34\xf0\x42\x30\xc9\xb7\xac\xca\xf9\x47\x9c\xdc\xca\xb7\xf7\x3f\x73\x08\x27\x96\x91\x1f\x2f\xc9\xd4\xc1\x18\xf6\xb0\xce\x2b\x14\x3f\xac\xe6\x32\x23\xaf\xca\xc3\x9f\x1e\x7e\xeb\xbc\x80\xa5\x95\x60\x18\x5f\x93\x6f\x6e\xdc\x6c\x51\x9e\x3f\x6d\xe3\x70\xac\x29\x39\x23\xf2\x2a\x4c\x78\xa5\xea\x5c\xb1\x42\xd5\x4e\xd4\x80\x43\x4d\xc5\x76\x43\x9e\x34\x6a\xca\xdd\xe1\xa0\x85\xcc\x32\x41\xa6\x7b\x43\x27\xb6\x45\x74\x93\x8e\xae\x01\x1d\x34\xe7\x33\x23\x8b\x9f\x38\x6c\x76\x15\x68\x1a\x90\xc7\x3c\x55\xe8\x14\x91\xad\x6e\x32\x2d\xc8\x92\xda\xdb\xfb\x5f\x39\x74\x81\x75\xc8\xd3\x29\xad\xcf\x77\x20\x4b\x60\xeb\x1d\x38\x69\x71\x1c\x69\x40\x05\x43\x52\x7c\x06\x22\xe3\xad\x72\x90\x70\xe2\xad\xbf\xbd\xff\xb1\x93\x54\x1c\x5a\xc2\x73\xee\xf9\x74\xae\x4f\x42\x79\x88\x03\x12\x28\x12\x45\xf2\xf9\x4f\x08\x5d\xcf\xb7\xf8\x46\xa4\x3b\x98\xe7\xa2\x10\x9b\x8d\x90\xfc\xa4\x3f\x78\x7b\xff\x1b\x87\x6d\x64\x79\x84\x70\x46\x6e\x65\xa8\xca\xc7\x9a\x6a\x43\x51\x15\xdc\x79\x4c\x88\x6c\xaa\x3d\x5a\xc1\x47\xd0\xa4\xdd\x04\x7c\xe4\x08\xf8\x21\x2d\xe0\x85\xcc\x60\xa3\xf4\x73\x94\x8e\xe3\x53\x11\xdb\x20\x5b\x9e\x58\x56\x1b\x51\x3f\xa6\xce\x35\xd0\x52\x36\xb4\xe9\x96\xce\xda\x42\x7a\x6d\x8d\x16\x12\x4a\xd4\xe0\x5d\x16\x17\x5a\x8b\x8b\x69\xcb\x1c\xf0\x2d\x8b\x41\x6c\x9f\x95\x63\x0c\x62\x87\x2f\xb6\x8d\x77\x4c\x6b\xa1\x39\x80\xcc\x1c\x7d\x1e\xdb\x32\xe7\x6a\x40\x0a\xc2\x2b\x2e\xf9\xbe\xe1\x28\x3e\x12\x55\x88\x3b\xd7\xa9\x20\x80\x29\x9c\x68\xe1\xec\x15\x47\x8b\xf3\xa4\x8b\xea\xdb\xd2\xd8\x0b\x49\x8f\xee\xe9\x8c\xcb\xfa\x39\x2f\xdc\xad\x7d\x85\x76\xbe\x99\xce\x39\x78\x7b\xae\x13\x10\x6f\xd0\x97\x97\x20\x81\xf4\x60\x0f\x9f\x3d\xfc\xda\x2d\x25\x79\x76\x52\xc2\xa7\x9d\xa6\xcf\x0b\xd0\xcd\x51\x1e\x25\x05\xe9\xc6\x36\x8e\xb9\xf4\x6d\x9f\xe9\xcf\x48\xf3\xec\xe3\x01\x8b\xc7\x93\x16\x20\xe9\xfd\x4c\x1c\xbe\x99\x65\x9e\x7d\xda\xf7\xf8\x4d\x71\x2c\x00\x17\xfc\xee\xa5\xc2\xc8\xb7\x7d\x4f\xcf\x3b\x11\x8c\xc8\xce\x7d\x1a\x08\x6d\x72\x4d\xe8\x30\x24\x17\x6d\xf6\xe8\xe5\x41\xc8\xc4\x7a\x9d\xbd\x3e\xdd\x54\xd1\xea\x0e\x34\x60\x4f\x06\x6c\x38\xda\x7c\x9a\xb4\xc9\xe8\xab\x52\xc8\xd6\xf6\x9e\x9a\x79\xd8\x77\xe6\xde\xb7\x2d\xcb\xf9\x25\x79\x20\xe7\xe2\x4d\xc7\xe3\x40\x60\x83\x69\x48\x17\x8a\x86\x70\x88\x09\x53\x28\xa0\x72\xd3\xdf\x76\xa5\x68\x48\x47\xb5\xc3\x66\x77\xc8\xdb\x74\x98\xf9\xd0\x8e\x67\x2f\x22\x52\xe8\x5d\xac\xd0\xf2\xa6\x62\xc5\xd6\xfa\xaf\xbf\x7b\x69\x05\x0f\xc1\x4d\xb6\x57\xe4\xca\x2e\x34\x7c\xf4\x9e\xc8\x01\x47\x1a\x50\x97\x73\xb2\x0e\x74\x09\x9b\xc3\x26\xed\xe8\xec\xa9\x5b\x03\x42\x60\x83\x69\x3c\x27\xb3\xf0\x81\xd2\x5f\x29\x0b\x3f\xb7\xb3\xf0\x57\x74\x48\x04\xf5\xdd\xa1\x90\x52\x57\x64\x64\x5c\x39\x01\xd3\x95\x15\x30\xf5\x46\xe4\xdb\x8e\xb6\xa2\xaa\xec\x77\xed\x12\x9d\x6b\xc7\xee\x23\xb0\xc1\x34\x1d\x92\x19\x3d\x0c\xcd\x32\xa8\x76\x0c\x8d\xb1\xde\x39\x21\xc5\xd0\x4a\xe9\x4d\xfb\xe4\x16\x4d\x55\x91\xaa\x3b\x5b\x07\x1f\xb1\xfa\x76\xd4\x48\x5b\xed\xa9\xd2\x6a\xd5\xb6\xd2\x08\x9d\xbb\x55\xde\xa9\x6d\x98\x83\x39\xe9\x4f\x03\xbe\x81\xd6\x5c\x76\x09\x0a\x11\xdb\x24\xbb\x20\xcb\x1e\x81\x58\x81\x86\xac\x69\x8b\x93\x3a\x55\x09\xf9\x7c\x7b\x6e\xcb\xc6\x85\x55\x13\x99\x8f\xc9\xc5\xcd\x61\x2d\x0e\x99\x86\x2e\xab\x43\x70\x93\xed\x86\x4c\xdc\xce\x0f\x6b\xdb\xb5\x45\x12\xd0\x20\xc9\x5a\xef\xd0\x79\x0e\x08\x6f\xf0\x45\x74\x2c\x17\x81\x90\x35\x1b\xf1\x82\x4b\xe8\x94\xe6\xb6\xa3\xb9\xa8\x47\xc7\x22\x7c\xb7\xca\x79\x51\x28\x51\x75\xda\x4f\x84\x37\xf9\xe8\x46\xd2\x48\x15\xaa\x54\x5f\x29\x5b\x1a\xd9\xed\xa4\x13\xfa\xc2\x44\x5a\xb0\x09\xc8\x75\xc7\x24\xc9\xc4\xbe\x31\xf1\x88\x94\x97\x71\x0e\x82\x25\x90\x93\xba\xf9\x5f\x9f\xfe\xd9\x09\x59\x46\x96\xac\x8c\x43\xd2\xf7\xc4\x8d\x5e\x8b\x2a\x6f\x93\x7c\x66\xad\xf4\x88\x17\x5a\x1e\xc6\x1b\xd0\x9d\x9f\xde\xe0\x3d\xe6\x07\x07\x1a\x48\x57\x8b\x2b\x32\x98\x02\xd9\x40\xdd\xb0\x3b\xa8\x1d\xcb\x88\x43\x0d\xac\x9b\x01\x79\x78\x37\xbc\xe4\xdd\x3a\xb4\x10\xd8\x5c\xff\x39\x1d\xf7\xdc\x66\x39\x48\x06\xed\x17\xf2\x95\xff\xf3\x5b\xce\x0e\x9d\xdb\x01\x4f\x9f\x4c\xa4\xf9\x20\xb3\x02\x52\x5e\xe5\x82\xd5\xb0\x76\xce\x0e\x87\x9b\x78\x21\xe9\x8b\x7c\x0d\x7b\xd1\x26\xe6\x34\xa7\x37\x2a\x74\xab\x34\xa1\xe5\xab\xc6\x23\xf2\x44\x7a\x50\x26\xea\xd0\xcc\x20\xe8\xc4\xf2\xbf\xbf\xff\x3d\x47\x87\x8c\xac\x43\x19\x4f\x4f\x28\xff\x52\x1d\x7a\x67\x9f\x12\xfe\x38\xd8\x40\x1b\x85\x74\x63\x8b\x56\xd0\x26\x4b\xd7\x27\x3a\x5b\xd6\xd2\xc9\xf8\x87\x96\xc8\xec\x5f\x92\x13\xef\xbf\x11\x89\x6a\xea\xf7\x4e\x1d\x87\x1b\x78\x83\x21\xe9\x17\x06\xd9\x6e\x53\x77\x2c\x7d\x22\xb4\xc9\x15\x93\x16\x6c\x50\xe7\x42\x6d\xda\x08\x59\x68\x47\x23\xe3\x60\x03\xed\xf5\x9c\x3e\xc2\xf3\xf9\x7b\x5a\x43\xe7\xd6\x16\x0c\x47\x64\x0e\x64\x98\x83\x04\x9c\xd5\x8a\xa7\x6e\x78\x33\xb2\xb2\x1c\xc3\x80\x9c\xd5\x63\x17\xe6\x93\x33\xc3\xb1\xe6\xbd\xa2\x9b\x6f\x47\x4a\x66\x6c\x8c\x7f\x75\x70\x67\x23\xa7\xef\x96\xae\x25\x45\x1b\xbe\x12\x50\xb0\xbe\x86\xad\x90\x19\x0b\x45\x96\xd7\x8e\xee\x78\x6d\x97\x8c\xc6\x03\x72\x73\xc7\x5c\xee\x9e\xd9\xd3\x38\x76\x7b\x1a\x11\xdc\x64\x1b\xd2\x6c\x3b\x9d\xed\xf6\x07\x15\x57\xa9\xf2\xc5\xd5\xa2\xa1\x45\x37\xa1\x33\xa0\x13\x91\x1c\x73\x71\x2f\x3f\x9e\x89\x9d\xff\x9c\xd2\xc9\x85\x29\xac\x8e\xfd\x0e\x50\xc3\xca\xed\x00\xb2\xd3\x07\xd3\xab\x53\xf5\xa4\xf4\x90\xad\xd2\xcd\xad\xd8\x81\x0b\x78\x65\xd7\x70\x5e\xd3\x2d\x45\xfc\x6b\xcf\x4a\x7d\x38\x1b\x81\xb8\x06\x51\x40\xc7\x34\x01\x94\x22\xe9\xb8\xe9\x81\x13\xf0\xcc\x48\xc3\x1e\x28\xbd\xe5\x6d\xe7\x5f\x87\x42\x35\x62\x1b\x64\xf3\x09\xdd\xc8\xa4\x0a\xd4\x5d\xfb\x42\xd5\x3b\x8a\x69\xff\xf7\x7b\x27\x1a\x98\x58\x7b\x18\x2e\x48\x4b\x1c\x36\xd5\xf3\xb2\xa2\x5f\x7c\xe3\x8b\x6f\x3e\x38\xa5\x1d\x44\x37\xe8\xa2\x25\xdd\x6b\xb4\x85\xbd\x60\x85\x28\x40\x66\x5c\x3a\x36\x16\x47\x1b\x70\xf1\xe5\x89\x56\xe9\x37\xe2\x31\x56\xab\x50\xf5\x3b\x88\x08\x60\x22\xd2\x7d\x56\xb1\x16\x52\xa4\x90\xb6\x7d\x56\xb1\x4a\x20\x53\xcf\xb8\x5e\xf1\x13\x55\x06\xbb\x0b\x6b\x31\x7c\x4d\x8a\xe3\xec\xd0\x68\xff\x1c\xfb\xb8\x70\xed\x23\xa2\x9b\x74\x37\x0b\x92\x4e\x37\xc7\xa0\xf3\xe4\x33\x5d\x38\x5c\x37\x0b\x4b\xad\xd3\x9f\x8d\xbb\x12\xbc\xfe\xb2\x3c\x44\xaf\xec\xed\xfd\xcf\x1d\x3d\x6f\x7f\x3e\x2e\xa6\xd9\xe2\x46\x8a\xea\xbd\x65\x91\xd8\x46\x5b\x78\x23\x72\x9f\xd6\x6d\xab\x13\x48\x96\xeb\xdd\x9d\x14\x27\x2a\xbb\x9f\x38\x7b\xe5\x8d\x2c\x3e\xba\xf9\x7f\xb1\x4f\xf8\x3b\x37\xfa\xa5\x7e\x6b\x61\xb7\xfe\xc7\x53\x32\x4c\xc0\x10\x0f\xc3\x1f\x24\xb4\x6b\x31\xc7\x0d\x9b\x5a\x81\xc2\x90\xee\xae\xf5\xb5\xa8\xd1\x78\x75\xf9\xdc\x9b\xdd\x5d\xbb\x9c\x92\x19\xcb\xe5\x31\x1d\xfd\x74\x73\x2d\x0e\x36\xd0\xfc\x98\xec\x81\xf7\x45\xbd\x52\x74\xeb\x1c\x22\xd8\xb1\x4d\xfc\xd8\xf9\xfe\x01\xfe\xf9\xf8\x83\xff\x06\x00\x00\xff\xff\x4a\x00\xeb\x13\x46\x3c\x00\x00")

func dataCurrenciesJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataCurrenciesJson,
		"data/currencies.json",
	)
}

func dataCurrenciesJson() (*asset, error) {
	bytes, err := dataCurrenciesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/currencies.json", size: 15430, mode: os.FileMode(420), modTime: time.Unix(1500893173, 0)}
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
	"data/currencies.json": dataCurrenciesJson,
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
	"data": &bintree{nil, map[string]*bintree{
		"currencies.json": &bintree{dataCurrenciesJson, map[string]*bintree{}},
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

