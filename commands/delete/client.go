package removerepo

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v89/github"
)

func NewDeleteClient() (client *github.Client, err error) {
	path := filepath.Join("./", "private-key.pem")
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("private key does not exist please provide private key")
	}
	itr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, 4516499, 151941873, path)
	if err != nil {
		return nil, err
	}
	client, err = github.NewClient(github.WithTransport(itr))
	if err != nil {
		return nil, err
	}
	return client, nil

}
