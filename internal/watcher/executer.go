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

	err = writeFile(logFile, stdout)
	if err != nil {
		return err
	}

	return nil
}

func writeFile(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.Write(data); err != nil {
		return err
	}
	return nil
}
