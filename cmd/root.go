package cmd

import (
	"fmt"
	"os"

	"github.com/Frug/goqu-table-generator/config"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var conf config.Config

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "scouter",
	Short: "Generates structs for use with goqu.",
	Long: `Scouter scans a database and writes corresponding structs for the tables it finds.
These structs can be used with goqu to provide some safety with table and column names.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().Int("port", 5432, "Database host port number")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	v := viper.New()
	v.SetConfigName("config")

	v.BindPFlag("DBPort", rootCmd.PersistentFlags().Lookup("port"))

	var err error
	conf, err = config.NewFromEnv(v)

	if err != nil {
		fmt.Printf("Failed to read configuration: %s.\n", err)
		os.Exit(1)
	}
}

func bindIfNonEmpty(v *viper.Viper, key string, flag *pflag.Flag) {

	v.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))

}
