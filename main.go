package main

import (
	"fmt"
	"os"

	"hacksim/cmd"
)

func main() {
	// Execute the root command
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
