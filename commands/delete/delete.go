package removerepo

import (
	"context"
	"flag"
	"fmt"
)

func DeleteRepository(args []string, ctx context.Context) {
	deleteCommandFlag := flag.NewFlagSet("delete", flag.ExitOnError)
	name := deleteCommandFlag.String("name", "", "enter the repo name which wants you delete")
	err := deleteCommandFlag.Parse(args)
	if err != nil {
		fmt.Printf("could not parse argument: %v\n", err)
		return
	}
	if *name == "" {
		fmt.Println("name is required for deleting repository")
		return
	}
	client, err := NewDeleteClient()
	if err != nil {
		fmt.Println("could not initialize client")
		return
	}
	_, err = client.Repositories.Delete(ctx, "meekailkhan", *name)
	if err != nil {
		fmt.Printf("could not delete the repository: %v\n", err)
		return
	}
	fmt.Printf("successfully delete repository: %s\n", *name)
}
