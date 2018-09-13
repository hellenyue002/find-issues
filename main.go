package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ghc-tdd/find-issues/issues"
	flags "github.com/jessevdk/go-flags"
)

type opts struct {
	Label string `long:"label" description:"Filter issues based on their labels."`
}

func main() {
	var options opts
	parser := flags.NewParser(&options, flags.HelpFlag|flags.PrintErrors)
	remainingArgs, err := parser.ParseArgs(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	var repo string
	if len(remainingArgs) == 0 {
		fmt.Fprintf(os.Stderr, "Missing repo. Pass in something like `ghc-tdd/find-issues`!")
		os.Exit(1)
	}

	repo = remainingArgs[1]

	httpClient := &http.Client{}

	service := issues.NewService(repo, httpClient)

	issues, err := service.Get(options.Label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	for _, issue := range issues {
		fmt.Printf("#%d: %s\n", issue.Number, issue.Title)
	}
}
