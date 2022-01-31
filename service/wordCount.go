package service

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Change directory path replace "/home/rizwana" to ypur path
const dirCount = "/home/rizwana/file-store-service/service/"

func WordCount(str string) error {
	filename := strings.Split(str, " ")
	fh, err := os.Open(dirCount + filename[2])
	if err != nil {
		fmt.Printf("Could not open file '%v': %v", filename[2], err)
		os.Exit(1)
	}
	reader := bufio.NewReader(fh)
	counter := make(map[string]int)
	for {
		line, _ := reader.ReadString('\n')
		fields := strings.Fields(line)
		for _, word := range fields {
			word = strings.ToLower(word)
			counter[word]++
		}
		if line == "" {
			break
		}
	}

	for word, cnt := range counter {
		fmt.Printf("%v %v\n", word, cnt)
	}

	words := make([]string, 0, len(counter))
	for word := range counter {
		words = append(words, word)
	}
	sort.Slice(words, func(i, j int) bool {
		return counter[words[i]] > counter[words[j]]
	})

	for _, word := range words {
		fmt.Printf("%v %v\n", word, counter[word])
	}
	return nil
}
