package contas

import "fmt"

// ContaCorrente é a entidade
type ContaCorrente struct {
	Titular string
	Agencia int
	Conta   int
	Saldo   float64
}

// Sacar subtrai o valor da conta caso haja saldo
func (c *ContaCorrente) Sacar(valorSaque float64) {
	podeSacar := valorSaque > 0 && valorSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorSaque
	} else {
		fmt.Println("saldo insuficiente")
	}
}

// Depositar atualiza o saldo com o valor enviado
func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	valido := valorDeposito > 0
	if valido {
		c.Saldo += valorDeposito
		return "depósito realizado", c.Saldo
	}
	return "depósito não aceito", c.Saldo
}

// Transferir envia um valor para a conta destino
func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor > c.Saldo && valor < 0 {
		return false
	}

	c.Saldo -= valor
	contaDestino.Depositar(valor)
	return true
}
