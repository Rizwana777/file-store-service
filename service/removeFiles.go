package service

import (
	"fmt"
	"os"
	"strings"
)

func RemoveFile(str string) error {
	filename := strings.Split(str, " ")
	err := os.Remove(filename[2])

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("File successfully deleted")
	return nil
}
