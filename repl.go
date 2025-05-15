package main

import(
	"fmt"
	"strings"
	)

func 	cleanInput(text string) []string {
	output := strings.TrimSpace(text)
	words := strings.Fields(output)
	return words

}