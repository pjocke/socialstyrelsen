package main

import (
	"fmt"
	"strings"
	"errors"
)

func isVowel(c string) bool {
	return strings.ContainsAny(strings.ToLower(c), "aeiouyåäö")
}

func getEnding(s string) (error, string) {
	word := []rune(s)

	l := len(word)
	for i := l; i > 0; i-- {
		if isVowel(string(word[i-1])){
			return nil, string(word[i-1:l])
		}
	}
	return errors.New("fel fel fel"), ""
}

func main() {
	//words := []string{"ambulans", "pomerans", "hästsvans", "resistans"}
	err, end := getEnding("resistans")
	if err != nil {
		fmt.Printf("ERR: %s\n", err)
	}

	fmt.Printf("%s\n", end)
}
