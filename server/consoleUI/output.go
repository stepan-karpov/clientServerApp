package ui

import (
	"fmt"
	common "main/common"
	db "main/server/db"
	"os"
	"os/exec"
	"runtime"
	"strings"
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

func OutputExperimentInfo() {
	queriesInfo, err := db.GetExperimentsInfo(common.DB_FILE_PATH)
	if err != nil {
		fmt.Println("Error retrieving queries:", err)
		return
	}

	fmt.Println("Registered Subscriptions:")
	fmt.Printf("%-15s %-15s %-15s\n", "ID", "NumberOfQueries", "MeanValue")
	fmt.Println(strings.Repeat("-", 42))
	
	for _, experiment := range queriesInfo {
		fmt.Printf("%-15d %-15d %-15f\n", experiment.ID, experiment.NumberOfQueries, experiment.MeanValue)
	}

	fmt.Println(strings.Repeat("\n", 10))
}

func OutputRegisteredStats() {
	clearTerminal()
	OutputExperimentInfo()

	subscriptionsInfo, err := db.GetAllSubscriptions(common.DB_FILE_PATH)
	if err != nil {
		fmt.Println("Error retrieving subscriptions:", err)
		return
	}

	fmt.Println("Registered Subscriptions:")
	fmt.Printf("%-5s %-15s %-20s\n", "ID", "IP", "Experiment Number")
	fmt.Println(strings.Repeat("-", 42))

	for _, subscription := range subscriptionsInfo {
		fmt.Printf("%-5d %-15s %-20d\n", subscription.ID, subscription.IP, subscription.ExperimentNumber)
	}

	fmt.Printf("\nDo you want to start an experiment? (yes/no)\n\n")
	fmt.Printf("Your input is: ")
}

func OutputQueries(experiment_number int) {
	clearTerminal()
	OutputExperimentInfo()

	queriesInfo, err := db.GetQueriesInfo(common.DB_FILE_PATH)
	if err != nil {
		fmt.Println("Error retrieving queries:", err)
		return
	}

	fmt.Println("Registered Subscriptions:")
	fmt.Printf("%-15s %-15s %-15s\n", "Query ID", "IP", "Query Value")
	fmt.Println(strings.Repeat("-", 42))

	for _, query := range queriesInfo {
		if query.ExperimentNumber == experiment_number {
			fmt.Printf("%-15d %-15s %-15d\n", query.ID, query.IP, query.QueryValue)

		}
	}

	fmt.Printf("\nDo you want to finish an experiment? (yes/no)\n\n")
	fmt.Printf("Your input is: ")
}
