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

func OutputWaitRegistration(already_won bool) {
	if !already_won {
		clearTerminal()
	}
	fmt.Println("Please, wait until registration starts!!")
}

func OutputGuessResult(response string) {
	fmt.Print("Result: " + response)
}

func OutputWaitRegistrationAgain(guessed bool) {
	if !guessed {
		clearTerminal()
		fmt.Print("Experiment ended! You didn't guess the number.\n")
	} else {
		fmt.Print("Experiment ended! You've managed to guess the number!\n")
	}
	fmt.Println("To register once again, wait until next registration starts!!")
}

func OutputRegistrationComplete() {
	clearTerminal()
	fmt.Println("You registration is successfull! Wait until experiment itself starts")
}

func OutputWaitQueryResponse() {
	fmt.Print("Please, input value: ")
}

func OutputGuessInterface() {
	clearTerminal()
	fmt.Println("Experiment started\n\n")
	OutputWaitQueryResponse()
}
