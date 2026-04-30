package main

import (
	"fmt"
	"os"

	"github.com/ead8/forge/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
