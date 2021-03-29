package cmd

import (
	"fmt"

	"github.com/Frug/goqu-table-generator/pkg"
	"github.com/spf13/cobra"
)

// generateCmd represents the setup command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Start the http generater.",
	Long:  `Starts the http generater listening on the host specified in config.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := pkg.Generate(conf)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
