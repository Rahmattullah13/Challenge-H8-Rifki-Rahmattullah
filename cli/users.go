package cli

import (
	"bufio"
	"fmt"
	"go-trial-class/config"
	"go-trial-class/entity"
	"go-trial-class/helpers"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func Register() {
	helpers.ClearScreen()
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("------Register------")

	var username string
	fmt.Println("Masukan username anda : ")
	username, _ = consoleReader.ReadString('\n')
	username = strings.TrimSpace(username)

	var existingUser entity.Users
	err := config.DB.Where("username = ?", username).First(&existingUser).Error
	if err == nil {
		fmt.Println("Username sudah digunakan. Silahkan gunakan username lain.")
		Register()
		return
	}

	fmt.Println("Masukan password anda : ")
	password, _ := consoleReader.ReadString('\n')
	password = strings.TrimSpace(password)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Terjadi kesalahan. Silahkan coba lagi.")
		Register()
		return
	}

	user := entity.Users{
		Username: username,
		Password: string(hashedPassword),
	}

	err = config.DB.Create(&user).Error
	if err != nil {
		fmt.Println("Terjadi kesalahan. Silahkan coba lagi.")
		Register()
		return
	}

	fmt.Println("Register Berhasil!!!")
	MainMenu()
}

func Login() {
	helpers.ClearScreen()

	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Println("------Login------")

	var username string
	fmt.Println("Masukan username anda : ")
	username, _ = consoleReader.ReadString('\n')
	username = strings.TrimSpace(username)

	var password string
	fmt.Println("Masukan password anda : ")
	password, _ = consoleReader.ReadString('\n')
	password = strings.TrimSpace(password)

	var user entity.Users
	err := config.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		fmt.Println("Username atau password salah. Silahkan login kembali")
		Login()
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println("Username atau password salah. Silahkan login kembali.")
		Login()
		return
	}

	fmt.Printf("Selamat datang, %s!!!\n", username)
	MainMenu()
}
