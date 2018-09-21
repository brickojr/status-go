// Code generated by go-bindata.
// sources:
// ../config/cli/fleet-eth.beta.json
// ../config/cli/fleet-eth.staging.json
// ../config/cli/fleet-eth.test.json
// ../config/cli/les-enabled.json
// ../config/cli/mailserver-enabled.json
// ../config/status-chain-genesis.json
// DO NOT EDIT!

package static

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

var _ConfigCliFleetEthBetaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x56\x5d\xaf\x1e\xc7\x6d\xbe\xf7\xaf\x38\xd0\x75\xf1\x6a\xf8\x4d\xea\xca\x8d\x61\xd7\x70\xda\x40\x89\x63\xd7\x6e\xd1\x0b\x0e\xc9\x91\xd5\xb8\x92\x2b\x1f\x09\x8d\x8b\xfc\xf7\x62\x4f\x4e\x91\xd8\xc9\xb9\x58\x60\x31\x3b\xbb\x7c\x96\xcf\xc7\xf0\x7f\x3f\xba\xbb\x7b\xf6\xbb\x79\xd3\xf3\xd3\x87\xb7\xef\x7f\x7c\xf6\xe2\xee\xfe\xdd\xfb\xf9\x87\x6b\xf5\x93\xef\xdf\xff\x78\x3f\xef\x3e\x79\xfb\xe6\xbc\x7e\xf5\xec\xc5\xdd\xb5\xf5\xee\xee\xd9\xa7\x6f\x72\x7f\x3f\xfd\x57\x3b\xef\xee\x9e\x7d\xf6\xfd\xcc\xfd\xb3\x17\x77\xcf\xe6\xfe\xbb\xdb\x9e\xfb\x7c\xf6\xf8\xe0\x57\x6f\xdf\xde\xff\xe6\x6d\xcf\xf5\xe5\x7f\x7f\x58\xba\xbb\x7b\x36\x6f\xde\xf6\xbc\x78\xfe\x7c\x3c\xad\x16\x6d\xf1\x00\x98\xf0\xbd\x5b\x35\xab\x36\xa6\x48\x8b\xa9\xd3\x21\xd9\x48\xfb\x44\x9f\x4a\x8f\x91\x41\xe6\xd9\x52\x45\x07\x65\x81\x6f\x06\xc0\xde\xb2\xec\x64\x25\xf1\xd9\x1a\xe7\x6c\xe6\x4d\x8a\xc7\x62\x3a\xc3\x52\x35\xd2\xfb\x60\x94\xe1\xa8\x30\x83\x9a\xf3\xc7\x6c\xb7\x80\x1b\x22\xdf\x48\x5e\xd0\xe2\xc5\x8f\xb8\xff\x0a\x24\x93\x56\xe9\x51\xe3\x40\x3f\x5d\x91\x71\x2c\x62\x1d\x0c\xe6\xb5\x70\xab\x4b\x43\x91\x9d\x85\x52\xb0\x87\x51\xc0\x65\xcb\x86\xb3\x62\xad\x33\x79\xa0\xaa\x30\x15\x68\xa1\xfa\x89\x58\xd7\xb6\xb4\x46\xc1\x43\x46\x0b\xcb\x89\x24\xb1\x20\x15\x8d\x74\x02\x09\x35\xa0\x8a\xf2\x63\x30\xbe\x01\xf9\x0d\x96\xdc\x90\xe9\x29\x9c\xc6\x68\x7d\x86\x7c\x37\xd7\xb1\x12\xdf\x3b\x94\xc1\x7c\xe9\xc9\x8d\x62\x8e\x53\x34\xba\x58\xd3\x97\x10\xd9\x5a\x88\xb5\x33\xd1\x41\x48\xc7\x5b\x79\xaa\x61\x2f\x1c\x38\x7e\x0c\x4d\x7d\x30\xa4\x97\xa2\xf8\x4e\xa6\x76\xa7\xc5\xbd\x97\xfa\xe8\x41\x06\xcb\xf1\x5d\x3b\x52\x3f\x86\xc5\x37\x10\xbe\xb9\xdf\x00\x9f\x84\xc9\x14\x6c\xae\x54\x27\x25\xb3\x01\xcc\x0f\x3b\x66\x91\xa4\xcf\xde\x01\xa0\xd5\xd3\x50\x48\xc7\x4e\xe4\x81\x64\xdb\xd9\xa7\x60\xc6\xce\xea\x8e\xa9\x25\x4c\x0c\x56\x45\x6c\x88\x92\x3a\x6c\xc3\x5a\xea\x46\x47\xd9\x44\x62\x98\x98\x58\x42\x4b\x78\xc0\x92\x9b\x12\xe6\x81\x73\xc1\xdb\xd5\x53\x5e\x4f\x81\x14\x0a\xc9\xdc\xe6\x44\x07\xa6\xb6\x1a\x6c\x89\x7d\x96\x20\xd4\xc1\x85\xc8\x67\x1c\x14\x4f\x51\xa3\x9a\xf4\xf0\x0c\xb7\x28\x69\x9a\x4c\x11\x36\xd0\xc5\x75\x81\x73\x1f\x6f\xe8\x3e\xe9\x8b\x82\x29\x96\xba\x63\xa7\x62\x72\x1f\xc6\xe6\x53\x75\xb4\x2f\xa6\x1c\x44\xd3\xed\x63\x5c\x7a\x03\x8f\x8b\xef\x9b\xd8\x53\x30\x67\xcf\xc9\x4d\xb1\x35\xf6\xde\xa3\xdc\x5e\xed\xba\xc7\x54\x36\xed\x21\xd1\xf6\xe2\x8d\xac\xba\x8e\x4a\x73\x10\x30\xe5\x2a\xe6\x43\x5e\x2e\x89\x62\xb4\xd6\x45\x82\x39\x8b\x04\x26\xec\x45\x84\x0e\x20\x8c\x13\x29\x8e\x0e\xe5\x24\x3b\x77\xb7\xc9\xc3\xcd\x56\x9e\x53\xf0\x31\xc9\x0d\x17\xde\x22\x2e\x0f\x3d\xa2\x7c\x00\xf9\x1f\x8f\xf6\xff\xfd\xbb\x2b\x43\xfa\x5f\xf2\xf5\xf7\x5f\xce\xbb\x0f\xf3\xee\xef\xe6\x40\xa7\xc2\xc4\x9c\xe3\x9a\xa2\x4a\xb4\x95\xe4\xb8\x5b\xfb\x0e\x98\x75\x8e\x20\xe9\xde\xb5\x64\x3b\x68\x38\x31\x06\x4e\x60\x9e\x1d\x18\x18\x5d\x47\xe7\xf4\xbe\xcc\x94\x64\xb1\xa8\xd3\xcf\x2c\x72\x6e\xc9\x95\x8e\x6c\xe3\x74\xa0\x06\x01\x32\xaf\xd7\xd7\xda\xda\xe8\x25\xec\x8f\x39\x00\xa2\xb7\xb8\x84\x2b\x7f\xaf\xd9\xc5\x78\x48\x3d\x91\x4e\x86\xcf\x08\xeb\x69\xbc\xe4\xb7\x4c\x62\x29\x12\x12\x72\xcc\x51\xb1\x46\x6d\x12\x4b\x33\x63\xa6\x9c\xba\x3c\x03\x1b\x23\x29\x11\xe7\xd0\xd8\x55\x77\xc0\x39\xa8\x3b\x8f\x40\x12\xec\x95\xd3\xba\x2c\x76\xeb\xec\x91\x1c\x27\x3f\x75\xf2\x50\x72\xfc\x4c\x13\xa0\xf8\x14\x4e\xeb\x89\x18\xae\x0d\x9b\x04\x69\x37\x6a\x25\x02\x92\x86\xf0\xd2\x4b\x0a\x56\x86\x90\xcd\x87\x46\xca\x11\xa6\x03\xd8\x41\x56\x8d\xd6\xe0\x18\x2a\x81\x21\x5d\xf9\x00\xeb\x64\xc1\x69\xc1\x4b\xc0\x7b\xf7\x21\x11\x99\x25\x64\x91\xc7\x42\x40\x06\xc0\xa2\x93\x2e\x3b\xf6\xbe\x44\x01\x57\x06\xc4\x0d\x61\x3d\x89\x92\x79\x85\xe7\xd6\x26\x5a\x9e\x47\xaa\x17\x05\xae\x9c\xd4\x55\xac\x0d\xba\xb1\x9a\x9a\x68\x1f\x52\xab\x9d\x07\x7a\x41\xe1\xa9\x5e\xaa\xe7\x92\x87\x8b\x69\xaf\x50\xf3\xb0\x6a\xeb\xbd\xd7\x28\x1d\x37\xaa\x5a\x9b\x8f\x0d\xef\x75\xa0\x0d\x11\x47\x37\xb1\x40\x9a\x77\xec\xce\x8b\xf2\xab\x97\xeb\x06\xf2\x24\xc6\x4c\x65\x6f\x1d\x17\x09\x59\x1b\x87\x9a\xf6\x41\x5c\xc5\xa1\xb3\x2a\x87\xb3\x4f\xf7\x45\x28\x8c\x2e\xbd\xaa\x80\x59\x4e\x05\xed\xd2\xea\x3a\x80\x8e\xb5\x79\x35\xa8\x68\x10\xce\x3e\xdd\xa2\x62\x18\x9d\xcc\x8b\xd4\xdb\x8a\x3b\xad\x77\x73\xaf\xc5\x1b\xb2\x16\xee\xe3\xbf\x60\x3c\x9e\xc2\xb9\x40\x06\xf1\x68\x35\x6e\xe6\xf2\x14\xd8\x6d\x89\x0f\x44\xa1\xc1\x5c\x02\x2d\x6b\x3b\x82\x4c\x68\x10\xaa\x92\xc6\xa1\x7a\xb0\x93\xe5\xf2\xc1\x80\xf0\x38\xe9\x9e\xd1\x15\xb8\xb9\xd5\xf5\x8c\xda\x10\x81\x8a\x88\x6b\x89\x6d\x03\xdf\x43\x11\x56\xa0\x18\x8d\xd7\x41\xfa\x10\x03\xa8\x37\x84\x1b\x3c\x09\xd2\x38\xc4\x86\x14\x72\x63\xac\xd1\x3c\x2c\x09\xc8\x42\xba\x2b\xb2\x67\x28\xce\x6e\x3c\x11\x92\x66\x59\xa3\xd3\xd6\x4b\x3a\x12\xca\x2a\x7c\x9b\x6f\xec\x23\xc7\x97\x41\x31\xc5\x8e\x5a\x79\x86\x53\xc3\x83\xa6\x2f\x36\x94\x88\x8d\x4e\xe8\x2e\x08\xe9\x3e\x44\x47\x6b\xd6\x7a\xcc\x7d\x14\xb9\x41\x3c\x49\xb9\xa7\xf2\xa6\x22\x8e\xc4\x59\x73\x38\x09\x27\x39\x74\xc5\x5c\xa6\x24\x52\xde\x03\x00\x0b\x85\x0a\x57\x76\x81\x25\xd5\xec\x5d\x14\x89\x10\x23\x4d\x03\xb4\x41\xa0\xd6\xcc\x48\xfb\xac\x13\xe9\x3b\xb1\x7a\xa1\x2e\xe0\x51\xdb\x0c\xc9\xb9\xad\xa1\xa5\x5b\xad\x12\xaf\x03\xfc\x17\x94\xfb\x53\x38\x85\x60\x50\x70\x2a\x54\xf7\x43\xb4\x89\x90\x17\xc4\x3e\x50\x3d\x46\x0e\x55\x16\x1c\x0b\xf5\x3a\x1c\x23\xb6\x4e\x84\x8f\x86\x98\x40\x66\x1f\xd4\xe6\x0a\x6f\x49\x9e\xdc\x67\x5b\xcd\x69\x82\xa2\x72\x6f\x5d\x2b\xcd\xe4\x00\x4f\x8b\x39\x08\xae\x74\x9f\x6a\x94\x4e\x2a\xfd\x33\xe5\xd7\x65\x37\xfb\x7f\xce\x7f\x96\xfc\x5f\xde\xe7\xfd\xeb\x7a\x72\xf4\x83\x70\xc3\x13\xbc\x61\xcc\xb4\x93\x12\x68\x50\xf2\xa4\xc1\x66\xeb\x93\x11\xa3\xe2\x79\x5a\x19\x6d\xf2\xf2\xe1\xa2\x42\xcc\x88\x03\x24\xb1\x70\x49\xba\xa3\x32\xd3\x84\x24\xd9\xcc\x40\x3b\xc8\x29\x26\xca\x63\xe9\x95\xd7\xd8\xb7\xfa\x58\x30\x35\x1c\x11\x75\x5e\x2c\xfb\xcf\x9a\x25\xbf\xe9\xba\x21\xe9\x0b\x5a\xb4\xe4\x6f\xdb\x1c\xba\xcc\xec\x2c\xd8\xd6\xd5\x79\x0d\x55\x04\x31\x94\x3d\x4d\x60\x09\x68\x47\x5d\x95\x60\x83\x2d\x4c\x03\xf5\x6c\x5e\xbe\xfd\xf8\x36\x05\x45\xc3\xf6\x75\x09\x31\xb7\xa5\xa4\xee\xe2\xdd\xcb\x72\x02\x6b\xd1\x20\x6d\xde\x5e\xcd\xfb\xb8\x78\xaf\x63\x0f\x53\xc4\x94\x69\xc5\x79\x14\xad\xf8\x0d\x81\x1e\x41\xfe\xac\xcb\x7f\x99\xdc\xff\xa6\xd3\xcf\x5f\xff\xc0\xcf\x7f\x31\x3b\x3e\xbf\xaf\x1f\x9e\xd3\xb2\x45\xcf\xe7\xfe\xbb\x0f\xfc\x1c\xf4\xab\xd7\xef\xf1\xf3\x7f\xfc\xaf\xdf\x7d\xfe\xf2\xa7\xcf\xe8\xdd\x6f\x5f\x89\xbc\x7c\xf5\xed\xcb\xfa\xed\x1f\xfe\xf8\xe1\xe5\xd7\x9f\xc7\x1b\xfc\xee\x5f\x7f\xfc\xf6\xe5\x77\x5f\x7d\xf1\xab\x7f\xd3\x3f\x7c\xfd\x9f\x5f\xbc\xea\xaf\xff\xd2\xa7\x87\x1a\x3f\x9f\x55\x9e\x2e\xf1\xcf\xff\xfd\xfb\x6f\xde\x7f\xcb\x5f\x6e\xfd\x27\xf4\xcf\x7f\xf3\xf6\x6d\x7e\xf6\xd5\x37\x5f\xfd\xfa\xa7\x1f\x7e\xfa\xf5\x37\x9f\xd4\xab\x3f\x7e\xf1\x3f\xaf\xf2\xd3\x4f\xf1\xb5\x7c\x78\x73\x1e\xff\xf2\xa3\xbb\xbb\x3f\x7d\xf4\xa7\x8f\xfe\x2f\x00\x00\xff\xff\x84\xfd\xd8\x40\xa5\x0c\x00\x00")

