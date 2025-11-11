package sub_test

import (
	"testing"

	"github.com/okieoth/pvault/cmd/sub"
)

func TestCheckForOutput(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		output    string
		overwrite bool
		stdout    bool
		wantStr   string
		wantOk    bool
	}{
		{
			name:      "explicit output (happy)",
			input:     "input.txt",
			output:    "out.txt",
			overwrite: false,
			stdout:    false,
			wantStr:   "out.txt",
			wantOk:    true,
		},
		{
			name:      "overwrite allowed (happy)",
			input:     "input.txt",
			output:    "",
			overwrite: true,
			stdout:    false,
			wantStr:   "input.txt",
			wantOk:    true,
		},
		{
			name:      "write to stdout (happy)",
			input:     "input.txt",
			output:    "",
			overwrite: false,
			stdout:    true,
			wantStr:   "stdout",
			wantOk:    true,
		},
		{
			name:      "missing all flags (unhappy)",
			input:     "input.txt",
			output:    "",
			overwrite: false,
			stdout:    false,
			wantStr:   "",
			wantOk:    false,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStr, gotOk := sub.CheckForOutput(tt.input, tt.output, tt.overwrite, tt.stdout)

			if gotStr != tt.wantStr {
				t.Errorf("(Test: %d) expected string %q, got %q", i, tt.wantStr, gotStr)
			}
			if gotOk != tt.wantOk {
				t.Errorf("(Test: %d) expected ok=%v, got %v", i, tt.wantOk, gotOk)
			}
		})
	}
}
