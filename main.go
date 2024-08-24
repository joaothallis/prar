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
	}

	return "./.prar.json", nil
}

func getProjectName() (string, error) {
	args := flag.Args()
	if len(args) > 0 {
		return os.Args[1], nil
	}

	cwd, err := os.Getwd()
	projectName := filepath.Base(cwd)
	return projectName, err
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

	repository, err := getProjectName()
	errorHandler(err)

	defer file.Close()

	var data map[string][]string
	err = json.NewDecoder(file).Decode(&data)
	errorHandler(err)

	return strings.Join(data[repository], ",")
}

func addReviewer(users string) {
	fmt.Println("Users: " + users)

	stdout, stderr, err := gh.Exec("pr", "edit", "--add-reviewer", users)
	if err != nil {
		fmt.Println(stderr.String())
		panic(err)
	}
	
	fmt.Println(stdout.String())
}

func main() {
	users := getUsers()
	addReviewer(users)
}
