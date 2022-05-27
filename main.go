package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
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
	app.Name = "embracing"
	app.Usage = "又是一个吹牛逼的项目"
	app.Version = "0.0.1"
}
func main() {
	core.VIPER = settings.SetConfig()

	app.Commands = []cli.Command{
		{
			Name:    "app",
			Aliases: []string{"run"},
			Usage:   "run app",
			Action: func(c *cli.Context) error {
				// Fiber instance
				apps := fiber.New()

				// middleware
				apps.Use(logger.New())  // 日志
				apps.Use(recover.New()) // recover
				apps.Use(cors.New())    // 跨域
				apps.Use(pprof.New())   // 开启pprof分析

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
				//core.Migrate()
				return nil
			},
		},
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	_ = app.Run(os.Args)
}
