// Code generated by go-bindata.
// sources:
// templates/call.tmpl
// templates/function.tmpl
// templates/header.tmpl
// templates/inline.tmpl
// templates/inputs.tmpl
// templates/message.tmpl
// templates/results.tmpl
// DO NOT EDIT!

package bindata

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

var _templatesCallTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x8e\xe1\x8a\xc2\x40\x0c\x84\x5f\x25\x94\xfe\xb8\x83\x92\x07\x38\xb8\x07\xb8\x3f\x87\xa8\xe8\xef\x65\x9b\xd6\x40\x5d\x25\x8d\x8a\x84\xbc\xbb\xbb\xa5\xba\xbf\x02\x33\x5f\x66\xc6\xac\xa7\x81\x13\x41\x13\xc3\x34\x35\xee\x66\x0f\xd6\x13\xe0\x96\x22\xf1\x9d\xa4\x28\x3c\x40\xba\x28\xe0\xdf\xbc\x53\xb9\x45\x75\x57\x45\x33\x4a\x7d\x71\xdf\x24\xa0\x7b\x55\xf1\x3f\x9c\xc9\xfd\xcb\x4c\x42\x1a\x09\x5a\xee\xa0\xa5\x09\x7e\x7e\x01\x37\x41\xb2\xa9\x24\xf3\x9a\xde\xb2\x7b\x07\x9f\xdf\xda\x77\x14\xd6\xb2\x21\xf7\x05\x19\xe7\x1a\xbf\x44\x94\xc6\x85\xc6\xfd\xf3\x4a\x19\x3f\x04\xe1\xd0\x73\xcc\x43\xb0\xb2\xcb\xf9\x5e\xef\x2b\x00\x00\xff\xff\x65\x08\xbc\x88\xf1\x00\x00\x00")

func templatesCallTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesCallTmpl,
		"templates/call.tmpl",
	)
}

