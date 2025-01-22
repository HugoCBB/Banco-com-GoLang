package main

import (
	"banco-com-golang/clientes"
	"banco-com-golang/contas"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	p1 := contas.ContaCorrente{}
	for {
		comando := Hub()
		Comando(comando, &p1)

		if comando == 0 {
			fmt.Println(p1)
			break
		}
	}
}

func Hub() int {
	fmt.Println("--------------------------------")
	fmt.Println("  BEM VINDO AO BANCO AGIOTAGENS")
	fmt.Println("--------------------------------")
	fmt.Println("DIGITE O SERVIÇO DESEJADO")
	fmt.Println("1 - Cadastrar Conta")
	fmt.Println("2 - Exibir dados")
	fmt.Println("3 - Depositar")
	fmt.Println("4 - Sacar")
	fmt.Println("5 - Transferir")
	fmt.Println("6 - Financiar")
	fmt.Println("7 - Obter saldo")
	fmt.Println("0 - Sair do programa")
	fmt.Println("--------------------------------")
	fmt.Print("Qual a sua escolha: ")

	var comando int
	fmt.Scan(&comando)
	return comando
}

func Comando(comando int, conta *contas.ContaCorrente) {
	switch comando {
	case 1:
		CadastrarConta(conta)
	case 2:
		ExibirDados(conta)
	case 3:
		Depositar(conta)
	case 4:
		Sacar(conta)
	case 5:
		Transferir(conta)
	case 6:
		Financiamento(conta)
	case 7:
		ObterSaldo(conta)
	case 0:
		fmt.Println("Saindo do programa")
		time.Sleep(2 * time.Second)
	default:
		fmt.Println("Comando inválido")
	}
}

func CadastrarConta(conta *contas.ContaCorrente) {
	var titular, telefone, cpf string
	var numeroDaConta, numeroDaAgencia int
	var input string

	fmt.Print("Digite o nome do titular: ")
	fmt.Scan(&titular)
	fmt.Print("Digite o telefone: ")
	fmt.Scan(&telefone)
	fmt.Print("Digite o CPF: ")
	fmt.Scan(&cpf)
	fmt.Print("Digite o número da conta: ")
	fmt.Scan(&numeroDaConta)
	fmt.Print("Digite o número da agência: ")
	fmt.Scan(&numeroDaAgencia)

	// Corrigido: atribuir os valores corretamente
	conta.Titular = clientes.Usuario{
		Titular:         titular,
		Telefone:        telefone,
		CPF:             cpf,
		NumeroDaConta:   numeroDaConta,
		NumeroDaAgencia: numeroDaAgencia,
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Conta cadastrada com sucesso!")
	fmt.Scan(&input)
	ClearTerminal()
}

func ExibirDados(conta *contas.ContaCorrente) {
	var input string
	fmt.Println("--------------------------------")
	fmt.Println("        DADOS DO USUARIO        ")
	fmt.Println("--------------------------------")
	fmt.Println("Titular:", conta.Titular.Titular)
	fmt.Println("Telefone:", conta.Titular.Telefone)
	fmt.Println("CPF:", conta.Titular.CPF)
	fmt.Println("Número da conta:", conta.Titular.NumeroDaConta)
	fmt.Println("Número da agência:", conta.Titular.NumeroDaAgencia)

	fmt.Scan(&input)
	ClearTerminal()

}

func Depositar(conta *contas.ContaCorrente) {
	var valorDeDeposito float64
	var input string

	fmt.Print("Digite o valor do depósito: ")
	fmt.Scan(&valorDeDeposito)

	resultado, _ := conta.Depositar(valorDeDeposito)

	time.Sleep(1 * time.Second)
	fmt.Println(resultado)

	fmt.Scan(&input)
	ClearTerminal()
}

func Sacar(conta *contas.ContaCorrente) {
	var valorDeSaque float64
	var input string

	fmt.Print("Digite o valor do saque: R$")
	fmt.Scan(&valorDeSaque)

	fmt.Println(conta.Sacar(valorDeSaque))

	time.Sleep(1 * time.Second)
	ObterSaldo(conta)
	fmt.Scan(&input)
	ClearTerminal()
}

func Transferir(conta *contas.ContaCorrente) {
	var valorDeTransferencia float64
	var contaDestino contas.ContaCorrente
	var titularDestino string
	var input string

	fmt.Print("Digite o valor da transferência: ")
	fmt.Scan(&valorDeTransferencia)

	fmt.Print("Digite o nome do titular da conta de destino: ")
	fmt.Scan(&titularDestino)

	contaDestino.Titular.Titular = titularDestino

	if conta.Transferir(valorDeTransferencia, &contaDestino) {
		fmt.Println("Transferência realizada com sucesso!")
	} else {
		fmt.Println("Falha na transferência. Saldo insuficiente.")
	}

	fmt.Scan(&input)
	ClearTerminal()
}

func Financiamento(conta *contas.ContaCorrente) {
	var valorDoFinanciamento float64
	var tempoDePagamento int
	var input string

	fmt.Print("Digite o valor do financiamento: R$")
	fmt.Scan(&valorDoFinanciamento)

	fmt.Print("Digite o tempo de pagamento: ")
	fmt.Scan(&tempoDePagamento)

	time.Sleep(1 * time.Second)
	fmt.Println(conta.Financiamento(valorDoFinanciamento, tempoDePagamento))

	fmt.Scan(&input)
	ClearTerminal()
}

func ObterSaldo(conta *contas.ContaCorrente) float64 {
	var input string

	mensagem, saldo := conta.ObterSaldo()
	fmt.Println(mensagem, saldo)

	time.Sleep(1 * time.Second)
	fmt.Scan(&input)
	ClearTerminal()
	return saldo
}

func ClearTerminal() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
