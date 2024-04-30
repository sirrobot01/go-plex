package common

import "regexp"

var (
	VideoMatch      = "(?i)(\\.)(YUV|WMV|WEBM|VOB|VIV|SVI|ROQ|RMVB|RM|OGV|OGG|NSV|MXF|MTS|M2TS|TS|MPG|MPEG|M2V|MP2|MPE|MPV|MP4|M4P|M4V|MOV|QT|MNG|MKV|FLV|DRC|AVI|ASF|AMV)$"
	SubtitleMatch   = "(?i)(\\.)(SRT|ASS|VTT|SUB|IDX|PGS)$"
	SeasonMatch     = "(?i)(?:season|s)[.\\-_\\s]?(\\d+)"
	EpisodeMatch    = "(?i)(?:episode|e)[.\\-_\\s]?(\\d+)"
	LanguageMatch   = "[\\x{1F1E6}-\\x{1F1FF}][\\x{1F1E6}-\\x{1F1FF}]"
	ResolutionMatch = "(2160|1080|720|480)(p|i)"
	GBSizeMatch     = "💾 ([0-9]+\\.?[0-9]+) GB"
	MBSizeMatch     = "💾 ([0-9]+\\.?[0-9]+) MB"
	SourceMatch     = "⚙️ (.*)(\\n|$)"
)

func RegexSearch(regex string, value string) []string {
	re := regexp.MustCompile(regex)
	match := re.FindStringSubmatch(value)
	if len(match) > 0 {
		return match
	} else {
		return nil
	}
}

func RegexFindAll(regex string, value string) []string {
	re := regexp.MustCompile(regex)
	match := re.FindAllString(value, -1)
	if len(match) > 0 {
		return match
	} else {
		return nil
	}
}

func RegexMatch(regex string, value string) bool {
	re := regexp.MustCompile(regex)
	return re.MatchString(value)
}

func RegexFind(regex string, value string) []string {
	re := regexp.MustCompile(regex)
	match := re.FindStringSubmatch(value)
	if len(match) > 0 {
		return match
	} else {
		return nil
	}
}
