package watcher

import (
	"os"
	"os/exec"
	"strings"
)

func executeCommand(command string, logFile string) error {
	cmd := exec.Command(command)
	commandAndArgs := strings.Split(command, " ")
	if len(commandAndArgs) > 1 {
		cmd = exec.Command(commandAndArgs[0], commandAndArgs[1:]...)
	}
	stdout, err := cmd.Output()
	if err != nil {
		return err
	}

	err = os.WriteFile(logFile, stdout, 0644)
	if err != nil {
		return err
	}

	return nil
}
