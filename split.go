package main

import (
	"fmt"
	"strings"
	"errors"
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
	words := []string{"ambulans", "pomerans", "hästsvans", "resistans", "kanelbulle", "höskulle", "bildrulle"}

	for _ ,w := range(words) {
		err, end := getEnding(w)
		if err != nil {
			fmt.Printf("ERR: %s\n", err)
		}

		fmt.Printf("%s\t\t%s\n", w, end)
	}
}
