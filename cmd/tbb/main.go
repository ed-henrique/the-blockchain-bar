package main

import (
	"ed-henrique/the-blockchain-bar/cmd/tbb/balances"
	"ed-henrique/the-blockchain-bar/cmd/tbb/tx"
	"ed-henrique/the-blockchain-bar/utils"
	"fmt"

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

func init() {
	tbbCmd.AddCommand(balances.BalancesCmd)
	tbbCmd.AddCommand(tx.TxCmd)
}

func main() {
	if err := tbbCmd.Execute(); err != nil {
		utils.Err(err)
		return
	}
}
