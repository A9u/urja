package pmset

import (
	"bytes"
	"fmt"
	"os/exec"
)

// BatteryStatus is a function used to get output of `pmset -g ps` command
func BatteryStatus() (batteryStatus string, err error) {
	output, err := executePmsetCommand("-g", "ps")

	if err != nil {
		return "", err
	}

	return output.String(), nil
}

// ExecuteCustomCommand is a function used to execute custom commands like `pmset -g therm`. To call this function, only
// the arguments, for example, ExecuteCustomCommand("-g", "therm")
func ExecuteCustomCommand(args ...string) (output bytes.Buffer, err error) {
	if len(args) == 0 {
		err = fmt.Errorf("executeCustomCommand: no args provided")
		return output, err
	}

	output, err = executePmsetCommand(args...)
	return output, err
}

// ExecuteHelpCommand() is a function used to see the manual page of pmset command
func ExecuteHelpCommand() (output bytes.Buffer, err error) {
	output, err = executeCommand("man", "pmset")

	return output, err
}

// executePmsetCommand is a util function to which we pass the options
func executePmsetCommand(args ...string) (output bytes.Buffer, err error) {
	output, err = executeCommand("pmset", args...)

	return output, err
}

// executeCommand is a util function in case in future we decide to support other commands
func executeCommand(command string, args ...string) (output bytes.Buffer, err error) {
	cmd := exec.Command(command, args...)

	cmd.Stdout = &output

	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("executeCommand: %s", output.String())
		return output, err
	}

	return output, nil
}
