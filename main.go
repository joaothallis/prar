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

func main() {
	dir := os.Args[1]
	homeDir, err := os.UserHomeDir()

	filePath := homeDir + "/.config/prar.json"

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
