package main

import (
	"cryptoguard/files"
	"cryptoguard/password"
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
	key, _ := password.GeneratePassword(10)
	files.CreatFile("file.txt")
	files.WriteFile("file.txt", []byte(key))
	res, err := files.PathExists("text.txt")
	fmt.Println(res, err)
}

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант")
	fmt.Println("Зашифровать файл или директори")
	fmt.Println("Расшифровать файл или дерикторию .secret")
	fmt.Println("Выход")
	fmt.Print("Ваш выбор: ")
	fmt.Scan(&variant)
	return variant
}
