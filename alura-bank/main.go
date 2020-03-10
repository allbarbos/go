package main

import (
	"fmt"

	"github.com/go/alura-bank/clientes"
	"github.com/go/alura-bank/contas"
)

func main() {
	clienteAllan := clientes.Titular{Nome: "Allan Barbosa", Cpf: "12345356545", Profissao: "Dev"}
	contaAllan := contas.ContaCorrente{Titular: clienteAllan, Agencia: 0001, Conta: 123456, Saldo: 1000}
	fmt.Println(contaAllan)
}
