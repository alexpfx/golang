package output

import "testing"

func Test_filter(t *testing.T) {
	type args struct {
		jsonInput string
		filter    []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{
				jsonInput: `{"cpfMassa": "87333851303","nitMassa": "26814418949","nomePfMassa":"FULANO IEVHSMRC HSCBUVJ"}`,
				filter:    []string{"cpfMassa cpf", "nitMassa nit", "nomePfMassa nomeTitular"},
			},
			want: `{"cpf":"87333851303","nit":"26814418949","nomeTitular":"FULANO IEVHSMRC HSCBUVJ"}`,
		},
		{
			name: "t2",
			args: args{
				jsonInput: `{"cpfMassa": "87333851303","nitMassa": "26814418949","nomePfMassa":"FULANO IEVHSMRC HSCBUVJ"}`,
				filter:    []string{"cpfMassa cpf", "nitMassa", "nomePfMassa nomeTitular"},
			},
			want: `{"cpf":"87333851303","nitMassa":"26814418949","nomeTitular":"FULANO IEVHSMRC HSCBUVJ"}`,
		},
		{
			name: "t3",
			args: args{
				jsonInput: `{"cpfMassa": "87333851303","nitMassa": "26814418949","nomePfMassa":"FULANO IEVHSMRC HSCBUVJ"}`,
				filter:    []string{"nitMassa", "nomePfMassa nomeTitular"},
			},
			want: `{"nitMassa":"26814418949","nomeTitular":"FULANO IEVHSMRC HSCBUVJ"}`,
		},

		{
			name: "t4",
			args: args{
				jsonInput: `{"cpfMassa": "87333851303","nitMassa": "26814418949","nomePfMassa":"FULANO IEVHSMRC HSCBUVJ"}`,
				filter:    []string{"xnitMassa", "xnomePfMassa nomeTitular"},
			},
			want: `{}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.jsonInput, tt.args.filter); got != tt.want {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
