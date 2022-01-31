package service

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func UpdateFile(str string) error {
	filename := strings.Split(str, " ")
	data, err := ioutil.ReadFile(filename[2])
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("\nFile Content: %s", data)

	// Read Write Mode
	file, err := os.OpenFile(filename[2], os.O_RDWR, 0644)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer file.Close()

	len, err := file.WriteAt([]byte{'G'}, 0) // Write at 0 beginning
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("\nLength: %d bytes", len)
	fmt.Printf("\nFile Name: %s", file.Name())

	data1, err := ioutil.ReadFile(filename[2])
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Printf("\nFile Content: %s", data1)

	fmt.Println("successfully updated file")
	return nil
}
