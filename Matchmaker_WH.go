package main

import (
	"fmt"
	"strconv"
)

// Next section will contain all of the necessary variables
// Stance is the stance or position the individual takes on a statement

const Stance1 = 1
const Stance2 = 2
const Stance3 = 3
const Stance4 = 4
const Stance5 = 5

// Multiplier is the multiplier for the statement indicating the level of importance levels 1 - 3
// Higher level = more importance for the statement

const Multi1 int = 1
const Multi2 int = 2
const Multi3 int = 3

//Next contain variables indicating the strength of the compatability score. 3 levels.

const KickStones = 69
const HangOut = 80
const Dateable = 90

// This will contain all of the statements that will be rated by the user

var Question1 string = "Olives are absolutely tasty."
var Question2 string = "I enjoy reading during my free time."
var Question3 string = "The sciences are more fascinating than the arts."
var Question4 string = "The Lord of The Rings is one of the best trilogy of both books and movies."
var Question5 string = "I would rather be hiking and exploring than sitting on a beach doing nothing."

// Next section will contain all of the functions necessary for the program to run
// The first function will validate the result for the compatability score

func valid(x string, Question1 string) int {
	xi, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("You need to type in an integer ranging 1 - 5")
		fmt.Println(Question1)
		fmt.Scanln(&x)
		xi = valid(x, Question1)
	}
	if xi > 5 || xi < 1 {
		fmt.Println("")
		fmt.Println("Again, you need to type in an integer ranging 1 - 5")
		fmt.Println(Question1)
		fmt.Scanln(&x)
		xi = valid(x, Question1)
	}
	return xi
}

// This will contain the main part of the code. This is where all of the witchcraft happens
func main() {
	// Print the Welcome statement and how the program works

	fmt.Println("Welcome to the MatchMaker!!!")
	fmt.Println("For every statement you will respond with an integer 1 - 5")
	fmt.Println("1 = This is Blasphemy")
	fmt.Println("2 = Somewhat Disagree")
	fmt.Println("3 = Mehh, I'm neutral")
	fmt.Println("4 = Somewhat Agree")
	fmt.Println("5 = This is a Law of Nature!")
	fmt.Println("Afterwards, the compatability score will indicate to you whether or not you are a match.")
	fmt.Println("Good Luck and Have Fun!")

	// This will display the first statement, Question1 specifically
	fmt.Println("")
	fmt.Println(Question1)
	var Answer1 string
	var Answer1i int
	var question1_comp int
	fmt.Scanln(&Answer1)
	Answer1i = valid(Answer1, Question1)
	Question1_comp = (prefansw5 - Answer1i)
	if question1_comp < 0 {
		question1_comp = question1_comp * -1
	}
	var Question1Score int
	Question1Score = question1_comp * Multi3

	// This part will display the second statement
	fmt.Println("")
	fmt.Println(Question2)
	var Answer2 string
	var Answer2i int
	var question2_comp int
	fmt.Scanln(&Answer2)
	Answer1i = valid(Answer2, Question2)
	question2_comp = (prefansw5 - Answer1i)
	if question2_comp < 0 {
		question2_comp = question2_comp * -1
	}
	var Question2Score int
	Question2Score = question2_comp * Multi3

	// This part will display the third statement
	fmt.Println("")
	fmt.Println(Question3)
	var Answer3 string
	var Answer3i int
	var question3_comp int
	fmt.Scanln(&Answer3)
	Answer3i = valid(Answer3, Question3)
	Question3_comp = (prefansw5 - Answer3i)
	if question3_comp < 0 {
		question3_comp = question3_comp * -1
	}
	var Question3Score int
	Question3Score = question3_comp * Multi3

	// This part will display the fourth statement
	fmt.Println("")
	fmt.Println(Question4)
	var Answer4 string
	var Answer4i int
	var question4_comp int
	fmt.Scanln(&Answer4)
	Answer4i = valid(Answer4, Question4)
	question4_comp = (prefansw5 - Answer4i)
	if question4_comp < 0 {
		question4_comp = question4_comp * -1
	}
	var Question4Score int
	Question4Score = question4_comp * Multi2

	//This part will display the fifth statement
	fmt.Println("")
	fmt.Println(Question5)
	var Answer5 string
	var Answer5i int
	var question5_comp int
	fmt.Scanln(&Answer5)
	Answer5i = valid(Answer5, Question5)
	question5_comp = (prefansw5 - Answer5i)
	if question5_comp < 0 {
		question5_comp = question5_comp * -1
	}
	var Question5Score int
	Question5Score = question5_comp * Multi1

	var totalComp int = question1_comp + question2_comp + question3_comp + question4_comp + question5_comp
	var final_comp int = 100 - Question1Score - Question2Score - Question3Score - Question4Score - Question5Score

	if final_comp >= Dateable {
		fmt.Println("")
		fmt.Println("It's a match, we should go on a date.")
	} else if final_comp >= HangOut && final_comp < Dateable {
		fmt.Println("")
		fmt.Println("Let's just hang out by the bar and put out the vibe.")
	} else if final_comp < HangOut || final_comp <= KickStones {
		fmt.Println("")
		fmt.Println("This is not gonna work out. Go kick stones.")
	}
}
