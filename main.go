package main

import (
	"fmt"

	"banco/contas"

	"banco/clientes"
)

func PagarBoleto (conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaMaysa := contas.ContaPoupanca {}
	contaMaysa.Depositar(100)
	PagarBoleto(&contaMaysa, 60)
	fmt.Println(contaMaysa.ObterSaldo())

	contaIna := contas.ContaCorrente {}
	contaIna.Depositar(200)
	PagarBoleto(&contaIna, 150)
	fmt.Println(contaIna.ObterSaldo())
	
}