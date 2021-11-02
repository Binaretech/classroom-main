package cmd

import (
	"github.com/Binaretech/classroom-main/internal/db"
	"github.com/Binaretech/classroom-main/internal/db/seeder"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate:fresh",
	Short: "Recreate the entire database",
	Run: func(cmd *cobra.Command, args []string) {
		if err := db.Drop(); err != nil {
			logrus.Error(err.Error())
			return
		}

		db.Migrate()

		if viper.GetBool("seed") {
			seeder.Run()
		}
	},
}

func init() {
	migrateCmd.PersistentFlags().BoolP("seed", "s", false, "run seeder after migrate")
	viper.BindPFlag("seed", migrateCmd.PersistentFlags().Lookup("seed"))
	rootCmd.AddCommand(migrateCmd)
}
