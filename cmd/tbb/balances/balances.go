package balances

import (
	"ed-henrique/the-blockchain-bar/cmd/tbb/balances/list"
	"ed-henrique/the-blockchain-bar/utils"

	"github.com/spf13/cobra"
)

var BalancesCmd = &cobra.Command{
	Use:   "balances",
	Short: "Interact with balances (list...).",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return utils.IncorrectUsageErr
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	BalancesCmd.AddCommand(list.BalancesListCmd)
}
