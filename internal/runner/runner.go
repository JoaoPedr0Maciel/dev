package runner

import (
	"bytes"
	"os/exec"
	"time"
)

type Result struct {
	Stdout   string
	Stderr   string
	Duration time.Duration
	Err      error
}

func Run(cmd string) Result {
	start := time.Now()

	var stdout, stderr bytes.Buffer
	c := exec.Command("sh", "-c", cmd)
	c.Stdout = &stdout
	c.Stderr = &stderr

	err := c.Run()

	return Result{
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
		Duration: time.Since(start),
		Err:      err,
	}
}
