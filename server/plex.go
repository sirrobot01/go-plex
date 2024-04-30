package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sirrobot01/goplex_debrid/common"
	"github.com/sirrobot01/goplex_debrid/config"
	"github.com/sirrobot01/goplex_debrid/schemas"
	"net/url"
	"strings"
)

var (
	conf    = config.CONFIG
	session = config.SESSION
)

func RefreshLibrary(library *schemas.Library, release schemas.Release) {
	// Refresh library
	sectionsUrl := fmt.Sprintf("%s/library/sections/?X-Plex-Token=%s", conf.PlexLibraryURL, conf.PlexLibraryToken)
	headers := map[string]string{
		"Accept": "application/json",
	}
	response, err := session.Get(sectionsUrl, headers)
	if err != nil {
		return
	}
	if !conf.PlexPartialRefresh {
		RefreshAllLibrary(library)
	} else {
		RefreshPartialLibrary(library, response, release)
	}

}

func RefreshAllLibrary(library *schemas.Library) {
	// Refresh all section's folders
	headers := map[string]string{
		"Accept": "application/json",
	}
	_, err := session.Get(library.Url, headers)
	if err != nil {
		return
	}
}

func RefreshPartialLibrary(library *schemas.Library, response map[string]any, release schemas.Release) {

	refreshUrls := make([]string, 0)

	// Refresh partial sections
	directories := response["MediaContainer"].(map[string]any)["Directory"].([]map[string]any)
	for _, directory := range directories {
		if directory["type"] == release.Type && directory["key"] == library.Id {
			locations := directory["Location"].([]map[string]any)
			for _, location := range locations {
				quotePath := url.PathEscape(location["path"].(string))
				refreshUrls = append(refreshUrls, fmt.Sprintf("%s&path=%s", library.Url, quotePath))
			}
		}
	}

	for _, refreshUrl := range refreshUrls {
		go func(url string) {
			headers := map[string]string{
				"Accept": "application/json",
			}
			_, err := session.Get(url, headers)
			if err != nil {
				return
			}
		}(refreshUrl)
	}
}

func Register(server *schemas.Server) string {
	// Register server
	conf := config.CONFIG
	errStr := fmt.Sprintf("Error Registering server: %s", server.Name)
	token, err := Authenticate(server, conf.PlexEmail, conf.PlexPassword)
	if err != nil {
		log.Error(errStr)
	}
	// Register server
	headers := make(map[string]string)
	headers["X-Plex-Token"] = token
	headers["X-Plex-Device-Name"] = server.Name
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"

	_url := fmt.Sprintf(
		"https://plex.tv/devices/%s?Connection[][uri]=%s&httpsEnabled=0&httpsRequired=0&dnsRebindingProtection=0&natLoopbackSupported=1",
		server.Hash, server.Url,
	)
	_, err = session.Put(_url, nil, headers)
	log.Infof("Server: %s successfully added", server.Url)
	return token
}

func Authenticate(server *schemas.Server, username string, password string) (string, error) {
	// Authenticate

	formdata := url.Values{
		"user[login]":    {username},
		"user[password]": {password},
	}
	// Authenticate
	post, err := session.Post("https://plex.tv/users/sign_in.json", formdata, server.Headers)
	if err != nil {
		return "", err
	}
	user := post["user"].(map[string]any)
	return user["authToken"].(string), nil
}

func Provider(server *schemas.Server, isMobile bool) fiber.Map {
	if isMobile {
		GetMobileProviderData(server)
	}
	return GetProviderData(server)
}

func Identify(server *schemas.Server, path string) (string, []int, int, string) {
	guid := common.RegexSearch(`guid=([^&]*)`, path)[0]
	if guid == "" {
		return "", nil, 0, ""
	}
	// Identify
	mediaType := "show" // Shows each season
	if strings.Contains(path, "movie") || strings.Contains(path, "type=1") {
		// Movie View
		mediaType = "movie"
	} else if strings.Contains(path, "episode") || strings.Contains(path, "type=4") {
		// Episode View
		mediaType = "episode"
	} else if strings.Contains(path, "season") || strings.Contains(path, "type=3") {
		// This is a Season View(It shows Episodes)
		mediaType = "show"
	}
	guids := strings.Split(guid, "/")
	idx := len(guids) - 1
	guid = guids[idx]

	s := common.RegexSearch(`season=([^&]*)`, path)
	episode := 0
	var seasons []int
	resp := GetMetadata(server, guid)
	if resp == nil {
		return "", nil, 0, ""
	}
	metadata, _ := resp["MediaContainer"].(map[string]any)["Metadata"].([]any)[0].(map[string]any)

	if len(s) < 1 && mediaType == "show" {
		cdCount := metadata["childCount"].(float64)
		childCount := int(cdCount)
		seasons = make([]int, childCount)
		for i := 0; i < childCount; i++ {
			seasons[i] = i + 1
		}
	} else if mediaType == "episode" {
		seasons = []int{metadata["parentIndex"].(int)}
		episode = metadata["index"].(int)
		guid = metadata["grandparentGuid"].(string)
		guids = strings.Split(guid, "/")
		idx = len(guids) - 1
		guid = guids[idx]
		resp = GetMetadata(server, guid)
		mediaType = "show"
	}
	// Get IMDB ID
	imdb := ""
	entries := metadata["Guid"].([]any)
	for _, e := range entries {
		entry := e.(map[string]any)
		if strings.HasPrefix(entry["id"].(string), "imdb://") {
			imdbs := strings.Split(entry["id"].(string), "/")
			idx := len(imdbs) - 1
			imdb = imdbs[idx]

		}
	}
	return mediaType, seasons, episode, imdb

}

func GetMetadata(server *schemas.Server, guid string) map[string]any {
	headers := map[string]string{
		"Accept": "application/json",
	}

	response, err := session.Get(fmt.Sprintf("https://metadata.provider.plex.tv/library/metadata/%s?X-Plex-Token=%s", guid, server.Token), headers)
	if err != nil {
		return nil
	}
	return response
}

func GetServer(serverId int) *schemas.Server {
	for _, server := range config.SERVERS {
		if server.Id == serverId {
			return &server
		}
	}
	return nil
}
