package debrid

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sirrobot01/goplex_debrid/common"
	"github.com/sirrobot01/goplex_debrid/config"
	"github.com/sirrobot01/goplex_debrid/schemas"
	"net/url"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var (
	session = config.SESSION
)

func CheckCached(releases []schemas.Release, mediaType string, seasons []int, episode int) []schemas.Release {
	hashes := make([]string, 0)
	conf := config.CONFIG

	for _, release := range releases {
		hashes = append(hashes, release.Hash)
	}
	cachedReleases := make([]schemas.Release, 0)

	// No cached versions in release
	if len(hashes) < 1 {
		return releases
	}
	hashesUrl := strings.Join(hashes, "/")
	url := fmt.Sprintf("https://api.real-debrid.com/rest/1.0/torrents/instantAvailability/%s", hashesUrl)
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", conf.RealdebridAPIKey),
	}
	rdReleases, _ := session.Get(url, headers)
	var wg sync.WaitGroup
	wg.Add(len(releases))
	for _, release := range releases {
		go func(r schemas.Release) {
			defer wg.Done()
			cacheRelease(rdReleases[r.Hash], r, &cachedReleases)
		}(release)
	}
	wg.Wait()

	return cachedReleases
}

func cacheRelease(fd any, release schemas.Release, cachedReleases *[]schemas.Release) {
	if fd == nil {
		return
	}
	pf, ok := fd.(map[string]any)
	if !ok {
		// assertion failed, probably an empty hash:[]
		return
	}
	parents := pf["rd"].([]any)

	if len(parents) < 1 {
		// Skip if no cached files
		return
	}
	versions := make([]schemas.Version, 0)
	for _, fs := range parents {
		files := fs.(map[string]any)
		version := schemas.Version{}
		for key, _f := range files {
			f := _f.(map[string]any)
			name := f["filename"].(string)
			s := common.RegexFind(common.SeasonMatch, name)
			season := 0
			if s != nil {
				season, _ = strconv.Atoi(s[1])
			}
			e := common.RegexFind(common.EpisodeMatch, name)
			episode := 0
			if e != nil {
				episode, _ = strconv.Atoi(e[1])
			}
			file := schemas.File{
				Name:     name,
				Size:     f["filesize"].(float64) / float64(8*1024*1024*1024),
				Id:       key,
				Video:    common.RegexMatch(common.VideoMatch, name),
				Subtitle: common.RegexMatch(common.SubtitleMatch, name),
				Season:   season,
				Episode:  episode,
			}
			version.Files = append(version.Files, file)
			version.Size += file.Size
			if file.Video {
				version.Videos += 1
			}
			if file.Subtitle {
				version.Subtitles += 1
			}
			if file.Season != 0 && file.Video && !slices.Contains(version.Seasons, file.Season) {
				version.Seasons = append(version.Seasons, file.Season)
			}
			if file.Episode != 0 && file.Video {
				version.Episodes += 1
			}
		}
		versions = append(versions, version)
	}

	sort.Slice(versions[:], func(i, j int) bool {
		a := versions[i]
		b := versions[j]
		if a.Videos != b.Videos {
			return a.Videos > b.Videos
		}
		ratioI := float64(a.Videos) / float64(len(a.Files))
		ratioJ := float64(b.Videos) / float64(len(b.Files))
		return ratioI > ratioJ
	})
	release.Versions = versions

	if len(release.Versions) > 0 {
		f := release.Versions[0]
		release.Size = f.Size
		release.Videos = f.Videos
		release.Seasons = f.Seasons
		release.Episodes = f.Episodes
		release.Cached = []string{"RD"}
	}
	*cachedReleases = append(*cachedReleases, release)
}

