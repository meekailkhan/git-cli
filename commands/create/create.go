package create

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/google/go-github/v89/github"
)

type Repository struct {
	Name        *string
	Description *string
	AutoInit    *bool
	Visibility  *string
}

func NewRepository(args []string, ctx context.Context) (htmlUrl, sshUrl string, err error) {
	createCommand := flag.NewFlagSet("create", flag.ExitOnError)
	name := createCommand.String("name", "", "Enter the repo name")
	description := createCommand.String("description", "Initialize Repository", "Message for initialize repo")
	visibility := createCommand.String("visibility", "public", "visibility for your repo default:public")
	autoInit := createCommand.Bool("auto_init", false, "Auto initilize your repo or not (default:false)")
	err = createCommand.Parse(args)
	if err != nil {
		return "", "", err
	}
	if *name == "" {
		return "", "", errors.New("Repo name is require for creating repository use '-name' flag to pass name")
	}

	client, err := NewClient()
	if err != nil {
		return "", "", err
	}
	repo := &github.Repository{
		Name:        github.Ptr(*name),
		AutoInit:    github.Ptr(*autoInit),
		Description: github.Ptr(*description),
		Visibility:  github.Ptr(*visibility),
	}
	gitRepo, _, err := client.Repositories.Create(ctx, "", repo)
	if err != nil {
		return "", "", err
	}
	return gitRepo.GetHTMLURL(), gitRepo.GetSSHURL(), nil

}

func CreateRepository(args []string, ctx context.Context) {
	htmlUrl, sshUrl, err := NewRepository(args, ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Github HTML url: %s\n", htmlUrl)
	fmt.Printf("Github SSH url: %s\n", sshUrl)
}
