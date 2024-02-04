package model

type Pesquisa struct {
	Id    uint64
	Email string
	Nota  uint
}

// mock dos dados, assim irei evitar de complicar essa prova de conceito com a utilizaçãop de um banco de dados
var Pesquisas = []Pesquisa{
	{Id: 1, Email: "exemplo1@example.com", Nota: 1},
	{Id: 2, Email: "exemplo2@example.com"},
	{Id: 3, Email: "exemplo3@example.com"},
	{Id: 4, Email: "exemplo4@example.com"},
	{Id: 5, Email: "exemplo5@example.com", Nota: 3},
	{Id: 6, Email: "exemplo6@example.com"},
	{Id: 7, Email: "exemplo7@example.com", Nota: 5},
	{Id: 8, Email: "exemplo8@example.com"},
	{Id: 9, Email: "exemplo9@example.com", Nota: 4},
	{Id: 10, Email: "exemplo10@example.com"},
}
