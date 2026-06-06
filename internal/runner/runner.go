package runner

import (
	"bytes"
	"os"
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

func RunLive(cmd string) error {
	c := exec.Command("sh", "-c", cmd)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	return c.Run()
}
