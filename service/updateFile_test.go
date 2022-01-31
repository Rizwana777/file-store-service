package service

import "testing"

func Test_UpdateFile(t *testing.T) {
	t.Run("Test something", func(t *testing.T) {
		if err := UpdateFile("store update file1.txt"); (err != nil) != false {
			t.Errorf("UpdateFile() error = %v", err)
		}
	})
}
