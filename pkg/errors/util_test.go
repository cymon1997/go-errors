package errors

import (
	"errors"
	"testing"
)

func TestIs(t *testing.T) {
	err := errors.New("some errors")
	errCustom := New(500).WithMessage("some errors")
	type args struct {
		err    error
		target error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case not same",
			args: args{
				err:    errors.New("other errors"),
				target: err,
			},
			want: false,
		},
		{
			name: "case not same (custom)",
			args: args{
				err:    New(400).WithMessage("other errors"),
				target: errCustom,
			},
			want: false,
		},
		{
			name: "case similar",
			args: args{
				err:    errors.New("some errors"),
				target: err,
			},
			want: true,
		},
		{
			name: "case similar (diff type)",
			args: args{
				err:    err,
				target: errCustom,
			},
			want: true,
		},
		{
			name: "case similar (custom)",
			args: args{
				err:    New(500).WithMessage("some errors"),
				target: errCustom,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is(tt.args.err, tt.args.target); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMessage(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case basic error",
			args: args{
				err: errors.New("basic error"),
			},
			want: "basic error",
		},
		{
			name: "case custom error",
			args: args{
				err: New(500).WithMessage("custom error"),
			},
			want: "custom error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMessage(tt.args.err); got != tt.want {
				t.Errorf("GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStatus(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case basic error",
			args: args{
				err: errors.New("basic error"),
			},
			want: 500,
		},
		{
			name: "case custom error",
			args: args{
				err: New(404).WithMessage("Not Found"),
			},
			want: 404,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStatus(tt.args.err); got != tt.want {
				t.Errorf("GetStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsShouldRetry(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case client error (4xx)",
			args: args{
				err: New(422).WithMessage("Invalid Request"),
			},
			want: false,
		},
		{
			name: "case server error (500)",
			args: args{
				err: New(500).WithMessage("Internal Server"),
			},
			want: true,
		},
		{
			name: "case unknown error",
			args: args{
				err: errors.New("some error"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsShouldRetry(tt.args.err); got != tt.want {
				t.Errorf("IsShouldRetry() = %v, want %v", got, tt.want)
			}
		})
	}
}
