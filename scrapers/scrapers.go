package scrapers

import (
	"github.com/sirrobot01/goplex_debrid/schemas"
	"sync"
)

type Scraper struct {
	Type string
	Url  string
	Lock sync.Mutex
}

type ScraperInterface interface {
	Search(query string) (string, error)
	Scrape(mediaType string, imdb string, season []int, episode string) ([]schemas.Release, error)
}
