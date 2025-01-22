package contas

import (
	"banco-com-golang/clientes"
)

type ContaCorrente struct {
	Titular clientes.Usuario
	saldo   float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	if valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito realizado com sucesso!", c.saldo
	} else {
		return "Depósito não realizado!", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDeTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDeTransferencia < c.saldo {
		c.saldo -= valorDeTransferencia
		contaDestino.Depositar(valorDeTransferencia)
		return true
	} else {
		return false
	}

}

func (c *ContaCorrente) ObterSaldo() (string, float64) {
	return "Saldo: R$", c.saldo
}

func (c *ContaCorrente) analiseDeCreditos() bool {
	if c.saldo >= 1000 {
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) calcularTaxaDeJuros(tempoDePagamento int, valorDoFinanciamento float64) float64 {
	if tempoDePagamento <= 10 || valorDoFinanciamento <= 1000 {
		return 0.5
	} else {
		return 0.1
	}

}

func (c *ContaCorrente) Financiamento(valorDoFinanciamento float64, tempoDePagamento int) (string, float64) {
	contador := c.analiseDeCreditos()

	if contador == true {
		valorDaParcela := (valorDoFinanciamento * c.calcularTaxaDeJuros(tempoDePagamento, valorDoFinanciamento)) + valorDoFinanciamento
		c.saldo += valorDoFinanciamento
		return "Financiamento aprovado! \nO valor da parcela é de: R$", valorDaParcela
	} else {
		return "Financiamento não aprovado!\nSaldo:", c.saldo
	}

}
