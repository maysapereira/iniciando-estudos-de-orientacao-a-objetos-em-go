package main

import (
	"fmt"
	
	"banco/contas"

	"banco/clientes"
	
)

func main() {
	contaMaysa := contas.ContaCorrente { Titular: clientes.Titular {
		Nome: "Maysa",
		CPF: "111.222.333.44",
		Profissao: "Desenvolvedora de Software"},
		NumeroAgencia: 789, NumeroConta: 654321, Saldo: 150.50 }

	fmt.Println(contaMaysa)
	
}