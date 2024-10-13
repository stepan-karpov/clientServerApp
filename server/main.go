package main

import (
	"fmt"
	common "main/common"
	db "main/server/db"
	logs "main/server/logs_writer"
	"net/http"
)

const (
  SUBSCRIPTION_STATE = "SUBSCRIPTION_STATE"
  EXPERIMENT_STATE = "EXPERIMENT_STATE"
)

var current_state = SUBSCRIPTION_STATE

func subscribeHandler(w http.ResponseWriter, r *http.Request) {
  if current_state != SUBSCRIPTION_STATE {
    fmt.Fprintf(w, "Bad for you")
    return
  }
	clientIP := r.RemoteAddr
	logs.LogDebug("Request from " + clientIP)

	fmt.Fprintf(w, "Привет, мир!")
}

func main() {
	db.ReinitializeDatabase(common.DB_FILE_PATH)

	http.HandleFunc(common.HANDLER_SUBSCRIBE, subscribeHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
