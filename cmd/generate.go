package cmd

import (
	"github.com/Frug/goqu-table-generator/pkg"
	"github.com/spf13/cobra"
)

// generateCmd represents the setup command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Start the http generater.",
	Long:  `Starts the http generater listening on the host specified in config.`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Generate(conf)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
