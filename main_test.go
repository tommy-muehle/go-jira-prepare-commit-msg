package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

const (
	TEST_FILE = "test.txt"
)

func TestHook(t *testing.T) {
	exec.Command("git", "checkout", "-b", "feature/MYPRO-123").Run()

	ioutil.WriteFile(TEST_FILE, []byte("Added new feature"), 0644)
	exec.Command("go", "run", "main.go", TEST_FILE).Run()

	msg, err := ioutil.ReadFile(TEST_FILE)

	if err != nil {
		t.Fail()
	}

	if string(msg) != "MYPRO-123: Added new feature" {
		t.Fail()
	}

	exec.Command("git", "checkout", "master").Run()
	exec.Command("git", "branch", "-d", "feature/MYPRO-123").Run()
	os.Remove(TEST_FILE)
}
