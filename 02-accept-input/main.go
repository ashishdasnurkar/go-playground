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
	fmt.Printf("%d", num)
}
