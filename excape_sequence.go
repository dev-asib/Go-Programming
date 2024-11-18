package main

import "fmt"

func main() {
	fmt.Println("Go Developer & Flutter Developer.")

	// "\n": NewLine (Line Break)
	fmt.Println("Go Developer\n &\nFlutter Developer.")
	fmt.Println("Go Developer \n\n& \n\nFlutter Developer.")

	// "\t": Horizontal Tab
	fmt.Println("Go Developer\t&\tFlutter Developer.")

	// "\r": Carriage return (moves cursor to the beginning of the line).
	fmt.Println("Go\rProgramming")
	fmt.Println("Hello\rWorld")

	// "\\": BackSlash
	fmt.Println("\\Go Developer\\")
	path := "C:\\Users\\Asib\\Documents"
	fmt.Println(path)

	// "\"": Double quote
	fmt.Println("\"Go\" Developer")

	// "\b": Backspace (moves the cursor back one position).
	fmt.Println("Go\bDeveloper")
	fmt.Println("Go \b Developer")

	// "\f": Form Feed
	fmt.Println("Go\fDeveloper")

	// "\u": Unicode (used for Unicode characters).
	fmt.Println("\u0048\u0049")

	// "\x": Hexadecimal representation.
	fmt.Println("\x48\x49")

}
