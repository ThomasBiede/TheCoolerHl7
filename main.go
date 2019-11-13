package main

import "hl7/parser"

func main() {
	p := parser.ADT_A01{}
	p.ParseFile("test.txt")
}
