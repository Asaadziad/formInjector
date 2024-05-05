package main

import "fmt"

type CMD struct {
	name        string
	description string
	callback    func([]string) error
}

func NewCallbacks() map[string](func([]string) error) {
	cmdHelp := func([]string) error {
		fmt.Println("help")
		return nil
	}
	return map[string]func([]string) error{
		"help": cmdHelp,
	}
}

func NewCommands() map[string]CMD {
	callbacks := NewCallbacks()
	return map[string]CMD{
		"help": {
			name:        "help",
			description: "usage",
			callback:    callbacks["help"],
		},
	}
}

func Command(name string) CMD {
	commands := NewCommands()
	val, ok := commands[name]
	if ok {
		return val
	}
	return commands["help"]
}
