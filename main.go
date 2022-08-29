package main

import (
	"GoBanco/clientes"
	"GoBanco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	clienteMatheus := clientes.Titular{"Pedrosa", "37448456806", "Dev"}
	contaMatheus := contas.ContaCorrente{Titular: clienteMatheus, NmAgencia: 123, NmConta: 123456}
	contaMatheus.Deposito(500)
	// fmt.Println(contaMatheus.ObterSaldo())
	contaDenis := contas.ContaPoupanca{}

	contaDenis.Deposito(100)
	PagarBoleto(&contaDenis, 50)
	PagarBoleto(&contaMatheus, 500)

	fmt.Println(contaDenis.ObterSaldo())
	fmt.Println(contaMatheus.ObterSaldo())

}
