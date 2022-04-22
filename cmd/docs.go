//go:build !production

package cmd

import (
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/swaggo/echo-swagger"

	_ "github.com/Binaretech/classroom-main/docs"
)

var docsServeCmd = &cobra.Command{
	Use:   "docs:serve",
	Short: "Serve the documentation",
	Run: func(cmd *cobra.Command, args []string) {
		app := echo.New()

		app.GET("/swagger/*", echoSwagger.WrapHandler)

		log.Fatal(app.Start(":8080"))

	},
}

var docsGenerateCmd = &cobra.Command{
	Use:   "docs:generate",
	Short: "Generate the documentation",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := exec.LookPath("swag"); err != nil {
			exec.Command("go", "install", "github.com/swaggo/swag/cmd/swag@latest").Run()
		}

		dir, _ := os.Getwd()

		source := path.Join(dir, "cmd", "service")
		output := path.Join(dir, "internal", "docs")

		exec.Command("swag", "init", "-g", source, "-o", output).Run()
	},
}

func init() {
	rootCmd.AddCommand(docsServeCmd)
	rootCmd.AddCommand(docsGenerateCmd)
}
