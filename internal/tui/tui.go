package tui

import (
	"fmt"
	"sort"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/JoaoPedr0Maciel/dev/internal/config"
	"github.com/JoaoPedr0Maciel/dev/internal/runner"
)

// ── styles ───────────────────────────────────────────────────────────────────

var (
	borderClr   = lipgloss.Color("238")
	titleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("75"))
	cursorStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212"))
	normalStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	dimStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	successStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("82"))
	failStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("196"))
	hintStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("238"))
	bdr         = lipgloss.NewStyle().Foreground(borderClr)
)

// ── box helpers ───────────────────────────────────────────────────────────────

const innerW = 38

func boxTop() string { return bdr.Render("┌" + strings.Repeat("─", innerW+2) + "┐") }
func boxBot() string { return bdr.Render("└" + strings.Repeat("─", innerW+2) + "┘") }
func boxSep() string { return bdr.Render("├" + strings.Repeat("─", innerW+2) + "┤") }

func boxRow(content string) string {
	vis := lipgloss.Width(content)
	pad := innerW - vis
	if pad < 0 {
		pad = 0
	}
	return bdr.Render("│") + " " + content + strings.Repeat(" ", pad) + " " + bdr.Render("│")
}

// ── model ─────────────────────────────────────────────────────────────────────

type viewState int

const (
	viewList viewState = iota
	viewRunning
	viewResult
)

type taskItem struct {
	name string
	task config.Task
}

type taskDoneMsg struct{ result runner.Result }

type model struct {
	tasks  []taskItem
	cursor int
	view   viewState
	active taskItem
	result *runner.Result
}

func newModel(cfg *config.Config) model {
	items := make([]taskItem, 0, len(cfg.Tasks))
	for name, task := range cfg.Tasks {
		items = append(items, taskItem{name: name, task: task})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].name < items[j].name
	})
	return model{tasks: items}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.view {
		case viewList:
			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.tasks)-1 {
					m.cursor++
				}
			case "enter":
				if len(m.tasks) == 0 {
					break
				}
				m.active = m.tasks[m.cursor]
				m.view = viewRunning
				return m, execTask(m.active.task.Cmd)
			}

		case viewRunning:
			// block input while running

		case viewResult:
			m.view = viewList
			m.result = nil
		}

	case taskDoneMsg:
		r := msg.result
		m.result = &r
		m.view = viewResult
	}

	return m, nil
}

func (m model) View() string {
	switch m.view {
	case viewList:
		return m.listView()
	case viewRunning:
		return m.runningView()
	case viewResult:
		return m.resultView()
	}
	return ""
}

// ── views ─────────────────────────────────────────────────────────────────────

func (m model) listView() string {
	var b strings.Builder

	b.WriteString(boxTop() + "\n")
	b.WriteString(boxRow(titleStyle.Render("Dev")) + "\n")
	b.WriteString(boxSep() + "\n")

	for i, item := range m.tasks {
		var row string
		if i == m.cursor {
			row = cursorStyle.Render("▶ " + item.name)
		} else {
			row = "  " + normalStyle.Render(item.name)
		}
		b.WriteString(boxRow(row) + "\n")
	}

	b.WriteString(boxSep() + "\n")
	b.WriteString(boxRow(dimStyle.Render("Description:")) + "\n")

	desc := ""
	if len(m.tasks) > 0 {
		desc = m.tasks[m.cursor].task.Description
	}
	b.WriteString(boxRow(normalStyle.Render(desc)) + "\n")
	b.WriteString(boxBot() + "\n")
	b.WriteString("\n" + hintStyle.Render("  ↑/↓ navigate  •  enter run  •  q quit"))

	return b.String()
}

func (m model) runningView() string {
	msg := "  Running " + cursorStyle.Render(m.active.name) + "..."

	var b strings.Builder
	b.WriteString(boxTop() + "\n")
	b.WriteString(boxRow("") + "\n")
	b.WriteString(boxRow(msg) + "\n")
	b.WriteString(boxRow("") + "\n")
	b.WriteString(boxBot() + "\n")

	return b.String()
}

func (m model) resultView() string {
	if m.result == nil {
		return ""
	}

	var b strings.Builder
	b.WriteString(boxTop() + "\n")
	b.WriteString(boxRow("") + "\n")

	if m.result.Err == nil {
		b.WriteString(boxRow(successStyle.Render("✓ Success")) + "\n")
		b.WriteString(boxRow("") + "\n")
		b.WriteString(boxRow(fmt.Sprintf("Task:     %s", normalStyle.Render(m.active.name))) + "\n")
		dur := m.result.Duration.Round(time.Millisecond).String()
		b.WriteString(boxRow(fmt.Sprintf("Duration: %s", normalStyle.Render(dur))) + "\n")

		if out := strings.TrimSpace(m.result.Stdout); out != "" {
			b.WriteString(boxSep() + "\n")
			for _, line := range strings.Split(out, "\n") {
				b.WriteString(boxRow(dimStyle.Render(line)) + "\n")
			}
		}
	} else {
		b.WriteString(boxRow(failStyle.Render("✗ Failed")) + "\n")
		b.WriteString(boxRow("") + "\n")
		b.WriteString(boxRow(fmt.Sprintf("Task: %s", normalStyle.Render(m.active.name))) + "\n")

		if errOut := strings.TrimSpace(m.result.Stderr); errOut != "" {
			b.WriteString(boxSep() + "\n")
			for _, line := range strings.Split(errOut, "\n") {
				b.WriteString(boxRow(failStyle.Render(line)) + "\n")
			}
		}
	}

	b.WriteString(boxRow("") + "\n")
	b.WriteString(boxSep() + "\n")
	b.WriteString(boxRow(hintStyle.Render("Press any key to return")) + "\n")
	b.WriteString(boxBot() + "\n")

	return b.String()
}

// ── commands ──────────────────────────────────────────────────────────────────

func execTask(cmd string) tea.Cmd {
	return func() tea.Msg {
		return taskDoneMsg{result: runner.Run(cmd)}
	}
}

// Start launches the TUI program.
func Start(cfg *config.Config) error {
	p := tea.NewProgram(newModel(cfg), tea.WithAltScreen())
	_, err := p.Run()
	return err
}