func ConfigCliFleetEthBetaJsonBytes() ([]byte, error) {
	return bindataRead(
		_ConfigCliFleetEthBetaJson,
		"../config/cli/fleet-eth.beta.json",
	)
}

func ConfigCliFleetEthBetaJson() (*asset, error) {
	bytes, err := ConfigCliFleetEthBetaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../config/cli/fleet-eth.beta.json", size: 3237, mode: os.FileMode(420), modTime: time.Unix(1537450683, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ConfigCliFleetEthStagingJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x94\x5d\x6f\x1d\xb7\x11\x86\xef\xfd\x2b\x0e\xce\x75\xb1\x9a\x21\x87\xc3\x19\x5d\xa9\x76\x25\x0b\x69\x1d\xb8\x72\xec\x06\x2e\x7a\x41\xce\x87\xa4\xca\x95\x9c\x73\x8e\x1c\x3b\x45\xfe\x7b\xb1\xae\x82\xd8\x49\x74\xb7\xe0\x72\xf9\x3e\xfb\xe0\xe5\xfc\xf7\xc9\x66\xb3\xbd\x88\x5b\x8f\x9f\x3e\xdc\xdd\xef\xb7\xc7\x9b\xc3\xee\x3e\xfe\xb4\xae\x3e\x7b\x77\xbf\x3f\xc4\xee\xd9\xdd\x6d\x5e\x5f\x6e\x8f\x37\xeb\xd6\xcd\x66\x7b\x7a\x3b\xe6\xbb\xf0\x2f\x76\x6e\x36\xdb\xb3\x77\x11\x87\xed\xf1\x66\x1b\x87\xab\x65\x7f\x18\x97\xd7\xb7\x97\xdb\x87\x77\x4f\xef\xee\x0e\xdf\xde\x79\xac\x87\xff\xf3\xf3\xd2\x66\xb3\x8d\xdb\x3b\x8f\xe3\xa3\x23\xae\x30\xa1\x52\xb1\x41\xa1\xad\x95\x6c\xd0\x91\x26\x9b\x44\x11\x67\x6d\x6d\x1a\xa4\x23\x45\xd7\x06\xa9\x75\x5a\x9d\x65\x8a\x99\x18\x66\x9d\xec\x08\xd5\x93\x39\x1b\x0e\xac\xde\x7b\x9d\xcd\x27\x24\x56\x60\xc6\x9c\xcd\xfa\x94\x1c\x05\x8d\x44\x14\x8c\xc9\xba\x4e\xc2\xd1\x78\x90\xc2\x09\xf5\x45\x71\x29\x45\x17\xa2\xe3\x0a\x04\xf4\xc0\xfd\x05\x64\x76\xcd\x59\x15\x35\x7b\xb1\xd1\x18\x86\x03\x55\x72\xb3\x2a\x7d\xcc\x0c\xc2\x00\x66\x2e\x80\x31\x7d\xf8\x90\xf0\x00\xe2\x42\x8d\x56\xa8\x98\xd0\xcc\x07\xb6\x2c\xd2\xbd\x18\x4d\x97\xe6\x43\x30\x21\x66\x69\x0e\x83\x84\xe7\x9c\x26\x9e\x54\xfa\xd4\x8e\xc3\x1a\xb6\x5a\xa7\x03\x64\x9c\x60\xa7\x05\xab\x2c\x08\x7d\x29\x55\x1f\xe3\x44\x18\x5d\x0c\xbb\x16\x1d\x1d\x50\x23\x69\x8c\x52\x48\xbd\x57\x28\xd9\x79\x84\x0c\xe0\x24\x98\xc5\x4d\x64\xf6\x59\x31\x66\x26\x0d\x2e\x35\xe7\x24\x9a\xeb\x63\x1f\x36\x47\x51\x36\xb4\xf0\xda\x7b\x73\xc5\x9c\x81\x42\x5c\x0d\x1b\x68\x1f\x3c\x2a\xa7\xcf\x62\x02\x94\x59\xd3\xda\xac\xed\xa4\xb6\xa5\x54\x59\x74\xa5\xfc\xc5\xe6\x67\xc8\x7f\x3d\x74\xe1\xbb\xdd\xda\x29\x7f\x31\xae\xdf\xbd\x8a\xdd\x87\xd8\xfd\x61\x29\x66\x27\x69\x8a\x9d\x4d\x7b\x43\xaf\x48\x23\x22\xd3\x0a\x87\x69\x0a\xf3\x20\x2c\xbd\x95\xe8\xee\x61\x8a\x13\x15\x50\x06\x4a\x74\xb3\x21\xce\xb5\x5b\x46\xb1\x70\x9b\xda\x4b\x4a\x4c\x26\x17\xe4\x9c\xde\x26\x85\xa8\xf5\x10\xeb\xe9\xdd\x53\x06\xd6\xa2\x39\xa8\xfa\x14\x98\x30\x33\xd6\x52\xb4\xb2\x28\x2c\xd8\xf8\xb8\x42\xfb\x23\xd9\xbc\xb6\x61\x8e\xd1\x13\x7b\x29\x8e\x88\x43\x4c\x8d\xc5\xaa\x8e\x8a\x54\x21\xd4\x1b\x77\xd6\x96\x8c\x20\x59\xd1\x6c\xb2\xb9\x24\x0c\xcf\x24\x55\x8c\x9e\x6e\x43\x72\xf4\x0e\xd1\xdb\x34\x19\x6d\x3d\xa7\x7b\x21\x56\xb0\x69\x02\x01\x20\xd8\x5b\x12\x18\xb6\xee\x9c\xdc\x45\x9c\xe4\xa4\x00\x2f\x28\xba\x14\x82\x05\x1f\xc5\x0c\x4a\x43\x30\x4c\x6e\x26\x23\x5c\xea\xb0\xc2\xd3\x70\xe6\x2c\x38\xa8\x99\xe1\x90\xd6\x60\x34\x81\xde\x4d\xbc\x78\x94\x01\x01\xe6\x28\x41\x90\x4e\x90\x3d\x38\xbb\x43\x71\x63\x73\x60\x95\x32\x01\xc9\x42\xc4\x39\x88\xa5\x97\x96\xab\x71\xac\x12\xda\x64\x25\x04\x80\x32\x7a\xfe\xbf\x13\xba\xa0\xd6\x85\xf0\x81\xf2\xab\x4e\xbc\x3a\x8c\xc3\xb5\x3d\x3a\x21\x14\xc9\x60\x56\xc8\xd2\xe7\x98\x15\x0c\x3d\xbd\xa2\x0f\xef\xdc\xca\x20\x4e\x1f\x5a\x3b\x34\x2a\x23\x02\x27\x70\x21\x95\x89\x95\x5a\x04\x28\x56\x46\x9a\x32\x35\x6a\x08\x71\x29\x21\x34\x7a\x81\x6a\x4d\x9a\x8c\x40\x57\x41\x4d\xae\x23\x2c\xb0\x46\x30\x4b\x50\x32\xb3\x00\x57\x15\x3d\x41\xee\x8b\xae\x3f\xb0\x20\xc9\x71\x85\x0a\xed\xf7\x9a\x8b\x8b\x76\xe3\x35\x42\x49\x53\xdd\x12\xa1\x74\x4d\x80\xd0\x29\xb5\x34\x43\xc9\xe8\x39\x5a\x71\x6e\xd2\x0a\x0c\xef\x33\x94\xa1\x9b\x54\x00\x99\x54\x26\xf0\x88\x74\x5d\xdb\x8b\xe9\x63\xb6\x8e\x59\xeb\x28\x43\xab\xd4\x4c\xed\xd6\x14\x34\xbc\x61\x72\x5d\x3f\xaa\xa4\x58\x4b\xed\xb1\x6a\x46\x2d\x0b\x2c\xc2\x0f\x8c\x5f\x49\xfe\x75\xc4\xff\x4e\xf4\xd1\xf5\x7b\x3a\xfa\xea\xe6\x1e\x1d\xec\xfd\x51\x85\x0e\xf5\x28\x0e\x57\x1f\xe8\x08\xf9\xf5\xf5\x7d\x39\xff\xf3\x7f\xf8\xb9\xee\xff\xf2\xe2\x66\xf7\x94\xbe\x1f\xed\xf4\xbc\xbe\xfd\x58\xfc\xd3\xfe\xd9\xc7\xb3\xcb\xa7\xdf\xbd\xba\xf8\xe9\xf2\xea\xda\xea\x5b\x3d\x5c\x9c\x5d\x7c\x7b\xfa\xab\xa4\xcf\x09\xbf\x99\x61\x8f\x67\xc8\xeb\xb7\xdf\xbf\x3e\x3f\x7d\xf9\x76\xf7\xfe\x9b\x69\x7f\xaf\x9f\xde\x9c\x9d\xf3\xeb\xc3\x5f\x77\x3f\xee\x9f\xf3\xbf\xcf\x69\x5c\xf7\xf2\x72\x3f\xff\x96\x6f\xce\xe6\x6f\x32\xbe\x1c\xe6\x8f\x07\x5c\xdc\x1e\xca\xe9\xa7\xbb\x9b\xdb\xab\x3a\xee\x3f\x5e\x51\x7e\xf3\xe3\xcd\xc5\xe5\x0f\x37\xe7\x78\xf9\xc3\xee\x1f\xcf\x63\x2f\x2f\x6f\xf0\xa6\xbe\x78\x43\x1f\xef\x1f\x2c\x3e\xd9\x6c\x7e\x7e\xf2\xf3\x93\xff\x05\x00\x00\xff\xff\xdf\x3e\xba\xe7\x2e\x07\x00\x00")

