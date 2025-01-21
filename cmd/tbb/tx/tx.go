package tx

import (
	"ed-henrique/the-blockchain-bar/cmd/tbb/tx/add"
	"ed-henrique/the-blockchain-bar/utils"

	"github.com/spf13/cobra"
)

var TxCmd = &cobra.Command{
	Use:   "tx",
	Short: "Interact with txs (add...).",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return utils.IncorrectUsageErr
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	TxCmd.AddCommand(add.TxAddCmd())
}
