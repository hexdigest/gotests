// Code generated by go-bindata.
// sources:
// internal/render/templates/call.tmpl
// internal/render/templates/function.tmpl
// internal/render/templates/header.tmpl
// internal/render/templates/inline.tmpl
// internal/render/templates/inputs.tmpl
// internal/render/templates/message.tmpl
// internal/render/templates/results.tmpl
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

var _internalRenderTemplatesCallTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x8e\xe1\x8a\xc2\x40\x0c\x84\x5f\x25\x94\xfe\xb8\x83\x92\x07\x38\xb8\x07\xb8\x3f\x87\xa8\xe8\xef\x65\x9b\xd6\x40\x5d\x25\x8d\x8a\x84\xbc\xbb\xbb\xa5\xba\xbf\x02\x33\x5f\x66\xc6\xac\xa7\x81\x13\x41\x13\xc3\x34\x35\xee\x66\x0f\xd6\x13\xe0\x96\x22\xf1\x9d\xa4\x28\x3c\x40\xba\x28\xe0\xdf\xbc\x53\xb9\x45\x75\x57\x45\x33\x4a\x7d\x71\xdf\x24\xa0\x7b\x55\xf1\x3f\x9c\xc9\xfd\xcb\x4c\x42\x1a\x09\x5a\xee\xa0\xa5\x09\x7e\x7e\x01\x37\x41\xb2\xa9\x24\xf3\x9a\xde\xb2\x7b\x07\x9f\xdf\xda\x77\x14\xd6\xb2\x21\xf7\x05\x19\xe7\x1a\xbf\x44\x94\xc6\x85\xc6\xfd\xf3\x4a\x19\x3f\x04\xe1\xd0\x73\xcc\x43\xb0\xb2\xcb\xf9\x5e\xef\x2b\x00\x00\xff\xff\x65\x08\xbc\x88\xf1\x00\x00\x00")

func internalRenderTemplatesCallTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalRenderTemplatesCallTmpl,
		"internal/render/templates/call.tmpl",
	)
}

