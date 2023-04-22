package tools

import (
	"ggt/config"
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantCf  config.Cfg
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				data: []byte(`
openai:
  token: abcd
`),
			},
			wantCf: config.Cfg{
				Open: config.OpenAI{
					Token: "abcd",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCf, err := Config(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCf, tt.wantCf) {
				t.Errorf("Config() = %v, want %v", gotCf, tt.wantCf)
			}
		})
	}
}
