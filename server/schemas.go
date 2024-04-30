package server

type PrefSetting struct {
	ID       string `xml:"id,attr" json:"id"`
	Label    string `xml:"label,attr" json:"label"`
	Summary  string `xml:"summary,attr" json:"summary"`
	Type     string `xml:"type,attr" json:"type"`
	Default  string `xml:"default,attr" json:"default"`
	Value    string `xml:"value,attr" json:"value"`
	Hidden   bool   `xml:"hidden,attr" json:"hidden"`
	Advanced bool   `xml:"advanced,attr" json:"advanced"`
	Group    string `xml:"group,attr" json:"group"`
}

type PrefMediaContainer struct {
	Size     int           `xml:"size,attr" json:"size"`
	Settings []PrefSetting `xml:"Setting" json:"Setting"`
}

type PrefMediaResponse struct {
	MediaContainer PrefMediaContainer `xml:"MediaContainer" json:"MediaContainer"`
}
