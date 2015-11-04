package main

import (
  "fmt"
  "github.com/google/go-github/github"
)

func main() {
  // Next step: read up https://godoc.org/github.com/google/go-github/github#NewClient
  client := github.NewClient(nil)
  opt := &github.RepositoryListOptions{Type: "owner", Sort: "updated", Direction: "desc"}
  repos, _, err := client.Repositories.List("mjbrender", opt)
  if err != nil {
    fmt.Printf("Error : %v\n", err)
  } else {
    fmt.Printf("%v\n", github.Stringify(repos))
  }

}
