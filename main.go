package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please input string you need hash...")
		os.Exit(0)
	}

	password := os.Args[1]
	hash, _ := HashPassword(password)

	fmt.Println("Input:   ", password)
	fmt.Println("Hash:    ", hash)

	//match := CheckPasswordHash(password, hash)
	//fmt.Println("Match:   ", match)
}
