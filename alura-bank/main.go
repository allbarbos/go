package main

import "fmt"

// ContaCorrente é a entidade
type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

// Sacar subtrai o valor da conta caso haja saldo
func (c *ContaCorrente) Sacar(valorSaque float64) {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		fmt.Println("saque realizado")
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

func main() {
	ccAllan := ContaCorrente{"Allan", 0001, 123456, 600}

	fmt.Println(ccAllan.saldo)
	ccAllan.Sacar(300)
	fmt.Println(ccAllan.saldo)
	status, saldo := ccAllan.Depositar(100)
	fmt.Println(status, "- saldo atual R$", saldo)
}
