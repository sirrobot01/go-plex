package config

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sirrobot01/goplex_debrid/common"
	"github.com/sirrobot01/goplex_debrid/schemas"
	"os"
)

type Config struct {
	PublicUrl           string   `json:"public_url"`
	PlexEmail           string   `json:"plex_email"`
	PlexPassword        string   `json:"plex_password"`
	PlexLibraryURL      string   `json:"plex_library_url"`
	PlexLibraryToken    string   `json:"plex_library_token"`
	PlexPartialRefresh  bool     `json:"plex_partial_refresh"`
	PlexRefreshSections []string `json:"plex_refresh_sections"`
	PlexRefreshDelay    int      `json:"plex_refresh_delay"`
	Versions            []struct {
		Name    string            `json:"name"`
		Results int               `json:"results"`
		Rules   []string          `json:"rules"`
		Filters map[string]string `json:"filters"`
	} `json:"versions"`
	RealdebridAPIKey  string   `json:"realdebrid_api_key"`
	TorrentioManifest string   `json:"torrentio_manifest"`
	Sorting           []string `json:"sorting"`
}

func GetConfig() *Config {
	// Load config from file
	jsonData, err := os.ReadFile("settings.json")
	if err != nil {
		log.Errorf("An error occured reading config file")
		panic(err)
	}
	var config *Config
	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		log.Errorf("An error occured parsing config file")
		panic(err)
	}

	return config
}

var (
	CONFIG  *Config
	SESSION common.Request
	SERVERS []schemas.Server
)

func InitConf() {
	CONFIG = GetConfig()
	SESSION = common.NewSession(10, 3, 5, []int{500, 502, 503, 504})

	servers := make([]schemas.Server, 0)
	counter := 1

	for _, version := range CONFIG.Versions {
		server := schemas.NewServer(counter, version.Name, version.Rules, version.Filters, version.Results, CONFIG.PublicUrl)
		servers = append(servers, server)
		counter += 1
	}

	SERVERS = servers
}