func internalRenderTemplatesCallTmpl() (*asset, error) {
	bytes, err := internalRenderTemplatesCallTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/render/templates/call.tmpl", size: 241, mode: os.FileMode(436), modTime: time.Unix(1511557124, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalRenderTemplatesFunctionTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x56\x5d\x4f\xdb\x4a\x10\x7d\x4e\x7e\xc5\x10\xc1\x95\x8d\x82\x79\x07\xf1\x70\xb9\xa0\xab\x3e\x94\x54\x09\x2a\x0f\x55\x55\x2d\xce\x38\x5d\xb1\x5e\xa7\xbb\x63\x10\xb2\xfc\xdf\x3b\xbb\xfe\xda\x24\x86\xf2\x52\x1e\x20\xfb\x31\x33\x67\xce\x39\x3b\xa4\xaa\xd6\x98\x49\x8d\x30\xcb\x4a\x9d\x92\x2c\xf4\xac\xae\xa7\x55\x75\x06\xc7\x19\x5c\x5c\x41\xc2\xab\xa9\x3b\x82\xaa\x4a\xee\xd1\xd2\x9d\xc8\xb1\xae\x23\x82\x53\xe2\x95\xd4\x9b\xe4\x3e\x86\x6a\x3a\x71\x21\x32\x03\x7f\xe7\x8b\x30\x7c\x8b\xd0\x58\x8e\x9e\xd0\xeb\x16\x41\x98\x8d\x05\x4b\xa6\x4c\xc9\xdd\xf6\xd7\x8d\xd0\x1b\x1c\x8b\x98\xb8\x73\xbf\xe5\xea\xfb\xca\x9c\xc3\x9f\xb8\x38\xd4\x6b\xf7\xb9\x9e\x86\x2b\x87\xc6\x3a\xc4\xdf\xbe\x07\x65\x34\x67\x75\x65\x19\x67\x1b\xfc\x22\xe9\x27\x24\x4b\x4c\x51\x3e\xa3\x69\xaa\x75\xd8\x3f\xd9\x55\x13\xca\xdb\xd0\xfe\x58\xa4\x72\x0b\x8e\x81\x28\x4f\xe1\x34\x97\x5a\xe6\x45\xfa\x94\xfc\x57\x68\x32\x85\x52\x68\xb8\xfd\xca\x77\xce\x18\x7d\x0a\xc1\x69\x4f\xab\xca\x03\x6b\xb1\x27\x5f\x85\x2a\x71\xa8\x86\xca\x62\xdf\x69\x07\xe6\xa0\xd9\xb0\xbf\x83\xcf\xe3\x64\x4f\x26\x9e\x69\xf7\x6b\x24\x26\x60\x7c\x89\xb6\x54\x64\xbb\x3a\x0f\x42\xd3\x7b\x64\xf7\x25\x97\xcc\x87\xd1\xf6\xd6\x98\xa2\x65\xef\x85\x43\x79\x09\x8f\x45\xa1\xf6\x14\x72\x12\x9c\x9f\xc3\xfd\xe2\x66\x71\x01\xff\xae\xd7\xe0\x54\x82\x54\x58\xb4\x89\x17\x30\x2b\x4c\xc3\x9e\x2e\xc8\xf1\x7f\x27\x9e\x90\x43\xe1\xc7\x1c\x88\x9c\x9a\x2d\x8d\x2d\xf2\x46\xe4\xaa\x17\xa7\x03\xb5\x2a\x1f\x9b\xa3\xba\xa6\x64\x59\xea\x88\x28\x71\xd2\xcf\x1b\xe1\xf6\xcc\x0a\x2d\x44\x38\x1b\x58\x1e\xb3\xc5\x81\x2f\x02\x5b\x00\xb0\x1b\x18\x5f\xef\x87\x3b\x7c\x19\x2c\x11\x51\x1c\xdc\xe4\x17\xc6\xe2\xe6\x69\xf2\x20\x24\x45\x24\x73\x4c\x56\x98\x16\x7a\x1d\xfb\x2a\xfb\x0e\xe0\xac\x0c\xdf\xfb\x8e\x2d\x17\xf7\x48\x3a\x25\x0e\x17\xad\xac\x23\x8f\xa8\xc3\xff\x60\x24\xf5\x6d\xed\x3c\x2e\xae\xf6\xcf\xe3\x2b\xd3\x93\x5c\x97\x19\xe3\xac\xea\x0f\x14\xe4\xa4\x82\xf9\x8b\xbc\x6a\x0b\xad\x5e\x43\x57\xc4\x87\xfb\x0b\x8d\xfe\x01\xc4\xd0\x23\x23\xcc\xb7\x4a\x10\x0f\x1e\xd3\x38\x71\xc6\x13\xc7\xfb\x6f\x38\x49\x85\x52\xcd\xf6\x5b\x28\x46\xec\x38\xe1\xed\xe6\x41\xee\x03\xe3\xec\xc8\x3e\xf5\xae\x1a\x2b\x72\xd9\xbb\x2d\x72\xf7\x8e\xae\x40\x4b\x15\xbb\xbf\xac\x47\xe7\xf2\xaa\xa1\x90\x12\x9f\x32\x8b\x66\x61\xae\x1c\xad\x15\x1b\x6c\x5b\x41\x77\x03\xae\xe0\xe4\x79\x0e\x5d\xf8\xc9\xf3\x6c\xbe\x53\x5e\xea\x6d\xd9\x37\xcf\x11\xf3\xa0\x58\xef\x8f\xe1\xb5\xef\x3c\x5b\x7f\xb6\xe7\x7f\xe3\x1b\xe6\x46\xfc\x84\x61\x97\xb1\xed\x4b\x6c\x1b\x1b\xf2\x0d\xcb\x77\x2d\x75\x58\xf2\x2d\x4f\x79\xd6\xff\x2f\xa8\x33\x55\xe0\xb1\x64\xe5\x07\x70\x14\x5f\x06\x57\x1a\x56\xc3\xc1\x33\xf8\x8e\xa1\xb7\x35\xae\x85\x95\xe9\x30\x0f\x07\x71\x8f\xb3\x31\x7f\xb9\x79\xbb\x83\x21\xe4\x59\xf1\xbf\xb9\x7d\xa1\x3f\x8c\xe7\x2f\xd5\x3f\x32\x98\x29\x4c\x29\xb9\x41\xdc\xde\xfe\x2a\x85\x8a\xfa\x0c\xf3\x5d\x40\x71\x88\xa8\x57\xef\x23\x3e\xec\x00\xb7\x60\x3f\xb3\x9a\x72\xab\x76\xc0\xb6\x78\x06\xaf\xfe\xc1\xa8\x6f\x82\x1c\xb7\xd4\xc8\x9c\x06\xd7\x50\x38\x87\xeb\x29\x7f\xd3\xe8\x7c\xfa\x3b\x00\x00\xff\xff\x30\xde\x62\xfa\x97\x08\x00\x00")

func internalRenderTemplatesFunctionTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalRenderTemplatesFunctionTmpl,
		"internal/render/templates/function.tmpl",
	)
}

func internalRenderTemplatesFunctionTmpl() (*asset, error) {
	bytes, err := internalRenderTemplatesFunctionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/render/templates/function.tmpl", size: 2199, mode: os.FileMode(436), modTime: time.Unix(1513585317, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalRenderTemplatesHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\x48\x4d\x4c\x49\x2d\x52\xaa\xad\xe5\xaa\xae\x2e\x4a\xcc\x4b\x4f\x55\xd0\x73\xce\xcf\xcd\x4d\xcd\x2b\x29\xae\xad\xad\xae\xd6\x03\x4b\xa4\xe6\xa5\x00\xe9\x82\xc4\xe4\xec\x44\xa0\x02\xa0\x68\x00\x84\x09\x14\xe4\xca\xcc\x2d\xc8\x2f\x2a\x51\xd0\x40\xe8\xf7\x04\x8b\x40\xb4\xfb\x25\xe6\x02\x55\x41\xb4\x94\x64\x20\x19\xa6\x09\x67\x01\x02\x00\x00\xff\xff\x18\xfd\x24\x71\x8c\x00\x00\x00")

func internalRenderTemplatesHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalRenderTemplatesHeaderTmpl,
		"internal/render/templates/header.tmpl",
	)
}

func internalRenderTemplatesHeaderTmpl() (*asset, error) {
	bytes, err := internalRenderTemplatesHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/render/templates/header.tmpl", size: 140, mode: os.FileMode(436), modTime: time.Unix(1511557124, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalRenderTemplatesInlineTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\xcc\xcb\x01\xd2\x4a\xb5\xb5\x0a\xd5\xd5\x25\xa9\xb9\x05\x39\x89\x25\x40\xd1\xe4\xc4\x9c\x1c\x25\x05\x3d\xb0\x68\x6a\x5e\x4a\x6d\x2d\x20\x00\x00\xff\xff\xaa\xeb\x41\xff\x31\x00\x00\x00")

func internalRenderTemplatesInlineTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalRenderTemplatesInlineTmpl,
		"internal/render/templates/inline.tmpl",
	)
}

func internalRenderTemplatesInlineTmpl() (*asset, error) {
	bytes, err := internalRenderTemplatesInlineTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/render/templates/inline.tmpl", size: 49, mode: os.FileMode(436), modTime: time.Unix(1511557124, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalRenderTemplatesInputsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8d\x41\x0a\x02\x31\x0c\x45\xaf\x12\x86\x59\x4a\x0e\x20\x78\x00\x77\x82\x27\xa8\x4c\x3a\x74\x61\x94\xf4\xcf\xea\xd3\xbb\x6b\xab\x8b\xae\x12\x1e\x2f\x2f\xe4\x66\xb9\xb8\xc9\x52\xfc\x7d\xa0\x2e\xad\x91\x6b\x96\xf3\x45\xb4\xaf\x25\x8b\xbf\x20\x7a\x3f\x1e\xb0\x8a\xda\x1a\xa0\x9e\x9e\x76\x12\xd2\x7c\xfb\x3b\x6b\xd6\x5b\x14\xc7\x75\x44\x3a\x8c\xe4\xbb\x0d\x9e\xe2\xab\xc3\xe2\x77\x9b\x62\xaf\x4a\x0e\xda\x5f\x4c\x9d\x79\x7c\x02\x00\x00\xff\xff\x8e\xbc\xcf\xda\x98\x00\x00\x00")

func internalRenderTemplatesInputsTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalRenderTemplatesInputsTmpl,
		"internal/render/templates/inputs.tmpl",
	)
}

