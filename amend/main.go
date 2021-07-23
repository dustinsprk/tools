package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	var edit bool
	var testRun bool
	flag.BoolVar(&edit, "edit", false, "edit commit message")
	flag.BoolVar(&edit, "e", false, "edit commit message (abbrv)")
	flag.BoolVar(&testRun, "testrun", false, "test run")
	flag.BoolVar(&testRun, "t", false, "test run (abbrv)")
	flag.Parse()
	now := time.Now().Format("Mon Jan 2 15:04:05 MST 2006")
	args := []string{"commit", "--amend", "--date", now}
	if !edit {
		args = append(args, "--no-edit")
	}
	cd := fmt.Sprintf("GIT_COMMITTER_DATE=\"%s\"", now)
	e := []string{cd}
	if testRun {
		fmt.Println(cd, "git", strings.Join(args, " "))
		return
	}
	cmd := exec.Command("git", args...)
	cmd.Env = append(e, os.Environ()...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Printf("error %v\n", err)
	}
}