func templatesCallTmpl() (*asset, error) {
	bytes, err := templatesCallTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/call.tmpl", size: 241, mode: os.FileMode(436), modTime: time.Unix(1511557124, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFunctionTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x56\xcf\x4f\xe3\x3a\x10\x3e\xb7\x7f\xc5\x50\xc1\x53\x82\x8a\xb9\x83\x38\x3c\x1e\xe8\x69\x0f\x4b\x57\x2d\x5a\x0e\xab\xd5\xca\xa4\x93\xae\x85\xe3\x74\x6d\x07\x84\xa2\xfc\xef\x3b\x76\x9c\xc4\x6d\x03\xcb\x65\x39\x40\xfd\x63\xe6\xfb\x66\xbe\xcf\x43\xeb\x7a\x8d\xb9\x50\x08\xb3\xbc\x52\x99\x15\xa5\x9a\x35\xcd\xb4\xae\xcf\xe0\x38\x87\x8b\x2b\x60\xb4\x9a\xba\x23\xa8\x6b\x76\x8f\xc6\xde\xf1\x02\x9b\x26\xb1\x70\x6a\x69\x25\xd4\x86\xdd\xa7\x50\x4f\x27\x2e\x44\xe4\xe0\xef\x7c\xe1\x9a\x6e\x59\xd4\x86\xa2\x27\xf6\x75\x8b\xc0\xf5\xc6\x80\xb1\xba\xca\xac\xbb\xed\xaf\x6b\xae\x36\x38\x16\x31\x71\xe7\x7e\xcb\xe1\x7b\x64\xca\xe1\x4f\x5c\x1c\xaa\xb5\xfb\xdc\x4c\xe3\x95\x63\x63\x1c\xe3\x6f\xdf\x23\x18\x45\x59\x1d\x2c\xf1\x0c\xc1\x2f\xc2\xfe\x04\xb6\xc4\x0c\xc5\x33\xea\x16\xad\xe3\xfe\xc9\xac\xda\x50\xda\x86\xf0\x63\xd0\x56\x5b\x70\x1d\x48\x8a\x0c\x4e\x0b\xa1\x44\x51\x66\x4f\xec\xbf\x52\x59\x5d\x4a\x89\x9a\xca\xaf\x7d\xe5\xc4\xd1\xa7\xe0\x94\xf6\xb4\xae\x3d\xb1\xc0\x9d\x7d\xe5\xb2\xc2\x01\x0d\xa5\xc1\xbe\xd2\x8e\xcc\x41\xb1\x71\x7d\x07\x9f\xc7\x9b\x3d\x99\xf8\x4e\xbb\x5f\x23\x31\x51\xc7\x97\x68\x2a\x69\x4d\x87\xf3\xc0\x95\x7d\xaf\xd9\x3d\xe4\x92\xfa\xa1\x95\xb9\xd5\xba\x0c\xdd\x7b\xa1\x50\x5a\xc2\x63\x59\xca\x3d\x85\x9c\x04\xe7\xe7\x70\xbf\xb8\x59\x5c\xc0\xbf\xeb\x35\x38\x95\x20\xe3\x06\x0d\xf3\x02\xe6\xa5\x6e\xbb\xa7\x4a\xeb\xfa\x7f\xc7\x9f\x90\x42\xe1\xc7\x1c\xac\x75\x6a\x86\x36\x06\xe6\xad\xc8\x75\x2f\x4e\x47\x6a\x55\x3d\xb6\x47\x4d\x63\xd9\xb2\x52\x89\xb5\xcc\x49\x3f\x6f\x85\xdb\x33\x2b\x04\x8a\x70\x16\xe9\x0c\x40\xf2\x12\x60\x2f\xf0\x1d\xbe\x0c\x1a\x27\x36\x8d\x6e\xd2\x93\x21\xb5\x8a\x8c\x3d\x70\x61\x13\x2b\x0a\x64\x2b\xcc\x4a\xb5\x4e\x3b\xd1\xc6\x5c\x76\x60\xb3\xb0\xbd\xef\x00\x22\x41\xf4\xbd\xef\xc8\x72\x69\x1f\xda\x29\x71\xb8\x08\xb2\x8e\x3c\xa2\x0e\xf0\x41\x0b\xdb\xf3\xd8\x79\x5c\x84\xf6\xcf\xe3\x2b\xb5\x87\x5d\x57\x39\x95\x55\x37\x1f\x00\xa4\xa4\x9c\xfa\x97\x78\xd5\x16\x4a\xbe\xc6\xae\x48\x0f\xf7\x17\x0a\xfd\x03\x48\xa1\x67\x66\xb1\xd8\x4a\x6e\x69\xf0\xe8\xd6\x89\x33\x9a\x38\xde\x7f\xc3\x49\xc6\xa5\x6c\xb7\xdf\x62\x31\x62\xc7\x09\x6d\xb7\x0f\x72\x9f\x18\x65\x47\xf2\xa9\x77\xd5\x18\xc8\x65\xef\xb6\xc4\xdd\x3b\xba\x02\x25\x64\xea\xfe\x92\x1e\x9d\xcb\xeb\xb6\x85\x96\xf9\x94\x79\x32\x8b\x73\x15\x68\x0c\xdf\x60\x28\x05\xdd\x0d\xb8\x82\x93\xe7\x39\x74\xe1\x27\xcf\xb3\xf9\x0e\xbc\x50\xdb\xaa\x2f\x9e\x22\xe6\x11\x58\xda\xc9\x35\xbc\xf6\x9d\x67\xeb\xcf\xf6\xfc\xaf\x7d\xc1\x54\x88\x9f\x30\x64\x4a\xb2\x7d\x85\xa1\xb0\x21\xdf\xb0\x7c\xd7\x52\x87\x90\x6f\x79\xca\x77\xfd\xff\xd2\x76\xa6\x8a\x3c\xc6\x56\x7e\x00\x27\xe9\x65\x74\xa5\xed\x6a\x3c\x78\x06\xdf\x11\xf5\x80\x71\xcd\x8d\xc8\x86\x79\x38\x88\x7b\x9c\x8f\xf9\xcb\xcd\xdb\x1d\x0e\x71\x9f\x25\xfd\x9b\xdb\x17\xfa\xc3\x7c\xfe\x12\xfe\x91\xc6\x5c\x62\x66\xd9\x0d\xe2\xf6\xf6\x57\xc5\x65\xd2\x67\x98\xef\x12\x4a\x63\x46\xbd\x7a\x1f\xf1\x61\x47\x38\x90\xfd\x4c\x6a\x8a\xad\xdc\x21\x1b\xf8\x0c\x5e\xfd\x83\x51\xdf\x24\x39\x6e\xa9\x91\x39\x0d\xae\xa0\x68\x0e\x53\x10\x7d\xd3\xe8\x7c\xfa\x3b\x00\x00\xff\xff\xa3\xe0\xed\x1b\x97\x08\x00\x00")

func templatesFunctionTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesFunctionTmpl,
		"templates/function.tmpl",
	)
}

func templatesFunctionTmpl() (*asset, error) {
	bytes, err := templatesFunctionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/function.tmpl", size: 2199, mode: os.FileMode(436), modTime: time.Unix(1511557717, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\x48\x4d\x4c\x49\x2d\x52\xaa\xad\xe5\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\xd0\x73\xce\xcf\xcd\x4d\xcd\x2b\x29\xae\xad\xad\xae\xd6\x03\x4b\xa4\xe6\xa5\x00\xe9\x82\xc4\xe4\xec\x44\xa0\x02\xa0\x68\x00\x84\x09\x14\xe4\xca\xcc\x2d\xc8\x2f\x2a\x51\xd0\x40\xe8\xf7\x04\x8b\x40\xb4\xfb\x25\xe6\x02\x55\x41\xb4\x94\x64\x20\x19\xa6\x09\x67\x01\x02\x00\x00\xff\xff\x18\xfd\x24\x71\x8c\x00\x00\x00")

func templatesHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHeaderTmpl,
		"templates/header.tmpl",
	)
}

func templatesHeaderTmpl() (*asset, error) {
	bytes, err := templatesHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/header.tmpl", size: 140, mode: os.FileMode(436), modTime: time.Unix(1511557124, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInlineTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\xcc\xcb\x01\xd2\x4a\xb5\xb5\x0a\xd5\xd5\x25\xa9\xb9\x05\x39\x89\x25\x40\xd1\xe4\xc4\x9c\x1c\x25\x05\x3d\xb0\x68\x6a\x5e\x4a\x6d\x2d\x20\x00\x00\xff\xff\xaa\xeb\x41\xff\x31\x00\x00\x00")

func templatesInlineTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesInlineTmpl,
		"templates/inline.tmpl",
	)
}

func templatesInlineTmpl() (*asset, error) {
	bytes, err := templatesInlineTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/inline.tmpl", size: 49, mode: os.FileMode(436), modTime: time.Unix(1511557124, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInputsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8d\x41\x0a\x02\x31\x0c\x45\xaf\x12\x86\x59\x4a\x0e\x20\x78\x00\x77\x82\x27\xa8\x4c\x3a\x74\x61\x94\xf4\xcf\xea\xd3\xbb\x6b\xab\x8b\xae\x12\x1e\x2f\x2f\xe4\x66\xb9\xb8\xc9\x52\xfc\x7d\xa0\x2e\xad\x91\x6b\x96\xf3\x45\xb4\xaf\x25\x8b\xbf\x20\x7a\x3f\x1e\xb0\x8a\xda\x1a\xa0\x9e\x9e\x76\x12\xd2\x7c\xfb\x3b\x6b\xd6\x5b\x14\xc7\x75\x44\x3a\x8c\xe4\xbb\x0d\x9e\xe2\xab\xc3\xe2\x77\x9b\x62\xaf\x4a\x0e\xda\x5f\x4c\x9d\x79\x7c\x02\x00\x00\xff\xff\x8e\xbc\xcf\xda\x98\x00\x00\x00")

func templatesInputsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesInputsTmpl,
		"templates/inputs.tmpl",
	)
}

