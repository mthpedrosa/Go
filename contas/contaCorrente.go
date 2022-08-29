package contas

import "GoBanco/clientes"

type ContaCorrente struct {
	Titular            clientes.Titular
	NmAgencia, NmConta int
	saldo              float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	aviso := valorDoSaque <= c.saldo && valorDoSaque > 0
	if aviso {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Deposito(valorDeposito float64) (string, float64) {
	aviso := valorDeposito > 0
	if aviso {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso!", c.saldo
	} else {
		return "O deposito tem que ser maior que zero!", c.saldo
	}
}

func (c *ContaCorrente) Tranferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Deposito(valorTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
