# travis-cli

To use you will need to add a go file with the following code:

```
package main

func myGithubToken() string {
	return "yourGithubToken"
}

func myTravisToken() string {
	return "yourTravisToken"
}

```

and replace the strings with your personal token 
