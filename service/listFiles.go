package service

import (
	"fmt"
	"io/ioutil"
)

const dir = "/home/rizwana/file-store-service"

func ListFiles() error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
	return nil
}
