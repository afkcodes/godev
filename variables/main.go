package main

import "fmt"

// Global variables

// var name = "hello world" // only in global namespace we can assign a value with a type
// var firstName string = "Ashish"
// var lastName string

var (
	name             = "hello world" // only in global namespace we can assign a value with a type
	firstName string = "Ashish"
	lastName  string
)

// const newVersion = "10.3.1" // constants are immutable

const (
	newVersion = "10.3.1"
)

func main() {
	version := 1 // infer
	otherVersion := "Bar"
	anotherVersion := 10.3
	var name = "Ashish"
	var versionNum int
	version = 20

	fmt.Println(version, otherVersion, anotherVersion, versionNum)
	fmt.Println(name)
	fmt.Println(version)

	fmt.Println(newVersion)
}
