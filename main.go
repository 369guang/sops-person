package main

import (
	"github.com/gofiber/fiber/v2"
	"person/core/logs"

	"github.com/urfave/cli"
	"log"
	"os"
	"person/core"
	"person/settings"
	"sort"
)

var app *cli.App

func init() {
	app = cli.NewApp()
	app.Name = "运维平台"
	app.Usage = "开源项目"
	app.Version = "0.0.1"
}
func main() {
	core.VIPER = settings.SetConfig()
	core.LOGGER = logs.Loggers()
	core.DATABASE = settings.InitGorm()

	app.Commands = []cli.Command{
		{
			Name:    "app",
			Aliases: []string{"run"},
			Usage:   "run app",
			Action: func(c *cli.Context) error {
				// Fiber instance
				apps := fiber.New()
				settings.LoadMiddleware(apps)
				settings.LoadRoutes(apps)
				// Start server
				log.Fatal(apps.Listen(":3000"))
				return nil
			},
		},
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "migrate database",
			Action: func(c *cli.Context) error {
				settings.MigrateTable(core.DATABASE)
				return nil
			},
		},
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	_ = app.Run(os.Args)
}
