package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ghc-tdd/find-issues/issues"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Missing repo. Pass in something like `ghc-tdd/find-issues`!")
		return
	}

	repo := os.Args[1]

	labelFilter := ""
	if len(os.Args) > 2 {
		if os.Args[2] == "--filter" {
			if len(os.Args) > 3 {
				labelFilter = os.Args[3]
			} else {
				fmt.Println("--filter flag passed with no argument")
				return
			}
		}
	}

	httpClient := &http.Client{}

	service := issues.NewService(repo, httpClient)

	issues, err := service.Get(labelFilter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	for _, issue := range issues {
		fmt.Printf("#%d: %s\n", issue.Number, issue.Title)
	}
}
