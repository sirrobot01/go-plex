package schemas

type File struct {
	Name     string  `json:"name"`
	Size     float64 `json:"size"`
	Id       string  `json:"id"`
	Video    bool    `json:"video"`
	Subtitle bool    `json:"subtitle"`
	Season   int     `json:"season"`
	Episode  int     `json:"episode"`
}

type Version struct {
	Size      float64 `json:"size"`
	Files     []File  `json:"files"`
	Videos    int     `json:"videos"`
	Episodes  int     `json:"episodes"`
	Subtitles int     `json:"subtitles"`
	Seasons   []int   `json:"seasons"`
}

type Release struct {
	Title      string    `json:"title"`
	Languages  []string  `json:"languages"`
	Resolution int       `json:"resolution"`
	Size       float64   `json:"size"`
	Seeders    int       `json:"seeders"`
	Source     string    `json:"source"`
	Magnet     string    `json:"magnet"`
	Hash       string    `json:"hash"`
	Type       string    `json:"type"`
	Cached     []string  `json:"cached"`
	Versions   []Version `json:"versions"`
	Seasons    []int     `json:"seasons"`
	Videos     int       `json:"videos"`
	Episodes   int       `json:"episodes"`
	Download   []string  `json:"download"`
	Files      []File    `json:"files"`
}

type Library struct {
	Id      string `json:"id"`
	Partial bool   `json:"partial"`
	Delay   int    `json:"delay"`
	Url     string `json:"url"`
}
