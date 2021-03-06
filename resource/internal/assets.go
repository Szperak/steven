package internal

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

func bindata_read(data, name string) ([]byte, error) {
	var empty [0]byte
	sx := (*reflect.StringHeader)(unsafe.Pointer(&data))
	b := empty[:]
	bx := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bx.Data = sx.Data
	bx.Len = len(data)
	bx.Cap = bx.Len

	gz, err := gzip.NewReader(bytes.NewBuffer(b))
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

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _assets_minecraft_texts_splashes_txt = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x55\x7f\x6e\x1c\x45\x13\xfd\xbf\x4f\x51\x93\x4f\xfa\xbc\x8b\xec\x59\x12\xdb\x28\xd9\x28\x42\x89\x13\xc2\x0a\x6f\x82\x70\x22\x40\x88\x44\xbd\x3d\x35\x3b\x9d\x9d\xe9\x1e\xfa\xc7\x8e\x57\x21\x77\xe1\x0c\x9c\x88\xa3\xf0\xaa\x17\x88\x25\x5b\x9a\x6e\x57\x57\xbd\x7a\xf5\x5e\x79\x45\x9b\x5e\x0f\x4c\x3f\xf9\xa0\x56\x29\x92\xf3\x89\x34\x6d\xf2\x96\x2c\x4e\x9a\x5a\xd6\x29\x07\xae\xd4\x73\xef\x4e\x12\x6d\x3d\x25\x4f\xff\x4b\x9d\x75\xbb\x53\x4a\x56\x42\xa2\xed\xfb\x03\x8d\xbd\x36\xac\xde\x70\x4c\xdc\x54\x34\x5b\x39\x1a\x83\x6f\xb2\x49\xd6\xbb\xb9\x7a\x85\xb4\xd6\x51\x34\xba\xd7\x55\x29\x84\xb8\x3d\xbb\x52\x6f\xec\xca\x77\x85\xa8\x89\x26\xde\x48\x14\x2a\xae\xd9\x4f\x0a\xf8\x38\x91\x77\x4c\xbe\x25\x63\x5d\x73\x38\x89\x64\x34\x12\x6c\x82\xdf\x31\x50\x56\xea\xc9\xbb\xfa\xdd\x13\xf5\xd7\x9f\x7f\xbc\xc7\xaf\xba\xf2\xc3\x60\x13\x05\xa4\x0c\xc0\x22\x65\x2f\xeb\xfa\xa2\xae\xcf\xeb\xba\x56\xd7\x3a\x01\x22\x01\xf8\x36\x70\xf9\xae\xd4\x2f\x6f\x3a\x9c\xe3\x88\x06\xe4\x0f\xd6\x25\x76\x02\x5b\x4b\x5f\x3d\xb7\x49\x38\x72\xbb\x5f\xd5\x1b\x69\xfb\x99\x4f\x4b\x42\xaa\x9a\xbe\x28\x67\xdf\x36\x48\xd4\x51\xa3\x07\x07\x38\xc7\x26\x2c\x2e\x06\x1f\x58\x98\x8c\x95\xfa\x97\xe5\xb5\xff\xa0\xdd\x16\x79\x98\x7a\x0f\x2e\x51\x2d\xf9\x54\xea\x14\xe2\xa3\xb1\x96\x74\x10\xf4\xae\xe1\x00\xf4\x1a\x11\x7c\x2b\x23\x88\xea\xda\x7b\x90\x8e\xde\x26\x1f\x76\x11\xa4\xd0\x70\xa0\x41\x1b\xa0\xe0\x5a\xbd\x1e\x41\xe7\x8d\xcf\xc1\x70\x45\x5d\x4a\x63\x5c\x2e\x16\x5b\x00\xc9\x9b\xda\xf8\x61\x91\xee\x80\x5d\x1c\xd9\x57\xcf\xb2\xed\xd3\x11\xed\x4b\x5f\xa9\x14\x0e\xf4\x91\x3e\x09\xbf\xa6\xa3\xd9\x8b\x5b\xc3\xa3\x10\x41\x3c\x97\x7b\x35\x6a\x67\xcd\x2c\xb0\xf1\xe0\x76\x36\x9f\xab\xc5\x82\x9e\x6e\xb4\x6b\x10\xd2\xf9\x91\x09\x9d\xd0\x81\x69\xea\x3c\x81\x42\x0e\xd4\xa1\x89\x32\xee\xde\x62\x58\xab\x93\x81\x82\xc6\x14\xb7\xb4\x47\xaa\xbe\xd7\xa2\xa6\xc8\xc7\x17\x46\x23\x8b\xde\x33\x25\xd0\x33\x78\x4c\x49\xc8\x53\x6f\xa3\xc4\x7f\xa6\xa6\xf5\xa1\x44\x14\x02\xf1\x76\x00\x4b\x25\x3b\x64\xab\x31\x6e\x34\x48\x1a\x61\x36\xc4\xa4\x64\x4c\xc0\x07\xb5\xa1\x1a\x35\x9e\x63\xad\xbe\xe5\x7e\xac\x0a\x96\x14\xf4\x38\x1e\x25\x22\x19\x23\x34\x1c\xbb\xc2\x77\xa5\xae\xad\xcb\xb7\x14\xf3\x38\xfa\x80\xe3\x8f\x90\x9e\x9f\xe2\xe7\x8b\xb5\x36\xff\x1d\x68\x76\xcc\xe0\xc3\x61\x2e\xfa\x63\x32\x9d\x16\xd5\xd2\xea\x87\xab\x4a\x7d\xe7\xfc\x24\x3d\xac\x73\x18\x3b\xd1\xef\xb5\x9e\x0a\x16\x71\x54\x07\x34\x6a\x8d\x09\x9a\xa0\xa1\xb4\x75\xee\x93\x3d\x83\x75\x0c\x47\x69\x7c\x09\x9d\xb3\xde\xc9\xf3\xd4\x05\x16\x72\xf0\x1d\xa9\xa4\x87\xe1\x6e\x8a\xf9\xd6\x40\xa1\x7b\xd1\xf7\x8d\x6f\xd3\xa4\x83\x08\x59\xb2\x07\xee\x59\x47\x6e\x6a\x5a\x25\x14\xc6\x7c\xfc\x84\x86\xc1\x3a\xc3\x65\x23\x64\xf3\x92\x1d\x07\xdd\xd3\xcd\x01\x9a\x18\xe8\x45\x08\x3e\x2c\xe9\xfb\xf2\x8c\xa2\x36\xc1\xb6\xd6\x08\xb9\x06\xba\x16\x2a\xd3\x84\x51\x75\xd6\xec\xd8\x89\x7a\x71\x0f\xb3\xb8\xcc\xd8\x10\x74\xf0\x99\x26\xed\x92\xdc\x6f\xa0\xad\x46\xde\xf5\x16\x5a\xf8\x5a\xc5\xdc\x78\x0a\x03\x9d\x85\x96\xce\xce\x9c\x47\x93\x1c\x39\xec\xf9\x2c\x78\x68\x7f\xa1\xee\xba\xa9\xde\x72\xba\xc2\x34\xe2\x6c\x2e\x9f\xdf\x58\xee\x9b\xd9\xbd\x08\xd1\xa4\xc3\xbd\x79\x1d\x39\xcd\xee\x86\x9f\x92\xcb\x7d\x3f\x7f\xfc\xd9\x7c\x23\x6c\x95\x40\x71\xa5\x9e\x0a\xe6\x0e\xde\x10\x70\x06\x7d\xb9\x58\x24\x36\xc8\xd8\x45\x33\x10\x06\x76\x0b\xa2\x07\x6b\x44\x36\x58\x3d\x22\x0e\x0d\x55\xa5\xd4\x43\x8a\x3a\x05\x7b\x7b\x4a\xf1\xb7\x2c\xc4\x1e\xe5\xe4\x73\x3a\x55\x49\xf6\x06\x7e\x60\x43\x48\xc9\xc5\xd1\x47\x3e\xa5\x3b\xb7\xc6\x67\x97\x6a\xb5\xd7\x81\x5a\x7a\x42\x6d\x76\x65\x1f\xce\xc4\x4e\x81\xe1\x6a\x47\xed\x63\xfa\xf4\x58\x65\x17\x75\xcb\xb5\xcc\x07\xe6\xe3\xd7\x9b\x0f\x6c\xd2\xec\xed\xf1\xd6\x08\x11\x68\x6e\x39\x9b\x7f\xa4\xe5\xef\xcb\xff\xe3\xc5\x52\xfd\x8c\x76\x86\x0c\x97\x8c\xfa\x20\x7c\xef\x2d\x4f\xc7\xda\x32\x11\x96\xba\xaf\x64\x5b\x67\xd3\x1d\x8a\xbc\x93\x1f\xad\x51\x57\x6c\x2c\xb9\x13\x2e\x0f\x23\x65\xf7\x8f\xec\x6b\x85\x7f\x01\x02\xbb\x2c\x22\x93\x72\xd9\x4a\x5a\x76\x41\x5d\x16\xd6\xfa\xf9\xa5\xf0\x34\x34\xef\x2f\x25\x8c\xef\x5f\x98\x56\x9b\xa6\xb9\xb8\x78\xa0\x1f\x5d\x9e\x9f\x5f\x9c\xf3\xa6\x79\x78\xf9\xe0\xd1\xfd\xf3\x87\x5f\x3d\xfc\x52\xfd\x1d\x00\x00\xff\xff\x7a\x5c\x78\x1c\x5a\x06\x00\x00"

func assets_minecraft_texts_splashes_txt_bytes() ([]byte, error) {
	return bindata_read(
		_assets_minecraft_texts_splashes_txt,
		"assets/minecraft/texts/splashes.txt",
	)
}

func assets_minecraft_texts_splashes_txt() (*asset, error) {
	bytes, err := assets_minecraft_texts_splashes_txt_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "assets/minecraft/texts/splashes.txt", size: 1626, mode: os.FileMode(436), modTime: time.Unix(1432374943, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assets_minecraft_textures_font_ascii_png = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\xd4\x6b\x54\x92\x09\x1e\xc7\xf1\x07\xcc\x32\xcb\xb2\x49\x1a\x45\x92\x4d\x3a\xdb\xa4\x43\xdb\x65\x4b\x1b\x0d\xc7\x06\x34\xcc\xb4\xbc\x80\x44\x81\x2d\x14\x95\xb9\xea\x74\xc6\x28\x44\x21\x4f\x53\xe8\xa8\x59\xa3\xa5\x19\x58\xed\xd6\x78\x43\xc5\x01\x73\x63\x04\x47\xf1\x32\x5e\xd0\xc1\x1b\xa0\xa4\x43\x47\x4a\x45\xbc\x44\x28\x12\x2e\x76\xd6\x1d\x9f\x73\x9e\xe7\xff\xea\xf7\x79\xf5\x9c\x2f\xe7\x64\x58\xb0\x93\x23\xd4\x11\x00\x00\x27\xec\x31\x74\x84\xed\xb2\x96\x5f\x87\xb5\xb6\x6f\x23\xea\x5c\xbb\xed\xac\xb9\x7a\x14\x1b\xe5\x60\x7b\x52\x1c\xa8\x77\x00\xc0\xbe\x0f\x8b\x0e\x8c\xa2\x17\xe9\xcf\x06\xde\xeb\xdd\xbb\x31\x4a\x34\xb6\x4b\x7c\xfb\x40\xd4\x29\xb1\xbd\xf5\xd7\x5b\xe1\x33\xf9\xa3\xc0\x93\x1f\xbc\xf3\xa9\xd8\x30\x81\x62\xf8\x81\x12\x4b\xcc\xbf\x72\x8d\x2a\xab\x79\x55\xc3\xbc\x97\x10\xee\x36\xf2\x22\x25\x4f\x3e\x4b\x41\xdf\x4f\x03\x59\x77\x78\xce\x62\xea\x37\x17\x5a\x3e\x9b\xad\xec\x94\x38\xb9\x82\xd7\xb8\x00\xbe\x5f\x00\x89\x27\x81\x91\x62\x40\x3a\xe3\xbc\x65\x61\xbe\xf6\x34\x8c\x6d\x99\x98\xdd\xbc\xc5\x0d\xe5\x6d\xfd\xb1\xf2\xa8\x55\x78\x9d\x6b\xc7\x9a\x2b\x07\x49\x03\x41\xac\x74\x30\xb0\xce\x1e\xd8\xe4\x08\xc0\xfc\xd7\x01\x3a\x9f\x09\x13\xe8\xcd\x42\xa0\x28\x5e\xfd\xf7\x75\x72\xdd\x5c\x5d\xbc\x5a\x24\x09\x52\x5d\x14\x11\xe8\xe6\x7a\x4f\xbe\x9f\xf4\xbb\x4d\x77\x6f\xd1\x66\xd4\x03\x35\xbb\xce\x10\xa9\xc7\x07\x0d\xdb\xdf\x41\x63\x93\x93\x4c\x03\x1a\xd2\x1b\x39\xd5\xc4\x84\x09\xe1\xb2\xa4\x26\x42\xa6\x8e\x30\xc4\x4c\x41\x30\x90\x7b\x32\xfd\xdc\x99\x5f\x5e\x02\xfd\x9c\xa9\x1b\x5d\x4c\xd6\xbf\xf1\x38\x11\xa5\xd6\x84\x28\x82\x37\x08\xcd\x5d\xc6\x2b\x4c\xfd\xc4\x97\x5c\xa7\xfd\x49\x53\x4c\x63\xd2\x9e\x34\xc7\x87\x9d\x80\xea\xf1\x73\xa4\x15\x85\xc5\xc7\xc3\x9f\xbd\xac\x70\x6f\xfa\xb5\xbf\xd8\xb3\xcb\xfa\x9f\xdd\xac\x72\xad\x72\x63\x62\xf2\x8e\xe7\x48\x6e\x8b\x84\xc6\xae\x40\xb6\xba\xc7\xf5\x92\x32\x5f\xbb\x7c\x8d\x61\xe1\x0a\x7b\xff\x5a\x5d\xeb\x17\x32\x13\x3f\xe3\xca\xf1\xb9\x77\x7d\x0a\xec\x7d\xc8\x3b\x75\x0e\x56\xfc\x95\x5d\x87\xf9\xf0\x6f\xae\xa7\x1f\x85\x96\x3e\xae\x23\xdd\x22\xbe\xaf\x84\xb8\x69\x5c\x28\xd1\x0d\x9a\xd4\x20\x91\x07\xbc\xd0\x10\x9a\x12\x8d\x19\x8c\x3e\xa9\x4d\xe0\xf4\x80\x63\x1e\x33\xf4\x88\xfb\x30\x48\x76\xed\xd9\x68\x0b\x01\xbd\xd8\xc5\x0f\x57\x52\x62\x7f\xfc\xbc\x4d\x5f\x83\x37\xf4\x1c\xd8\x50\x2c\xec\xf1\x40\x08\xfd\x06\x4b\x19\x1d\x75\x6d\x17\x80\x04\x3f\x24\x86\xa1\xa2\x46\x88\x21\xbc\x6b\x1b\x62\x72\x41\xa5\xe6\x8d\x39\x92\xa3\x80\xd2\x65\x88\xd2\xaf\xc9\x68\x54\xd3\x64\xc3\xb0\x40\x5c\x4a\xc9\xcb\x1a\xf0\xc3\x51\x04\xb3\xb3\xe9\x41\x56\xc0\x81\x36\x9b\xfb\x37\xce\xf1\xc8\x87\x34\x30\x32\xef\x86\x82\xd6\x8b\xda\x2a\xa9\x5d\x72\xd0\x19\x25\xd0\x6b\x77\x40\xdb\xae\xf1\x2c\x3e\xb9\x43\x28\xad\x30\x4b\x45\x74\xbb\xd0\xbc\xd4\x6a\x7d\xab\x2b\xb0\x1e\x64\x64\xf2\x84\x3d\x1a\xe2\x55\xd7\x6f\xbb\x49\x99\x61\xb1\x10\x03\x27\xa0\x31\x6e\xeb\x4f\xc6\x27\x42\xc5\x58\x97\x3b\x59\xdd\xd2\x73\xe3\x7c\x7f\x8b\x91\xc0\x32\x5e\x1a\xf6\x40\xcb\x0a\xa4\x5e\xf4\xfd\x4d\x21\xaf\x0b\xaa\x19\xa8\xe3\x9c\xe6\xa5\x70\x80\xbb\x1e\x2a\x54\x06\xf0\xcf\x1c\xe4\x28\x5f\x9d\x41\xf3\xfe\x95\x91\x43\x7c\x56\x4f\x33\xb7\x4e\x23\x2c\x0f\x40\x41\x5e\x6a\x28\xba\x88\x9a\xa4\xeb\xf8\x43\xf6\x9b\xfb\xce\xf4\x34\xb0\x2f\x85\x6d\xf0\x05\x11\xc4\xa6\x7d\x47\x6e\x9c\x0f\x42\x57\xa9\x73\x04\xb7\x9b\x89\x1f\x90\x70\x8e\x5c\x4b\xfe\x86\x18\x8b\x69\xd1\x19\xab\xff\x58\x5b\x7e\x27\x87\x7e\x3e\xc6\x9e\xa4\x1e\x73\x64\x4c\xec\x7b\x2e\xc9\x67\xbe\x26\x9b\x06\xe3\x28\x55\x63\xcd\x37\x07\x50\x93\x94\x0b\x37\xf1\xd3\x2a\xa3\xd2\x8a\x30\x54\x1c\x22\x35\xd0\x17\xb3\xa5\xe3\x94\x9d\xb8\xcd\x56\xbf\xf5\x1d\x93\xe7\x3c\x16\x5f\x4d\x05\xea\x51\xca\x9c\x6f\xb4\x5e\x77\x21\xc7\xdc\x23\xdf\x62\x1b\xcb\xf8\xa7\x88\x86\xed\x28\x86\x0e\x6d\x09\xd6\xe0\x13\xa1\xe4\xdc\x7a\x6f\x41\xab\xc3\x0b\xcb\x77\xd3\xe3\xe9\x9e\x9b\xfd\xef\xc3\xb2\xf6\x3d\x72\xe3\x17\x3d\x52\x58\xf7\x52\x21\xfe\xbb\x55\xa7\x87\x42\xa1\x15\xaa\x6d\x7b\xb6\xb2\x7d\x30\x59\x8e\x59\x7d\x44\x56\x8b\x66\xf6\xc8\x1c\xea\x77\x3e\x66\x67\x72\xbb\x21\xe4\xdc\x90\xe0\x0b\xc4\x78\x77\x86\x6e\x6f\xe4\xf9\xbe\xe8\x92\x2e\xef\x88\x91\xcb\x3c\xfc\x36\xba\x87\xb8\x43\x72\x87\x70\x30\xb7\xe0\xfd\x47\x8f\xe9\xb1\xdd\x28\xc1\x62\xf7\x65\x46\x99\x73\xfb\x3c\x6f\xf7\x34\xee\xf1\xd5\x1f\xc6\x2b\xf3\x07\x9a\x63\x70\xc1\x17\x31\xd9\x2f\x1e\x66\x6c\x94\x0a\xe4\x66\x5e\x75\x6d\xb1\x0c\xae\x97\xd4\x1a\xd4\x72\x71\xc0\xac\xb1\xbb\x88\x91\x3e\x94\x97\x30\xcc\x51\x21\x0a\xc7\x68\x12\x3e\x13\xfa\x80\xb9\xc9\xd4\x04\x97\x95\xa6\x7a\x8f\x58\xf3\xaa\x32\x40\xd4\xdf\x13\x7b\x97\xf6\xdc\xbd\xe7\x42\x07\xc2\x25\x10\xf9\x7c\xd9\x1a\xaa\xed\xaf\xc9\xaa\xaa\xed\x9b\x44\x6a\xb4\x01\xfb\xb7\xe0\xbe\x3f\x6c\xf7\x2c\x34\x3a\xb7\xc1\x3f\xd5\xc4\xc3\xcf\xbf\x3c\x82\xbb\xc0\x3a\x60\x26\xaf\x0f\xc5\x88\x5d\xcd\x5f\x95\xfe\xa2\xa8\x58\xb2\x7b\xe7\x05\xc6\x6b\xbb\x30\x11\x90\xa7\x97\x06\x05\x6a\x5d\x58\xdc\x8b\x54\x32\xcb\xe8\xfa\xb3\x9f\x9b\xb0\xa4\x2a\x81\x1c\x6c\xf9\x29\xeb\xc3\x60\xe4\x47\x0d\x54\x72\xe8\xfb\x53\x3b\xe2\xe2\xb7\xf7\xef\x2a\xd1\xdf\x8e\xc9\x7d\xf4\xb6\xf3\xdb\xa7\xf6\x86\xf5\x02\x21\x1c\xc2\xd6\x4f\x28\x2f\x6a\xb3\xa9\x9e\xf1\x57\xb8\xa8\x5f\xac\x5e\xdc\x9d\xd9\x27\x6f\xb8\x94\x47\x2f\xe0\x3f\xc7\x16\x2e\x65\x84\xa4\x50\xc9\xc2\xc4\x32\x51\x5b\x8b\xa6\xee\xd8\xe4\xbe\xce\xf4\x53\xec\x08\x76\xf2\xcc\xd3\x8f\x71\x1f\xb8\x33\x8c\xa9\xd9\x1e\x18\x51\xa4\x52\x45\x8d\x8c\xbe\x13\x9f\x95\x12\xd6\x55\x9c\xa0\x3a\xe3\xd8\x0a\x51\xcf\x1a\xff\xcb\x0a\x92\x02\xdd\x9d\xf3\x6f\x38\x6f\x57\x73\xc1\x04\x09\x95\xf9\x3e\xf2\x6b\x19\x80\x33\x3d\xe9\x73\x7b\x26\xfb\xe7\xb9\x78\x38\xee\x0a\x1e\x61\x64\x82\xf0\x69\x92\x23\x0b\x42\x7f\x89\x8f\x73\x51\x5a\x01\xd9\x61\xc7\x3f\x6c\x55\xfa\x33\x3b\xa9\x09\x7f\xf1\x5d\x72\xae\x33\x99\xd4\xf7\xc9\xed\x73\x30\xa0\x88\xe6\xf4\xa9\x45\xff\x6f\xd8\xc8\x61\xc0\x57\x7e\x15\x48\xe4\xea\x80\x11\xff\xb8\x95\xac\xfd\x0f\x90\x83\xa4\x9f\x4a\xb6\x5c\x2e\x82\xdd\xaa\x94\xd9\xe6\xf4\x4d\x40\xd7\x32\xa8\x5b\x06\xa0\xab\x44\xdb\xbc\x61\x05\x5c\x06\xfa\x57\x8b\xb6\xf9\x0a\x78\xdd\x06\xd4\xac\x5d\x25\xda\xe6\xd3\x2b\xe0\x32\x70\x70\x95\xb8\x3c\xb7\x81\x96\x84\xcf\x4e\x5c\x8c\x09\xb2\x8f\xc9\xb3\x05\x1f\xc0\x62\xc2\xd0\xfc\xa3\xb1\x37\xff\x1b\x00\x00\xff\xff\xcf\xb1\x4d\x92\x1e\x06\x00\x00"

func assets_minecraft_textures_font_ascii_png_bytes() ([]byte, error) {
	return bindata_read(
		_assets_minecraft_textures_font_ascii_png,
		"assets/minecraft/textures/font/ascii.png",
	)
}

func assets_minecraft_textures_font_ascii_png() (*asset, error) {
	bytes, err := assets_minecraft_textures_font_ascii_png_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "assets/minecraft/textures/font/ascii.png", size: 1566, mode: os.FileMode(436), modTime: time.Unix(1432837142, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assets_steven_blockstates_missing_block_json = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xe6\x52\x00\x02\xa5\xb2\xc4\xa2\xcc\xc4\xbc\x92\x62\x25\x2b\x05\x88\x08\x58\x34\x2f\xbf\x28\x37\x31\x07\x24\xa6\xa0\x94\x9b\x9f\x92\x0a\x62\x2a\xe5\x66\x16\x17\x67\xe6\xa5\xc7\x27\xe5\xe4\x27\x67\x2b\x29\xd4\x82\x95\xd7\x72\xd5\x72\x01\x02\x00\x00\xff\xff\x11\x91\x43\x4e\x4b\x00\x00\x00"

func assets_steven_blockstates_missing_block_json_bytes() ([]byte, error) {
	return bindata_read(
		_assets_steven_blockstates_missing_block_json,
		"assets/steven/blockstates/missing_block.json",
	)
}

func assets_steven_blockstates_missing_block_json() (*asset, error) {
	bytes, err := assets_steven_blockstates_missing_block_json_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "assets/steven/blockstates/missing_block.json", size: 75, mode: os.FileMode(436), modTime: time.Unix(1432490704, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assets_steven_models_block_missing_block_json = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x90\xc1\xaa\x83\x30\x10\x45\xd7\x2f\x5f\x11\xe6\x6d\xb3\x68\x37\x5d\xf4\x57\x8a\x14\xb1\x63\x2b\xc4\xa4\x98\x04\x0b\xe2\xbf\x37\x33\xb1\x96\xd6\x40\x23\x0a\x7a\xe7\xcc\xc9\xc5\x49\xc8\x78\x81\xc7\x87\x0f\x03\x3a\x38\xca\x94\x70\x5a\x6b\x1d\x03\xe8\x3b\xe7\x3a\x73\x3d\x2f\x10\x30\x30\x2b\xf1\x07\xa8\xb1\x47\xe3\x69\xed\xb4\xae\x4d\xb4\xda\x0e\xb6\xa7\x54\xee\x14\xdf\xb2\x52\x2b\x90\x4e\xb4\x3c\xde\x1f\xd4\xf2\x6c\x88\xb6\x6e\xbe\x0a\xad\xa3\x8b\x1d\x4d\x9c\xc4\xa3\x5e\xcd\xa9\xe7\x3f\xf5\x55\x12\x9a\xa0\x35\x2d\x53\xc6\x24\x95\xdd\x38\xc2\x9d\x0c\x25\x8e\x48\x66\x0d\xc6\x0e\xfe\x46\x05\x7f\x1a\x12\x99\x95\x38\x1b\x0a\x25\x89\xcc\x4a\x46\x74\xbe\xec\x7f\x30\x99\x75\x60\x5d\xea\x60\x52\xce\x1f\x8a\xf7\x57\x7a\xab\xc4\x2c\x9e\x01\x00\x00\xff\xff\xef\x0a\x21\x1d\x5c\x02\x00\x00"

func assets_steven_models_block_missing_block_json_bytes() ([]byte, error) {
	return bindata_read(
		_assets_steven_models_block_missing_block_json,
		"assets/steven/models/block/missing_block.json",
	)
}

func assets_steven_models_block_missing_block_json() (*asset, error) {
	bytes, err := assets_steven_models_block_missing_block_json_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "assets/steven/models/block/missing_block.json", size: 604, mode: os.FileMode(436), modTime: time.Unix(1432491461, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assets_steven_textures_gui_cog_png = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xea\x0c\xf0\x73\xe7\xe5\x92\xe2\x62\x60\x60\xe0\xf5\xf4\x70\x09\x02\xd2\x22\x20\xcc\xc1\x06\x24\x7b\x3b\x65\x79\x81\x14\x4b\xb1\x93\x67\x08\x07\x10\xd4\x70\xa4\x74\x00\xf9\xd7\x3d\x5d\x1c\x43\x2c\x7a\x8f\x86\xdc\x16\x3c\xac\xc0\x13\x5c\x24\xd9\x20\x28\xd1\x28\x28\xa1\xba\xd8\x90\x4b\xd9\x80\x49\xb9\xa2\xf9\xcb\xe9\x95\x8f\x13\x5e\xc5\x15\x87\xff\x3d\x18\xe9\x54\x97\xf0\x40\x5e\x63\xa9\xcc\xd2\x29\x87\x85\x77\x57\x1c\x79\x52\x79\xff\xbb\x66\xfa\xe6\x13\xf3\x3b\xc2\xde\x98\xf1\xd5\x0b\x57\xbd\x0f\x3e\x1e\xdc\x2e\xb5\xb7\xfc\x4a\x94\x7e\xb9\x65\xfa\x11\x09\xdf\x73\xba\xba\x2f\x7d\xba\x7b\x5d\x4a\x33\xc3\xfa\x8d\x95\x98\x04\xe7\x3e\x8f\x7a\xbd\x56\xc4\xbc\x7d\x9a\xb1\x92\x53\x76\x2e\x63\xc3\x94\x4b\x82\x4e\x33\xac\x62\x97\xce\x68\x13\x61\x6e\xb8\x76\x4b\x69\xae\xd7\x86\x63\x06\x4f\xe3\x37\x30\x7d\x3b\xd3\x2a\x30\x21\xc8\xea\xee\xeb\x29\xaa\x11\xc9\x37\x9e\x06\xca\xfe\x64\x64\xbb\x7e\xfe\xea\x53\x09\xe5\xa2\x73\xe2\xcb\xc3\x8c\x4d\xf4\x5e\xed\xb2\xbc\x77\xb0\xac\x76\xef\xc4\xe0\x7d\xb3\x0f\x8a\x3c\x09\x5f\x9b\x9b\xdc\x7b\xeb\xf1\xeb\xd3\xad\xf9\x62\x6f\x1e\x57\x30\x9e\x2d\xeb\xad\x8f\x98\xc1\x62\x0d\xf4\x22\x83\xa7\xab\x9f\xcb\x3a\xa7\x84\x26\x40\x00\x00\x00\xff\xff\xa7\x92\x0b\x68\x20\x01\x00\x00"

func assets_steven_textures_gui_cog_png_bytes() ([]byte, error) {
	return bindata_read(
		_assets_steven_textures_gui_cog_png,
		"assets/steven/textures/gui/cog.png",
	)
}

func assets_steven_textures_gui_cog_png() (*asset, error) {
	bytes, err := assets_steven_textures_gui_cog_png_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "assets/steven/textures/gui/cog.png", size: 288, mode: os.FileMode(436), modTime: time.Unix(1432387105, 0)}
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
	"assets/minecraft/texts/splashes.txt":           assets_minecraft_texts_splashes_txt,
	"assets/minecraft/textures/font/ascii.png":      assets_minecraft_textures_font_ascii_png,
	"assets/steven/blockstates/missing_block.json":  assets_steven_blockstates_missing_block_json,
	"assets/steven/models/block/missing_block.json": assets_steven_models_block_missing_block_json,
	"assets/steven/textures/gui/cog.png":            assets_steven_textures_gui_cog_png,
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

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"assets": {nil, map[string]*_bintree_t{
		"minecraft": {nil, map[string]*_bintree_t{
			"texts": {nil, map[string]*_bintree_t{
				"splashes.txt": {assets_minecraft_texts_splashes_txt, map[string]*_bintree_t{}},
			}},
			"textures": {nil, map[string]*_bintree_t{
				"font": {nil, map[string]*_bintree_t{
					"ascii.png": {assets_minecraft_textures_font_ascii_png, map[string]*_bintree_t{}},
				}},
			}},
		}},
		"steven": {nil, map[string]*_bintree_t{
			"blockstates": {nil, map[string]*_bintree_t{
				"missing_block.json": {assets_steven_blockstates_missing_block_json, map[string]*_bintree_t{}},
			}},
			"models": {nil, map[string]*_bintree_t{
				"block": {nil, map[string]*_bintree_t{
					"missing_block.json": {assets_steven_models_block_missing_block_json, map[string]*_bintree_t{}},
				}},
			}},
			"textures": {nil, map[string]*_bintree_t{
				"gui": {nil, map[string]*_bintree_t{
					"cog.png": {assets_steven_textures_gui_cog_png, map[string]*_bintree_t{}},
				}},
			}},
		}},
	}},
}}

// Restore an asset under the given directory
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
