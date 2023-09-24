package errors

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		status int
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "case normal",
			args: args{
				status: 400,
			},
			want: &Error{
				status: 400,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.status); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_WithCode(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "case normal",
			args: args{
				code: "INVALID_REQUEST",
			},
			want: &Error{
				code: "INVALID_REQUEST",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{}
			if got := e.WithCode(tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_WithMessage(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "case normal",
			args: args{
				message: "Invalid Request",
			},
			want: &Error{
				message: "Invalid Request",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{}
			if got := e.WithMessage(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Code(t *testing.T) {
	type fields struct {
		code string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "case normal",
			fields: fields{
				code: "INVALID_REQUEST",
			},
			want: "INVALID_REQUEST",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				code: tt.fields.code,
			}
			if got := e.Code(); got != tt.want {
				t.Errorf("Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Message(t *testing.T) {
	type fields struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "case normal",
			fields: fields{
				message: "Invalid Request",
			},
			want: "Invalid Request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				message: tt.fields.message,
			}
			if got := e.Message(); got != tt.want {
				t.Errorf("Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Status(t *testing.T) {
	type fields struct {
		status int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "case normal",
			fields: fields{
				status: 400,
			},
			want: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				status: tt.fields.status,
			}
			if got := e.Status(); got != tt.want {
				t.Errorf("Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	type fields struct {
		code    string
		message string
		status  int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "case message",
			fields: fields{
				code:    "INVALID_REQUEST",
				message: "Invalid Request",
				status:  400,
			},
			want: "Invalid Request",
		},
		{
			name: "case no message",
			fields: fields{
				status: 400,
			},
			want: "Bad Request",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				code:    tt.fields.code,
				message: tt.fields.message,
				status:  tt.fields.status,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_ShouldRetry(t *testing.T) {
	type fields struct {
		status int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "case retry",
			fields: fields{
				status: 500,
			},
			want: true,
		},
		{
			name: "case not retry",
			fields: fields{
				status: 400,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				status: tt.fields.status,
			}
			if got := e.IsRetry(); got != tt.want {
				t.Errorf("IsShouldRetry() = %v, want %v", got, tt.want)
			}
		})
	}
}
