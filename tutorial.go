package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game")

	fmt.Printf("Enter your name: ")
	var name string

	// handling errors
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("An error occured, please enter the right input")
		return
	}
	fmt.Printf("Hello, %v welcome to the game\n\n", name)

	instructions := `choose the correct answer from the options a to d.
Your answer should be in lower case.
If you want to quit the quiz enter "q" also lowercase`

	fmt.Println(instructions)

	score := 0
	number_of_questions := 4
	num_of_ques_answered := 0

	fmt.Println("\nWho is the first nigerian female to drive a car?")
	fmt.Printf("(a)Ngozi Okonjo-Iweala (b)Patience Jonathan (c) Funmilayo Ransome-Kuti (d) i don't care\n")
	var answer1 string
	_, err1 := fmt.Scan(&answer1)

	if err1 != nil {
		fmt.Println("An error occured, please enter the right input")
		return
	}

	if answer1 == "c" {
		fmt.Println("Correct!")
		score++
	} else if answer1 == "q" {
		fmt.Println("You exited the game")
		fmt.Println("score: ", score, ", number of questions:", num_of_ques_answered)
		return
	} else {
		fmt.Println("Incorrect")
	}

	num_of_ques_answered++ // this is to increase the variable to show that a question
	// was answered

	fmt.Println("\nQuestion 2 : who is the current president of nigeria")
	fmt.Printf("(a)M. buhari (b)Peter Obi (c) Goodluck Jonathan (d)You\n")

	var answer2 string

	// handling error with inputing the answer
	_, err2 := fmt.Scan(&answer2)
	if err2 != nil {
		fmt.Println("An error occured, please enter the right input")
		return
	}
	// handling error with inputing the answer

	if answer2 == "a" {
		fmt.Println("Correct!")
		score++
	} else if answer2 == "q" {
		fmt.Println("You exited the game")
		fmt.Println("score: ", score, ", number of questions:", number_of_questions)
		return
	} else {
		fmt.Println("Incorrect")
	}
	num_of_ques_answered++ // this is to increase the variable to show that a question
	// was answered

	fmt.Println("\nQuestion3 : when did Nigeria get her independence")
	fmt.Println("(a)1000 (b)2015 (c) 100BC (d) 1960")
	var answer3 string

	// handling error with inputing the answer
	_, err3 := fmt.Scan(&answer3)

	if err3 != nil {
		fmt.Println("An error occured, please enter the right input")
		return
	}

	if answer3 == "d" {
		fmt.Println("Correct!")
		score++
	} else if answer3 == "q" {
		fmt.Println("You exited the game")
		fmt.Println("score: ", score, ", number of questions:", num_of_ques_answered)
		return
	} else {
		fmt.Println("Incorrect")
	}
	num_of_ques_answered++ // this is to increase the variable to show that a question
	// was answered

	fmt.Println("\nQuestion1 : who colonized Nigeria")
	fmt.Println("(a)Ghana (b)The british (c) The french (d) Asians")

	var answer4 string
	_, err4 := fmt.Scan(&answer4)
	if err4 != nil {
		fmt.Println("An error occured, please enter the right input")
		return
	}

	if answer4 == "b" {
		fmt.Println("Correct!")
		score++
	} else if answer4 == "q" {
		fmt.Println("You exited the game")
		fmt.Println("score: ", score, ", number of questions:", number_of_questions)
		return
	} else {
		fmt.Println("Incorrect")
	}

	fmt.Println("score: ", score, "number of questions:", number_of_questions)
	percent := (float32(score) / float32(number_of_questions)) * 100
	fmt.Printf("You answered %v%% of the questions correctly", percent)

}
