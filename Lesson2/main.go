package main

import (
	"fmt"

	Utils "./utilities"
)

func main() {
	requiredString := "Big Boss"
	fmt.Println(Utils.SplitString(requiredString, 3, 7))
	fmt.Println(Utils.ReverseString(requiredString))
	Utils.IsSTringAPalindrome(requiredString)
	fmt.Println(Utils.ConvertAStringToASlice(requiredString))
}
