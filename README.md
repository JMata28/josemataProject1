# josemataProject1

This is Jose Amilcar Mata Calidonio's Project #1 for the COMP 510 course.

The instructions for this assignment are shown below:

Project 1 : fortune fun


Summary: lets practice with files and go routines, making a multithreaded fortune program



Due: Tuesday Feb 20th at 11:59pm (the day before our Monday on Wednesday class)

This one will go into your project grade this time.

Details:

Create a new project. Make sure it has your name in it somewhere in the project name.


(granted I am pushing the limits to try and get a simple program which also practices with channels and goroutines)


Write go go program that has two functions



The main function will

1. create a channel and then spawn the fortune function as a go routine taking the channel as a parameter.

2. Enter a forever loop and ask the user if they want another fortune. If the user says yes (with any capitalization) then send a message down the channel to the for the fortune function to select a fortune and display it.

3. If the user answers no (with any capitalization), then end the program

4. if the user answers anything else, then ask again


The fortune function will

1. Open the fortunes.txt file  and read it in.
    1. be sure to use go:embed to embed the file into the final executable
2. split the contents of the fortunes file on the %% string to get a slice of strings

3. loop forever

    1. wait for a message on the channel

    2. when a message is received, randomly select one of the fortunes from your slice

    3. and print it to the screen.


Using goroutines and channels will feature heavily in this project’s grading


Submitting the project.


submit the project by putting it on github and sending me a collaboration invite. As before.

the min/max opportunity of the month