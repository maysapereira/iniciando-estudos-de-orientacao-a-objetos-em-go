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

	var contaIna *ContaCorrente
	contaIna = new(ContaCorrente)
	contaIna.titular = "Iná"
	contaIna.numeroAgencia = 325
	contaIna.numeroConta = 333444
	contaIna.saldo = 300

	fmt.Println(contaIna)
}