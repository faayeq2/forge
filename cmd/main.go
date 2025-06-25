package main

import (
	"bufio"
	"fmt"
	"forge/internal/config"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

func user_input() error {

	// for git add {file(s)/.}

	gitStatus := exec.Command("git", "status", "--porcelain") // get changed files
	output, err := gitStatus.Output()
	if err != nil {
		return fmt.Errorf("failed to get git status: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var files []string

	if len(lines) == 0 || (len(lines) == 1 && lines[0] == "") {
		fmt.Println("No changes to stage.")
		return nil
	}

	// add "Add all files" option
	files = append(files, "Add all files")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			file := parts[1]
			files = append(files, file)
		}
	}

	// prompt user to select a file or add all
	filePrompt := promptui.Select{
		Label: "Select file to add",
		Items: files,
	}

	selectedIndex, selectedItem, err := filePrompt.Run()
	if err != nil {
		return fmt.Errorf("file selection cancelled: %w", err)
	}

	// run git add
	var add *exec.Cmd
	if selectedIndex == 0 {
		add = exec.Command("git", "add", ".")
	} else {
		add = exec.Command("git", "add", selectedItem)
	}

	add.Stdout = os.Stdout
	add.Stderr = os.Stderr
	if err := add.Run(); err != nil {
		return fmt.Errorf("git add failed: %w", err)
	}

	// ui, commit type
	comm_type := promptui.Select{
		Label: "Commit type",
		Items: config.Gitverbs,
	}

	_, selected_type, err := comm_type.Run()
	if err != nil {
		return fmt.Errorf("failed to select commit type: %w", err)
	}

	// read commit message
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter commit message: ")
	msg, _ := reader.ReadString('\n')
	msg = strings.TrimSpace(msg)

	// create commit
	fullMsg := fmt.Sprintf("%s: %s", selected_type, msg)
	cmd := exec.Command("git", "commit", "-m", fullMsg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git commit failed: %w", err)
	}

	// codeberg push, hardcoded via remote "origin"
	pushCodeberg := exec.Command("git", "push", "origin")
	pushCodeberg.Stdout = os.Stdout
	pushCodeberg.Stderr = os.Stderr
	if err := pushCodeberg.Run(); err != nil {
		return fmt.Errorf("push to Codeberg failed: %w", err)
	}

	// github push, hardcoded via remote "github"
	pushGithub := exec.Command("git", "push", "github")
	pushGithub.Stdout = os.Stdout
	pushGithub.Stderr = os.Stderr
	if err := pushGithub.Run(); err != nil {
		return fmt.Errorf("push to GitHub failed: %w", err)
	}

	return nil
}

func main() {
	if err := user_input(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
