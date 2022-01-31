package service

import "testing"

func Test_ListFiles(t *testing.T) {
	t.Run("Test something", func(t *testing.T) {
		if err := CreateFile("store add file1.txt"); (err != nil) != false {
			t.Errorf("CreateFile() error = %v", err)
		}
	})
}