func FilterReleases(releases []schemas.Release, mediaType string, seasons []int, episode int) []schemas.Release {
	filteredReleases := make([]schemas.Release, 0)
	slices.Sort(seasons)              // Sort
	seasons = slices.Compact(seasons) // Remove duplicate

	var wg sync.WaitGroup
	wg.Add(len(releases))

	if mediaType == "movie" {
		for _, release := range releases {
			go func(release schemas.Release) {
				defer wg.Done()
				if release.Videos == 0 {
					return
				}
				versions := make([]schemas.Version, 0)
				for _, version := range release.Versions {
					if version.Videos == 0 {
						continue
					}
					versions = append(versions, version)
				}
				release.Versions = versions
				release.Type = "movie"
				filteredReleases = append(filteredReleases, release)
			}(release)
		}
	} else if len(seasons) > 1 {
		// Show with multiple seasons
		for _, release := range releases {
			go func(release schemas.Release) {
				defer wg.Done()
				rs := release.Seasons
				slices.Sort(rs)         // Sort
				rs = slices.Compact(rs) // Remove duplicate
				diff := common.SetDifference(seasons, rs)
				if len(diff) > len(seasons)/2.0 || release.Episodes <= 1 {
					return
				}
				versions := make([]schemas.Version, 0)
				for _, version := range release.Versions {
					rs := version.Seasons
					slices.Sort(rs)         // Sort
					rs = slices.Compact(rs) // Remove duplicate
					diff := common.SetDifference(seasons, rs)
					if len(diff) > len(seasons)/2.0 || release.Episodes <= 1 {
						continue
					}
					versions = append(versions, version)
				}
				release.Versions = versions
				release.Type = "show"
				filteredReleases = append(filteredReleases, release)
			}(release)
		}
	} else if episode == 0 {
		// No episode specified
		for _, release := range releases {
			go func(release schemas.Release) {
				defer wg.Done()
				rs := release.Seasons
				slices.Sort(rs)
				rs = slices.Compact(rs)
				diff := common.SetDifference(seasons, rs)
				if len(diff) > 0 || release.Episodes <= 1 {
					return
				}
				versions := make([]schemas.Version, 0)
				for _, version := range release.Versions {
					rs := version.Seasons
					slices.Sort(rs)         // Sort
					rs = slices.Compact(rs) // Remove duplicate
					diff := common.SetDifference(seasons, rs)
					if len(diff) > 0 || release.Episodes <= 1 {
						continue
					}
					versions = append(versions, version)
				}
				release.Versions = versions
				release.Type = "show"
				filteredReleases = append(filteredReleases, release)
			}(release)
		}
	} else {
		for _, release := range releases {
			go func(release schemas.Release) {
				defer wg.Done()
				rs := release.Seasons
				slices.Sort(rs)
				rs = slices.Compact(rs)
				diff := common.SetDifference(seasons, rs)
				if len(diff) > 0 || !(release.Episodes == 1) {
					return
				}
				versions := make([]schemas.Version, 0)
				for _, version := range release.Versions {
					rs := version.Seasons
					slices.Sort(rs)         // Sort
					rs = slices.Compact(rs) // Remove duplicate
					diff := common.SetDifference(seasons, rs)
					if len(diff) > 0 || release.Episodes == 1 {
						continue
					}
					versions = append(versions, version)
				}
				release.Versions = versions
				release.Type = "show"
				filteredReleases = append(filteredReleases, release)
			}(release)
		}
	}
	wg.Wait()
	if len(filteredReleases) > 0 {
		releases = filteredReleases
	}
	// Remove duplicate
	return slices.CompactFunc(releases, func(release schemas.Release, release2 schemas.Release) bool {
		return release.Magnet == release2.Magnet
	})
}

func Download(release *schemas.Release) bool {

	conf := config.CONFIG

	payload := url.Values{
		"magnet": {release.Magnet},
	}
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", conf.RealdebridAPIKey),
	}
	response, err := session.Post(fmt.Sprintf("https://api.real-debrid.com/rest/1.0/torrents/addMagnet"), payload, headers)
	if err != nil {
		return false
	}
	torrentId := response["id"].(string)
	var lastVersion schemas.Version

	for _, version := range release.Versions {
		fileIds := make([]string, 0)
		for _, file := range version.Files {
			fileIds = append(fileIds, file.Id)
		}

		session.Post(
			fmt.Sprintf("https://api.real-debrid.com/rest/1.0/torrents/selectFiles/%s", torrentId),
			url.Values{
				"files": {strings.Join(fileIds, ",")},
			},
			headers,
		)

		// Fetch info

		response, _ := session.Get(
			fmt.Sprintf("https://api.real-debrid.com/rest/1.0/torrents/info/%s", torrentId),
			headers,
		)

		links := response["links"].([]string)
		if len(links) == len(fileIds) {
			release.Title = response["filename"].(string)
			release.Download = links
			break
		} else {
			session.Delete(
				fmt.Sprintf("https://api.real-debrid.com/rest/1.0/torrents/delete/%s", torrentId),
				headers,
			)
		}
		lastVersion = version

	}

	if len(release.Download) > 0 {
		log.Infof("Added %s to Realdebrid", release.Title)
		for _, link := range release.Download {
			go func(_url string) {
				session.Post(
					"https://api.real-debrid.com/rest/1.0/unrestrict/link",
					url.Values{
						"link": {_url},
					},
					headers,
				)
			}(link)
		}
		release.Files = lastVersion.Files
		return true
	}
	return false
}
