package config

import (
	"os"
	"reflect"
	"testing"
)

func TestReadServerConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    ServerEnvConfig
		wantErr bool
	}{
		{
			name: "ReadServerConfig Happy Path",
			want: ServerEnvConfig{
				Port:       8080,
				Hostname:   "localhost",
				AutoTLS:    false,
				LogLevel:   "info",
				LogFormat:  "text",
				DBLocation: "data/",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadServerConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadServerConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadServerConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadServerBadEnvVar(t *testing.T) {
	// Set an invalid env var
	err := os.Setenv("TINYMONITOR_PORT", "not an int")
	if err != nil {
		t.Errorf("Error setting env var: %v", err)
	}
	_, err = ReadServerConfig()
	if err == nil {
		t.Errorf("Expected error when reading invalid env var, got nil")
	}
	return
}
