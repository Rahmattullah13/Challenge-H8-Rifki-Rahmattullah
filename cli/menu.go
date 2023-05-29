package cli

import (
	"fmt"
	"go-trial-class/helpers"
	"os"
)

func MainMenu() {
	helpers.ClearScreen()
	fmt.Println("Selamat Datang di Challenge Hactive8!!!")
	fmt.Println("--------------------------------")

	var input string
	fmt.Println("Tekan (R) untuk mendaftar")
	fmt.Println("Tekan (L) untuk Login")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err.Error())
	}

	switch input {
	case "r":
		Register()
	case "l":
		Login()
	case "q":
		fmt.Println("Terima kasi telah menggunakan aplikasi ini")
		os.Exit(1)
	default:
		MainMenu()
	}
}
