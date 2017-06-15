package algorithms

import (
	"reflect"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	type fields struct {
		Elements []int
		Max      int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				Elements: tt.fields.Elements,
				Max:      tt.fields.Max,
			}
			if got := s.Pop(); got != tt.want {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStack(t *testing.T) {
	type args struct {
		opts []func(*Stack)
	}
	tests := []struct {
		name string
		args args
		want *Stack
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetMax(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want func(*Stack)
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetMax(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Push(t *testing.T) {
	type fields struct {
		Elements []int
		Max      int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				Elements: tt.fields.Elements,
				Max:      tt.fields.Max,
			}
			s.Push(tt.args.i)
		})
	}
}

func TestStack_IsEmpty(t *testing.T) {
	type fields struct {
		Elements []int
		Max      int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				Elements: tt.fields.Elements,
				Max:      tt.fields.Max,
			}
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("Stack.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_IsFull(t *testing.T) {
	type fields struct {
		Elements []int
		Max      int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				Elements: tt.fields.Elements,
				Max:      tt.fields.Max,
			}
			if got := s.IsFull(); got != tt.want {
				t.Errorf("Stack.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Peep(t *testing.T) {
	type fields struct {
		Elements []int
		Max      int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				Elements: tt.fields.Elements,
				Max:      tt.fields.Max,
			}
			if got := s.Peep(); got != tt.want {
				t.Errorf("Stack.Peep() = %v, want %v", got, tt.want)
			}
		})
	}
}
