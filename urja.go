package urja

import (
	"bytes"
	"github.com/A9u/urja/internal/pmset"
)

func GetBatteryStatus() (status string, err error) {
	status, err = pmset.BatteryStatus()

	return status, err
}

func RunCustomCommand(args ...string) (output bytes.Buffer, err error) {
	output, err = pmset.ExecuteCustomCommand(args...)

	return output, err
}

func GetHelp() (output bytes.Buffer, err error) {
	output, err = pmset.ExecuteHelpCommand()

	return output, err
}
