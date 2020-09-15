package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func ConsoleClear() {
	currOS := runtime.GOOS
	switch currOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
