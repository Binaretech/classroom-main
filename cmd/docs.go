//go:build !production

package cmd

import (
	"log"
	"os"
	"os/exec"
	"path"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	"github.com/spf13/cobra"

	_ "github.com/Binaretech/classroom-main/docs"
)

var docsServeCmd = &cobra.Command{
	Use:   "docs:serve",
	Short: "Serve the documentation",
	Run: func(cmd *cobra.Command, args []string) {
		app := fiber.New()

		app.Get("/*", swagger.HandlerDefault)

		log.Fatal(app.Listen(":8080"))

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
