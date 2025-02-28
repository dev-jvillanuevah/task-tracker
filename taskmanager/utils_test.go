package taskmanager

import (
	"testing"

	"github.com/dev-jvillanuevah/task-tracker/common"
	"github.com/stretchr/testify/assert"
)

func TestParseUserInput(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		want       *common.UserCommand
		wantErr    bool
		errMessage string
	}{
		{
			name:  "valid input - add",
			input: `add "Buy groceries"`,
			want: &common.UserCommand{
				Command:     common.CommandAdd,
				Input1:      "Buy groceries",
				Description: nil,
			},
			wantErr:    false,
			errMessage: "",
		},
		{
			name:  "valid input - list",
			input: "list",
			want: &common.UserCommand{
				Command: common.CommandList,
			},
			wantErr:    false,
			errMessage: "",
		},
		{
			name:  "valid input - exit",
			input: "exit",
			want: &common.UserCommand{
				Command: common.CommandExit,
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
