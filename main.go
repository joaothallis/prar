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

func main() {
	globalFlag := flag.Bool("global", false, "When you choose to using ./config/prar.json")
	flag.Parse()
	fmt.Println(flag.Args())

	dir, err := getProjectName()
	if err != nil {
		panic(err)
	}
	filePath, err := getPrarFilePath(*globalFlag)
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
