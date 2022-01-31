package service

import (
	"fmt"
	"os"
	"strings"
)

func CreateFile(str string) error {
	filename := strings.Split(str, " ")
	if _, err := os.Stat(filename[2]); err == nil {
		fmt.Printf("File already exists\n")
		return err
	}

	file, err := os.Create(filename[2])
	if err != nil {
		fmt.Println("os.Create:", err)
		return err
	}

	defer file.Close()

	fmt.Fprintf(file, "[%s]: ", filename)
	fmt.Fprintf(file, "Hello %s\n", filename)

	fmt.Println("Files added successfully")
	return nil
}
