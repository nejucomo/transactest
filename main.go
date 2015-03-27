package main

import "os"


func main() {
	if len(os.Args) > 1 {
		runPath(os.Args[1])
	} else {
		runReader(os.Stdin)
	}
}
