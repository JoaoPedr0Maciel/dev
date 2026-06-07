package formatter

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	Register("env", "get", envGetImpl)
}

func envGetImpl(args string) (string, error) {
	key := strings.Trim(args, `"'`)
	if key == "" {
		return "", fmt.Errorf("empty variable name for env.get")
	}

	_ = godotenv.Load()

	return os.Getenv(key), nil
}
