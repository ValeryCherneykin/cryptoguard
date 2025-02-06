package main

import (
	"cryptoguard/password"
	"fmt"
)

func main() {
	// TODO Реализовать UI / Меню
	// TODO Реализовать генератор пароля
	key, _ := password.GeneratePassword(10)
	fmt.Println(key)
	// TODO Реализовать записать пароля в файл .txt
	// TODO Реализовать зашифровку файла
	// TODO Реализовать расшифровку файла

}
