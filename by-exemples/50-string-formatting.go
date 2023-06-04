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
	// this prints an instance of our point struct.
	fmt.Printf("struct1: %v\n", p)

	// If the value is a struct, the %+v variant will include the struct’s field names.
	fmt.Printf("struct2: %+v\n", p)

	// The %#v variant prints a Go syntax representation of the value, i.e. the source code snippet that would produce that value.
	fmt.Printf("struct3: %#v\n", p)

	// To print the type of a value, use %T.
	fmt.Printf("type: %T\n", p)

	// Formatting booleans is straight-forward.
	fmt.Printf("bool: %t\n", true)

	// Use %d for standard, base-10 formatting.
	fmt.Printf("int: %d\n", 123)

	// This prints a binary representation.
	fmt.Printf("bin: %b\n", 123)

	// This prints the character corresponding to the given integer.
	fmt.Printf("char: %c\n", 33)

	// %x provides hex encoding.
	fmt.Printf("hex: %x\n", 33)

	// For basic decimal formatting use %f.
	fmt.Printf("float1: %f\n", 78.9)

	// %e and %E format the float in (slightly different versions of) scientific notation.
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// For basic string printing use %s.
	fmt.Printf("str1: %s\n", "\"string\"")

	// To double-quote strings as in Go source, use %q.
	fmt.Printf("str2: %q\n", "\"string\"")

	// %x renders the string in base-16, with two output characters per byte of input.
	fmt.Printf("str3: %x\n", "hex this")

	// To print a representation of a pointer, use %p.
	fmt.Printf("pointer: %p\n", &p)

	// To specify the width of an integer, use a number after the % in the verb. By default the result will be right-justified and padded with spaces.
	fmt.Printf("width1: |%6d|%6d|%6d|\n", 12, 5213123345, 833)

	// You can also specify the width of printed floats, though usually you’ll also want to restrict the decimal precision at the same time with the width.precision syntax.
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// To left-justify, use the - flag.
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// For basic right-justified width.
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// To left-justify use the - flag as with numbers.
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
