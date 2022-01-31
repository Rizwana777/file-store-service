package input

import (
	"file-store-service/service"
	"strings"
)

func Command(str string) {
	if strings.HasPrefix(str, "store add") {
		service.CreateFile(str)
	}
	if strings.HasPrefix(str, "store rm") {
		service.RemoveFile(str)
	}
}
