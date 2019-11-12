package main

import "hl7/parser"

func main() {
	p := parser.NewMainParser("test.txt")
	p.ParseFile()
}
