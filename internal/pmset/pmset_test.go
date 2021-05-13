package pmset

import "testing"

func TestExecuteCustomcommand(t *testing.T) {
	defaultErrMsg := `Usage: pmset <options>
See pmset(1) for details: 'man pmset'`

	_, err := ExecuteCustomCommand("abc", "xyz")
	expectedErrMsg := "executeCommand:" + defaultErrMsg

	if err != nil {
		t.Logf("ExecuteCustomCommand(\"abc\") success, expected error: %v, got %v", expectedErrMsg, err)
	} else {
		t.Errorf("ExecuteCustomCommand(\"abc\") failure, expected error: %v, got %v", "executeCommand", err)
	}

	_, err = ExecuteCustomCommand()
	if err != nil {
		t.Logf("ExecuteCustomCommand() success, expected error: %v, got %v", "executeCustomCommand: no args provided", err)
	} else {
		t.Errorf("ExecuteCustomCommand() failure, expected error: %v, got %v", "executeCustomCommand: no args provided", err)
	}

	output, err := ExecuteCustomCommand("-g")
	if err != nil {
		t.Errorf("ExecuteCustomCommand(\"--help\") failure, expected success with value: %s", "System-wide power settings")
	} else {
		t.Logf("ExecuteCustomCommand(\"--help\") success, expected value: %v, got %v", "System-wide power settings", output.String())
	}
}

func TestBatteryStatus(t *testing.T) {
	message, err := BatteryStatus()

	if err != nil {
		t.Errorf("BatteryStatus() failure, expected error: %v, got %v", "Now drawing from 'Battery Power'", err)
	} else {
		t.Logf("BatteryStatus() success, expected error: %v, got %v", "Now drawing from 'Battery Power'", message)
	}
}

func TestExecuteHelpCommand(t *testing.T) {
	output, err := ExecuteHelpCommand()

	if err != nil {
		t.Errorf("BatteryStatus() failure, got error: %v", err)
	} else {
		t.Logf("BatteryStatus() success, got %v", output.String())
	}
}
