package tests

import (
	"os"
	"testing"
	"time"

	"github.com/JoaoPedr0Maciel/dev/internal/formatter"
)

func TestInterpolate(t *testing.T) {
	os.Setenv("PORT", "8080")
	defer os.Unsetenv("PORT")

	tests := []struct {
		cmd     string
		enabled []string
		want    string
		wantErr bool
	}{
		{"echo @time.now('YYYY')", []string{"time"}, "echo " + time.Now().Format("2006"), false},
		{"echo @env.get('PORT')", []string{"env"}, "echo 8080", false},
		{"echo @time.now('YYYY')", []string{}, "", true}, // not enabled
	}

	for _, tt := range tests {
		got, err := formatter.Interpolate(tt.cmd, tt.enabled)
		if (err != nil) != tt.wantErr {
			t.Fatalf("Interpolate(%q) error = %v, wantErr = %v", tt.cmd, err, tt.wantErr)
		}
		if !tt.wantErr && got != tt.want {
			t.Errorf("Interpolate(%q) = %q, want %q", tt.cmd, got, tt.want)
		}
	}
}

