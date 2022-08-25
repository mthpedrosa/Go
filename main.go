package main

import "fmt"

type ContaCorrente struct {
	titular   string
	nmAgencia int
	nmConta   int
	saldo     float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	aviso := valorDoSaque <= c.saldo && valorDoSaque > 0
	if aviso {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente"
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

func main() {
	contaMatheus := ContaCorrente{titular: "Matheus", nmAgencia: 0061, nmConta: 585858, saldo: 100.50}

	contaGustavo := ContaCorrente{titular: "Gustavo", nmAgencia: 0061, nmConta: 585859, saldo: 300.50}

	fmt.Println(contaGustavo, contaMatheus)

	status := contaGustavo.Tranferir(200, &contaMatheus)
	fmt.Println(status)

	fmt.Println(contaGustavo, contaMatheus)

}
