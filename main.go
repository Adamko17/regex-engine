package main

import (
	"fmt"

	"./rgx"
)

func main() {
	ctx := rgx.Parse(`[a-zA-Z][a-zA-Z0-9_.]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,}`)
	nfa := rgx.ToNfa(ctx)
	
	email := "test@example.com"
	result := nfa.Check(email, -1)
	fmt.Printf("Email '%s' is valid: %t\n", email, result)
}