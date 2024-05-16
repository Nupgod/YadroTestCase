package tests

import (
	"testing"
	"YadroTestCase/pkg/logger"
)

func TestParseFile(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		wantErr  bool
	}{
		{
			name:     "Valid file",
			filepath: "testdata/valid.txt",
			wantErr:  false,
		},
		{
			name:     "Invalid number of tables",
			filepath: "testdata/invalid_tables.txt",
			wantErr:  true,
		},
		{
			name:     "Invalid start and end times",
			filepath: "testdata/invalid_times.txt",
			wantErr:  true,
		},
		{
			name:     "Invalid hourly rate",
			filepath: "testdata/invalid_rate.txt",
			wantErr:  true,
		},
		{
			name:     "Invalid event format",
			filepath: "testdata/invalid_event.txt",
			wantErr:  true,
		},
		{
			name:     "Invalid event time",
			filepath: "testdata/invalid_event_time.txt",
			wantErr:  true,
		},
		{
			name:     "Invalid client name",
			filepath: "testdata/invalid_client.txt",
			wantErr:  true,
		},
		{
			name:     "Invalid table num",
			filepath: "testdata/invalid_table.txt",
			wantErr:  true,
		},
		{
			name:     "Non-existent file",
			filepath: "testdata/nonexistent.txt",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := logger.ParseFile(tt.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
		})
	}
}