package contas

import "GoBanco/clientes"

type ContaPoupanca struct {
	Titular                      clientes.Titular
	NmAgencia, NmConta, Operacao int
	saldo                        float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	aviso := valorDoSaque <= c.saldo && valorDoSaque > 0
	if aviso {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Deposito(valorDeposito float64) (string, float64) {
	aviso := valorDeposito > 0
	if aviso {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso!", c.saldo
	} else {
		return "O deposito tem que ser maior que zero!", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
