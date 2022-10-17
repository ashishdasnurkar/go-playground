package main

import "fmt"

func main() {

	fmt.Println("Hello world!")
	fmt.Println("Hello", "world", "!") // adds spaces and a new line
	fmt.Print("Hello", "world", "!")   // no spaces and no new line
	fmt.Println("")
	fmt.Println(`multiline
string
is possible with backticks 
in go`)
}
