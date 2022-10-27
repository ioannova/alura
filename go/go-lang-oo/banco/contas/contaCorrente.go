package contas

type ContaCorrente struct {
	Titular string
	Agencia int
	Conta   int
	Saldo   float64
}

func (c *ContaCorrente) Sacar(saque float64) string {
	if saque > 0 && saque <= c.Saldo {
		c.Saldo -= saque
		return "Saque realizado com sucesso!"
	}
	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(deposito float64) string {
	if deposito > 0 {
		c.Saldo += deposito
		return "Deposito efetuado com sucesso!"
	}
	return "Valor inválido"
}

func (c *ContaCorrente) Transferir(valor float64, ccDestino *ContaCorrente) bool {
	if valor > 0 && valor <= c.Saldo {
		c.Saldo -= valor
		ccDestino.Saldo += valor
		return true
	}
	return false
}
