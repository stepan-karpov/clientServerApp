package logs

import "fmt"

var debug_logs bool = false
var info_logs bool = false
var error_logs bool = false

func LogDebug(output string) {
	if !debug_logs {
		return
	}
	fmt.Println("[ DEBUG ] " + output)
}

func LogInfo(output string) {
	if !info_logs {
		return
	}
	fmt.Println("[ INFO ] " + output)
}

func LogError(output string) {
	if !error_logs {
		return
	}
	fmt.Println("[ ERROR!!! ] " + output)
}
