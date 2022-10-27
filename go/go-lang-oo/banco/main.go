package main

import (
	"fmt"

	cc "../../../../srv/alura/go"
)

func main() {
	ccNelson := cc.ContaCorrente{
		Titular: "Nelson",
		Agencia: 589,
		Conta:   123456,
		Saldo:   1000,
	}
	ccEleonor := cc.ContaCorrente{
		Titular: "Eleonor",
		Agencia: 588,
		Conta:   12123,
		Saldo:   500}
	// var ccPaloma *ContaCorrente
	// ccPaloma = new(ContaCorrente)
	// ccPaloma.Titular = "Paloma"

	fmt.Println("Saldo anterior:", ccNelson.Saldo, "Saldo atual:", ccNelson.Sacar(300))

	fmt.Println("\nSaldo anterior:", ccNelson.Saldo, "Saldo atual:", ccNelson.Depositar(384.50))

	if ccNelson.Transferir(500, &ccEleonor) {
		fmt.Println("\nTransferência de R$ 500,00 realizada com sucesso!")
		fmt.Println("Saldo de ", ccNelson.Titular, " R$ ", ccNelson.Saldo, "\nSaldo de ", ccEleonor.Titular, " R$ ", ccEleonor.Saldo)
	} else {
		fmt.Println("\nTransferência de R$ 500 não realizada!")
		fmt.Println("Saldo de ", ccNelson.Titular, " R$ ", ccNelson.Saldo, "\nSaldo de ", ccEleonor.Titular, " R$ ", ccEleonor.Saldo)
	}
}
