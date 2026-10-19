package hooks

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"go.yaml.in/yaml/v4"
	// "path/filepath"
)

type Cmd struct {
	Cmd  string   `yaml:"cmd"`
	Args []string `yaml:"args"`
}

type Hook struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	EnvFile     string   `yaml:"env_file"`
	DependsOn   []string `yaml:"depends_on"`
	Execs       []Cmd    `yaml:"execs"`
	Script      string   `yaml:"script"`
	NeedsArgs   bool     `yaml:"needs_args"`
}

type Env struct {
	Vars []EnvItem `yaml:"vars"`
}

type EnvItem struct {
	Key   string `yaml:"key"`
	Value string `yaml:"value"`
}

var firedHooks = make(map[string]bool)

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

func RunHook(hooksDir, hookName string, hookArg string) error {

	filePath := filepath.Join(hooksDir, hookName+".hook.yml")
	hook, err := LoadHook(filePath)
	if err != nil {
		return err
	}
	fmt.Println("Running hook:", hook.Name)
	if hook.NeedsArgs && hookArg == "" {
		return fmt.Errorf("hook %s requires an argument", hook.Name)
	}
	firedHooks[hook.Name] = true
	for _, dep := range hook.DependsOn {
		if firedHooks[dep] {
			return fmt.Errorf("loop detected for dependency: %s", dep)
		}
		firedHooks[dep] = true
		err := RunHook(hooksDir, dep, hookArg)
		if err != nil {
			return err
		}
	}

	env, err := loadEnv(hook.EnvFile)
	if err != nil {
		return err
	}

	if hook.NeedsArgs {
		env.Vars = append(env.Vars, EnvItem{
			Key:   "HOOK_ARG",
			Value: hookArg,
		})
	}

	err = runExecs(hook.Execs, env)
	if err != nil {
		return err
	}
	err = runScript(hook.Script, env)
	if err != nil {
		return err
	}
	return nil
}

func runExecs(cmds []Cmd, env *Env) error {
	for _, cmd := range cmds {
		if len(cmd.Args) == 0 {
			continue
		}

		c := exec.Command(cmd.Cmd, cmd.Args...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		if env != nil {
			environment := os.Environ()
			for _, item := range env.Vars {
				environment = append(environment, fmt.Sprintf("%s=%s", item.Key, item.Value))
			}
			c.Env = environment
		}
		if err := c.Run(); err != nil {
			return err
		}
	}
	return nil
}

func runScript(script string, env *Env) error {
	if script == "" {
		return nil
	}

	c := exec.Command("bash", "-c", script)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if env != nil {
		environment := os.Environ()
		for _, item := range env.Vars {
			environment = append(environment, fmt.Sprintf("%s=%s", item.Key, item.Value))
		}
		c.Env = environment
	}
	if err := c.Run(); err != nil {
		return err
	}
	return nil
}

func loadEnv(envFile string) (*Env, error) {
	var envs Env
	if envFile == "" {
		return &envs, nil
	}
	data, err := os.ReadFile(envFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &envs)
	if err != nil {
		return nil, err
	}
	return &envs, nil
}
