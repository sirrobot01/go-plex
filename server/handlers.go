package server

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirrobot01/goplex_debrid/common"
	"github.com/sirrobot01/goplex_debrid/config"
	"github.com/sirrobot01/goplex_debrid/debrid"
	"github.com/sirrobot01/goplex_debrid/schemas"
	"github.com/sirrobot01/goplex_debrid/scrapers"
	"github.com/webtor-io/go-jackett"
	"net/url"
	"strconv"
	"strings"
)

var (
	DB = make(map[string][]schemas.Release)
)

func ProvidersHandler(c *fiber.Ctx) error {
	server := c.Locals("server").(*schemas.Server)
	path, _ := url.QueryUnescape(c.OriginalURL())
	if strings.Contains(path, "includePreferences=1") {
		return c.XML(GetMobileProviderData(server))
	}
	return ProcessResponse(c, GetProviderData(server))

}

func MetadataHandler(c *fiber.Ctx) error {
	server := c.Locals("server").(*schemas.Server)
	return ProcessResponse(c, GetMetadataData(server))
}

func AvailabilityHandler(c *fiber.Ctx) error {
	server := c.Locals("server").(*schemas.Server)
	servers := c.Locals("servers").([]schemas.Server)

	path, _ := url.QueryUnescape(c.OriginalURL())
	guid := common.RegexSearch(`guid=([^&]*)`, path)[0]
	mediaType, seasons, episode, imdb := Identify(server, path)
	if imdb == "" {
		return c.JSON(fiber.Map{"error": "not found"})
	}
	scraper := scrapers.Scraper{
		Type: "torrentio",
		Url:  fmt.Sprintf("https://torrentio.strem.fun/%s/stream/", config.CONFIG.TorrentioManifest),
	}
	releases := scraper.Scrape(mediaType, imdb, seasons, episode)
	releases = debrid.CheckCached(releases, mediaType, seasons, episode)
	releases = debrid.FilterReleases(releases, mediaType, seasons, episode)

	showTitle := false
	if len(servers) == 1 {
		showTitle = true
		// A single server, try to sort and filter releases
	}
	metadata := make([]map[string]any, 0)
	for _, s := range servers {
		_uuid := uuid.NewString()
		// sort release by resolution using server's preference
		tempRelease := debrid.SortReleases(releases, s)
		if len(tempRelease) < 1 {
			tempRelease = releases
		}
		tempRelease = tempRelease[:s.Results]
		DB[_uuid] = tempRelease // Store releases in a temp db, expires on restarts
		for i, release := range tempRelease {
			var media []map[string]string
			res := strconv.Itoa(release.Resolution)
			if showTitle {
				res = release.Title
			}
			media = append(media, map[string]string{"videoResolution": res})
			meta := map[string]any{
				"ratingKey":           "2785",
				"key":                 fmt.Sprintf("/download/%s/%d", url.PathEscape(_uuid), i),
				"librarySectionID":    2,
				"librarySectionKey":   "/library/sections/2",
				"guid":                guid,
				"librarySectionTitle": s.Name,
				"Media":               media,
			}
			metadata = append(metadata, meta)
		}
	}
	results := fiber.Map{
		"MediaContainer": fiber.Map{
			"size":            1,
			"allowSync":       false,
			"identifier":      "com.plexapp.plugins.library",
			"mediaTagPrefix":  "/system/bundle/media/flags/",
			"mediaTagVersion": 1655122614,
			"Metadata":        metadata,
		},
	}

	res := ProcessResponse(c, results)
	return res

}

func SearchHandler(c *fiber.Ctx) error {
	query := c.Query("query")
	_ = c.Locals("server").(*schemas.Server)
	_ = c.Locals("servers").([]schemas.Server)
	if query == "" {
		return c.JSON(fiber.Map{"error": "no query"})
	}

	ctx := context.Background()
	j := jackett.NewJackett(&jackett.Settings{
		ApiURL: "http://localhost:9117/",
		ApiKey: "il3qviy5b5qoia86lj5ituxfaijadiky",
	})
	resp, err := j.Fetch(ctx, &jackett.FetchRequest{
		Categories: []uint{5000},
		Query:      query,
	})
	if err != nil {
		panic(err)
	}
	return c.JSON(resp.Results)

}

func DownloadHandler(c *fiber.Ctx) error {
	uid := c.Params("id")
	num := c.Params("num")
	releases := DB[uid]
	if len(releases) < 1 {
		return c.JSON(fiber.Map{"error": "no releases"})
	}
	i, _ := strconv.Atoi(num)
	release := releases[i]
	debrid.Download(&release)
	return c.SendString("Hello, World!")
}

func AgentsHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func TranscodeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func PrefsHandler(c *fiber.Ctx) error {
	server := c.Locals("server").(*schemas.Server)
	return ProcessResponse(c, PreferenceData(server))

}

func WildcardHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Activities(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
