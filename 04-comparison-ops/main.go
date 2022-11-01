package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	num := 42
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	if strings.TrimSpace(line) == "hi" {
		fmt.Println("hi back..")
	} else {
		fmt.Println("whatever")
	}

	if num == 42 {
		fmt.Println("num is 42")
	}
}
