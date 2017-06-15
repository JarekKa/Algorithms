package algorithms

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	type args struct {
		opts []func(*Queue)
	}
	tests := []struct {
		name string
		args args
		want *Queue
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueue(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueue() = %v, want %v", got, tt.want)
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
		want func(*Queue)
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

func TestQueue_Dequeue(t *testing.T) {
	type fields struct {
		Elements []int
		Max      int
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Queue{
				Elements: tt.fields.Elements,
				Max:      tt.fields.Max,
			}
			got, err := s.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.Dequeue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Queue.Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Peep(t *testing.T) {
	type fields struct {
		Elements []int
		Max      int
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Queue{
				Elements: tt.fields.Elements,
				Max:      tt.fields.Max,
			}
			got, err := s.Peep()
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.Peep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Queue.Peep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_IsFull(t *testing.T) {
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
			s := &Queue{
				Elements: tt.fields.Elements,
				Max:      tt.fields.Max,
			}
			if got := s.IsFull(); got != tt.want {
				t.Errorf("Queue.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}
