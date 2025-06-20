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

//example code

func user_input() error {

	fmt.Println("type of commit: ")
	comm_type := promptui.Select{
		Label: "commit type:",
		Items: config.Gitverbs,
	}

	_, selected_type, err := comm_type.Run()
	if err != nil {
		return fmt.Errorf("failed to execute selection list: %w", err)
	}

	fmt.Println("")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter commit message: ")

	msg, _ := reader.ReadString('\n')
	msg = strings.TrimSpace(msg)

	fullMsg := fmt.Sprintf("%s: %s", selected_type, msg)
	cmd := exec.Command("git", "commit", "-m", fullMsg)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("git commit failed: %w", err)
	}

	return nil
}

func main() {
	user_input()
}
