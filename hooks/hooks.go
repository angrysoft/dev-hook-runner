package hooks

import (
	"fmt"
	"os"
	"path/filepath"

	"go.yaml.in/yaml/v4"
)

type Cmd struct {
	Exec []string `yaml:"exec"`
}

type Hook struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	EnvFile     string `yaml:"env_file"`
	Execs       []Cmd  `yaml:"execs"`
	Script      string `yaml:"script"`
}

func ListHooks(hooksDir string) {
	if _, err := os.Stat(hooksDir); err == nil {
		files, err := filepath.Glob(filepath.Join(hooksDir, "*.hook.yml"))
		if err != nil {
			fmt.Println("Error listing hooks:", err)
			return
		}
		for _, file := range files {
			PrintHook(file)
		}
	} else {
		fmt.Println("Hooks directory does not exist:", hooksDir)
	}
}

func LoadHook(filePath string) (*Hook, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var hook Hook
	err = yaml.Unmarshal(data, &hook)
	if err != nil {
		return nil, err
	}
	return &hook, nil
}

func PrintHook(filePath string) {
	hook, err := LoadHook(filePath)
	if err != nil {
		fmt.Println("Error loading hook:", err)
		return
	}
	fmt.Printf("%s : %s\n", hook.Name, hook.Description)
}
