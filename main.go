package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirrobot01/goplex_debrid/config"
	"github.com/sirrobot01/goplex_debrid/schemas"
	"github.com/sirrobot01/goplex_debrid/server"
)

func main() {

	app := fiber.New(fiber.Config{
		UnescapePath: true,
	})

	app.Use(recover.New())

	app.Use(logger.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(server.GetServerContext)

	app.Use(server.UpdateHeader)
	config.InitConf()

	for i, s := range config.SERVERS {
		go func(i int, s schemas.Server) {
			token := server.Register(&s)
			config.SERVERS[i].Token = token
		}(i, s)
	}

	server.SetRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
