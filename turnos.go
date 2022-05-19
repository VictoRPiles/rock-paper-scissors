package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"math/rand"
	"strconv"
	"syscall"
)

/* Genera un número aleatorio entre 1 y 3 */
func turnoMaquina() int {
	return rand.Intn(3-1+1) + 1
}

/*
Para el modo un jugador
Pide un número al usuario, debe estar entre 1 y 3
*/
func turnoUsuarioUnJugador() int {
	var turno int

	fmt.Printf("TU TURNO\n")
	fmt.Printf("1) PIEDRA\n")
	fmt.Printf("2) PAPEL\n")
	fmt.Printf("3) TIJERA\n")
	fmt.Printf("> ")
	_, err := fmt.Scanln(&turno)
	if err != nil {
		return 0
	}

	if turno < 1 || turno > 3 {
		fmt.Printf("OPCIÓN INCORRECTA\n")
		turno = turnoUsuarioUnJugador()
	}

	return turno
}

/*
Para el modo un Multijugador
Pide un número al usuario, debe estar entre 1 y 3
*/
func turnoUsuarioMultijugador() int {
	var turnoString string
	var turno int

	fmt.Printf("TU TURNO\n")
	fmt.Printf("1) PIEDRA\n")
	fmt.Printf("2) PAPEL\n")
	fmt.Printf("3) TIJERA\n")
	fmt.Printf("> ")

	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return 0
	}

	turnoString = string(bytePassword)

	/* Conversión de string a entero */
	turno, err = strconv.Atoi(turnoString)
	if err != nil {
		return 0
	}

	if turno < 1 || turno > 3 {
		fmt.Printf("OPCIÓN INCORRECTA\n")
		turno = turnoUsuarioUnJugador()
	}

	return turno
}
