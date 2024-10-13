package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	ui "main/client/consoleUI"
	logs "main/client/logs_writer"
	common "main/common"
	"net/http"
	"os"
	"time"
)

func GetCurrentServerState() string {
	resp, err := http.Get("http://localhost:8080/polling_state")

	if err != nil {
		logs.LogError(fmt.Sprint("Error while polling state request:", err))
		return "None"
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.LogError(fmt.Sprint("Error while response parsing:", err))
		return "None"
	}
	answer := string(body)
	resp.Body.Close()
	return answer
}

func WaitUntilStateHappens(state string) {
	for {
		current_state := GetCurrentServerState()
		if current_state == state {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func Register() {
	_, err := http.Get("http://localhost:8080/subscribe")

	if err != nil {
		logs.LogError(fmt.Sprint("Error while registrating client", err))
	}
}

func TryToSubscribe() {
	WaitUntilStateHappens(common.SUBSCRIPTION_STATE)
	Register()
}

func WaitExperimentStart() {
	WaitUntilStateHappens(common.EXPERIMENT_STATE)
}

func InitializeGuessProcess() {
  ui.OutputGuessInterface()
	var guessed bool = false
	reader := bufio.NewReader(os.Stdin)

	for !guessed {
		var current_state string = GetCurrentServerState()
		if current_state != common.EXPERIMENT_STATE {
			break
		}

		input, _ := reader.ReadString('\n')
		requestBody := bytes.NewBufferString(input[:len(input) - 1])

		resp, _ := http.Post("http://localhost:8080/submit", "text/plain", requestBody)
    logs.LogInfo("Query request sent")
		body, _ := io.ReadAll(resp.Body)

		ui.OutputGuessResult(string(body))
	}
	ui.OutputWaitRegistrationAgain(guessed)
}

func main() {
	ui.OutputWaitRegistration()
	TryToSubscribe()
	ui.OutputRegistrationComplete()
	WaitExperimentStart()
	InitializeGuessProcess()
}
