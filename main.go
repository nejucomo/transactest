package main

import (
	"github.com/nejucomo/transactest/run"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		run.Path(os.Args[1])
	} else {
		run.Reader(os.Stdin)
	}
}
