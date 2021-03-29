package cmd

import (
	"fmt"
	"os"

	"github.com/Frug/goqu-table-generator/config"
	"github.com/spf13/cobra"
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

	rootCmd.PersistentFlags().IntP("port", "p", 5432, "database host port number")
	rootCmd.PersistentFlags().StringP("host", "a", "", "database host address")
	rootCmd.PersistentFlags().StringP("name", "n", "public", "database name")
	rootCmd.PersistentFlags().StringP("user", "u", "root", "database connection username")
	rootCmd.PersistentFlags().String("pass", "", "database connection password")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable verbose output")

	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
// Flags override environment variables.
func initConfig() {
	v := viper.New()
	v.SetConfigName("config")

	// map flags to the keys found in config/config.go
	v.BindPFlag("DBPort", rootCmd.PersistentFlags().Lookup("port"))
	v.BindPFlag("DBHost", rootCmd.PersistentFlags().Lookup("host"))
	v.BindPFlag("DBUser", rootCmd.PersistentFlags().Lookup("user"))
	v.BindPFlag("DBPass", rootCmd.PersistentFlags().Lookup("pass"))
	v.BindPFlag("DBName", rootCmd.PersistentFlags().Lookup("name"))

	var err error
	conf, err = config.NewFromEnv(v)

	if err != nil {
		fmt.Printf("Failed to read configuration: %s.\n", err)
		os.Exit(1)
	}
}
