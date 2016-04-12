package main

import (
	"reflect"

	"github.com/HPI-BP2015H/go-travis/travis"
)

func test() error {
	travisToken := myTravisToken()

	travisClient := travis.NewClient(travisToken, nil)

	repos := travisClient.Repository.List()
	println("These are your current Repositories:")
	println(reflect.TypeOf(repos).String())
	println(len(repos))
	println(reflect.TypeOf(repos[0].Name).String())
	//repo0 := repos[0]
	println(*repos[0].ID)
	return nil
}
