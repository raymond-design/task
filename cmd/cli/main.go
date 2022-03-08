package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/raymond-design/task/cli"
)

var (
	apiflag  = flag.String("api", "http://localhost:8080", "API URL")
	helpflag = flag.Bool("help", false, "Show help")
)

func main() {
	flag.Parse()
	s := cli.CreateSwitch(*apiflag)

	if *helpflag || len(os.Args) == 1 {
		s.Help()
		return
	}

	err := s.Switch()
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(2)
	}
}
