package cmd

import (
	"github.com/Binaretech/classroom-main/db"
	"github.com/Binaretech/classroom-main/db/seeder"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate:fresh",
	Short: "Recreate the entire database",
	Run: func(cmd *cobra.Command, args []string) {
		conn, _ := db.Connect()

		if err := db.Drop(conn); err != nil {
			logrus.Error(err.Error())
			return
		}

		db.Migrate(conn)

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
