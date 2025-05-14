package main

import (
	"fmt"
	"forge/internal/config"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
	//"os/exec"
)

func user_input() {
	fmt.Println("type of commit: ")
	comm_type := promptui.Select{
		Label: "commit type:",
		Items: config.Gitverbs,
	}

	comm_type.Run()

	fmt.Println("")

	fmt.Print("enter commit message")
	var msg string
	fmt.Scanln(&msg)
	cmd := exec.Command("git", "commit", "-m", msg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("git commit failed: ", err)
		return
	}
	fmt.Println("commit succesful")
}

func main() {
	user_input()
}
