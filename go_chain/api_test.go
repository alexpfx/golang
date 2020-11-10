package go_chain

import (
	"reflect"
	"testing"
)

func Test_replaceAll(t *testing.T) {
	type args struct {
		sourceStr   string
		replaceStrs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{
				sourceStr:   "{{ baseUrl }}/sibews/rest/requerimento/incluir/bi/inicial",
				replaceStrs: []string{"https://localhost:7001"},
			},
			want: "https://localhost:7001/sibews/rest/requerimento/incluir/bi/inicial",
		},
		{
			name: "t2",
			args: args{
				sourceStr:   "{{baseUrl}}:{{port}}/sibews/rest/requerimento/incluir/bi/inicial",
				replaceStrs: []string{"https://localhost", "7001"},
			},
			want: "https://localhost:7001/sibews/rest/requerimento/incluir/bi/inicial",
		},
		{
			name: "t3",
			args: args{
				sourceStr:   "{{baseUrl}}:{{port}}/sibews/rest/requerimento/incluir/bi/{{inicial}}",
				replaceStrs: []string{"https://localhost", "7001"},
			},
			want: "https://localhost:7001/sibews/rest/requerimento/incluir/bi/{{inicial}}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceAll(tt.args.sourceStr, tt.args.replaceStrs); got != tt.want {
				t.Errorf("replaceAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *Request
		wantErr bool
	}{
		{
			name: "t1",
			args: args{
				filePath: "go_chain_test/data/simplificado.toml",
			},
			want: &Request{
				Method: "post",
				Input:  []string{"cpf", "nome"},
				Json:   `{"cpf": 123,}`,
				EndpointReplaces: map[string][]string{
					"localhost": {"localhost", "7001"},
					"blue":      {"192.168.1.160", "7001"},
				},
				Endpoint: "{{ baseUrl }}:{{ port }}/sibews/rest/requerimento/incluir/bi/inicial",
				Output:   []string{"nb"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseTomlFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseTomlFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTomlFile() \ngot: %v\nwant: %v", got, tt.want)
			}
		})
	}
}

func Test_replaceInput(t *testing.T) {
	type args struct {
		json  string
		input map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "t1",
			args:    args{
				json:  `{"cpf":"100200300-40"}`,
				input: map[string]string{"cpf": "111111111-20"},
			},
			want:    `{"cpf":"111111111-20"}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := replaceInput(tt.args.json, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("replaceInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("replaceInput() got = %v, want %v", got, tt.want)
			}
		})
	}
}
