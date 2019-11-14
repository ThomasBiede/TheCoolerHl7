package main

import (
	"encoding/xml"
	"fmt"
	"hl7/parser"
)

func main() {
	p := &parser.ADT_A01{}
	p.ParseFile("test.txt")

	out, err := xml.Marshal(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
