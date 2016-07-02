package main

import (
	"fmt"

	"github.com/google/go-github/github"
)

func GetOcto() {
	client := github.NewClient(nil)
	cat, _, err := client.Octocat("")
	if err != nil {
		fmt.Printf("Error on Octocatting: %s\n", err)
	} else {
		fmt.Printf("%s", cat)
	}
}

func main() {
	GetOcto()
}
