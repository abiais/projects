package main

import "fmt"

func fibo(n int) int {
	if n < 0 {
		print("")
	}

	if n == 0 || n == 1 {
		return n
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func main() {

	fmt.Println("Please pick an integer that is positive or nul: ")

	var i int

	fmt.Scanln(&i)

	fmt.Print(fibo(i))

}
