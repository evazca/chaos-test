package exec

import (
	"bytes"
	"os/exec"
)

func Command(cmdStr string)  (*exec.Cmd, *bytes.Buffer, *bytes.Buffer){
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	return cmd,&stdout,&stderr
}
