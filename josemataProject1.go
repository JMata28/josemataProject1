/*
This is Project #1 by José Amilcar Mata Calidonio for the COMP 510 course.
*/
package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

/*
The following two lines embed the contents of the "fortunes.txt" file into the running Go program
and store it in the string named "contentInAString"
*/
//go:embed fortunes.txt
var contentInAString string

func main() {
	var answer string //Create variable where answer of the user will be stored. It is initialized with a zero value here.
	for {
		fmt.Println("Would you like another fortune? Type yes or no") //Ask for user input
		fmt.Scanln(&answer)                                           //Save user input in the variable "answer"
		answerLowercase := strings.ToLower(answer)                    //Make "answer" all lowercase (this will ignore any capitalization in the user input)
		if answerLowercase == "yes" {
			fmt.Println("You answered yes.")
			fortune()
		} else if answerLowercase == "no" {
			fmt.Println("You answered no. Program terminated.")
			os.Exit(0) //This function terminates the program. The parameter 0 signals an exit with no error.
		}
	}
}

func fortune() {
	contentInASlice := strings.Split(contentInAString, "%%") //Creates a slice of strings separated by the "%%" symbols
	fmt.Print(contentInASlice[0])
}
