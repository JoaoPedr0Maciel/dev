package formatter

import (
	"fmt"
	"regexp"
)

// ActionFunc defines the signature for any macro function
type ActionFunc func(args string) (string, error)

// registry stores: Namespace -> Action -> Function
var registry = make(map[string]map[string]ActionFunc)

// Register exposes the ability for formatters to register themselves
func Register(namespace, action string, fn ActionFunc) {
	if registry[namespace] == nil {
		registry[namespace] = make(map[string]ActionFunc)
	}
	registry[namespace][action] = fn
}

var macroRegex = regexp.MustCompile(`@([a-zA-Z0-9_]+)\.([a-zA-Z0-9_]+)\(([^)]*)\)`)

// Interpolate parses the command string, validates enabled formatters, and executes macros
func Interpolate(cmd string, enabled []string) (string, error) {
	var firstErr error

	enabledMap := make(map[string]bool, len(enabled))
	for _, f := range enabled {
		enabledMap[f] = true
	}

	result := macroRegex.ReplaceAllStringFunc(cmd, func(match string) string {
		if firstErr != nil {
			return match
		}

		parts := macroRegex.FindStringSubmatch(match)
		if len(parts) != 4 {
			return match
		}

		ns, act, args := parts[1], parts[2], parts[3]
		if !enabledMap[ns] {
			firstErr = fmt.Errorf("formatter %q is not enabled in dev.yaml", ns)
			return match
		}

		actions, exists := registry[ns]
		if !exists {
			firstErr = fmt.Errorf("formatter %q not found", ns)
			return match
		}

		fn, exists := actions[act]
		if !exists {
			firstErr = fmt.Errorf("action %q not found in formatter %q", act, ns)
			return match
		}

		val, err := fn(args)
		if err != nil {
			firstErr = fmt.Errorf("evaluating @%s.%s: %w", ns, act, err)
			return match
		}

		return val
	})

	return result, firstErr
}



