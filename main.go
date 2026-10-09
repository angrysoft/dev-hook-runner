package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/angrysoft/dhr/hooks"
)

func main() {
	var listFlag bool
	flag.BoolVar(&listFlag, "list", false, "List all hooks")
	flag.BoolVar(&listFlag, "l", false, "List all hooks (shorthand)")
	var dirFlag string
	flag.StringVar(&dirFlag, "dir", ".", "Directory containing hooks")
	flag.StringVar(&dirFlag, "d", ".", "Directory containing hooks (shorthand)")
	hooksDir := filepath.Join(dirFlag, ".hooks.d")

	flag.Parse()

	if listFlag {
		hooks.ListHooks(hooksDir)
		return
	}

	args := flag.Args()

	if flag.NArg() > 0 {
		// Tutaj można dodać logikę obsługi innych argumentów
		for _, arg := range args {
			fmt.Println("Handling argument:", arg)
		}
	}
}
