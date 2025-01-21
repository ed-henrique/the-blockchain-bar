package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Versioning
const (
	Major  = "0"
	Minor  = "1"
	Fix    = "0"
	Verbal = "TX Add & Balances List"
)

var tbbCmd = &cobra.Command{
	Version: fmt.Sprintf("%s.%s.%s-beta %s", Major, Minor, Fix, Verbal),
	Use:     "tbb",
	Short:   "The Blockchain Bar CLI",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func main() {
	if err := tbbCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
