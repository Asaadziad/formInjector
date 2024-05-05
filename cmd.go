package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

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
	cmdGenerate := func(args []string) error {
		jsonF, err := os.Open("./data/forms.json")
		if err != nil {
			return err
		}
		defer jsonF.Close()

		bytes, _ := io.ReadAll(jsonF)
		var forms []Form
		json.Unmarshal(bytes, &forms)

		for _, e := range forms {
			generated := "<div>\n"
			generated += fmt.Sprintf("<h1>%v</h1>\n", e.Title)
			for _, f := range e.Fields {
				generated += fmt.Sprintf("<label for='%v'>%v</label>\n", f.Name, f.Label)
				generated += fmt.Sprintf("<input type='%v' id='%v' name='%v'></input>\n", f.FType, f.Name, f.Name)
			}
			generated += "</div>\n"
			nFName := fmt.Sprintf("%v.html", e.Title)
			err := os.WriteFile(nFName, []byte(generated), 0644)
			if err != nil {
				return nil
			}
		}
		return nil
	}
	return map[string]func([]string) error{
		"help":     cmdHelp,
		"generate": cmdGenerate,
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
		"generate": {
			name:        "generate",
			description: "generates new form",
			callback:    callbacks["generate"],
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
