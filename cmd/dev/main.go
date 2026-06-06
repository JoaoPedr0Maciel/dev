package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/JoaoPedr0Maciel/dev/internal/config"
	"github.com/JoaoPedr0Maciel/dev/internal/runner"
	"github.com/JoaoPedr0Maciel/dev/internal/tui"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "path", "dev.yaml", "path to the yaml configuration file")
	flag.Parse()

	cfg, err := config.Load(configPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	if len(cfg.Tasks) == 0 {
		fmt.Fprintln(os.Stderr, "Error: no tasks found")
		os.Exit(1)
	}

	// Direct execution: dev <task-name>
	args := flag.Args()
	if len(args) > 0 {
		runDirect(cfg, args[0])
		return
	}

	// TUI mode
	if err := tui.Start(cfg); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
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
	result := runner.Run(task.Cmd)

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
