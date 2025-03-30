package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/atotto/clipboard"
)

const lowercases string = "abcdefghijklmnopqrstuvwxyz"
const uppercases string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers string = "0123456789"
const symbols string = "!@#$%^&*()-_=+"

func createCharset(complexity int8) string {
	switch complexity {
	case 1:
		return lowercases
	case 2:
		return lowercases + uppercases
	case 3:
		return lowercases + uppercases + numbers
	case 4:
		return lowercases + uppercases + numbers + symbols
	default:
		println("Complexity value was set to default. Lowercases")
		return lowercases
	}

}

func generatePassword(length, complexity int8) string {
	charset := createCharset(complexity)
	password := make([]byte, length)
	for i := range password {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[randomIndex.Int64()]
	}
	return string(password)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <lenght> <complexity>")
		return
	}

	length, err := strconv.Atoi(os.Args[1])
	if err != nil || length <= 0 {
		fmt.Println("Error: Lenght must be a positive number (grather than 0).")
		return
	}

	complexity, err := strconv.Atoi(os.Args[2])
	if err != nil || complexity < 1 || complexity > 4 {
		fmt.Println("Error: Complexity is not set between allowed values 1 to 4. \n 1: Lowercases \n 2: Lowercases + uppercases \n 3: Lowercases + uppercases + numbers \n 4: Lowercases + uppercases + numbers + symbols")
		return
	}
	password := generatePassword(int8(length), int8(complexity))
	// Copy to clipboard
	err = clipboard.WriteAll(password)
	if err != nil {
		fmt.Println("Error when copy password to clipboard:", err)
		return
	}

	fmt.Println("Password successfully copied to clipboard! Remember to store in your favorite password manager.")
}
