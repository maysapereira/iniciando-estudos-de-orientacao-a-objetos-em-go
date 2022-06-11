package main

import (
	"fmt"
	
	"banco/contas"
)

func main() {
	contaMaysa := contas.ContaCorrente { Titular: "Maysa", NumeroAgencia: 789, NumeroConta: 654321, Saldo: 150.50 }
	contaKasper := contas.ContaCorrente {"Kasper", 222, 111222, 200 }
	contaIna := contas.ContaCorrente {"Iná", 325, 3344, 300 }
	contaNikole := contas.ContaCorrente{"Nikole", 444, 112233, 350}

	fmt.Println(contaMaysa)
	fmt.Println(contaKasper)
	fmt.Println(contaIna)
	fmt.Println(contaNikole)

	status := contaIna.Transferir(100, &contaMaysa)

	fmt.Println(status)
	fmt.Println(contaIna)
	fmt.Println(contaMaysa)
	
}