package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ProcessCase(s string) string {
	upper := "[[:upper:]]+"
    lower := "[[:lower:]]+"
	nonword := "[^[:word]]+"

	for _, c := range s{
		if unicode.IsUpper(c){
			return upper
		} else if unicode.IsLower(c){
			return lower
		} else if !unicode.IsLetter(c) {
			return nonword
		}
	}
	return "none"
}

func ProcessString(s string)string{
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabetSplit := strings.Split(alphabet, "")
	alphabetMap := make(map[string]bool)
	inputString := strings.Split(s, "")

	for _, alphabetChar := range alphabetSplit{
		alphabetMap[alphabetChar] = true
		alphabetMap[strings.ToUpper(alphabetChar)] = true
	}

	var result string

	for _, inputChar := range inputString{
		if _, exists := alphabetMap[inputChar]; exists {
			result += ProcessCase(inputChar)
		} else {
			result += "[^[:word]]+"
		}
	}
	return result
}


func main(){
	var word string
	fmt.Println("Enter word or sentence")
	fmt.Scanln(&word)
	
	sendIt := ProcessString(word)
	fmt.Println("regex:", sendIt)
}