func ConfigCliFleetEthStagingJsonBytes() ([]byte, error) {
	return bindataRead(
		_ConfigCliFleetEthStagingJson,
		"../config/cli/fleet-eth.staging.json",
	)
}

func ConfigCliFleetEthStagingJson() (*asset, error) {
	bytes, err := ConfigCliFleetEthStagingJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../config/cli/fleet-eth.staging.json", size: 1838, mode: os.FileMode(420), modTime: time.Unix(1537450683, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ConfigCliFleetEthTestJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x94\x3d\x8f\x14\x36\x10\x86\xfb\xfb\x15\xab\xad\xd1\x32\x9e\xef\xd9\x0a\x41\x48\x1a\x92\x86\xa4\x8a\x52\x78\x3c\x63\x82\x74\xe2\xa4\x63\xa1\x89\xf8\xef\xd1\x72\x57\x10\xc2\xb5\xaf\xc7\xf2\xa3\x77\x1e\xf9\x9f\x9b\xc3\xe1\xf8\xe6\xee\xdd\x9b\xfe\xdc\xb7\xc7\xf3\xe1\xf8\xd3\xeb\x97\x7f\xfc\x72\x7c\x76\x8d\x5f\xdd\x7e\xfa\x78\xe9\xfb\x57\x77\x1f\xf6\xfb\x77\xc7\xf3\xe1\x3a\x7b\x38\x1c\x5f\x7f\x98\x79\xdb\x75\x3c\x1f\x2e\xf7\x9f\xfa\xd9\x43\xf8\xf3\x6d\xf7\xe5\x7a\xbf\x2f\x7f\x9f\x2e\xfd\xf1\x72\x7c\x3c\x78\x79\x77\x77\xf9\xed\xae\xfa\xe3\xf1\x7c\xf8\xf3\x6b\x74\x38\x1c\xfb\xc3\x5d\xf5\xf9\xf9\xf3\x9a\xb3\xb1\x0d\x1d\xa1\x5d\x3b\x18\xf7\xc4\xe9\x73\x5a\x69\x47\x08\x17\x03\xd3\x34\x21\x76\x2a\xcf\x22\xf2\x1e\x9a\xed\xb8\x76\x28\x52\x60\xc9\x82\x1e\xb3\xc5\x16\x95\x05\x2d\x2a\xaa\xaa\xf2\x5d\xe2\x44\x81\x8a\xcd\x48\x5c\x2b\x54\x77\x08\xad\x32\x4a\x21\xd9\xb2\xe7\x0b\xb6\x93\xe0\x69\xb8\x9f\x06\xc7\x99\x80\x81\x1f\xc1\xbf\xa1\x8c\x86\x70\x17\x93\x4e\xb3\x61\x0b\xa5\x7a\x1a\xee\x1a\x63\x59\x92\x99\x5a\x2d\x88\x35\xa6\xa9\xeb\xb6\x85\xbd\xc4\xac\x08\xbc\x90\xaf\x03\xab\x53\x4d\xaa\x99\x86\xb1\xf1\x1c\xd3\x0d\x9a\xcd\x1d\x8d\x30\xd4\x36\xef\x69\x2e\x09\x98\x33\xa4\x54\x23\x69\x14\x2b\x57\x2f\x78\x81\xa0\xa7\xe1\x71\x42\xa6\xd3\x50\x7e\x8a\x73\x8d\x16\x18\xee\x6e\xcb\x95\x4a\xf9\xfa\x5e\x6a\xe4\xd6\x61\xa2\x03\xdc\x1c\x85\x09\x9a\x19\x86\x11\x6d\xc9\x69\xb6\x0c\x2a\x07\x23\x69\x6d\xf2\xb1\x7b\x27\x74\xf6\x98\x8b\x11\x83\x33\x1a\x75\x64\x36\x12\x55\x56\x3a\x35\xe1\x12\xd7\xa5\x3a\x1b\x75\x79\xb5\xc1\x4a\x5e\x2f\x48\x1e\xaa\x54\x3f\x0d\xb2\x47\xcc\xaf\x94\x7f\x3d\xda\xf0\xfb\xfd\x55\xa9\xfa\x75\xbe\xbf\x7d\xdb\xf7\x9f\xfb\xfe\x87\x5a\x84\xf0\x02\x55\xa0\xa9\x6d\x22\xb9\x77\x47\x04\xea\x90\xcd\x26\xe2\xec\x59\x33\xe6\xac\x9a\xc6\x15\x08\x73\xd2\xa8\x51\xde\xbc\x15\x10\x6b\x89\x68\xad\xa9\xa6\xbe\x7d\xc2\x0e\x96\xd8\x62\x69\x18\x02\x51\x49\xcb\x93\x32\x1d\xe6\xda\x59\x3e\x71\x20\x81\xdb\xd6\x95\x65\x59\xdf\x6a\x11\x7a\x26\x90\x1f\xd5\xdd\xec\x2a\xbb\x75\xe1\xbc\xee\x5e\x94\xa6\x32\x5b\x04\x94\x77\xac\x56\x43\x65\x9e\x4d\x0d\x8e\x57\x07\xc8\x7b\x8d\x7d\xdd\x08\x74\x6f\x25\x84\x05\x36\xa5\x14\x16\xa5\x62\x6c\x29\xe6\xe0\x1d\x54\x9a\xae\xd3\x58\x26\xe4\x56\x9e\x89\x21\x99\x5b\xc5\x02\x86\xcd\x5a\xda\xe5\xdf\x69\x31\x9e\xe2\x34\xb0\x16\x63\xa1\xb9\x8a\x9a\xdd\x17\x73\x46\x41\x0f\x0b\x13\xb2\xd1\xb8\x7d\x5f\x05\x45\xe9\xd9\x92\x73\xcd\xc8\x58\xde\xa0\xcb\x75\x55\xdb\xb2\xe0\xa9\x0b\x9b\x14\x81\x72\xc7\x16\x5d\xb3\x3d\xa1\x05\x36\x25\xd1\xe2\x85\xa9\xc1\xd3\x72\x76\xcf\x61\xc2\xac\xbc\x9a\x9b\xbe\x6a\x11\x78\x42\x8c\xd3\x30\x7c\xc4\xfc\x8f\x16\x6f\x2f\xf3\xf2\x7e\x3d\xf9\x4d\x8c\x1a\x41\x4a\xd2\x30\x24\x86\xef\x74\xc9\x9c\xdb\x8c\xaf\x9e\x0f\xdc\x2a\x65\xb0\x34\xcc\x04\x74\xb8\xf5\x06\x66\x84\x32\xee\x05\xba\xa2\x7d\x43\xad\x14\xeb\x09\x8c\xdb\xa5\xb6\x1b\x13\xd5\xcc\x61\x73\xa0\x42\x97\x57\xf5\xc8\xda\x51\x5a\x52\x58\xcd\x69\xb9\xbd\x23\xe8\xbb\xa6\xe9\x4c\x40\x20\xff\x6f\x7a\x4b\xd0\x44\x33\x1a\xb9\x60\x7b\x27\xb8\x37\x16\x05\x22\x2e\x2c\x89\xda\x19\xb9\x21\x29\x04\xca\xa6\xa3\x97\x8c\xf6\x99\xd1\xe0\x7b\xef\xac\x88\xa1\xd3\x71\x57\x04\xad\x31\xc1\x61\x89\x2d\xcc\x32\xe8\xd2\x45\xba\xd9\x63\x86\x46\xb5\x86\xcd\xbd\x83\xc0\xbd\xb2\x7b\xcc\x78\x68\x9a\x4f\x34\x4e\x03\xfc\x91\xf2\xa1\xe8\x9b\xc3\xe1\xcb\xcd\x97\x9b\x7f\x03\x00\x00\xff\xff\x0e\xad\x46\x5b\xef\x05\x00\x00")

