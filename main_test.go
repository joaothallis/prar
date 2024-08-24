package main

import (
	"flag"
	"os"
	"path/filepath"
	"testing"
)

func TestGetPrarFilePath(t *testing.T) {
	t.Run("Global Flag False", func(t *testing.T) {
		path, err := getPrarFilePath()
		if err != nil {
			t.Errorf("error: %v", err)
		}

		expectedPath := "./.prar.json"
		if path != expectedPath {
			t.Errorf("Expected path %s, but got %s", expectedPath, path)
		}
	})

	t.Run("Global Flag True", func(t *testing.T) {
		os.Args = []string{"cmd", "-global"}
		flag.CommandLine = flag.NewFlagSet(os.Args[1], flag.ExitOnError)

		path, err := getPrarFilePath()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		homeDir, _ := os.UserHomeDir()
		expectedPath := homeDir + "/.config/prar.json"
		if path != expectedPath {
			t.Errorf("Expected path %s, but got %s", expectedPath, path)
		}
	})
}

func TestGetProjectName(t *testing.T) {
	t.Run("No Argument Provided", func(t *testing.T) {
		projectName, err := getProjectName()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		cwd, _ := os.Getwd()
		expectedName := filepath.Base(cwd)
		if projectName != expectedName {
			t.Errorf("Expected project name %s, but got %s", expectedName, projectName)
		}
	})

	t.Run("Argument Provided", func(t *testing.T) {
		os.Args = []string{"cmd", "myproject"}
		flag.CommandLine = flag.NewFlagSet(os.Args[1], flag.ExitOnError)
		flag.Parse()

		projectName, err := getProjectName()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		expectedName := "myproject"
		if projectName != expectedName {
			t.Errorf("Expected project name %s, but got %s", expectedName, projectName)
		}
	})
}

func TestGetUsers(t *testing.T) {
	expectedResult := "claudionts,joaothallis"
	result := getUsers()
	if result != expectedResult {
		t.Errorf("Expected %s, but got %s", expectedResult, result)
	}
}
