package tui

import (
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/JoaoPedr0Maciel/dev/internal/config"
)

// ── styles ───────────────────────────────────────────────────────────────────

var (
	borderClr   = lipgloss.Color("238")
	titleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("75"))
	cursorStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212"))
	normalStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("252"))
	dimStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
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

type taskItem struct {
	name string
	task config.Task
}

type model struct {
	tasks       []taskItem
	cursor      int
	selectedCmd string
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
			m.selectedCmd = m.tasks[m.cursor].task.Cmd
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
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

// Start launches the TUI and returns the command selected by the user, or an
// empty string if the user quit without selecting anything.
func Start(cfg *config.Config) (string, error) {
	p := tea.NewProgram(newModel(cfg), tea.WithAltScreen())
	final, err := p.Run()
	if err != nil {
		return "", err
	}
	return final.(model).selectedCmd, nil
}
