package schema

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _schema_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xb1\x0a\x02\x31\x0c\x86\xf7\x3c\x45\xdc\xf4\x15\xba\x29\x38\x3a\xe8\x0d\xce\xd1\x86\xbb\x42\xdb\x9c\x4d\x0f\xee\x90\xbe\xbb\x58\xed\x81\xe0\x94\x10\xbe\xff\xff\xa2\xf7\x81\x03\xe1\x13\x10\x1f\x13\xa7\xc5\xe0\xf9\x3d\x00\x31\x4c\x99\xb2\x93\x68\xf0\xf4\xdd\xa0\x00\xe4\x65\x64\xf4\x4c\xca\xfb\x8e\xa2\x3d\xc8\x7c\x61\x1d\x25\x2a\xd7\x8e\xc0\xaa\xd4\xb3\xc1\x2e\x27\x17\x7b\x40\x54\x8a\xf6\x26\xf3\x7a\x69\x1d\x55\x53\x33\x03\x7b\x2f\x57\x49\xde\x36\x68\x03\xe5\x03\x35\x73\xe5\x7e\xac\xdb\x63\x20\xe7\xd7\xc0\xce\xfc\x7f\x0a\x0a\xbc\x02\x00\x00\xff\xff\x98\xe2\x83\xfc\xe2\x00\x00\x00")

func schema_graphql() ([]byte, error) {
	return bindata_read(
		_schema_graphql,
		"schema.graphql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"schema.graphql": schema_graphql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"schema.graphql": &_bintree_t{schema_graphql, map[string]*_bintree_t{
	}},
}}
