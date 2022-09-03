package env

import (
	"os"
	"testing"
)

func TestPort(t *testing.T) {
	tests := []struct {
		name   string
		port   string
		result string
	}{
		{
			name:   "test env Port success",
			port:   ":9000",
			result: ":9000",
		},
		{
			name:   "test env Port give default",
			port:   ":900x",
			result: ":5000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("Port", tt.port)
			got := Port()
			if got != tt.result {
				t.Errorf("Port() want: %s, got: %s", tt.result, got)
			}
		})
	}
}
