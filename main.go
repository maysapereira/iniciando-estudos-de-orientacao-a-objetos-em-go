package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float64
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64){
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito

		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "É obrigatório que o valor depositado seja maior que zero", c.saldo
	}
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

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)

		return true
	} else {
		return false
	}

}

func main() {
	contaMaysa := ContaCorrente { titular: "Maysa", numeroAgencia: 789, numeroConta: 654321, saldo: 150.50 }
	contaKasper := ContaCorrente {"Kasper", 222, 111222, 200 }
	contaIna := ContaCorrente {"Iná", 325, 3344, 300 }
	contaNikole := ContaCorrente{"Nikole", 444, 112233, 350}

	fmt.Println(contaMaysa)
	fmt.Println(contaKasper)
	fmt.Println(contaIna)
	fmt.Println(contaNikole)

	status := contaIna.Transferir(100, &contaMaysa)

	fmt.Println(status)
	fmt.Println(contaIna)
	fmt.Println(contaMaysa)
	
}