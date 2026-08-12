package auth

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func SaveToken(args []string) (savedToken string, err error) {
	data := []byte("token=\"" + args[0] + "\"\n")

	path := filepath.Join("./", ".env")

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return "", errors.New("no file fould ====>")
	}
	fmt.Printf("token save successfully: %s\n", args[0])
	return args[0], nil

}
