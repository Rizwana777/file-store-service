package service

import "testing"

func Test_Removefiles(t *testing.T) {
	t.Run("Test something", func(t *testing.T) {
		if err := RemoveFile("store rm file1.txt"); (err != nil) != false {
			t.Errorf("RemoveFile() error = %v", err)
		}
	})
}
