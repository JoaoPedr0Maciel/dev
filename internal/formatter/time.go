package formatter

import (
	"fmt"
	"strings"
	"time"
)

func init() {
	Register("time", "now", timeNowImpl)
}

var formatReplacer = strings.NewReplacer(
	"YYYY", "2006",
	"YY", "06",
	"MM", "01",
	"DD", "02",
	"hh", "15",
	"mm", "04",
	"ss", "05",
)

func timeNowImpl(args string) (string, error) {
	if args == "" {
		return "", fmt.Errorf("empty argument for time.now")
	}

	// Remove possible quotes
	args = strings.Trim(args, `"'`)
	goFormat := formatReplacer.Replace(args)

	return time.Now().Format(goFormat), nil
}
