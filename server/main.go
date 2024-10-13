package main

import (
	"bufio"
	"fmt"
	common "main/common"
	ui "main/server/consoleUI"
	db "main/server/db"
	logs "main/server/logs_writer"
	"net/http"
	"os"
)

var current_state = common.SUBSCRIPTION_STATE
var current_experiment_number = 0

func subscribeHandler(w http.ResponseWriter, r *http.Request) {
	if current_state != common.SUBSCRIPTION_STATE {
		fmt.Fprintf(w, "You can not register at the experiment right now!")
		return
	}
	clientIP := r.RemoteAddr

	logs.LogDebug("Request from " + clientIP)
	err := db.SubscribeUser(common.DB_FILE_PATH, clientIP, current_experiment_number)

	if err != nil {
		fmt.Fprintf(w, "You already registered!")
		return
	}

	fmt.Fprintf(w, "Registration success")
	ui.OutputRegisteredStats()
}

func pollingStateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, current_state)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "UNIMPLEMENTED")
}

func main() {
	db.ReinitializeDatabase(common.DB_FILE_PATH)

	http.HandleFunc(common.HANDLER_SUBSCRIBE, subscribeHandler)
	http.HandleFunc(common.HANDLER_POLLING_STATE, pollingStateHandler)
	http.HandleFunc(common.HANDLER_SUBMIT, submitHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")

	ui.OutputRegisteredStats()

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ := reader.ReadString('\n')
			if input == "yes" {
        current_state = common.EXPERIMENT_STATE
      } else {
        ui.OutputRegisteredStats()
      }
		}
	}()

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
