package main

import (
	"fmt"

	"banco/contas"

	"banco/clientes"
)

func main() {

	contaMaysa := contas.ContaPoupanca {}

	contaMaysa.Depositar(100)
	contaMaysa.Sacar(88)

	fmt.Println(contaMaysa.ObterSaldo())
	
}