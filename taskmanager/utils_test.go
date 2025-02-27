package taskmanager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUserInput(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		want       *UserCommand
		wantErr    bool
		errMessage string
	}{
		{
			name:  "valid input - add",
			input: `add "Buy groceries"`,
			want: &UserCommand{
				Command:     CommandAdd,
				Input1:      "Buy groceries",
				Description: nil,
			},
			wantErr:    false,
			errMessage: "",
		},
		{
			name:  "valid input - exit",
			input: "exit",
			want: &UserCommand{
				Command: CommandExit,
			},
			wantErr:    false,
			errMessage: "",
		},
		{
			name:       "invalid input - asd",
			input:      `asd "Buy groceries"`,
			want:       nil,
			wantErr:    true,
			errMessage: "invalid command: asd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUserInput(tt.input)
			if tt.wantErr {
				assert.NotNil(t, tt.errMessage)
				assert.ErrorContains(t, err, tt.errMessage)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
