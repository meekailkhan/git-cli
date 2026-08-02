package auth

import (
	"fmt"
	"os"
	"path/filepath"
)

func SaveToken(args []string) (savedToken string, err error) {
	data := []byte("token=\"" + args[0] + "\"\n")

	path := filepath.Join("./", ".env")
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return "", err
	}
	fmt.Printf("token save successfully: %s\n", args[0])
	return args[0], nil

}
