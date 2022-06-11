package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque

		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
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

	contaNikole := ContaCorrente{"Nikole", 444, 112233, 350}

	fmt.Println(contaNikole.saldo)

	fmt.Println(contaNikole.Sacar(200))

	fmt.Println(contaNikole.saldo)

}