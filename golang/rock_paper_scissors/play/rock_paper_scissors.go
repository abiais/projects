package main

import (
	"fmt"

	"rps_functions"
)

func main() {

	score_player := 0
	score_computer := 0

	options := map[int]string{
		1: "Rock",
		2: "Paper",
		3: "Scissors",
	}

	options_arr := []string{options[1], options[2], options[3]}

	options_reversed := map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}

	for score_player < 3 && score_computer < 3 {
		// Input by the user
		var player_input string
		fmt.Println("Rock, Paper, Scissors ?:")
		fmt.Scan(&player_input)

		for !rps_functions.Stringinslice(player_input, options_arr) {
			fmt.Println("Please pick from Rock, Paper, Scissors :")
			fmt.Scan(&player_input)
		}

		fmt.Println("You played :", player_input)

		player_choice := options_reversed[player_input]

		// Computer play
		computer_choice, computer_input := rps_functions.Computerplay(options)
		fmt.Println("The computer played :", computer_input)

		// Results
		round_winner := rps_functions.Results(player_choice, computer_choice)

		if round_winner == 1 {
			score_player += 1
		} else if round_winner == 2 {
			score_computer += 1
		}

		fmt.Println("Your score: ", score_player, "vs. Computer score ", score_computer)

	}

	if score_player == 3 {
		fmt.Println("You win !!!")
	} else {
		fmt.Println("The computer wins :)")
	}

}
