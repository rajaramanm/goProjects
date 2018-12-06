package main

import (
	"fmt"

	Utils "./utilities"
)

const (
	fileName       = "sample.txt"
	location       = "/Users/itrm1/GoProjects/Lesson2/"
	requiredString = "Big Boss"
)

var mainOption, subOption int
var continueOrQuit bool

func main() {

	fmt.Println("Utilities - using Golang")
	fmt.Println("************************")
MAINOPTION:
	fmt.Println("Enter a choice from the below options")
	fmt.Println("1. String Utilities")
	fmt.Println("2. File Utilities")
	fmt.Println("3. QUIT")
	fmt.Scan(&mainOption)

	switch mainOption {
	case 1:
	SUBOPTION1:
		fmt.Println("1. Split the String")
		fmt.Println("2. Reverse the String")
		fmt.Println("3. Check if a String is a Palindrome")
		fmt.Println("4. Convert a String to an Array")
		fmt.Println("5. Concatenate Strings")
		fmt.Scan(&subOption)

		switch subOption {
		case 1:
			fmt.Println(Utils.SplitString(requiredString, 3, 7))
			break

		case 2:
			fmt.Println(Utils.ReverseString(requiredString))
			break

		case 3:
			Utils.IsSTringAPalindrome(requiredString)
			break

		case 4:
			fmt.Println(Utils.ConvertAStringToASlice(requiredString))
			break

		case 5:
		default:
			fmt.Println(Utils.ConcatenateStrings(fileName, location, requiredString))
			break
		}
		fmt.Println("Do you want Continue or go back to Main Menu?")
		fmt.Scan(&continueOrQuit)
		if continueOrQuit == true {
			goto SUBOPTION1
		} else {
			goto MAINOPTION
		}
	case 2:
	SUBOPTION2:
		fmt.Println("1. Create a New File")
		fmt.Println("2. Write to a File")
		fmt.Println("3. Read From a File")
		fmt.Println("4. Copy a File to new location")
		fmt.Println("5. Rename a File")
		fmt.Println("6. Delete a File")
		fmt.Scan(&subOption)
		switch subOption {
		case 1:
			Utils.CreateANewFile(fileName, location)
			break

		case 2:
			Utils.WriteToFile(fileName, location, "Rajaraman")
			break

		case 3:
			Utils.ReadContentsFromFile(fileName, location)
			break

		case 4:
			Utils.CopyFile(fileName, location, "sample2.txt", "/Users/itrm1/GoProjects/Lesson2/")
			break

		case 5:
			Utils.RenameFile(fileName, "trial.txt", location)
			break

		case 6:
			Utils.DeleteFile("trail.txt", "/Users/itrm1/GoProjects/Lesson2/")
			break
		}
		fmt.Println("Do you want Continue or go back to Main Menu?")
		fmt.Scan(&continueOrQuit)
		if continueOrQuit == true {
			goto SUBOPTION2
		} else {
			goto MAINOPTION
		}
	case 3:
		break
	}
}
