package main

import (
	"fmt"
	"strings"
)

func Parse(text string) ([]string, error) {
	if len(text) == 0 {
		return nil, fmt.Errorf("error: didn't specify a command")
	}
	command := strings.Split(text, " ")
	return command, nil
}
