package main

//typeof pck reflect

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

// func ProcessCase(a string, b string ) string {
func ProcessCase(a string) string {
    upper := "[[:upper:]]+"
    lower := "[[:lower:]]+"
    for _, r := range a {
        if !unicode.IsLower(r) && unicode.IsLetter(r) {
            return upper
        }
    }
    return lower
}

func ProcessString(s string) string {
	
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabetSplit := strings.Split(alphabet, "")

	inputLetters := strings.Split(s, "")
	fmt.Println("2", alphabetSplit)
	
	
    var result string
    for _, userInputSplit := range s {
		fmt.Println("user", reflect.TypeOf(inputLetters))
		// result nonWordCharacter
		for _, letterOfAlphabet := range alphabetSplit{
			// nonWordCharacter := 0
			fmt.Println("letters", reflect.TypeOf(letterOfAlphabet))
			
			if letterOfAlphabet == alphabetSplit[1] {
				// nonWordCharacter = 0
				result += ProcessCase(string(userInputSplit))
				break
			}
		}
    }
	return ""
}

func main() {
    var word string
    fmt.Println("enter word or sentence")
    fmt.Scanln(&word)
    sendIt := ProcessString(word)
    fmt.Println("Result:", sendIt)
}