func ConfigCliFleetEthTestJsonBytes() ([]byte, error) {
	return bindataRead(
		_ConfigCliFleetEthTestJson,
		"../config/cli/fleet-eth.test.json",
	)
}

func ConfigCliFleetEthTestJson() (*asset, error) {
	bytes, err := ConfigCliFleetEthTestJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../config/cli/fleet-eth.test.json", size: 1519, mode: os.FileMode(420), modTime: time.Unix(1537450683, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ConfigCliLesEnabledJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\x52\x50\x50\x50\x50\xf2\xc9\x4c\xcf\x28\x71\x2d\xc9\x70\xce\xcf\x4b\xcb\x4c\x57\xb2\x52\x80\x88\x83\xe5\x5c\xf3\x12\x93\x72\x52\x53\x94\xac\x14\x4a\x8a\x4a\x53\xc1\xe2\xb5\x5c\xb5\x5c\x80\x00\x00\x00\xff\xff\xa6\x74\x24\x05\x3a\x00\x00\x00")

func ConfigCliLesEnabledJsonBytes() ([]byte, error) {
	return bindataRead(
		_ConfigCliLesEnabledJson,
		"../config/cli/les-enabled.json",
	)
}

func ConfigCliLesEnabledJson() (*asset, error) {
	bytes, err := ConfigCliLesEnabledJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../config/cli/les-enabled.json", size: 58, mode: os.FileMode(420), modTime: time.Unix(1537450683, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ConfigCliMailserverEnabledJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xe6\x52\x50\x50\x50\x50\x0a\xcf\xc8\x2c\x2e\x48\x2d\x72\xce\xcf\x4b\xcb\x4c\x57\xb2\x52\x80\x08\x83\xa5\x5c\xf3\x12\x93\x72\x52\x7d\x13\x33\x73\x82\x53\x8b\xca\x52\x8b\x94\xac\x14\x4a\x8a\x4a\x53\x75\x10\x2a\x10\x72\x01\x89\xc5\xc5\xe5\xf9\x45\x29\x4a\x56\x0a\x4a\xc5\x25\x89\x25\xa5\xc5\xba\xf9\x69\x69\x39\x99\x79\xa9\xba\x99\x79\x49\xf9\x15\x4a\x60\x4d\xb5\x5c\xb5\x5c\x80\x00\x00\x00\xff\xff\x84\xf6\x09\xc4\x78\x00\x00\x00")

