package cli

import (
	"flag"
	"os"

	"github.com/raymond-design/task/cli"
)

var (
	apiFlag  = flag.String("api", "http://localhost:8080", "API URL")
	helpFlag = flag.Bool("help", false, "Show help")
)

func main() {
	flag.Parse()
	s := cli.CreateSwitch(*apiFlag)

	if *helpFlag || len(os.Args) == 1 {
		s.Help()
		return
	}
}
