package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
)

const (
	JIRA_ISSUE_FORMAT = "[A-Z]{2,}-[0-9]{1,}"
)

func main() {
	if 0 == len(os.Args[1]) {
		return
	}

	messageFile := os.Args[1]
	message, _ := ioutil.ReadFile(messageFile)

	branch, err := exec.Command("git", "symbolic-ref", "--short", "HEAD").CombinedOutput()

	if err != nil {
		return
	}

	regex := regexp.MustCompile(JIRA_ISSUE_FORMAT)
	name := regex.Find(branch)

	if 0 == len(name) {
		return
	}

	ioutil.WriteFile(messageFile, []byte(fmt.Sprintf("%s: %s", string(name), string(message))), 0644)
}
