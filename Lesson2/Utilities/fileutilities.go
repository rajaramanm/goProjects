package utilities

import (
	"io/ioutil"
	"os"
)

type file struct {
	header *string
	body   *[]string
	footer *string
}

// OpenFile ... opens a specified file in a location and read the contents of the file as byte
func OpenFile(fileName string, location string) *os.File {
	file, err := os.OpenFile(ConcatenateStrings(location, fileName), os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	return file
}

// WriteToFile ... Write contents to file in a new line
func WriteToFile(fileName string, location string, contentsToWrite string) {
	_, err := OpenFile(fileName, location).WriteString(contentsToWrite)
	check(err)
}

// ReadContentsFromFile ... Reads the contents of a file in bytes
func ReadContentsFromFile(fileName string, location string) []byte {
	file, err := ioutil.ReadFile(ConcatenateStrings(location, fileName))
	check(err)
	return file
}

// LockFile ... Lock a file and prevent writing to that file
func LockFile(fileName string) {

}

// DeleteFile ... Deletes a file in a specified location
func DeleteFile(fileName string, location string) {
	os.Remove(ConcatenateStrings(location, fileName))
}

// CopyFile ... Copies a file from a source to a destination
func CopyFile(fileName string, sourceLocation string, newFileName string, destinationLocation string) {
	CreateANewFile(newFileName, destinationLocation).Write(ReadContentsFromFile(fileName, sourceLocation))
}

// RenameFile ... Renames a file in a location
func RenameFile(fileName string, newName string, location string) {
	os.Rename(ConcatenateStrings(location, fileName), newName)
}

// CreateANewFile ... creates a new file of specified type/extension in the specified location with a header, body and a footer
func CreateANewFile(fileName string, location string) *os.File {
	file, err := os.Create(ConcatenateStrings(location, fileName))
	check(err)
	return file
}

// check ... checks if error is nil
func check(err error) {
	if err != nil {
		panic(err)
	}
}
