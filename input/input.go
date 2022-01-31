package input

import (
	"file-store-service/service"
	"os"
	"strings"
)

func Command(str string) {
	if strings.HasPrefix(str, "store add") {
		service.CreateFile(str)
	}
	if strings.HasPrefix(str, "store rm") {
		service.RemoveFile(str)
	}
	if strings.HasPrefix(str, "store ls") {
		service.ListFiles()
	}
	if strings.HasPrefix(str, "store update") {
		service.UpdateFile(str)
	}
	if strings.HasPrefix(str, "store wc") {
		service.WordCount(str)
	}
	if strings.HasPrefix(str, "exit") {
		os.Exit(1)
	}
}
