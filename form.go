package main

type Field struct {
	Name  string
	FType string
	Label string
}

type Form struct {
	Id     int
	Title  string
	Fields []Field
}
