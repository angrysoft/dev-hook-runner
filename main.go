package main

import (
	"flag"
	"fmt"
	"os"
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

	var hookArg string
	flag.StringVar(&hookArg, "arg", "", "Argument to pass to the hook")
	flag.StringVar(&hookArg, "a", "", "Argument to pass to the hook (shorthand)")

	flag.Parse()
	hooksDir := filepath.Join(dirFlag, ".hooks.d")

	if listFlag {
		hooks.ListHooks(hooksDir)
		return
	}

	args := flag.Args()

	if flag.NArg() > 0 {
		if !isDestDirExist(hooksDir) {
			fmt.Println("Hooks directory does not exist:", hooksDir)
			return
		}
		err := os.Chdir(dirFlag)
		if err != nil {
			fmt.Println("Error changing directory:", err)
			return
		}

		hookName := args[0]
		err = hooks.RunHook(hooksDir, hookName, hookArg)
		if err != nil {
			fmt.Println("Error running hook:", err)
		}
	}
}

func isDestDirExist(dir string) bool {
	if info, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	} else if !info.IsDir() {
		return false
	}
	return true
}
