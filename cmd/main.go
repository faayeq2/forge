package main

import (
	"fmt"
	"forge/internal/config"
	//"os/exec"
)

func disp_verbs() {
	for i := 0; i < len(config.Gitverbs); i++ {
		fmt.Println(config.Gitverbs[i])
	}
}

func main() {
	disp_verbs()
	//cmd := exec.Command(str)
	//cmd.Run()
}
