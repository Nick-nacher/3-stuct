package storage

import (
	"3-struct/bins"
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(name string) (*[]bins.Bin, error) {

	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var bin []bins.Bin
	err = json.Unmarshal(data, &bin)
	if err != nil {
		return nil, err
	}
	return &bin, nil

}

func WriteFile(content []byte, name string) error {

	err := os.WriteFile(name, content, 0644)
	if err != nil {
		return err
	}
	return nil

}
