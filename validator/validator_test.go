package validator

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want Validator
	}{
		{
			name: "case normal",
			want: &validatorImpl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidator_Missing(t *testing.T) {
	type fields struct {
		missing []string
	}
	type args struct {
		param string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "case first",
			fields: fields{
				missing: nil,
			},
			args: args{
				param: "id",
			},
			want: []string{
				"id",
			},
		},
		{
			name: "case normal",
			fields: fields{
				missing: []string{"id"},
			},
			args: args{
				param: "name",
			},
			want: []string{
				"id", "name",
			},
		},
		{
			name: "case empty first",
			fields: fields{
				missing: nil,
			},
			args: args{
				param: "",
			},
			want: []string{
				"",
			},
		},
		{
			name: "case empty",
			fields: fields{
				missing: []string{""},
			},
			args: args{
				param: "",
			},
			want: []string{
				"", "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &validatorImpl{
				missing: tt.fields.missing,
			}
			v.Missing(tt.args.param)
			assert.Equal(t, tt.want, v.missing)
		})
	}
}

func TestValidator_Message(t *testing.T) {
	type fields struct {
		errs []string
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "case first",
			fields: fields{
				errs: nil,
			},
			args: args{
				message: "id",
			},
			want: []string{
				"id",
			},
		},
		{
			name: "case normal",
			fields: fields{
				errs: []string{"id"},
			},
			args: args{
				message: "name",
			},
			want: []string{
				"id", "name",
			},
		},
		{
			name: "case empty first",
			fields: fields{
				errs: nil,
			},
			args: args{
				message: "",
			},
			want: []string{
				"",
			},
		},
		{
			name: "case empty",
			fields: fields{
				errs: []string{""},
			},
			args: args{
				message: "",
			},
			want: []string{
				"", "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &validatorImpl{
				errs: tt.fields.errs,
			}
			v.Message(tt.args.message)
			assert.Equal(t, tt.want, v.errs)
		})
	}
}

func TestValidator_Error(t *testing.T) {
	type fields struct {
		missing []string
		errs    []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantMsg string
		wantErr bool
	}{
		{
			name: "case nil",
			fields: fields{
				missing: nil,
				errs:    nil,
			},
			wantErr: false,
		},
		{
			name: "case empty",
			fields: fields{
				missing: []string{},
				errs:    []string{},
			},
			wantErr: false,
		},
		{
			name: "case missing",
			fields: fields{
				missing: []string{"id"},
				errs:    nil,
			},
			wantMsg: "Missing Param(s) [id]",
			wantErr: true,
		},
		{
			name: "case missing multiple",
			fields: fields{
				missing: []string{"id", "name"},
				errs:    nil,
			},
			wantMsg: "Missing Param(s) [id, name]",
			wantErr: true,
		},
		{
			name: "case message",
			fields: fields{
				missing: nil,
				errs:    []string{"Some Error"},
			},
			wantMsg: "Some Error",
			wantErr: true,
		},
		{
			name: "case message multiple",
			fields: fields{
				missing: nil,
				errs:    []string{"Some Error", "Other Error"},
			},
			wantMsg: "Some Error; Other Error",
			wantErr: true,
		},
		{
			name: "case combined multiple",
			fields: fields{
				missing: []string{"id", "name"},
				errs:    []string{"Some Error", "Other Error"},
			},
			wantMsg: "Missing Param(s) [id, name]; Some Error; Other Error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &validatorImpl{
				missing: tt.fields.missing,
				errs:    tt.fields.errs,
			}
			err := v.Error()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tt.wantMsg, err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
