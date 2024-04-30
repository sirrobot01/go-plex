package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirrobot01/goplex_debrid/config"
	"github.com/sirrobot01/goplex_debrid/schemas"
	"strconv"
)

func ProcessResponse(c *fiber.Ctx, response any) error {
	accept := c.Accepts("json", "xml")
	switch accept {
	case "xml":
		return c.XML(response)
	case "json":
		return c.JSON(response)
	default:
		return c.JSON(response)
	}
}

func GetServerContext(c *fiber.Ctx) error {
	servers := make([]schemas.Server, 0)
	serverId := c.Params("server")
	var server *schemas.Server
	if serverId != "" {
		serverInt, _ := strconv.Atoi(serverId)
		server = GetServer(serverInt)
		if server == nil {
			// server not found, use first server
			server = &config.SERVERS[0]
			servers = config.SERVERS
		} else {
			servers = append(servers, *server)
		}
	}
	if server == nil {
		// No server id
		servers = config.SERVERS
		server = &servers[0]
	}
	c.Locals("server", server)
	c.Locals("servers", servers)

	return c.Next()
}

func UpdateHeader(c *fiber.Ctx) error {
	// Set default headers here
	// Continue stack execution
	err := c.Next()
	c.Set("Vary", "Origin, X-Plex-Token")
	c.Set("X-Plex-Protocol", "1.0")
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("X-Plex-Content-Compressed-Length", c.Get("Content-Length"))

	return err
}
