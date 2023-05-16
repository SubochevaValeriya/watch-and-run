package worker

import (
	"os"
	"os/exec"
	"strings"
)

type Executer interface {
	ExecuteCommand()
}

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

	err = writeToFile(stdout, logFile)
	if err != nil {
		return err
	}

	return nil
}

func writeToFile(data []byte, logfile string) error {
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		return err
	}
	//if err := f.Close(); err != nil {
	//	return err
	//}
	return nil
}
