package pmset

import (
	"bytes"
	"fmt"
	"os/exec"
)

func BatteryStatus() (batteryStatus string, err error) {
	output, err := executePmsetCommand("-g", "ps")

	if err != nil {
		return "", err
	}

	return output.String(), nil
}

func ExecuteCustomCommand(args ...string) (output bytes.Buffer, err error) {
	if len(args) == 0 {
		err = fmt.Errorf("executeCustomCommand: no args provided")
		return output, err
	}

	output, err = executePmsetCommand(args...)
	return output, err
}

func ExecuteHelpCommand() (output bytes.Buffer, err error) {
	output, err = executeCommand("man", "pmset")

	return output, err
}

func executePmsetCommand(args ...string) (output bytes.Buffer, err error) {
	output, err = executeCommand("pmset", args...)

	return output, err
}

func executeCommand(command string, args ...string) (output bytes.Buffer, err error) {
	cmd := exec.Command(command, args...)

	cmd.Stdout = &output

	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("executeCommand: %s", output)
		return output, err
	}

	return output, nil
}
