package contas

import (
	"fmt"

	"github.com/go/alura-bank/clientes"
)

// ContaCorrente é a entidade
type ContaCorrente struct {
	Titular clientes.Titular
	Agencia int
	Conta   int
	saldo   float64
}

// Sacar subtrai o valor da conta caso haja saldo
func (c *ContaCorrente) Sacar(valorSaque float64) {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
	} else {
		fmt.Println("saldo insuficiente")
	}
}

// Depositar atualiza o saldo com o valor enviado
func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	valido := valorDeposito > 0
	if valido {
		c.saldo += valorDeposito
		return "depósito realizado", c.saldo
	}
	return "depósito não aceito", c.saldo
}

// Transferir envia um valor para a conta destino
func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor > c.saldo && valor < 0 {
		return false
	}

	c.saldo -= valor
	contaDestino.Depositar(valor)
	return true
}

// Saldo informa o valor em conta
func (c *ContaCorrente) Saldo() float64 {
	return c.saldo
}
