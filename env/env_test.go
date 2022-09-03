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
			result: DefaultPort,
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

func TestAppName(t *testing.T) {
	tests := []struct {
		name    string
		appName string
		expName string
	}{
		{
			name:    "test env AppName success",
			appName: "some app name",
			expName: "some app name",
		},
		{
			name:    "test env AppName trim success",
			appName: "  trim me please    ",
			expName: "trim me please",
		},
		{
			name:    "test env AppName give default",
			appName: " ",
			expName: DefaultAppName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("AppName", tt.appName)
			got := AppName()
			if got != tt.expName {
				t.Errorf("AppName() want: %s, got: %s", tt.expName, got)
			}
		})
	}
}

func TestAppEnv(t *testing.T) {
	tests := []struct {
		name   string
		appEnv string
		expEnv string
	}{
		{
			name:   "test env AppEnv success",
			appEnv: "env-name",
			expEnv: "env-name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("AppEnv", tt.appEnv)
			got := AppEnv()
			if got != tt.expEnv {
				t.Errorf("AppName() want: %s, got: %s", tt.expEnv, got)
			}
		})
	}
}
