package main

import (
	"fmt"
	"strings"
	"errors"
	"os"
	"bufio"
)

func getEnding(s string) (error, string) {
	word := []rune(s)

	l := len(word)

	levels := 1
	if strings.ContainsAny(strings.ToLower(string(word[l-1])),  "aeiouyåäö") {
		levels = 2
	}

	level := 0
	for i := l; i > 0; i-- {
		if strings.ContainsAny(strings.ToLower(string(word[i-1])),  "aeiouyåäö") {
			level++
		}
		if level == levels {
			return nil, string(word[i-1:l])
		}
	}

	return errors.New("fel fel fel"), ""
}

func main() {
	f, err := os.Open("./substantiv.txt")
	if err != nil {
		fmt.Printf("Unable to open file.")
		return
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		w := scanner.Text()
		err, end := getEnding(w)
		if err != nil {
			fmt.Printf("%s,%s\n", w, err)
		}

		fmt.Printf("%s,%s\n", w, end)
	}
}
