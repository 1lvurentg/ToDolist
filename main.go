package main

import (
	"github.com/1lvurentg/ToDo-list/scanner"
	"github.com/1lvurentg/ToDo-list/todo"
)

func main() {
	list := todo.NewList()

	scanner := scanner.NewScanner(list)

	scanner.Start()
}
