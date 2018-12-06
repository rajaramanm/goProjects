package utilities

import (
	"fmt"
	"strings"
)

// SplitString ... splits a given string from a starting and ending position
func SplitString(stringToSplit string, startingPosition int, endingPosition int) string {
	if (startingPosition > len(stringToSplit)) || (endingPosition > len(stringToSplit)) {
		fmt.Println("Invalid inputs.")
		return stringToSplit
	} else {
		stringArray := make([]string, len(stringToSplit))
		newStringArray := make([]string, 0)
		stringArray = strings.Split(stringToSplit, "")
		for i := startingPosition; i <= endingPosition; i++ {
			newStringArray = append(newStringArray, stringArray[i-1])
		}
		result := strings.Join(newStringArray, "")
		return result
	}
}

// ReverseString ... reverses a given string
func ReverseString(stringToReverse string) string {
	stringArray := make([]string, len(stringToReverse))
	newStringArray := make([]string, 0)
	stringArray = strings.Split(stringToReverse, "")
	for i := 1; i <= len(stringToReverse); i++ {
		newStringArray = append(newStringArray, stringArray[len(stringToReverse)-i])
	}
	result := strings.Join(newStringArray, "")
	return result
}

// IsSTringAPalindrome ... check if a given string is a palindrome
func IsSTringAPalindrome(stringToCheckIfPalindrome string) bool {
	if stringToCheckIfPalindrome == ReverseString(stringToCheckIfPalindrome) {
		fmt.Println("Yup. Its a Palindrome!!")
		return true
	} else {
		fmt.Println("Nope. Its not a palindrome!!")
		return false
	}
}

// ConvertAStringToASlice ... converts a string to a slice/array
func ConvertAStringToASlice(stringToConvertToArray string) []string {
	stringArray := make([]string, len(stringToConvertToArray))
	stringArray = strings.Split(stringToConvertToArray, "")
	return stringArray
}

// ConcatenateStrings ... concatenates the strings passed in the function
func ConcatenateStrings(stringsToAdd ...string) string {
	var newString string
	for _, v := range stringsToAdd {
		newString += v
	}
	return newString
}
