package task_test

import (
	"testing"

	"github.com/helmecke/taskgopher/internal/parser"
	"github.com/helmecke/taskgopher/internal/task"
)

func TestItem_Matches(t *testing.T) {
	t.Parallel()
	type fields struct {
		Filter *parser.Filter
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"true", fields{Filter: &parser.Filter{IDs: []int{0}}}, true},
		{"false", fields{Filter: &parser.Filter{IDs: []int{3}}}, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := &task.Item{}
			if got := i.Matches(tt.fields.Filter); got != tt.want {
				t.Errorf("Item.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItem_HasTag(t *testing.T) {
	t.Parallel()
	type fields struct {
		Tags []string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"true", fields{Tags: []string{"test"}}, true},
		{"false", fields{Tags: []string{}}, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := &task.Item{
				Tags: tt.fields.Tags,
			}
			if got := i.HasTag(); got != tt.want {
				t.Errorf("Item.HasTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItem_HasProject(t *testing.T) {
	t.Parallel()
	type fields struct {
		Project string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"true", fields{Project: "test"}, true},
		{"false", fields{Project: ""}, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := &task.Item{
				Project: tt.fields.Project,
			}
			if got := i.HasProject(); got != tt.want {
				t.Errorf("Item.HasProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItem_IsPending(t *testing.T) {
	t.Parallel()
	type fields struct {
		Status string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"true", fields{Status: "pending"}, true},
		{"false", fields{Status: "test"}, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := &task.Item{
				Status: tt.fields.Status,
			}
			if got := i.IsPending(); got != tt.want {
				t.Errorf("Item.IsPending() = %v, want %v", got, tt.want)
			}
		})
	}
}
