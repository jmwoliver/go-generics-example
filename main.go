package main

import (
	"flag"
	"fmt"

	"github.com/jmwoliver/go-generics-example/generics"
	"github.com/jmwoliver/go-generics-example/nogenerics"
)

func main() {
	var useGenerics bool
	flag.BoolVar(&useGenerics, "generics", false, "Choose to use generic example or non-generic example. Defaults to non-generic example.")
	flag.Parse()
	if useGenerics {
		fmt.Println("Example with generics:")
		generics.Example()
	} else {
		fmt.Println("Example without generics:")
		nogenerics.Example()
	}
}
