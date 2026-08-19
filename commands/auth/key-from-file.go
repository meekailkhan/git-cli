package auth

import (
	"fmt"
	"io/fs"
	"os"
)

func SaveKeyFromFile(dirPath string) {
	if !fs.ValidPath(dirPath) {
		fmt.Printf("Invalid directory\n")
	}
	data, err := os.ReadFile(dirPath)

	if err != nil {
		fmt.Printf("could not open the file %v\n", data)
	}

	os.WriteFile("private-key-test.pem", data, 440)

}
