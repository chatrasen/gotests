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

var _templatesCallTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x8e\x41\x6a\xc3\x40\x0c\x45\xaf\x22\x8c\x17\x2d\x18\x1d\xa0\xd0\x03\x78\x53\x4a\x5b\xda\xb5\x18\xcb\xae\xc0\x9e\x16\x8d\x92\x10\x84\xee\x1e\xc6\x38\x99\xd5\x87\x3f\x6f\xde\x97\xfb\xc4\xb3\x64\x86\x2e\xd1\xba\x76\x11\xee\x17\xb1\x5f\xc0\x0f\x4e\x2c\x67\xd6\xda\xc8\x0c\xf9\xcf\x00\xc7\xf2\x69\x7a\x4a\x16\x61\x86\xee\x9c\xa7\xfa\x7a\x27\x01\x23\x5a\x8b\x6f\xb4\x71\xc4\x93\xbb\x52\x5e\x18\x7a\x19\xa0\xe7\x15\x5e\x5e\x01\xdf\x49\x69\x63\x63\x2d\x87\xbd\x97\x88\x01\x1e\x7f\xdb\xde\x8f\x8a\xd5\x1b\xcc\x90\x74\x29\x4d\xbf\x2b\xea\xe2\x4e\xe3\xd7\xf5\x9f\x71\x2c\xdf\xa4\x42\x93\xa4\x08\xc4\xc6\xee\xf1\x7c\xe4\x2d\x00\x00\xff\xff\x65\x08\xbc\x88\xf1\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/call.tmpl", size: 241, mode: os.FileMode(436), modTime: time.Unix(1531946760, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFunctionTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x4d\x4f\xe3\x3a\x14\x5d\x27\xbf\xe2\x52\x01\x6a\x9f\x8a\xd9\x17\x75\xf1\x10\xbc\xa7\xb7\x78\x74\x54\xd0\xb0\x18\x8d\x46\x26\xbd\xe9\x58\xe3\x3a\x1d\xfb\x06\x54\x59\xfe\xef\x23\x3b\xce\x67\x53\x86\xcd\xb0\x20\xc9\xf5\xc7\x39\x3e\xf7\x1c\x83\xb5\x1b\xcc\x85\x42\x98\xe4\xa5\xca\x48\x14\x6a\xe2\x5c\x6a\xed\x15\x9c\xe7\xb0\x58\x02\x73\x2e\x4d\xfd\x10\x58\xcb\x9e\xd0\xd0\x03\xdf\xa1\x73\x53\x82\xbf\x08\x0d\x09\xb5\x65\x4f\x33\xb0\x69\xe2\x97\xbc\x09\xfa\x0e\x6c\x8d\x19\x8a\x57\xd4\xce\xa5\x49\x28\x8b\x1c\xd8\x7f\xe6\x91\x74\x99\x51\x28\x36\xd5\x7f\x04\xca\x8d\xa9\x6a\x09\x1d\xf6\x08\x79\xa8\x80\x09\x93\xfd\xbe\x71\xb6\xe6\x6a\x8b\x83\x05\x89\xb5\xe1\xdb\x93\x0c\xf4\x0e\x7b\x8c\x43\x7e\x09\xaa\x4d\xfc\x6a\x30\xeb\x52\xe7\x7d\xf0\xea\x59\xf9\x63\x7e\xe2\x9a\xef\x90\x50\x07\xb0\x40\x8d\xeb\x6d\x8f\x58\x87\xd6\xf1\x8a\x00\x18\x4a\x47\xec\x3a\x88\x7d\x7c\x2f\xa8\xf1\xa2\x7f\xf9\xda\x81\x51\x7c\x87\x1e\x56\xa8\x6d\x5c\x3c\x22\x73\xcd\x9d\xab\x4d\xab\xf5\x40\xae\x28\x6d\xf5\x68\x14\x91\xa6\xd5\xac\xde\xf2\x58\xd0\x13\xe2\xbd\x23\x59\x92\x04\xbd\xfc\xaf\x91\x35\x1d\xdd\xd6\x68\x4a\x49\xa6\xc6\x79\xe6\x8a\xde\x93\xac\x81\x5c\x23\x95\x5a\x99\x7b\xad\x8b\xa8\xc1\x1b\x57\x74\xaf\x35\xbc\x14\x85\x1c\xe8\xec\x85\xbc\xbe\x86\xa7\xd5\xdd\x6a\x01\x7f\x6f\x36\xe0\xb5\x86\x8c\x1b\x34\x2c\xb4\x21\x2f\x34\x58\x2b\x72\x50\x05\x79\x05\x1f\xf8\x0f\xdc\x38\x07\xdf\xe6\x40\xe4\x7b\x62\x6d\xd8\x2b\x32\xaf\x5a\x65\x53\x88\x3f\x35\xa9\xc7\xf2\xa5\x1a\x72\x8e\xd8\xba\x54\x53\x22\xe6\x1b\x38\x07\x9f\xa1\x61\x6a\x20\x52\x84\xab\x56\xe5\xb1\xe6\x9e\x48\xd1\x51\xcb\x02\xcd\xd0\x8e\xc3\x1e\xc3\x64\xae\x9d\xbb\x8c\xd4\xa3\xa0\xec\x33\x97\x25\x3a\x67\xeb\x2d\x4e\x84\x2b\xb1\x96\x55\x61\x5f\x00\x11\xab\x6c\xc3\x3a\x91\x9b\xb7\x1b\xb4\x51\x4b\x46\xf2\x77\xf4\x11\xf1\x46\x12\x53\x1f\xf3\x59\x0b\x6a\x4e\xdf\x4b\xd2\x62\x09\x97\x2f\x07\x42\xc3\x6e\xcb\x3c\x47\x6d\x3f\x02\x18\x93\x31\x0d\xcd\x5d\x29\x79\xe8\x9a\x67\x76\x5c\x5f\x29\x0c\x2a\xcd\xa0\x61\x46\xb8\xdb\x4b\x4e\x08\x13\x5d\x19\x76\x02\xe7\x79\xb0\x69\x3b\x92\x71\x29\xab\xf2\x29\x16\x23\xae\x4d\x44\x1e\xbb\x36\x24\xe6\x1c\xa0\xd6\x55\x57\xc7\x40\x6e\x1a\x53\x4e\xfd\xbc\xb3\x25\x28\x21\x67\xfe\x49\xc4\xea\x30\xc4\x36\x13\x0b\x5b\xe6\xd3\x49\x77\xaf\x1d\x1a\xc3\xb7\x18\x8f\x82\x7e\x06\x2c\xe1\xe2\x75\x0e\xf5\xf2\x8b\xd7\xc9\xbc\x07\x2f\xd4\xbe\x6c\x0e\x8f\x5a\xcf\x3b\x60\xb3\xd6\x11\xf5\xa5\xd0\x4b\x77\x18\x1b\xc4\x44\x87\x03\x5b\x5b\x5d\x44\x59\xa1\x48\xa8\x12\xe3\xc1\xc6\x1c\xf6\xae\xa5\x8e\x21\x4f\x79\x2a\xa8\xfe\x6f\x41\x6d\x70\x1a\x8f\xb1\xc7\x70\xdb\x4e\x67\x37\x9d\x29\x95\xaa\xdd\xfb\xa9\xf5\x9d\x34\x18\x31\x6e\xb9\x11\x59\xe7\xef\x50\xd3\xdc\xf3\x7c\xcc\x5f\x3e\x94\x3d\x0e\x5d\x9d\xa5\x50\x38\x6c\xf4\x87\xf9\xfc\x21\xfc\x33\x8d\xb9\xc4\x8c\xd8\x1d\xe2\xfe\xfe\x67\xc9\xe5\xb4\xd9\x61\xde\x27\x34\xeb\x32\x6a\xba\xf7\x11\x1f\xd6\x84\x23\xd9\xff\x4b\x49\x62\x2f\x7b\x64\x23\x9f\xd6\xab\xbf\x31\xea\x49\x92\xa7\xff\x3f\x18\xf8\x14\xfc\x81\xba\xd7\xb5\x4b\x5d\x9a\xd6\x3e\xfd\x15\x00\x00\xff\xff\x12\xba\x7c\xa5\x47\x09\x00\x00")

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

	info := bindataFileInfo{name: "templates/function.tmpl", size: 2375, mode: os.FileMode(436), modTime: time.Unix(1531946760, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xcc\x31\x0e\xc2\x30\x0c\x85\xe1\xdd\xa7\xb0\x3a\xc1\x40\x2e\xc1\xc4\x82\xb8\x82\x45\x1e\x6d\x85\xe2\x56\x21\x9b\xf5\xee\x8e\x94\x08\xba\xfd\xb2\xf5\xbe\x88\x8c\xd7\xea\xd0\x69\x81\x65\xd4\x89\x94\x88\x6a\x3e\x43\xd3\x75\x2b\x05\xde\x3e\x64\x44\xea\x0f\x78\xd6\x0b\x29\xbb\x3d\xdf\x36\x43\x23\xd2\x63\x24\x29\xb2\x96\x7d\xab\x4d\x4f\x87\x70\xeb\x97\x01\xdc\xad\x80\x1c\x93\xb6\xfc\x38\x52\xce\xff\xfa\x06\x00\x00\xff\xff\xc6\xcc\xc3\x3d\x8e\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/header.tmpl", size: 142, mode: os.FileMode(436), modTime: time.Unix(1531948110, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInlineTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\xcc\xcb\xc9\xcc\x4b\x55\xaa\xad\x55\xa8\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\x4a\x4e\xcc\xc9\x51\x52\xd0\x03\x8b\xa6\xe6\xa5\xd4\xd6\x02\x02\x00\x00\xff\xff\xaa\xeb\x41\xff\x31\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/inline.tmpl", size: 49, mode: os.FileMode(436), modTime: time.Unix(1531431964, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInputsTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8d\x31\x0a\x02\x41\x0c\x45\xaf\x12\x96\x2d\x25\x07\x10\x3c\x80\x9d\xe0\x09\x22\x9b\x59\xa6\xd8\x28\x99\x6c\xf5\xc9\xdd\x65\x46\x8b\xa9\x12\x1e\xff\xbf\x0f\x6c\x5a\xaa\x29\x2d\xd5\x3e\x67\xb4\x25\x13\x58\x0b\x5d\x6f\xc4\xfd\xad\x85\xec\x1d\xc4\xcf\xf3\x15\xda\xa2\x65\x46\xb0\xc9\xa1\x17\x02\xd4\xb6\x7f\x66\x2d\xfc\xf0\x6a\x71\x1f\x92\x0e\x5d\x6c\xd7\xc1\xc5\xe5\xd0\x50\xff\x75\xc5\xf7\xc6\xc0\xa0\x7d\x62\xf2\xcc\xe7\x1b\x00\x00\xff\xff\x8e\xbc\xcf\xda\x98\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/inputs.tmpl", size: 152, mode: os.FileMode(436), modTime: time.Unix(1531946760, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMessageTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8d\xe1\x6a\x83\x40\x10\x84\xff\xfb\x14\x8b\x28\xb4\xa0\xfb\x00\x85\x3e\x40\xff\x14\x69\x4b\xff\x5f\xe2\x68\x16\xf4\x62\xee\x4e\x43\x58\xf6\xdd\x83\x07\xf1\xd7\x0c\x33\xc3\x37\xaa\x3d\x06\xf1\xa0\x72\x46\x8c\x6e\x44\x49\xad\x59\xa1\x2a\x03\xf9\x6b\x22\xfe\x5d\x4f\x09\x31\x45\xb3\xfa\xc6\xa4\x0a\xdf\x9b\xa9\xde\x25\x5d\x88\x7f\x70\x86\x6c\x08\x7b\xc2\x7f\x8f\x05\xfc\xef\xa6\x15\x66\x7c\x0c\xf9\xdb\xcd\x30\x7b\xcb\x44\xee\x82\xf8\xf4\xe5\x97\x75\x07\xaa\x06\xe7\x47\x50\x25\x0d\x55\x98\xe8\xe3\x93\xb8\x73\xc1\xcd\x48\x08\xb9\x97\x81\x2a\x31\x6b\x5e\xbf\xf5\x76\x70\xb3\xbc\x17\xaa\x2d\x65\xfb\x0c\x00\x00\xff\xff\x90\x2e\xb9\x52\xc9\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/message.tmpl", size: 201, mode: os.FileMode(436), modTime: time.Unix(1531946760, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResultsTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8d\x4d\x0a\xc2\x40\x0c\x85\xaf\xf2\x28\xb3\x2c\x3d\x80\xe0\x52\xdc\x7b\x03\xa1\x19\x09\x94\x0c\xbc\x99\xae\x42\xee\x2e\xa9\x45\xc1\x55\x7e\xbe\x2f\x79\xee\xab\x54\x35\xc1\x44\xe9\xfb\x36\xfa\x14\x01\x77\x3e\xed\x25\x28\x3a\xa3\xc8\x86\xcb\x15\xcb\xe3\x83\x23\xdc\xb5\xa2\x68\xc4\x0c\x77\xb1\x35\x37\xf7\x36\xb0\x64\x73\xce\x5a\xf3\x60\xec\xb4\x7e\x23\x1b\x53\x16\xf2\xe4\x38\x84\xc6\xef\xd3\x7f\x39\x03\x7f\xee\x51\xdf\x01\x00\x00\xff\xff\xb0\x4f\xcf\x61\xa8\x00\x00\x00")

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

	info := bindataFileInfo{name: "templates/results.tmpl", size: 168, mode: os.FileMode(436), modTime: time.Unix(1531946760, 0)}
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

