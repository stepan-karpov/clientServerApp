package ui

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func clearTerminal() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func OutputWaitRegistration() {
	clearTerminal()
	fmt.Println("Please, wait until registration starts!!")
}

func OutputRegistrationComplete() {
	clearTerminal()
	fmt.Println("You registration is successfull! Wait until experiment itself starts")
}
