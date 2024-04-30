package server

import (
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {

	app.Get("/:server/media/providers/", ProvidersHandler)
	app.Get("/media/providers/", ProvidersHandler)

	app.Get("/:server/library/metadata/:guid/", MetadataHandler)

	app.Get("/:server/library/all/", AvailabilityHandler)
	app.Get("/library/all/", AvailabilityHandler)

	app.Get("/:server/hubs/search/", SearchHandler)
	app.Get("/hubs/search/", SearchHandler)

	app.Get("/:server/download/:id/:num", DownloadHandler)
	app.Get("/download/:id/:num", DownloadHandler)

	app.Get("/:server/system/agents/", AgentsHandler)
	app.Get("/system/agents/", AgentsHandler)

	app.Get("/:server/photo/*/transcode/", TranscodeHandler)
	app.Get("/photo/*/transcode/", TranscodeHandler)

	app.Get("/:server/activities/", Activities)
	app.Get("/activities/", Activities)

	prefsPaths := make([]string, 0)
	prefsPaths = append(prefsPaths, "/library/sections/2/prefs")
	prefsPaths = append(prefsPaths, "/*/prefs")
	prefsPaths = append(prefsPaths, "/:server?/library/sections/2/prefs")
	prefsPaths = append(prefsPaths, "/library/sections/2/prefs")
	prefsPaths = append(prefsPaths, "/:server?/*/prefs")

	for _, url := range prefsPaths {
		app.Get(url, PrefsHandler)
	}

	extraUrls := make([]string, 0)
	extraUrls = append(extraUrls, "/*/websockets/notifications")
	extraUrls = append(extraUrls, "/updater/status")
	extraUrls = append(extraUrls, "/updater/check")
	extraUrls = append(extraUrls, "/accounts/1")
	extraUrls = append(extraUrls, "/myplex/account")
	extraUrls = append(extraUrls, "/:server?/*/websockets/notifications")
	extraUrls = append(extraUrls, "/:server/updater/status")
	extraUrls = append(extraUrls, "/:server/updater/check")
	extraUrls = append(extraUrls, "/:server/accounts/1")
	extraUrls = append(extraUrls, "/:server/myplex/account")

	for _, url := range extraUrls {
		app.Get(url, WildcardHandler)
	}

	app.Options("/:server/updater/check", WildcardHandler)
	app.Get("/:server/status/sessions", WildcardHandler)
}
