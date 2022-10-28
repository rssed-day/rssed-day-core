package test

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestCliPipe(t *testing.T) {
	arg := []string{"pipe", "-c", "configs/dev/plugins.yaml"}
	cmd := exec.Command("rsseday_cli.exe", arg...)
	_, err := cmd.Output()
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	fmt.Println(cmd.Stdout)
}
