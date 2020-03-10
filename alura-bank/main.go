package main

import (
	"fmt"

	"github.com/go/alura-bank/contas"
)

func main() {
	ccAllan := contas.ContaCorrente{Titular: "Allan", Agencia: 0001, Conta: 123456, Saldo: 600}
	ccBile := contas.ContaCorrente{Titular: "Bile", Agencia: 0001, Conta: 654321, Saldo: 100}

	ccAllan.Sacar(300)
	ccAllan.Depositar(100)
	ccAllan.Transferir(50, &ccBile)

	fmt.Println("Saldo Allan:", ccAllan.Saldo)
	fmt.Println("Saldo Bile:", ccBile.Saldo)
}
