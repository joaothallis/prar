package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/cli/go-gh/v2"
)

// JSON example
// {"ar": ["joaothallis"], "elixir": ["josevalim", "eksperimental"]}

func getPrarFilePath(globalArg string) (string, error) {
	switch globalArg {
	case "--global":
		homeDir, err := os.UserHomeDir()
		return homeDir + "/.config/prar.json", err
	default:
		homeDir, err := os.Getwd()
		return homeDir + "/.prar.json", err
	}
}

func main() {
	args := os.Args[1:]
	dir := os.Args[1]
	var globalArg string

	if len(args) > 1 {
		globalArg = os.Args[2]
	} else {
		globalArg = ""
	}

	filePath, err := getPrarFilePath(globalArg)
	fmt.Printf(filePath)
	if err != nil {
		panic(err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data map[string][]string
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		panic(err)
	}
	users := strings.Join(data[dir], ",")

	fmt.Println("Users: " + users)

	stdout, stderr, err := gh.Exec("pr", "edit", "--add-reviewer", users)
	if err != nil {

		stringStderr := stderr.String()

		fmt.Println(stringStderr)

		panic(err)
	}
	stringStdout := stdout.String()
	fmt.Println(stringStdout)
}
