package main

import (
	"fmt"
	"is_palindrome_func"
	"log"
)

func main() {

	log.SetPrefix("is_palindrom error: ")
	log.SetFlags(0)

	fmt.Println("Please enter a string to check if it is a palindrome : ")

	var text string

	fmt.Scanln(&text)

	result, err := is_palindrome_func.Is_palindrome(text)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(result)

}
