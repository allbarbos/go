package main

import (
	"fmt"

	"github.com/go/alura-bank/clientes"
	"github.com/go/alura-bank/contas"
)

func main() {
	clienteAllan := clientes.Titular{Nome: "Allan Barbosa", Cpf: "12345356545", Profissao: "Dev"}
	cc := contas.ContaCorrente{Titular: clienteAllan, Agencia: 0001, Conta: 123456}
	cc.Depositar(1000)
	fmt.Println(cc.Saldo())

	pp := contas.ContaPoupanca{}
	pp.Depositar(35)
	pp.Sacar(5)
	fmt.Println(pp.Saldo())
}
