package rps_functions

import (
	"fmt"
	"math/rand"
	"time"
)

// Initiates the seed for the random package
func init() {
	rand.Seed(time.Now().UnixNano())
}

// string in slice checks if a string belongs to a slice
func Stringinslice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// computer_play returns a random key value pair from a map object
func Computerplay(options map[int]string) (int, string) {
	random_key := rand.Intn(len(options)) + 1
	return random_key, options[random_key]
}

// results returns a the winner of a round
func Results(a, b int) int {

	var round_winner int

	if a == b {
		round_winner = 0
	} else if a+b == 5 || a+b == 3 {
		if a > b {
			round_winner = 1
		} else {
			round_winner = 2
		}
	} else if a+b == 4 {
		if a < b {
			round_winner = 1
		} else {
			round_winner = 2
		}
	}

	if round_winner == 1 {
		fmt.Println("You win this round")
	} else if round_winner == 2 {
		fmt.Println("You loose this round")
	} else {
		fmt.Println("Draw")
	}

	return round_winner
}
