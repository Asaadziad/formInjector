package main

import "fmt"

func Eval(command []string) {
	if len(command) < 1 {
		fmt.Println("Error")
		return
	}
	cmd := Command(command[0])
	err := cmd.callback(nil)
	if err != nil {
		fmt.Println("Error occurred")
	}
}
