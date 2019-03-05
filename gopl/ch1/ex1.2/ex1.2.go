package main

import (
	"fmt"
	"os"
)

func main() {
	for k, v := range os.Args {
		fmt.Printf("%v: %v\n", k, v)
	}
}
