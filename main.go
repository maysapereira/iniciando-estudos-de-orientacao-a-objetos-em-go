package main

import (
	"fmt"

	"banco/contas"

	"banco/clientes"
)

func main() {

	clienteMaysa := clientes.Titular {"Maysa", "111.222.333.44", "Desenvolvedora de Software"}

	contaMaysa := contas.ContaCorrente { clienteMaysa, 789, 654321, 800.50 }

	fmt.Println(contaMaysa)
	
}