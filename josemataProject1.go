/*
This is Project #1 by José Amilcar Mata Calidonio for the COMP 510 course.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var answer string //Create variable where answer of the user will be stored. It is initialized with a zero value here.
	for {
		fmt.Println("Would you like another fortune? Type yes or no") //Ask for user input
		fmt.Scanln(&answer)                                           //Save user input in the variable "answer"
		answerLowercase := strings.ToLower(answer)                    //Make "answer" all lowercase (this will ignore any capitalization in the user input)
		if answerLowercase == "yes" {
			fmt.Println("You answered yes")
			fortune()
		} else if answerLowercase == "no" {
			fmt.Println("You answered no. Program terminated.")
			os.Exit(0) //This function terminates the program. The parameter 0 signals an exit with no error.
		}
	}
}

func fortune() {
	/*
		The following three lines of code open up the fortunes.txt file.
		The if statement in the second line below checks for errors when opening the .txt file.
	*/
	content, err := os.ReadFile("fortunes.txt") //This line of code was taken from the code used in our in-class Go program created during the lecture on January 29, 2024 and modified to be used in this program.
	if err != nil {
		log.Fatal("Help file error", err)
	} //This is statement was taken from the code used in our in-class Go program created during the lecture on January 29, 2024.
	contentInAString := string(content)                      //The contents of the file into a string called contentstring, so it is no longer handled as byte data.
	contentInASlice := strings.Split(contentInAString, "%%") //Creates a slice of strings separated by the "%%" symbols
	fmt.Print(contentInASlice[0])
}
