package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Magenta("   ____________  ______ __________  ________  _____    ____  ____ \n" +
		"  / ____/ __ \\ \\/ / __ /_  __/ __ \\/ ____/ / / /   |  / __ \\/ __ \\\n" +
		" / /   / /_/ /\\  / /_/ // / / / / / / __/ / / / /| | / /_/ / / / /\n" +
		"/ /___/ _, _/ / / ____// / / /_/ / /_/ / /_/ / ___ |/ _, _/ /_/ / \n" +
		"\\____/_/ |_| /_/_/    /_/  \\____/\\____/\\____/_/  |_/_/ |_/_____/  \n")

	// Основной код:
Menu:
	for {
		variant := getMenu()

		switch variant {
		case 1:
			encrypt()
		case 2:
			decrypt()
		case 3:
			break Menu
		default:
			fmt.Println("Некорректный выбор. Попробуйте снова.")
		}
	}
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант")
	fmt.Println("1. Зашифровать файл или директори")
	fmt.Println("2. Расшифровать файл или дерикторию .secret")
	fmt.Println("3. Выход")
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&variant)
	return variant
}

func encrypt() {
}

func decrypt() {
	fmt.Print("Расшифровка")
}
