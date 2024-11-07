package intrct

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

type testWriter struct{}

func (tw testWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func TestInteraction_PromptBool(t *testing.T) {
	type fields struct {
		Stdin           io.Reader
		Stdout          io.Writer
		UserInputPrefix string
	}
	type args struct {
		question string
		def      bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:   "users answer: y",
			fields: fields{Stdin: strings.NewReader("y\n"), Stdout: testWriter{}},
			args:   args{},
			want:   true,
		},
		{
			name:   "users answer: yEs",
			fields: fields{Stdin: strings.NewReader("yEs\n"), Stdout: testWriter{}},
			args:   args{},
			want:   true,
		},
		{
			name:   "users answer: n",
			fields: fields{Stdin: strings.NewReader("n\n"), Stdout: testWriter{}},
			args:   args{},
			want:   false,
		},
		{
			name:   "users answer: nO",
			fields: fields{Stdin: strings.NewReader("nO\n"), Stdout: testWriter{}},
			args:   args{},
			want:   false,
		},
		{
			name:   "users answer: wat?",
			fields: fields{Stdin: strings.NewReader("wat?\n"), Stdout: testWriter{}},
			args:   args{},
			want:   false,
		},
		{
			name:   "using default: true",
			fields: fields{Stdin: strings.NewReader("\n"), Stdout: testWriter{}},
			args:   args{def: true},
			want:   true,
		},
		{
			name:   "using default: false",
			fields: fields{Stdin: strings.NewReader("\n"), Stdout: testWriter{}},
			args:   args{def: false},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Interaction{
				Stdin:           tt.fields.Stdin,
				Stdout:          tt.fields.Stdout,
				UserInputPrefix: tt.fields.UserInputPrefix,
			}
			got, err := i.PromptBool(tt.args.question, tt.args.def)
			if (err != nil) != tt.wantErr {
				t.Errorf("PromptBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PromptBool() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteraction_PromptList(t *testing.T) {
	type fields struct {
		Stdin           io.Reader
		Stdout          io.Writer
		UserInputPrefix string
	}
	type args struct {
		question string
		opts     []string
		def      []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:   "single option",
			fields: fields{Stdin: strings.NewReader("2\n"), Stdout: testWriter{}},
			args:   args{opts: []string{"opt1", "opt2"}},
			want:   []string{"opt2"},
		},
		{
			name:   "several options",
			fields: fields{Stdin: strings.NewReader("1 3\n"), Stdout: testWriter{}},
			args:   args{opts: []string{"opt1", "opt2", "opt3"}},
			want:   []string{"opt1", "opt3"},
		},
		{
			name:   "using default",
			fields: fields{Stdin: strings.NewReader("\n"), Stdout: testWriter{}},
			args:   args{opts: []string{"opt1", "opt2"}, def: []string{"opt2"}},
			want:   []string{"opt2"},
		},
		{
			name:    "name incorrect option",
			fields:  fields{Stdin: strings.NewReader("3\n"), Stdout: testWriter{}},
			args:    args{opts: []string{"opt1", "opt2"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Interaction{
				Stdin:           tt.fields.Stdin,
				Stdout:          tt.fields.Stdout,
				UserInputPrefix: tt.fields.UserInputPrefix,
			}
			got, err := i.PromptList(tt.args.question, tt.args.opts, tt.args.def)
			if (err != nil) != tt.wantErr {
				t.Errorf("PromptList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PromptList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteraction_PromptString(t *testing.T) {
	type fields struct {
		Stdin           io.Reader
		Stdout          io.Writer
		UserInputPrefix string
	}
	type args struct {
		question string
		def      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "regular input",
			fields: fields{Stdin: strings.NewReader("test answer\n"), Stdout: testWriter{}},
			args:   args{"test question", ""},
			want:   "test answer",
		},
		{
			name:   "empty input",
			fields: fields{Stdin: strings.NewReader("\n"), Stdout: testWriter{}},
			args:   args{"test question", "default answer"},
			want:   "default answer",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Interaction{
				Stdin:           tt.fields.Stdin,
				Stdout:          tt.fields.Stdout,
				UserInputPrefix: tt.fields.UserInputPrefix,
			}
			got, err := i.PromptString(tt.args.question, tt.args.def)
			if (err != nil) != tt.wantErr {
				t.Errorf("PromptString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PromptString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseOptionsIndices(t *testing.T) {
	type args struct {
		opts string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "single option",
			args: args{"2"},
			want: []int{1},
		},
		{
			name: "multiple options",
			args: args{"1 2 4"},
			want: []int{0, 1, 3},
		},
		{
			name:    "not a number",
			args:    args{"1 a"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseOptionsIndices(tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseOptionsIndices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOptionsIndices() got = %v, want %v", got, tt.want)
			}
		})
	}
}
