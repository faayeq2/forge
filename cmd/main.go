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

func user_input() {
	fmt.Println("type of commit: ")
	comm_type := promptui.Select{
		Label: "commit type:",
		Items: config.Gitverbs,
	}

	_, selected_type, _ := comm_type.Run()

	fmt.Println("")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter commit message: ")

	msg, _ := reader.ReadString('\n')
	msg = strings.TrimSpace(msg)

	fullMsg := fmt.Sprintf("%s: %s", selected_type, msg)
	cmd := exec.Command("git", "commit", "-m", fullMsg)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("git commit failed: ", err)
		return
	}
	fmt.Println("commit successful")
}

func main() {
	user_input()
}
