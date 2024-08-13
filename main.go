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

func getPrarFilePath(globalArg bool) (string, error) {
	if globalArg {
		homeDir, err := os.UserHomeDir()
		return homeDir + "/.config/prar.json", err
	} else {
		homeDir, err := os.Getwd()
		return homeDir + "/.prar.json", err
	}
}

func getProjectName() (string, error) {
	args := flag.Args()
	if len(args) >= 1 {
		return os.Args[1], nil
	} else {
		cwd, err := os.Getwd()
		dirName := filepath.Base(cwd)
		return dirName, err
	}
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func getUsers(filePath string, dir string) string {
	file, err := os.Open(filePath)
	errorHandler(err)

	defer file.Close()

	var data map[string][]string
	err = json.NewDecoder(file).Decode(&data)
	errorHandler(err)

	return strings.Join(data[dir], ",")
}

func ghPrRequest(users string) {
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
	globalFlag := flag.Bool("global", false, "When you choose to using ./config/prar.json")
	flag.Parse()

	dir, err := getProjectName()
	errorHandler(err)

	filePath, err := getPrarFilePath(*globalFlag)
	errorHandler(err)

	users := getUsers(filePath, dir)
	ghPrRequest(users)
}
