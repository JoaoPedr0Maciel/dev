package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/JoaoPedr0Maciel/dev/internal/config"
	"github.com/JoaoPedr0Maciel/dev/internal/formatter"
	"github.com/JoaoPedr0Maciel/dev/internal/runner"
	"github.com/JoaoPedr0Maciel/dev/internal/tui"
)

var Version = "dev"

func main() {
	var configPath string
	flag.StringVar(&configPath, "path", "dev.yaml", "path to the yaml configuration file")
	flag.Parse()

	cfg, err := config.Load(configPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	args := flag.Args()
	if len(args) > 0 {
		switch args[0] {
		case "version", "--version", "-v":
			fmt.Println("dev", Version)
			return
		case "update":
			runUpdate()
			return
		default:
			if len(cfg.Tasks) == 0 {
				fmt.Fprintln(os.Stderr, "Error: no tasks found")
				os.Exit(1)
			}
			runDirect(cfg, args[0])
			return
		}
	}

	if len(cfg.Tasks) == 0 {
		fmt.Fprintln(os.Stderr, "Error: no tasks found")
		os.Exit(1)
	}

	selectedCmd, err := tui.Start(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if selectedCmd == "" {
		return
	}

	selectedCmd, err = formatter.Interpolate(selectedCmd, cfg.Formatters)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if err := runner.RunLive(selectedCmd); err != nil {
		os.Exit(1)
	}
}

func runDirect(cfg *config.Config, name string) {
	task, ok := cfg.Tasks[name]
	if !ok {
		fmt.Fprintf(os.Stderr, "Error: task %q not found\n", name)
		os.Exit(1)
	}

	fmt.Printf("Running %s...\n\n", name)
	
	cmdStr, err := formatter.Interpolate(task.Cmd, cfg.Formatters)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	result := runner.Run(cmdStr)

	if result.Err != nil {
		fmt.Printf("✗ Failed\n\n")
		if result.Stderr != "" {
			fmt.Println(result.Stderr)
		}
		os.Exit(1)
	}

	dur := result.Duration.Round(time.Millisecond)
	fmt.Printf("✓ Success\n\nTask:     %s\nDuration: %s\n", name, dur)

	if out := result.Stdout; out != "" {
		fmt.Printf("\nOutput:\n%s", out)
	}
}

func runUpdate() {
	fmt.Println("Updating dev...")
	c := exec.Command("sh", "-c", `curl -fsSL https://raw.githubusercontent.com/JoaoPedr0Maciel/dev/main/install.sh | sh`)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error: update failed")
		os.Exit(1)
	}
}
