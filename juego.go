package main

import (
	"fmt"
	"strings"
)

/* Permite seleccionar el modo de juego */
func menu() int {
	var op int
	/* IMPORTANTE: Cambiar estos valores cuando se añaden o se quitan entradas en el menu */
	const MinOp = 1
	const MaxOp = 3

	fmt.Printf("1) Un jugador\n")
	fmt.Printf("2) Multijugador\n")
	fmt.Printf("3) Salir\n")
	fmt.Printf("> ")

	_, err := fmt.Scanln(&op)
	if err != nil {
		op = 0
	}

	if op < MinOp || op > MaxOp {
		fmt.Printf("ERROR: Opción no válida\n")
		return menu()
	}

	return op
}

/* Imprime el veredicto de la jugada */
func jugar(maquina int, usuario int, nombreUsuario1 string, nombreUsuario2 string) {

	if maquina == usuario {
		fmt.Printf("EMPATE\n")
	}
	if maquina == 1 && usuario == 2 {
		fmt.Printf(nombreUsuario2 + " GANA\n")
	}
	if maquina == 1 && usuario == 3 {
		fmt.Printf(nombreUsuario1 + " GANA\n")
	}
	if maquina == 2 && usuario == 1 {
		fmt.Printf(nombreUsuario1 + " GANA\n")
	}
	if maquina == 2 && usuario == 3 {
		fmt.Printf(nombreUsuario2 + " GANA\n")
	}
	if maquina == 3 && usuario == 1 {
		fmt.Printf(nombreUsuario2 + " GANA\n")
	}
	if maquina == 3 && usuario == 2 {
		fmt.Printf(nombreUsuario1 + " GANA\n")
	}
}

/* Juego con una entrada para el jugador y cálculo del turno de la máquina */
func juegoUnJugador() {
	var seguir = "s"
	for strings.EqualFold(seguir, "s") {
		var maquina int
		var usuario int

		maquina = turnoMaquina()
		usuario = turnoUsuarioUnJugador()

		dibujar(maquina, usuario, "MAQUINA", "TU JUGADA")
		jugar(maquina, usuario, "MAQUINA", "TU JUGADA")

		fmt.Printf("SEGUIR JUGANDO? (S / N) > ")
		_, err := fmt.Scanln(&seguir)
		if err != nil {
			return
		}
	}
}

/* Juego con dos entradas, una para cada jugador */
func juegoMultijugador() {
	var seguir = "s"
	for strings.EqualFold(seguir, "s") {
		var usuario1 int
		var usuario2 int

		fmt.Printf("===== Jugador 1 =====\n")
		usuario1 = turnoUsuarioMultijugador()
		fmt.Printf("===== Jugador 2 =====\n")
		usuario2 = turnoUsuarioMultijugador()

		dibujar(usuario1, usuario2, "JUGADOR 1", "JUGADOR 2")
		jugar(usuario1, usuario2, "JUGADOR 1", "JUGADOR 2")

		fmt.Printf("SEGUIR JUGANDO? (S / N) > ")
		_, err := fmt.Scanln(&seguir)
		if err != nil {
			return
		}
	}
}
