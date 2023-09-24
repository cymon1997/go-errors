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

func Test_validatorImpl_Required(t *testing.T) {
	//var empty *string
	type fields struct {
		required []string
	}
	type args struct {
		val  interface{}
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "case interface",
			args: args{
				val:  nil,
				name: "interface",
			},
			want: []string{"interface (<nil>)"},
		},
		{
			name: "case pointer",
			args: args{
				val: func() *string {
					return nil
				}(),
				name: "pointer",
			},
			want: []string{"pointer (<nil>)"},
		},
		{
			name: "case string",
			args: args{
				val:  "",
				name: "string",
			},
			want: []string{"string ()"},
		},
		{
			name: "case number",
			args: args{
				val:  0,
				name: "number",
			},
			want: []string{"number (0)"},
		},
		{
			name: "case array str",
			args: args{
				val:  [0]string{},
				name: "array",
			},
			want: []string{"array ([])"},
		},

		{
			name: "case slice",
			args: args{
				val: func() []string {
					return nil
				}(),
				name: "slice",
			},
			want: []string{"slice ([])"},
		},
		{
			name: "case slice str",
			args: args{
				val:  []string{},
				name: "slice",
			},
			want: []string{"slice ([])"},
		},
		{
			name: "case slice int",
			args: args{
				val:  []int{},
				name: "slice",
			},
			want: []string{"slice ([])"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &validatorImpl{
				required: tt.fields.required,
			}
			v.Required(tt.args.val, tt.args.name)
			assert.Equal(t, tt.want, v.required)
		})
	}
}

func Test_validatorImpl_Positive(t *testing.T) {
	type fields struct {
		positives []string
	}
	type args struct {
		val  int
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "case positive",
			args: args{
				val:  10,
				name: "num",
			},
			want: nil,
		},
		{
			name: "case negative",
			args: args{
				val:  -1,
				name: "num",
			},
			want: []string{"num (-1)"},
		},
		{
			name: "case zero",
			args: args{
				val:  0,
				name: "num",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &validatorImpl{
				positives: tt.fields.positives,
			}
			v.Positive(tt.args.val, tt.args.name)
			assert.Equal(t, tt.want, v.positives)
		})
	}
}

func Test_validatorImpl_Negative(t *testing.T) {
	type fields struct {
		negatives []string
	}
	type args struct {
		val  int
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "case positive",
			args: args{
				val:  10,
				name: "num",
			},
			want: []string{"num (10)"},
		},
		{
			name: "case negative",
			args: args{
				val:  -1,
				name: "num",
			},
			want: nil,
		},
		{
			name: "case zero",
			args: args{
				val:  0,
				name: "num",
			},
			want: []string{"num (0)"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &validatorImpl{
				negatives: tt.fields.negatives,
			}
			v.Negative(tt.args.val, tt.args.name)
			assert.Equal(t, tt.want, v.negatives)
		})
	}
}

func Test_validatorImpl_Custom(t *testing.T) {
	type fields struct {
		errs []string
	}
	type args struct {
		val  string
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "case normal",
			args: args{
				val:  "val",
				name: "field",
			},
			want: []string{"field (val)"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &validatorImpl{
				errs: tt.fields.errs,
			}
			v.Custom(tt.args.val, tt.args.name)
			assert.Equal(t, tt.want, v.errs)
		})
	}
}

func TestValidator_Error(t *testing.T) {
	type fields struct {
		required  []string
		positives []string
		negatives []string
		errs      []string
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
				required:  nil,
				positives: nil,
				negatives: nil,
				errs:      nil,
			},
			wantErr: false,
		},
		{
			name: "case empty",
			fields: fields{
				required:  []string{},
				positives: []string{},
				negatives: []string{},
				errs:      []string{},
			},
			wantErr: false,
		},
		{
			name: "case required",
			fields: fields{
				required:  []string{"id"},
				positives: nil,
				negatives: nil,
				errs:      nil,
			},
			wantMsg: "Required Param(s) [id]",
			wantErr: true,
		},
		{
			name: "case required multiple",
			fields: fields{
				required:  []string{"id", "name"},
				positives: nil,
				negatives: nil,
				errs:      nil,
			},
			wantMsg: "Required Param(s) [id, name]",
			wantErr: true,
		},
		{
			name: "case positive",
			fields: fields{
				required:  nil,
				positives: []string{"id"},
				negatives: nil,
				errs:      nil,
			},
			wantMsg: "Positive Param(s) [id]",
			wantErr: true,
		},
		{
			name: "case positive multiple",
			fields: fields{
				required:  nil,
				positives: []string{"id", "name"},
				negatives: nil,
				errs:      nil,
			},
			wantMsg: "Positive Param(s) [id, name]",
			wantErr: true,
		},
		{
			name: "case negative",
			fields: fields{
				required:  nil,
				positives: nil,
				negatives: []string{"id"},
				errs:      nil,
			},
			wantMsg: "Negative Param(s) [id]",
			wantErr: true,
		},
		{
			name: "case negative multiple",
			fields: fields{
				required:  nil,
				positives: nil,
				negatives: []string{"id", "name"},
				errs:      nil,
			},
			wantMsg: "Negative Param(s) [id, name]",
			wantErr: true,
		},
		{
			name: "case message",
			fields: fields{
				required:  nil,
				positives: nil,
				negatives: nil,
				errs:      []string{"Some Error"},
			},
			wantMsg: "Some Error",
			wantErr: true,
		},
		{
			name: "case message multiple",
			fields: fields{
				required:  nil,
				positives: nil,
				negatives: nil,
				errs:      []string{"Some Error", "Other Error"},
			},
			wantMsg: "Some Error; Other Error",
			wantErr: true,
		},
		{
			name: "case combined multiple",
			fields: fields{
				required:  []string{"id", "name"},
				positives: nil,
				negatives: nil,
				errs:      []string{"Some Error", "Other Error"},
			},
			wantMsg: "Required Param(s) [id, name]; Some Error; Other Error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &validatorImpl{
				required:  tt.fields.required,
				positives: tt.fields.positives,
				negatives: tt.fields.negatives,
				errs:      tt.fields.errs,
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
