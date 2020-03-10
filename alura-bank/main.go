package main

import (
	"fmt"

	"github.com/go/alura-bank/clientes"
	"github.com/go/alura-bank/contas"
)

type verificarConta interface {
	Sacar(valor float64)
}

// PagarBoleto realiza o pagamento
func PagarBoleto(conta verificarConta, valor float64) {
	conta.Sacar(valor)
}

func main() {
	clienteAllan := clientes.Titular{Nome: "Allan Barbosa", Cpf: "12345356545", Profissao: "Dev"}
	cc := contas.ContaCorrente{Titular: clienteAllan, Agencia: 0001, Conta: 123456}
	cc.Depositar(1000)
	PagarBoleto(&cc, 500)
	fmt.Println(cc.Saldo())

	pp := contas.ContaPoupanca{}
	pp.Depositar(35)
	PagarBoleto(&pp, 30)
	fmt.Println(pp.Saldo())
}
