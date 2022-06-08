package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli"
	"log"
	"os"
	"person/core"
	"person/core/logs"
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

	db, _ := core.DATABASE.DB()
	defer db.Close()

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
				if err := apps.Listen(":3000"); err != nil {
					core.LOGGER.Error(err.Error())
					log.Fatal(err)
				}
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
		{
			Name:    "CreateApp",
			Aliases: []string{"c"},
			Usage:   "创建APP",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "CreateAppModelApi",
			Aliases: []string{"c"},
			Usage:   "基于model创建API请求函数",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	_ = app.Run(os.Args)
}
