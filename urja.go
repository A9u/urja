package urja

import (
	"bytes"
	"log"
	"os/exec"
)

func GetBatteryStatus() string {
	cmd := exec.Command("pmset", "-g", "ps")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return out.String()
}
