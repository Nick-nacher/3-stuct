package file

import (
	"os"
	"path/filepath"
)

func ReadFile(name string) (*[]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func IsJson(path string) bool {
	const ext string = ".json"
	return filepath.Ext(path) == ext
}
