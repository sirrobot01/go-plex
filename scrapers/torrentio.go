package scrapers

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sirrobot01/goplex_debrid/common"
	"github.com/sirrobot01/goplex_debrid/config"
	"github.com/sirrobot01/goplex_debrid/schemas"
	"strconv"
	"strings"
	"sync"
)

var (
	session = config.SESSION
)

func (scraper *Scraper) Search(query string) (string, []int, int, string) {
	return TorrentioSearch(query)
}

func scrapeSeason(releases *[]schemas.Release, scraper *Scraper, mediaType string, imdb string, season int, episode int) {
	log.Infof("Scraping %s imdb:%s %d %d\n", mediaType, imdb, season, episode)
	extras := ""
	if mediaType == "series" {
		extras = fmt.Sprintf(":%d:%d", season, episode)
	}

	url := fmt.Sprintf("%s%s/%s%s.json", scraper.Url, mediaType, imdb, extras)
	response, err := session.Get(url, nil)
	if err != nil {
		return
	}
	streams := response["streams"].([]any)
	for _, stream := range streams {
		var (
			size       float64
			resolution int
			seeders    int
			source     string
		)
		title := stream.(map[string]any)["title"].(string)
		languages := make([]string, 0)
		languages = append(languages, "EN")
		m := common.RegexFindAll(common.LanguageMatch, title)
		if len(m) > 0 {
			languages = m
		}
		r := common.RegexFindAll(common.ResolutionMatch, title)
		if len(r) > 0 {
			resolution, _ = strconv.Atoi(strings.Replace(r[0], "p", "", 1))
		}

		//if resolution == 0 || !slices.Contains(resolutions, resolution) {
		//	// Skip early
		//	continue
		//}

		s := common.RegexFindAll(common.GBSizeMatch, title)
		if len(s) > 0 {
			size, _ = strconv.ParseFloat(strings.Split(s[0], " ")[1], 64)
		} else {
			s := common.RegexFindAll(common.MBSizeMatch, title)
			if len(s) > 0 {
				_s := strings.Split(s[0], " ")
				if len(_s) < 1 {
					continue
				}
				sz, _ := strconv.ParseFloat(_s[1], 64)

				size = sz / float64(1024)
			}
		}
		hash := stream.(map[string]any)["infoHash"].(string)
		magnet := fmt.Sprintf("magnet:?xt=urn:btih:%s&dn=&tr=", hash)
		hash = strings.ToLower(hash)
		se := common.RegexFindAll(`👤 ([0-9]+)`, title)
		if len(se) > 0 {
			seeders, _ = strconv.Atoi(strings.Replace(se[0], "👤 ", "", 1))
		} else {
			seeders = 0
		}
		so := common.RegexFindAll(common.SourceMatch, title)
		if len(so) > 0 {
			source = strings.Replace(so[0], "⚙️ ", "", 1)
		} else {
			source = "unknown"
		}

		release := schemas.Release{
			Title:      strings.Replace(strings.Split(title, "\n")[0], " ", ".", -1),
			Languages:  languages,
			Resolution: resolution,
			Size:       size,
			Seeders:    seeders,
			Source:     source,
			Magnet:     magnet,
			Hash:       hash,
			Type:       scraper.Type,
		}
		*releases = append(*releases, release)
	}
	log.Infof("Finished Scrapping imdb:%s %d:%d", imdb, season, episode)
}

func (scraper *Scraper) Scrape(mediaType string, imdb string, seasons []int, episode int) []schemas.Release {
	if mediaType == "" || mediaType == "show" {
		mediaType = "series"
	}

	releases := make([]schemas.Release, 0)
	var wg sync.WaitGroup

	if len(seasons) < 1 {
		seasons = append(seasons, 1)
	}
	if episode == 0 {
		episode = 1
	}

	seasons = []int{1} // For now, only scrap like plex_debrid

	wg.Add(len(seasons))
	for _, season := range seasons {
		go func(s int) {
			defer wg.Done()
			scrapeSeason(&releases, scraper, mediaType, imdb, s, episode)
		}(season)
	}
	wg.Wait()

	return releases
}

func TorrentioSearch(query string) (string, []int, int, string) {
	return "", nil, 0, ""
}
