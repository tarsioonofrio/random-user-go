package main


type Pessoa struct {
	ID			int
	Genero 		string
	Email		string
	Nome		string
	Username	string
	Password  	string
}

type Endereco struct {
	ID			int
	PessoaID 	int
	Pessoa   	*Pessoa //`pg:"fk:id"`
	Rua		 	string
	Numero  	string
	Cidade  	string
	Estado  	string
}
