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

func main() {
	ccAllan := ContaCorrente{"Allan", 0001, 123456, 600}
	ccBile := ContaCorrente{"Bile", 0001, 654321, 100}

	ccAllan.Sacar(300)
	ccAllan.Depositar(100)
	ccAllan.Transferir(50, &ccBile)

	fmt.Println("Saldo Allan:", ccAllan.saldo)
	fmt.Println("Saldo Bile:", ccBile.saldo)
}
