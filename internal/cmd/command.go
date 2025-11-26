package cmd

import (
	"fmt"
	"os/exec"
	"strings"
)

func HandleCommands(command []byte) ([]byte, error) {
	fullCommand := strings.TrimSpace(string(command))

	if fullCommand == "" {
		return nil, nil
	}

	formattedLine := strings.Fields(fullCommand)

	return handle(formattedLine)
}

func handle(commands []string) ([]byte, error) {
	if len(commands) == 0 {
		return nil, fmt.Errorf("empty command")
	}

	output, err := exec.Command(commands[0], commands[1:]...).CombinedOutput()
	if err != nil {
		return output, fmt.Errorf("error to execute the command: %w", err)
	}

	return output, nil
}
