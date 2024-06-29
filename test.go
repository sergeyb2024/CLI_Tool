package main

import (
	"fmt"
	"unicode"
)

func IsLower(s string) string {
	var upper = "/[[:upper:]]+/g" 
	var lower = "/[[:lower:]]+/g"
    for _, r := range s {
        if !unicode.IsLower(r) && unicode.IsLetter(r) {
            return upper
        }
    }
    return lower
}

func putIntoArray(s string) string {
	var arr []string
	var characters string
	for _, c := range s {
		arr = append(arr, string(c))
	}
	for _, arrayItem := range arr {
		characters = arrayItem
	}
	fmt.Println(characters)
	return IsLower(characters)
}


func main() {
	var word string
	for _, c := range word  {
		IsLower(string(c))
	}
	
	fmt.Scanln(&word)
	putIntoArray(word)
}