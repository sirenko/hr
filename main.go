package main

import (
	"flag"
	"fmt"
	"os"
)

type hr struct {
	BodyHTML string `json:"body_html"`
}

var (
	contest       = flag.String("contest", "master", "the contest containing the challenge")
	debug         = flag.Bool("debug", false, "debug the chatter")
	overwriteMain = flag.Bool("-m", false, "allow overwriting main")
)

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Fprintln(os.Stderr, "Missing challenge name (as first argument)")
		os.Exit(1)
	}

	if flag.Args()[0] == "submit" {
		if len(flag.Args()) < 2 {
			fmt.Fprintln(os.Stderr, "Missing challenge name (as first argument)")
			os.Exit(1)
		}

		f, err := os.OpenFile("./main.go", os.O_RDONLY, 0640)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()

		err = submit(flag.Args()[1], f)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}

	err := get(*contest, flag.Args()[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
