package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	ConfigPath string = "/.push-me-config.yml"
)

func NewConfig() (*Config, error) {
	home, _ := os.UserHomeDir()
	fullConfigPath := home + ConfigPath

	configFile, err := os.ReadFile(fullConfigPath)
	if err != nil {
		return nil, fmt.Errorf("Error reading config file: %s", err)
	}

	// Unmarshal the config file into the Config struct
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling config file: %v", err)
	}

	// Add home prefix to repos
	for i, repo := range config.Repos {
		config.Repos[i] = home + "/" + repo
	}

	return &config, nil
}

type Config struct {
	Repos []string `yaml:"repos"`
}

// Git command wrapper for `git add`
func Add(repo string) error {
	out, err := exec.Command("git", "-C", repo, "add", ".").Output()
	if err != nil {
		fmt.Println("`git add` exited abnormally")
		return err
	}

	output := string(out)

	fmt.Print(output)
	return nil
}

// Git command wrapper for `git commit`
func Commit(repo string) error {

	timestamp := time.Now()

	message := "auto commit: " + timestamp.Format("20060102150405")
	out, err := exec.Command("git", "-C", repo, "commit", "-m", message).Output()
	if err != nil {
		fmt.Println("`git commit` exited abnormally")
		return err
	}

	output := string(out)

	fmt.Print(output)
	return nil
}

// Git command wrapper for `git push`
func Push(repo string) error {
	out, err := exec.Command("git", "-C", repo, "push").Output()
	if err != nil {
		fmt.Println("`git push` exited abnormally")
		return err
	}

	output := string(out)

	fmt.Print(output)
	return nil
}

func main() {
	config, err := NewConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Loop through the repos array
	for _, repo := range config.Repos {
		fmt.Println("Repository:", repo)

		if err := Add(repo); err != nil {
			fmt.Println("Git add returned an error")
		}
		if err := Commit(repo); err != nil {
			fmt.Println("Git commit returned an error")
		}
		if err := Push(repo); err != nil {
			fmt.Println("Git push returned an error")
		}
	}
}