func templatesInputsTmpl() (*asset, error) {
	bytes, err := templatesInputsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/inputs.tmpl", size: 152, mode: os.FileMode(436), modTime: time.Unix(1511557124, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMessageTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x3c\x8e\x41\x8a\xc3\x30\x0c\x45\xf7\x73\x0a\x11\x12\x98\x81\x44\x07\x18\x98\x03\xcc\xa6\x94\xb6\x74\xef\x36\x3f\xa9\x21\x71\x53\xdb\x49\x29\x42\x77\xaf\x63\x68\x56\x12\x5f\x4f\x4f\x12\x69\xd1\x59\x07\x2a\x46\x84\x60\x7a\x14\xd4\xa8\x7e\x89\xd8\x8e\xdc\x3d\x12\x1f\xe7\x4b\x44\x88\x41\xb5\x7a\x30\x89\xc0\xb5\xaa\x22\x4f\x1b\x6f\xc4\x07\x5c\x61\x17\xf8\x35\xe1\xd3\x6b\x02\x9f\xcd\x30\x43\x95\x37\x90\x77\x66\x4c\xc1\x77\x36\xf2\xde\x5b\x17\xff\xdd\x34\xaf\x42\x11\x6f\x5c\x0f\x2a\x6d\x4d\x25\x06\xfa\xfd\x4b\x80\xf1\x89\x8f\xf0\x79\x9e\x56\x4a\xab\x5a\x7f\xee\x56\xcb\xe6\xcd\xe5\x27\x3d\xda\x50\x6e\xdf\x01\x00\x00\xff\xff\x90\x2e\xb9\x52\xc9\x00\x00\x00")

func templatesMessageTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesMessageTmpl,
		"templates/message.tmpl",
	)
}

func templatesMessageTmpl() (*asset, error) {
	bytes, err := templatesMessageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/message.tmpl", size: 201, mode: os.FileMode(436), modTime: time.Unix(1511557124, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResultsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\xcd\x41\x0a\xc2\x40\x0c\x05\xd0\xab\x7c\x4a\x97\xa5\x07\x10\x5c\x8a\x7b\x6f\x20\x34\x95\x81\x92\x81\x3f\xd3\x55\xc8\xdd\x4d\x6a\x51\x70\x35\x93\xfc\x97\xc4\x6c\x91\xb5\xa8\x60\xa0\xb4\x7d\xeb\x6d\x70\x87\x19\x9f\xfa\x12\x8c\x65\xc2\x28\x1b\x2e\x57\xcc\x8f\x4f\xec\x6e\x56\xd6\x48\xdc\xa7\x70\xa2\x4b\x76\xee\xb5\x63\xce\xcf\x59\x87\x88\x81\xbe\x53\xdb\x8d\xac\x4c\x2c\xe4\x99\xe3\x00\x95\xdf\xa5\xff\x38\x0f\xfe\xec\xf1\xbe\x03\x00\x00\xff\xff\xb0\x4f\xcf\x61\xa8\x00\x00\x00")

func templatesResultsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesResultsTmpl,
		"templates/results.tmpl",
	)
}

func templatesResultsTmpl() (*asset, error) {
	bytes, err := templatesResultsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/results.tmpl", size: 168, mode: os.FileMode(436), modTime: time.Unix(1511557124, 0)}
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
	"templates/call.tmpl": templatesCallTmpl,
	"templates/function.tmpl": templatesFunctionTmpl,
	"templates/header.tmpl": templatesHeaderTmpl,
	"templates/inline.tmpl": templatesInlineTmpl,
	"templates/inputs.tmpl": templatesInputsTmpl,
	"templates/message.tmpl": templatesMessageTmpl,
	"templates/results.tmpl": templatesResultsTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"call.tmpl": &bintree{templatesCallTmpl, map[string]*bintree{}},
		"function.tmpl": &bintree{templatesFunctionTmpl, map[string]*bintree{}},
		"header.tmpl": &bintree{templatesHeaderTmpl, map[string]*bintree{}},
		"inline.tmpl": &bintree{templatesInlineTmpl, map[string]*bintree{}},
		"inputs.tmpl": &bintree{templatesInputsTmpl, map[string]*bintree{}},
		"message.tmpl": &bintree{templatesMessageTmpl, map[string]*bintree{}},
		"results.tmpl": &bintree{templatesResultsTmpl, map[string]*bintree{}},
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

