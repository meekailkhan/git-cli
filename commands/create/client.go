package create

import (
	"errors"
	"os"

	"github.com/google/go-github/v89/github"
)

func NewClient() (client *github.Client, err error) {
	token := os.Getenv("token")
	if token == "" {
		return nil, errors.New("token not provide please set token via --save-token option")
	}
	client, err = github.NewClient(github.WithAuthToken(token))
	if err != nil {
		return nil, err
	}
	return client, nil
}
