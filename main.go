package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ghc-tdd/spike/issues"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Missing repo. Pass in something like `ghc-tdd/spike`!")
		return
	}

	repo := os.Args[1]

	httpClient := &http.Client{}

	service := issues.NewService(repo, httpClient)

	issues, err := service.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	for _, issue := range issues {
		fmt.Printf("#%d: %s\n", issue.Number, issue.Title)
	}
}
