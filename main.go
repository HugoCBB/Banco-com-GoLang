package main

import (
	"caixa/contas"
	"fmt"
)

func main() {

	p1 := contas.ContaCorrente{}
	fmt.Print(p1)
	// p1.Depositar(1000)
	// fmt.Println(p1.Financiamento(100, 20))
	// comando := Hub()
	// Comando(comando)
}

func Hub() int {
	fmt.Println("--------------------------------")
	fmt.Println("  BEM VINDO AO BANCO AGIOTAGENS")
	fmt.Println("--------------------------------")
	fmt.Println("DIGITE O SERVIÇO DESEJADO")
	fmt.Println("1 - Depositar")
	fmt.Println("2 - Sacar")
	fmt.Println("3 - Transferir")
	fmt.Println("4 - Financiar")
	fmt.Println("0 - Sair do programa")
	fmt.Println("--------------------------------")
	fmt.Print("Qual a sua escolha: ")

	var comando int
	fmt.Scan(&comando)
	return comando

}

func Comando(comando int) {
	switch comando {
	case 1:
		fmt.Println("Depositar")
	case 2:
		fmt.Println("Sacar")
	case 3:
		fmt.Println("Transferir")
	case 4:
		fmt.Println("Financiar")
	case 0:
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("Comando inválido")
	}

}
