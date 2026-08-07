package main

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/meekailkhan/meekcli/commands/auth"
	"github.com/meekailkhan/meekcli/commands/create"
	removerepo "github.com/meekailkhan/meekcli/commands/delete"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(length int) (string, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i := range bytes {
		bytes[i] = charset[int(bytes[i])%len(charset)]
	}

	return string(bytes), nil
}

func main() {
	if _, err := os.Stat(".env"); errors.Is(err, os.ErrNotExist) {
		os.Create(".env")
	}
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	ctx := context.Background()
	switch os.Args[1] {
	case "create":
		create.CreateRepository(os.Args[2:], ctx)
	case "--save-token":
		fmt.Println("dflkadsfklksd")
		auth.SaveToken(os.Args[2:])
	case "delete":
		removerepo.DeleteRepository(os.Args[2:],ctx)
	}

}
