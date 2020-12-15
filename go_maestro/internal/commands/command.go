package commands

type Cmd struct {
	// Binary define o arquivo executável do comando. ex: client.sh
	Binary string

	// Name define o nome do comando a ser usado como título no menu
	Name string

	// Desc representa a descrição do comando
	Desc string

	// Args define os argumentos que serão passados para o comando
	Args []string

	// UserInput fixme
	UserInput map[string]string

	// CopyOutput indica se a saída do comando deve ser copiada para a área de transferência do sistema
	CopyOutput bool

	// FilterOutput permite filtrar uma saída json para mostrar apenas os campos desejados. pode-se também criar um
	// alias para o campo.
	//	ex:
	// para o json:
	//		{
	//			"nome": "jose",
	//			"cpfPessoa": "1111111111",
	//			"nasc": "01/01/1987",
	//		}
	//		FilterOutput = [] {"nome", "cpfPessoa cpf"}
	//o filtro acima mostraria a seguinte saída:
	//
	//		{
	//			"nome": "jose",
	//			"cpf": "1111111111"
	//		}
	FilterOutput []string

	// FormatOutput permite transformar a saída json em uma string utilizando as chaves e valores do json (consultar lib: github.com/tidwall/gjson)
	FormatOutput []string
	// OutputConverter função que converte a saída de um comando para que possa ser utilizada como argumento do próximo comando (CallNext).
	OutputConverter func(out string) (string, []string)
	// CallNext recebe uma função utilizada para preparar um novo comando a ser executado em sequência
	CallNext func(...string) *Cmd
}
