package list

import (
	"ed-henrique/the-blockchain-bar/database"
	"ed-henrique/the-blockchain-bar/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var BalancesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all balances.",
	Run: func(cmd *cobra.Command, args []string) {
		state, err := database.NewStateFromDisk()
		if err != nil {
			utils.Err(err)
			return
		}
		defer state.Close()

		fmt.Println("---------------------")
		fmt.Println("| Accounts Balances |")
		fmt.Println("---------------------")
		fmt.Println("")

		for account, balance := range state.Balances {
			fmt.Printf("%s: %d\n", account, balance)
		}
	},
}
