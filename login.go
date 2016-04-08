package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func login() error {

	//this needs the myToken.go file. This should be changed when its ready to ship...
	token := myToken()
	if token == "" {
		println("I need your github Access Token to log you in. Please paste it here:")
		fmt.Scanln(&token)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
		//&oauth2.Token{AccessToken: myToken()},
		//&oauth2.Token{AccessToken: "... your access token ..."},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	github := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := github.Repositories.List("", nil)
	user, _, err := github.Users.Get("")
	if err != nil {
		return err
	}
	color.Green("Success! You are now logged into the account " + *(user.Login) + ".\n")
	println("These are your current Repositories:")
	for _, repo := range repos {
		color.Blue("     " + *repo.FullName)
	}

	return nil
}
