package add

import (
	"ed-henrique/the-blockchain-bar/database"
	"ed-henrique/the-blockchain-bar/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	flagFrom  = "from"
	flagTo    = "to"
	flagValue = "value"
)

func TxAddCmd() *cobra.Command {
	var txAddCmd = &cobra.Command{
		Use:   "add",
		Short: "Adds new TX to database.",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.ValidateRequiredFlags()
			if err != nil {
				utils.Err(err)
				return
			}

			from, _ := cmd.Flags().GetString(flagFrom)
			to, _ := cmd.Flags().GetString(flagTo)
			value, _ := cmd.Flags().GetUint(flagValue)

			fromAcc := database.NewAccount(from)
			toAcc := database.NewAccount(to)

			tx := database.NewTx(fromAcc, toAcc, value, "")

			state, err := database.NewStateFromDisk()
			if err != nil {
				utils.Err(err)
				return
			}
			defer state.Close()

			err = state.Add(tx)
			if err != nil {
				utils.Err(err)
				return
			}

			err = state.Persist()
			if err != nil {
				utils.Err(err)
				return
			}

			fmt.Println("TX successfully added to the ledger")
		},
	}

	txAddCmd.Flags().String(flagFrom, "", "From what account to send tokens")
	txAddCmd.Flags().String(flagTo, "", "To what account to send tokens")
	txAddCmd.Flags().Uint(flagValue, 0, "How many tokens to send")
	txAddCmd.MarkFlagRequired(flagFrom)
	txAddCmd.MarkFlagRequired(flagTo)
	txAddCmd.MarkFlagRequired(flagValue)

	return txAddCmd
}
