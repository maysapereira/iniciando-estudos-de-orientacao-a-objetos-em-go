package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float64
}

func main() {
	contaMaysa := ContaCorrente { titular: "Maysa", numeroAgencia: 789, numeroConta: 654321, saldo: 150.50 }

	contaKasper := ContaCorrente {"Kasper", 222, 111222, 200 }

	fmt.Println(contaMaysa)
	fmt.Println(contaKasper)
}