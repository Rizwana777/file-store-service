package service

import "testing"

func Test_WordCount(t *testing.T) {
	t.Run("Test something", func(t *testing.T) {
		if err := WordCount("store wc file1.txt"); (err != nil) != false {
			t.Errorf("WordCount() error = %v", err)
		}
	})
}
