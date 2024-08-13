package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cli/go-gh/v2"
)

// JSON example
// {"ar": ["joaothallis"], "elixir": ["josevalim", "eksperimental"]}

func getPrarFilePath() (string, error) {
	globalFlag := flag.Bool("global", false, "When you choose to using ./config/prar.json")
	flag.Parse()
	if *globalFlag {
		homeDir, err := os.UserHomeDir()
		return homeDir + "/.config/prar.json", err
	} else {
		return "./.prar.json", nil
	}
}

func getProjectName() (string, error) {
	args := flag.Args()
	if len(args) > 0 {
		return os.Args[1], nil
	} else {
		cwd, err := os.Getwd()
		projectName := filepath.Base(cwd)
		return projectName, err
	}
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func getUsers() string {
	filePath, err := getPrarFilePath()
	errorHandler(err)

	file, err := os.Open(filePath)
	errorHandler(err)

	dir, err := getProjectName()
	errorHandler(err)

	defer file.Close()

	var data map[string][]string
	err = json.NewDecoder(file).Decode(&data)
	errorHandler(err)

	return strings.Join(data[dir], ",")
}

func addReviewer(users string) {
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

func main() {
	users := getUsers()
	addReviewer(users)
}
