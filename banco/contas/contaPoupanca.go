package contas

import (
	"github.com/alura-golang/banco/clientes"
)

type ContaPoupanca struct {
	Titular  clientes.Titular
	Agencia  int
	Conta    int
	Operacao int
	saldo    float64
}

func (c *ContaPoupanca) Sacar(saque float64) string {
	if saque >= 0 && saque <= c.saldo {
		c.saldo -= saque
		return "Saque realizado com sucesso!"
	}
	return "Saldo insuficiente"
}

func (c *ContaPoupanca) Depositar(deposito float64) string {
	if deposito >= 0 {
		c.saldo += deposito
		return "Deposito efetuado com sucesso!"
	}
	return "Valor inválido"
}

func (c *ContaPoupanca) Transferir(valor float64, ccDestino *ContaCorrente) bool {
	if valor >= 0 && valor <= c.saldo {
		c.saldo -= valor
		ccDestino.saldo += valor
		return true
	}
	return false
}

func (c *ContaPoupanca) Saldo() float64 {
	return c.saldo
}
