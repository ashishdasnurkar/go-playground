package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Println(line)
	line = strings.TrimSpace(line)
	num, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d \n", num)

	var num1 int
	fmt.Scanln(&num1)
	fmt.Printf("Num + 5 is %d \n", num1+5)

}
