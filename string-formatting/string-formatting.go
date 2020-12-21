package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	// Prints an instance of the struct
	fmt.Printf("%v\n", p)
	// Include the struct's field names
	fmt.Printf("%+v\n", p)
	// Print a Go syntax representation
	// i.e. the source code snippet that would produce the value
	fmt.Printf("%#v\n", p)
	// Print the type
	fmt.Printf("%T\n", p)
	// Format boolean
	fmt.Printf("%t\n", true)
	// Format integers
	// %d is standard, base-10 formatting
	fmt.Printf("%d\n", 123)
	// Binary representation
	fmt.Printf("%b\n", 14)
	// Character corresponding to given integer
	fmt.Printf("%c\n", 33)
	// Hex encoding
	fmt.Printf("%x\n", 456)
	// Format a float
	// Basic formatting
	fmt.Printf("%f\n", 78.9)
	// Format a float in scientific notation
	fmt.Printf("%e\n", 1234000000.0)
	fmt.Printf("%E\n", 1234000000.0)
	// Basic string
	fmt.Printf("%s\n", "\"string\"")
	// Double-quote strings
	fmt.Printf("%q\n", "\"string\"")
	// %x is hex encoding, otherwise known as base-16
	fmt.Printf("%x\n", "hex this")
	// Representation of a pointer
	fmt.Printf("%p\n", &p)
	// Control width and precision of a number
	// By default, result is right-justified and padded with spaces
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	// Control width when formatting strings
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	// Left-align
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	// Format string without printing it
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	// Use io.Writers to send formatted string somewhere other than stdout
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