func internalRenderTemplatesInputsTmpl() (*asset, error) {
	bytes, err := internalRenderTemplatesInputsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/render/templates/inputs.tmpl", size: 152, mode: os.FileMode(436), modTime: time.Unix(1511557124, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalRenderTemplatesMessageTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x3c\x8e\x41\x8a\xc3\x30\x0c\x45\xf7\x73\x0a\x11\x12\x98\x81\x44\x07\x18\x98\x03\xcc\xa6\x94\xb6\x74\xef\x36\x3f\xa9\x21\x71\x53\xdb\x49\x29\x42\x77\xaf\x63\x68\x56\x12\x5f\x4f\x4f\x12\x69\xd1\x59\x07\x2a\x46\x84\x60\x7a\x14\xd4\xa8\x7e\x89\xd8\x8e\xdc\x3d\x12\x1f\xe7\x4b\x44\x88\x41\xb5\x7a\x30\x89\xc0\xb5\xaa\x22\x4f\x1b\x6f\xc4\x07\x5c\x61\x17\xf8\x35\xe1\xd3\x6b\x02\x9f\xcd\x30\x43\x95\x37\x90\x77\x66\x4c\xc1\x77\x36\xf2\xde\x5b\x17\xff\xdd\x34\xaf\x42\x11\x6f\x5c\x0f\x2a\x6d\x4d\x25\x06\xfa\xfd\x4b\x80\xf1\x89\x8f\xf0\x79\x9e\x56\x4a\xab\x5a\x7f\xee\x56\xcb\xe6\xcd\xe5\x27\x3d\xda\x50\x6e\xdf\x01\x00\x00\xff\xff\x90\x2e\xb9\x52\xc9\x00\x00\x00")

func internalRenderTemplatesMessageTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalRenderTemplatesMessageTmpl,
		"internal/render/templates/message.tmpl",
	)
}

func internalRenderTemplatesMessageTmpl() (*asset, error) {
	bytes, err := internalRenderTemplatesMessageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/render/templates/message.tmpl", size: 201, mode: os.FileMode(436), modTime: time.Unix(1511557124, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalRenderTemplatesResultsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\xcd\x41\x0a\xc2\x40\x0c\x05\xd0\xab\x7c\x4a\x97\xa5\x07\x10\x5c\x8a\x7b\x6f\x20\x34\x95\x81\x92\x81\x3f\xd3\x55\xc8\xdd\x4d\x6a\x51\x70\x35\x93\xfc\x97\xc4\x6c\x91\xb5\xa8\x60\xa0\xb4\x7d\xeb\x6d\x70\x87\x19\x9f\xfa\x12\x8c\x65\xc2\x28\x1b\x2e\x57\xcc\x8f\x4f\xec\x6e\x56\xd6\x48\xdc\xa7\x70\xa2\x4b\x76\xee\xb5\x63\xce\xcf\x59\x87\x88\x81\xbe\x53\xdb\x8d\xac\x4c\x2c\xe4\x99\xe3\x00\x95\xdf\xa5\xff\x38\x0f\xfe\xec\xf1\xbe\x03\x00\x00\xff\xff\xb0\x4f\xcf\x61\xa8\x00\x00\x00")

func internalRenderTemplatesResultsTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalRenderTemplatesResultsTmpl,
		"internal/render/templates/results.tmpl",
	)
}

func internalRenderTemplatesResultsTmpl() (*asset, error) {
	bytes, err := internalRenderTemplatesResultsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/render/templates/results.tmpl", size: 168, mode: os.FileMode(436), modTime: time.Unix(1511557124, 0)}
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
	"internal/render/templates/call.tmpl": internalRenderTemplatesCallTmpl,
	"internal/render/templates/function.tmpl": internalRenderTemplatesFunctionTmpl,
	"internal/render/templates/header.tmpl": internalRenderTemplatesHeaderTmpl,
	"internal/render/templates/inline.tmpl": internalRenderTemplatesInlineTmpl,
	"internal/render/templates/inputs.tmpl": internalRenderTemplatesInputsTmpl,
	"internal/render/templates/message.tmpl": internalRenderTemplatesMessageTmpl,
	"internal/render/templates/results.tmpl": internalRenderTemplatesResultsTmpl,
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
	"internal": &bintree{nil, map[string]*bintree{
		"render": &bintree{nil, map[string]*bintree{
			"templates": &bintree{nil, map[string]*bintree{
				"call.tmpl": &bintree{internalRenderTemplatesCallTmpl, map[string]*bintree{}},
				"function.tmpl": &bintree{internalRenderTemplatesFunctionTmpl, map[string]*bintree{}},
				"header.tmpl": &bintree{internalRenderTemplatesHeaderTmpl, map[string]*bintree{}},
				"inline.tmpl": &bintree{internalRenderTemplatesInlineTmpl, map[string]*bintree{}},
				"inputs.tmpl": &bintree{internalRenderTemplatesInputsTmpl, map[string]*bintree{}},
				"message.tmpl": &bintree{internalRenderTemplatesMessageTmpl, map[string]*bintree{}},
				"results.tmpl": &bintree{internalRenderTemplatesResultsTmpl, map[string]*bintree{}},
			}},
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

