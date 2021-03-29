package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// powerlvlCmd represents the setup command
var powerlvlCmd = &cobra.Command{
	Use:   "powerlvl",
	Short: "Check goku's power level.",
	Long:  `What does the scouter say about his power level?`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Estimated power level: 9000+\n")
	},
}

func init() {
	rootCmd.AddCommand(powerlvlCmd)
}