func ConfigCliMailserverEnabledJsonBytes() ([]byte, error) {
	return bindataRead(
		_ConfigCliMailserverEnabledJson,
		"../config/cli/mailserver-enabled.json",
	)
}

func ConfigCliMailserverEnabledJson() (*asset, error) {
	bytes, err := ConfigCliMailserverEnabledJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../config/cli/mailserver-enabled.json", size: 120, mode: os.FileMode(420), modTime: time.Unix(1537450683, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _ConfigStatusChainGenesisJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x41\x6f\x13\x31\x10\x85\xef\xf9\x15\x96\xcf\x1c\xc6\xf6\x78\x66\xbd\x37\x4a\x59\x81\xc4\x8d\x5f\x30\xb6\xc7\xcd\xaa\x9b\x6c\x94\x6c\xa4\x16\xd4\xff\x8e\x92\x2c\xa1\x42\x9c\xf0\xcd\x7e\x33\xef\x7b\x7e\x3f\x37\xc6\x58\x99\xa6\xb9\xd8\xde\x5c\x2e\xc6\x18\x9b\x07\x47\x58\x24\xa0\x0b\x9e\x04\x42\x8e\xc8\x05\xe2\x43\xc0\x90\xfd\x67\xef\x74\xf8\xa8\x1e\x73\xfa\xb3\x62\x6c\x96\x49\xf6\x45\x6d\x6f\x2c\xbc\x38\x18\xe0\xd3\x00\x84\x8f\x8f\x31\x79\xb8\x1c\x7b\x1d\x7c\xfb\xb0\x22\x3c\x44\x89\x15\x94\x7d\x6b\x14\x81\x53\x2b\x92\x33\x0a\x05\x0d\xa1\xfa\x4e\x84\xaa\x36\x5f\xfe\x03\xb1\x59\x31\xb6\xcc\xfb\x36\x3e\xdd\x1d\x6c\xd9\xca\xb8\xff\x5a\x6d\x6f\x98\xf9\x16\xc4\x6e\xe7\x9d\x9e\x16\x95\xfa\x30\xcd\xe5\xd9\xf6\x06\x56\xa1\xca\x3c\xcc\xc7\xe7\xef\xe7\xc3\x61\x3e\x2e\xb6\x37\xcb\xf1\xac\xab\x96\x5f\x7f\xc8\x7e\x19\xcf\xbb\xbf\x97\x74\x3c\xb8\x08\x5f\xe4\xb4\xbd\x85\x8c\xa0\xa4\x35\xa3\x26\xa8\x89\x1c\x49\xe9\x72\x63\xe7\x92\x6a\xe0\x82\x80\x5d\x41\x27\x15\x52\xf0\x1d\x31\xa9\xcb\x21\xd5\x42\x9d\x80\x96\x96\x43\xb5\xef\x7d\xe3\x3f\x69\xdd\xfd\xf5\xfe\xef\x3a\xb6\x36\x96\xf3\xb4\xbc\xfe\xae\xea\xda\xce\x55\xd3\x97\xe5\x28\x8f\xb2\xc8\x1a\x30\x30\x92\x63\xe4\xc8\x01\x03\x75\xe4\x28\x91\x7a\x88\xc0\x9e\x12\xd3\x45\xa3\xe8\x21\x22\x45\xbe\xcc\x26\x52\x62\x0f\xa8\x14\x19\x99\xa9\xb1\xa7\x7c\xb3\x7e\x92\xd3\xb7\x71\x37\x2e\xef\xa1\x60\x37\x6f\x9b\x5f\x01\x00\x00\xff\xff\x91\xc6\xb3\x58\x64\x02\x00\x00")

