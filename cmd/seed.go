//go:build !production

package cmd

import (
	"github.com/Binaretech/classroom-main/db/seeder"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var seedCmd = &cobra.Command{
	Use:   "db:seed",
	Short: "Fill database with testing data",
	Run: func(cmd *cobra.Command, args []string) {
		seeder.Run()
	},
}

func init() {
	seedCmd.PersistentFlags().StringP("name", "n", "", "Run concrete seeder")
	viper.BindPFlag("name", seedCmd.PersistentFlags().Lookup("name"))
	rootCmd.AddCommand(seedCmd)
}
