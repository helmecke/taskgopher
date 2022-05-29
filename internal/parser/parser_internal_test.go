package parser

import (
	"testing"
)

func Test_isFilterCommand(t *testing.T) {
	t.Parallel()
	type args struct {
		cmd string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"false", args{cmd: "bla"}, false},
		{"complete", args{cmd: "complete"}, true},
		{"delete", args{cmd: "delete"}, true},
		{"list", args{cmd: "list"}, true},
		{"modify", args{cmd: "modify"}, true},
		{"mod", args{cmd: "mod"}, true},
		{"show", args{cmd: "show"}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isFilterCommand(tt.args.cmd); got != tt.want {
				t.Errorf("isFilterCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isModifyCommand(t *testing.T) {
	t.Parallel()
	type args struct {
		cmd string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"false", args{cmd: "bla"}, false},
		{"add", args{cmd: "add"}, true},
		{"complete", args{cmd: "complete"}, true},
		{"modify", args{cmd: "modify"}, true},
		{"mod", args{cmd: "mod"}, true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isModifyCommand(tt.args.cmd); got != tt.want {
				t.Errorf("isModifyCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
