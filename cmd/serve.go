//go:build !production

package cmd

import (
	"fmt"

	"github.com/Binaretech/classroom-main/server"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Fatalln(server.App().Listen(fmt.Sprintf(":%s", viper.GetString("port"))))
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
