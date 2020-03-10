package contas

import (
	"fmt"

	c "github.com/go/alura-bank/clientes"
)

// ContaPoupanca é a entidade
type ContaPoupanca struct {
	Titular                  c.Titular
	Agencia, Conta, Operacao int
	saldo                    float64
}

// Sacar subtrai o valor da conta caso haja saldo
func (c *ContaPoupanca) Sacar(valorSaque float64) {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
	} else {
		fmt.Println("saldo insuficiente")
	}
}

// Depositar atualiza o saldo com o valor enviado
func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	valido := valorDeposito > 0
	if valido {
		c.saldo += valorDeposito
		return "depósito realizado", c.saldo
	}
	return "depósito não aceito", c.saldo
}

// Saldo informa o valor em conta
func (c *ContaPoupanca) Saldo() float64 {
	return c.saldo
}
