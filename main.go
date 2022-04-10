package main

import (
	"os"
)

func main() {
	for _, args := range os.Args {
		run(Function(args))
	}
}
