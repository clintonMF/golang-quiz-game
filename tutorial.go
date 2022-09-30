package main

import "fmt"

func answer(user_answer, answer string, score, num_of_ques_answered *int) string {
	// this function figures if the answer is correct or wrong and gives appropriate
	// response for each case or input the user makes
	options := [4]string{"a", "b", "c", "d"}
	found := false
	for _, v := range options {
		if user_answer == v {
			found = true
		}
	}

	if found == false && user_answer != "q" {
		// this handles the case in which the user puts an answer that isn't
		// an option, i want to notify the user that they have done so.
		// this will help if the user is unaware of the instructions
		fmt.Println("You can only choose from options a to d in small letters")
	}
	if user_answer == answer {
		fmt.Println("Correct!")
		*score++
		return "continue"
	} else if user_answer == "q" {
		fmt.Println("You exited the game")
		fmt.Println("score: ", *score, ", number of questions:", *num_of_ques_answered)
		return "exit"
	} else {
		fmt.Println("Incorrect")
		return "continue"
	}
}

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

	question1 := answer(answer1, "c", &score, &num_of_ques_answered)

	if question1 == "exit" {
		return
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

	question2 := answer(answer2, "a", &score, &num_of_ques_answered)

	if question2 == "exit" {
		return
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

	question3 := answer(answer3, "d", &score, &num_of_ques_answered)

	if question3 == "exit" {
		return
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

	question4 := answer(answer4, "b", &score, &num_of_ques_answered)

	if question4 == "exit" {
		return
	}

	fmt.Println("score: ", score, "number of questions:", number_of_questions)
	percent := (float32(score) / float32(number_of_questions)) * 100
	fmt.Printf("You answered %v%% of the questions correctly", percent)

}
