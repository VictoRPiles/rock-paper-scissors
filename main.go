package main

import "fmt"
import "math/rand"
import "strings"

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
	fmt.Scanln(&turno)

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
			"---.__(___)            " + "            (___)__.---\n")
	}
	if usuario == 1 && maquina == 2 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    _______            " + "           ________    \n" +
			"---'   ____)           " + "        __(____    '---\n" +
			"      (_____)          " + "      _(_______        \n" +
			"      (_____)          " + "     (_________        \n" +
			"      (____)           " + "      (________        \n" +
			"---.__(___)            " + "         (_________.---\n")
	}
	if usuario == 1 && maquina == 3 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    _______            " + "           ________    \n" +
			"---'   ____)           " + "      ____(____    '---\n" +
			"      (_____)          " + "     (_________        \n" +
			"      (_____)          " + "     (_________        \n" +
			"      (____)           " + "           (____)      \n" +
			"---.__(___)            " + "            (___)__.---\n")
	}
	if usuario == 2 && maquina == 1 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    ________           " + "            _______    \n" +
			"---'    ____)__        " + "           (____   '---\n" +
			"        _______)_      " + "          (_____)      \n" +
			"        _________)     " + "          (_____)      \n" +
			"        ________)      " + "           (____)      \n" +
			"---._________)         " + "            (___)__.---\n")
	}
	if usuario == 2 && maquina == 2 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    ________           " + "           ________    \n" +
			"---'    ____)__        " + "        __(____    '---\n" +
			"        _______)_      " + "      _(_______        \n" +
			"        _________)     " + "     (_________        \n" +
			"        ________)      " + "      (________        \n" +
			"---._________)         " + "         (_________.---\n")
	}
	if usuario == 2 && maquina == 3 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    ________           " + "           ________    \n" +
			"---'    ____)__        " + "      ____(____    '---\n" +
			"        _______)_      " + "     (_________        \n" +
			"        _________)     " + "     (_________        \n" +
			"        ________)      " + "           (____)      \n" +
			"---._________)         " + "            (___)__.---\n")
	}
	if usuario == 3 && maquina == 1 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    ________           " + "            _______    \n" +
			"---'    ____)____      " + "           (____   '---\n" +
			"        _________)     " + "          (_____)      \n" +
			"        _________)     " + "          (_____)      \n" +
			"      (____)           " + "           (____)      \n" +
			"---.__(___)            " + "            (___)__.---\n")
	}
	if usuario == 3 && maquina == 2 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    ________           " + "           ________    \n" +
			"---'    ____)____      " + "        __(____    '---\n" +
			"        _________)     " + "      _(_______        \n" +
			"        _________)     " + "     (_________        \n" +
			"      (____)           " + "      (________        \n" +
			"---.__(___)            " + "         (_________.---\n")
	}
	if usuario == 3 && maquina == 3 {
		fmt.Printf("" +
			"TU JUGADA              " + "                MAQUINA\n" +
			"    ________           " + "           ________    \n" +
			"---'    ____)____      " + "      ____(____    '---\n" +
			"        _________)     " + "     (_________        \n" +
			"        _________)     " + "     (_________        \n" +
			"      (____)           " + "           (____)      \n" +
			"---.__(___)            " + "            (___)__.---\n")
	}
}

/* Imprime el veredicto de la jugada, llana a la función dibujar */
func jugar(maquina int, usuario int) {

	dibujar(maquina, usuario)

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

func main() {
	var seguir string = "s"

	fmt.Printf("===================================\n")
	fmt.Printf("====  PIEDRA - PAPEL - TIJERA  ====\n")
	fmt.Printf("===================================\n")

	for strings.EqualFold(seguir, "s") {
		jugar(turnoMaquina(), turnoUsuario())

		fmt.Printf("SEGUIR JUGANDO? (S / N) > ")
		fmt.Scanln(&seguir)
	}
}
