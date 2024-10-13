package main

import (
	"fmt"
	"io"
	ui "main/client/consoleUI"
	logs "main/client/logs_writer"
	common "main/common"
	"net/http"
	"time"
)

func WaitUntilStateHappens(state string) {
  for {
		resp, err := http.Get("http://localhost:8080/polling_state")

		if err != nil {
			logs.LogError(fmt.Sprint("Error while polling state request:", err))
			time.Sleep(1 * time.Second)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logs.LogError(fmt.Sprint("Error while response parsing:", err))
			time.Sleep(1 * time.Second)
			continue
		}

		if string(body) != state {
			resp.Body.Close()
			continue
		}
    return

		// Polling interval is 1 second
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

func main() {
  ui.OutputWaitRegistration()
  TryToSubscribe()
  ui.OutputRegistrationComplete()
  WaitExperimentStart()
}
