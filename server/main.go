package main

import (
	"bufio"
	"fmt"
	"io"
	common "main/common"
	ui "main/server/consoleUI"
	db "main/server/db"
	logs "main/server/logs_writer"
	utils "main/server/utils"
	"net/http"
	"os"
	"strconv"
)

var current_state utils.AtomicString
var current_experiment_number = 0
var currrent_guessed_number = -1

func subscribeHandler(w http.ResponseWriter, r *http.Request) {
	state := current_state.Load()
	if state != common.SUBSCRIPTION_STATE {
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
	state := current_state.Load()
	fmt.Fprint(w, state)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	clientIP := r.RemoteAddr
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	bodyStr := string(body)
	bodyInt, err := strconv.Atoi(bodyStr)
	if err != nil {
		http.Error(w, "Invalid body format. Expected an integer.", http.StatusBadRequest)
		return
	}

	db.WriteSubmission(common.DB_FILE_PATH, clientIP, current_experiment_number, bodyInt)
	ui.OutputQueries(current_experiment_number)

	if bodyInt == currrent_guessed_number {
		fmt.Fprintln(w, "Hoourray!! You won!!")
	} else if bodyInt < currrent_guessed_number {
		fmt.Fprintln(w, "Value "+string(body)+" less than expected")
	} else {
		fmt.Fprintln(w, "Value "+string(body)+" more than expected")
	}
}

func WaitUntilExperimentStarts() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		if input == "yes\n" {
			InitializeExperimentState()
			return
		} else {
			ui.OutputRegisteredStats()
		}
	}
}

func WaitUntilSubscriptionStarts() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		if input == "yes\n" {
			InitializeSubscriptionState()
			return
		} else {
			ui.OutputQueries(current_experiment_number)
		}
	}
}

func InitializeSubscriptionState() {
	current_experiment_number += 1
	current_state.Store(common.SUBSCRIPTION_STATE)
	ui.OutputRegisteredStats()
	go WaitUntilExperimentStarts()
}

func InitializeExperimentState() {
	current_state.Store(common.EXPERIMENT_STATE)
	currrent_guessed_number = 1234
	ui.OutputQueries(current_experiment_number)
	go WaitUntilSubscriptionStarts()
}

func main() {

	db.ReinitializeDatabase(common.DB_FILE_PATH)

	http.HandleFunc(common.HANDLER_SUBSCRIBE, subscribeHandler)
	http.HandleFunc(common.HANDLER_POLLING_STATE, pollingStateHandler)
	http.HandleFunc(common.HANDLER_SUBMIT, submitHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")

	InitializeSubscriptionState()

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
