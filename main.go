package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/pflag"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Define command-line flags with both long and short forms
	password := pflag.StringP("password", "p", "", "Password to hash")
	costFactor := pflag.StringP("cost", "c", "12", "Cost factor")

	// Parse command-line flags
	pflag.Parse()

	if *password == "" {
		fmt.Println("Password is required")
		pflag.Usage()
		return
	}

	costFactorInt, err := strconv.Atoi(*costFactor)
	if err != nil {
		fmt.Printf("Invalid cost factor: %v\n", err)
		return
	}

	// Generate the bcrypt hash
	hash, err := bcrypt.GenerateFromPassword([]byte(*password), costFactorInt)
	if err != nil {
		fmt.Printf("Error generating hash: %v\n", err)
		return
	}

	fmt.Printf("{bcrypt}%s\n", string(hash))
}
