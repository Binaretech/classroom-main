package cmd

import (
	"fmt"
	"os"

	"github.com/Binaretech/classroom-main/server"
	"github.com/Binaretech/classroom-main/storage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "classroom",
	Short: "Classroom main server",
	Run: func(cmd *cobra.Command, args []string) {
		storage.OpenStorage()

		logrus.Fatalln(server.App().Listen(fmt.Sprintf(":%s", viper.GetString("port"))))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
