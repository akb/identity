package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/akb/go-cli"

	"github.com/akb/identify/cmd/config"
	"github.com/akb/identify/internal/token"
)

type deleteTokenCommand struct {
	id      *string
	tokenID *string
}

func (deleteTokenCommand) Help() {
	fmt.Println("identify - authentication and authorization service")
	fmt.Println("")
	fmt.Println("Usage: identify delete token <id>")
	fmt.Println("")
	fmt.Println("Delete a token.")
}

func (c *deleteTokenCommand) Flags(f *flag.FlagSet) {
	c.id = f.String("id", "", "identity to authenticate")
	c.tokenID = f.String("token-id", "", "id of token to delete")
}

func (c deleteTokenCommand) Command(ctx context.Context) int {
	tokenSecret, err := config.GetTokenSecret()
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	tokenDBPath, err := config.GetTokenDBPath()
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	tokenStore, err := token.NewLocalStore(tokenDBPath, tokenSecret)
	if err != nil {
		fmt.Println("An error occurred while opening token database file:")
		fmt.Println(err.Error())
		return 1
	}
	defer tokenStore.Close()

	var id, tokenID string
	if len(*c.id) > 0 && len(*c.tokenID) > 0 {
		id = *c.id
		tokenID = *c.tokenID

	} else if len(*c.id) > 0 || len(*c.tokenID) > 0 {
		fmt.Println("both an identity and a token must be specified")
		return 1

	} else {
		credsPath, err := config.GetCredentialsPath()
		if err != nil {
			println("error getting credentials path")
			fmt.Println(err)
			return 1
		}

		credsJSON, err := ioutil.ReadFile(credsPath)
		if err != nil {
			fmt.Println(err)
			return 1
		}

		var creds Credentials
		err = json.Unmarshal(credsJSON, &creds)
		if err != nil {
			println("error unmarshaling creds json")
			fmt.Println(err)
			return 1
		}

		t, err := token.Parse(creds.Access, tokenSecret)
		if err != nil {
			println("error parsing token")
			fmt.Println(err)
			return 1
		}

		id = t.Identity()
		tokenID = t.ID()
	}

	if err := tokenStore.Delete(id, tokenID); err != nil {
		fmt.Println(err.Error())
		return 1
	}

	fmt.Println("Token successfully deleted")
	return 0
}

func (deleteTokenCommand) Subcommands() cli.CLI {
	return nil
}
