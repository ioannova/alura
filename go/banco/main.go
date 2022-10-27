package main

import (
	"fmt"

	"github.com/alura/go/banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaLuisa := contas.ContaPoupanca{}
	ccNelson := contas.ContaCorrente{}

	ccNelson.Depositar(200)
	ccNelson.Sacar(55)

	PagarBoleto(&ccNelson, 60)
	fmt.Println(contaLuisa, ccNelson)
	// ccNelson := contas.ContaCorrente{
	// 	Titular: clientes.Titular{
	// 		Nome:      "Nelson",
	// 		CPF:       "123.456.789-1",
	// 		Profissao: "Dev."},
	// 	Agencia: 589,
	// 	Conta:   123456}

	// cliEleonor := clientes.Titular{
	// 	Nome:      "Eleonor",
	// 	CPF:       "123.456.789-1",
	// 	Profissao: "Prof."}

	// ccEleonor := contas.ContaCorrente{
	// 	Titular: cliEleonor,
	// 	Agencia: 588,
	// 	Conta:   12123}

	// ccEleonor.Depositar(100)
	// ccNelson.Depositar(1000)
	// // var ccPaloma *ContaCorrente
	// // ccPaloma = new(ContaCorrente)
	// // ccPaloma.Titular = "Paloma"

	// fmt.Println("Saldo anterior:", ccNelson.Saldo(), "Saldo atual:", ccNelson.Sacar(300))

	// fmt.Println("\nSaldo anterior:", ccNelson.Saldo(), "Saldo atual:", ccNelson.Depositar(384.50))

	// if ccNelson.Transferir(500, &ccEleonor) {
	// 	fmt.Println("\nTransferência de R$ 500,00 realizada com sucesso!")
	// 	fmt.Println("Saldo de ", ccNelson.Titular, " R$ ", ccNelson.Saldo(), "\nSaldo de ", ccEleonor.Titular, " R$ ", ccEleonor.Saldo)
	// } else {
	// 	fmt.Println("\nTransferência de R$ 500 não realizada!")
	// 	fmt.Println("Saldo de ", ccNelson.Titular, " R$ ", ccNelson.Saldo(), "\nSaldo de ", ccEleonor.Titular, " R$ ", ccEleonor.Saldo)
	// }
}
