package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo float64
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64){
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito

		return "Depósito realizado com sucesso", c.saldo
	} else {
		return "É obrigatório que o valor depositado seja maior que zero", c.saldo
	}
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque

		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino *ContaPoupanca) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)

		return true
	} else {
		return false
	}

}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}