package main

import (
	"bufio"
	"file-store-service/input"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		text := in.Text()
		input.Command(text)
	}
}
