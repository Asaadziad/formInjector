package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
	fmt.Print("cmd > ")
	for scanner.Scan() {
		command, err := Parse(scanner.Text())
		if err != nil {
			fmt.Println(err.Error())
			fmt.Print("cmd > ")
			continue
		}
		Eval(command)
		fmt.Print("cmd > ")
	}
}
