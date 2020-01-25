package command

import (
	"testing"

	"github.com/urfave/cli"
)

func TestCmdInit(t *testing.T) {
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CmdInit(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("CmdInit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
