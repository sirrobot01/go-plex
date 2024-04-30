package schemas

import (
	"fmt"
	"github.com/sirrobot01/goplex_debrid/common"
)

type ServerVersion struct {
	Name    string   `json:"name"`
	Results int      `json:"results"`
	Rules   []string `json:"rules"`
	Filters []string `json:"filters"`
}

type Server struct {
	Id      int               `json:"id"`
	Name    string            `json:"name"`
	Rules   []string          `json:"rules"`
	Filters map[string]string `json:"filters"`
	Results int               `json:"results"`
	Hash    string            `json:"hash"`
	Url     string            `json:"url"`
	Token   string            `json:"token"`
	Headers map[string]string `json:"headers"`
}

func NewServer(id int, name string, rules []string, filters map[string]string, results int, url string) Server {
	hash := common.GenerateHash(name)
	return Server{
		Id:      id,
		Name:    name,
		Rules:   common.InvertArray(rules), // invert array
		Filters: filters,
		Results: results,
		Url:     fmt.Sprintf("%s/%d", url, id),
		Hash:    hash,
		Headers: map[string]string{
			"X-Plex-Device-Name":       name,
			"X-Plex-Product":           "Plex Media Server",
			"X-Plex-Version":           "1.32.5.7349-8f4248874",
			"X-Plex-Client-Identifier": hash,
			"X-Plex-Platform":          "Windows",
			"X-Plex-Platform-Version":  "10.0 (Build 19045)",
			"X-Plex-Device":            "PC",
			"X-Plex-Model":             "x64-x64",
			"X-Plex-Device-Vendor":     "Microsoft",
			"X-Plex-Provides":          "server",
			"X-Plex-Protocol":          "1.0",
		},
	}
}