func ConfigStatusChainGenesisJsonBytes() ([]byte, error) {
	return bindataRead(
		_ConfigStatusChainGenesisJson,
		"../config/status-chain-genesis.json",
	)
}

func ConfigStatusChainGenesisJson() (*asset, error) {
	bytes, err := ConfigStatusChainGenesisJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../config/status-chain-genesis.json", size: 612, mode: os.FileMode(420), modTime: time.Unix(1537450683, 0)}
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
	"../config/cli/fleet-eth.beta.json": ConfigCliFleetEthBetaJson,
	"../config/cli/fleet-eth.staging.json": ConfigCliFleetEthStagingJson,
	"../config/cli/fleet-eth.test.json": ConfigCliFleetEthTestJson,
	"../config/cli/les-enabled.json": ConfigCliLesEnabledJson,
	"../config/cli/mailserver-enabled.json": ConfigCliMailserverEnabledJson,
	"../config/status-chain-genesis.json": ConfigStatusChainGenesisJson,
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
	"..": &bintree{nil, map[string]*bintree{
		"config": &bintree{nil, map[string]*bintree{
			"cli": &bintree{nil, map[string]*bintree{
				"fleet-eth.beta.json": &bintree{ConfigCliFleetEthBetaJson, map[string]*bintree{}},
				"fleet-eth.staging.json": &bintree{ConfigCliFleetEthStagingJson, map[string]*bintree{}},
				"fleet-eth.test.json": &bintree{ConfigCliFleetEthTestJson, map[string]*bintree{}},
				"les-enabled.json": &bintree{ConfigCliLesEnabledJson, map[string]*bintree{}},
				"mailserver-enabled.json": &bintree{ConfigCliMailserverEnabledJson, map[string]*bintree{}},
			}},
			"status-chain-genesis.json": &bintree{ConfigStatusChainGenesisJson, map[string]*bintree{}},
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

