package main

import (
	"fmt"
	"math/rand"
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

/* Genera un número aleatorio entre 1 y 3 */
func turnoMaquina() int {
	return rand.Intn(3-1+1) + 1
}

/* Pide un número al usuario, debe estar entre 1 y 3 */
func turnoUsuario() int {
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
		turno = turnoUsuario()
	}

	return turno
}

/* Dibuja las manos acordes con las jugadas */
func dibujar(maquina int, usuario int) {
	if usuario == 1 && maquina == 1 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    _______            " + "            _______    \n" +
			"---'   ____)           " + "           (____   '---\n" +
			"      (_____)          " + "          (_____)      \n" +
			"      (_____)          " + "          (_____)      \n" +
			"      (____)           " + "           (____)      \n" +
			"---.__(___)            " + "            (___)__.---\n\n")
	}
	if usuario == 1 && maquina == 2 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    _______            " + "           ________    \n" +
			"---'   ____)           " + "        __(____    '---\n" +
			"      (_____)          " + "      _(_______        \n" +
			"      (_____)          " + "     (_________        \n" +
			"      (____)           " + "      (________        \n" +
			"---.__(___)            " + "         (_________.---\n\n")
	}
	if usuario == 1 && maquina == 3 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    _______            " + "           ________    \n" +
			"---'   ____)           " + "      ____(____    '---\n" +
			"      (_____)          " + "     (_________        \n" +
			"      (_____)          " + "     (_________        \n" +
			"      (____)           " + "           (____)      \n" +
			"---.__(___)            " + "            (___)__.---\n\n")
	}
	if usuario == 2 && maquina == 1 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    ________           " + "            _______    \n" +
			"---'    ____)__        " + "           (____   '---\n" +
			"        _______)_      " + "          (_____)      \n" +
			"        _________)     " + "          (_____)      \n" +
			"        ________)      " + "           (____)      \n" +
			"---._________)         " + "            (___)__.---\n\n")
	}
	if usuario == 2 && maquina == 2 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    ________           " + "           ________    \n" +
			"---'    ____)__        " + "        __(____    '---\n" +
			"        _______)_      " + "      _(_______        \n" +
			"        _________)     " + "     (_________        \n" +
			"        ________)      " + "      (________        \n" +
			"---._________)         " + "         (_________.---\n\n")
	}
	if usuario == 2 && maquina == 3 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    ________           " + "           ________    \n" +
			"---'    ____)__        " + "      ____(____    '---\n" +
			"        _______)_      " + "     (_________        \n" +
			"        _________)     " + "     (_________        \n" +
			"        ________)      " + "           (____)      \n" +
			"---._________)         " + "            (___)__.---\n\n")
	}
	if usuario == 3 && maquina == 1 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    ________           " + "            _______    \n" +
			"---'    ____)____      " + "           (____   '---\n" +
			"        _________)     " + "          (_____)      \n" +
			"        _________)     " + "          (_____)      \n" +
			"      (____)           " + "           (____)      \n" +
			"---.__(___)            " + "            (___)__.---\n\n")
	}
	if usuario == 3 && maquina == 2 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    ________           " + "           ________    \n" +
			"---'    ____)____      " + "        __(____    '---\n" +
			"        _________)     " + "      _(_______        \n" +
			"        _________)     " + "     (_________        \n" +
			"      (____)           " + "      (________        \n" +
			"---.__(___)            " + "         (_________.---\n\n")
	}
	if usuario == 3 && maquina == 3 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    ________           " + "           ________    \n" +
			"---'    ____)____      " + "      ____(____    '---\n" +
			"        _________)     " + "     (_________        \n" +
			"        _________)     " + "     (_________        \n" +
			"      (____)           " + "           (____)      \n" +
			"---.__(___)            " + "            (___)__.---\n\n")
	}
}

/* Imprime el veredicto de la jugada */
func jugar(maquina int, usuario int) {

	if maquina == usuario {
		fmt.Printf("EMPATE\n")
	}
	if maquina == 1 && usuario == 2 {
		fmt.Printf("GANAS\n")
	}
	if maquina == 1 && usuario == 3 {
		fmt.Printf("PIERDES\n")
	}
	if maquina == 2 && usuario == 1 {
		fmt.Printf("PIERDES\n")
	}
	if maquina == 2 && usuario == 3 {
		fmt.Printf("GANAS\n")
	}
	if maquina == 3 && usuario == 1 {
		fmt.Printf("GANAS\n")
	}
	if maquina == 3 && usuario == 2 {
		fmt.Printf("PIERDES\n")
	}
}

func juegoUnJugador() {
	var seguir = "s"
	for strings.EqualFold(seguir, "s") {
		var maquina int
		var usuario int

		maquina = turnoMaquina()
		usuario = turnoUsuario()

		dibujar(maquina, usuario)
		jugar(maquina, usuario)

		fmt.Printf("SEGUIR JUGANDO? (S / N) > ")
		_, err := fmt.Scanln(&seguir)
		if err != nil {
			return
		}
	}
}

func juegoMultijugador(cantidadJugadores int) {

}

func main() {

	fmt.Printf("===================================\n")
	fmt.Printf("====  PIEDRA - PAPEL - TIJERA  ====\n")
	fmt.Printf("===================================\n")

	switch menu() {
	case 1:
		juegoUnJugador()
		break
	case 2:
		var cantidadJugadores int

		_, err := fmt.Scanln(&cantidadJugadores)
		if err != nil {
			return
		}

		juegoMultijugador(cantidadJugadores)
	}

}
