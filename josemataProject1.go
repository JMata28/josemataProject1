/*
This is Project #1 by José Amilcar Mata Calidonio for the COMP 510 course.
*/
package main

import (
	_ "embed"
	"fmt"
	"math/rand"
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
	channel1 := make(chan int)     //create channel named "channel1" that takes integers. This will be used to send a value that signals the fortune function to print out the fortune.
	channel2 := make(chan int)     //create a channel named "channel2" that takes integers. This will be used to send a value that signals the UI that it can continue after the fortune function prints the fortune.
	go fortune(channel1, channel2) //run the fortune function as a go routine and pass the channels as the parameters. Since it is a go routine, it will run "in the background" as its own lightweight thread.
	var answer string              //Create variable where answer of the user will be stored. It is initialized with the zero value (an empty string) here.
	for {                          //infinite for loop
		fmt.Println("Would you like another fortune? Please type yes or no.") //Ask for user input
		fmt.Scanln(&answer)                                                   //Save user input in the variable "answer"
		answerLowercase := strings.ToLower(answer)                            //Make the string saved in variable "answer" all lowercase (this is used to ignore any capitalization in the user input in the if statements below)
		if answerLowercase == "yes" {
			channel1 <- 1          //Sends a value of 1 into the channel named "channel1"
			message2 := <-channel2 //Wait until channel2 is loaded with data (which will only happen after the fortune is printed) and once it is, store it in the variable "message2"
			if message2 == 1 {     //This is just an empty if statement so the variable "message2" is used.
			}
		} else if answerLowercase == "no" {
			os.Exit(0) //This function terminates the program. The parameter 0 signals an exit with no error.
		}
	}
}

/*
The "fortune" function below receives both channels as parameters. The arrows show that channel1
will be read from only in this function and channel2 will only have data sent to it in this function.
*/
func fortune(channel1 <-chan int, channel2 chan<- int) {
	contentInASlice := strings.Split(contentInAString, "%%") //Creates a slice of strings separated by the "%%" symbols. This separates each fortune.
	for {                                                    //infinite loop
		numberOfFortunes := len(contentInASlice) //Determines how many strings are in this slice, that is, the number of fortunes in the "fortunes.txt" file
		/*
			The line below generates a random number from 0 up to (but NOT including) the number of fortunes in the
			"fortunes.txt" file. This is important because slices are zero indexed, so we don't want to exceed the number
			of items in the slice. According to https://pkg.go.dev/math/rand#Rand.Intn, "Intn returns, as an int, a non-negative pseudo-random number in the
			half-open interval [0,n). It panics if n <= 0."
		*/
		randomnumber := rand.Intn(numberOfFortunes)
		message := <-channel1 //Wait until the channel is loaded with data and once it is, store it in the variable "message"
		/*
			The value will only be 1 when the user has entered "yes", and since the value from a channel can only be
			read once, the user will have to enter "yes" every time they want a new fortune.
		*/
		if message == 1 {
			fmt.Println(contentInASlice[randomnumber]) //Print the fortune with the random number as its index. This picks a random fortune every time.
			channel2 <- 1                              //Send to channel2 the value of 1. This only happens after the fortune is printed out.
		}
	}
}
