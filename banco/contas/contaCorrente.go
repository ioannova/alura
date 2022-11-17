package contas

import (
	"github.com/alura-golang/banco/clientes"
)

type ContaCorrente struct {
	Titular clientes.Titular
	Agencia int
	Conta   int
	saldo   float64
}

func (c *ContaCorrente) Sacar(saque float64) string {
	if saque >= 0 && saque <= c.saldo {
		c.saldo -= saque
		return "Saque realizado com sucesso!"
	}
	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(deposito float64) string {
	if deposito >= 0 {
		c.saldo += deposito
		return "Deposito efetuado com sucesso!"
	}
	return "Valor inválido"
}

func (c *ContaCorrente) Transferir(valor float64, ccDestino *ContaCorrente) bool {
	if valor >= 0 && valor <= c.saldo {
		c.saldo -= valor
		ccDestino.saldo += valor
		return true
	}
	return false
}

func (c *ContaCorrente) Saldo() float64 {
	return c.saldo
}
