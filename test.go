package main

import "github.com/HPI-BP2015H/go-travis/travis"

func test() error {
	travisToken := travis.myTravisToken()

	travisClient := travis.NewClient(travisToken, nil)

	repos := travisClient.Repository.List()
	println("These are your current Repositories:")

	println(len(repos))
	//repo0 := repos[0]
	//println(*repo0.Name)

	return nil
}
