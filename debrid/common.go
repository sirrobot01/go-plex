package debrid

import (
	"github.com/sirrobot01/goplex_debrid/schemas"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func SortReleases(releases []schemas.Release, server schemas.Server) []schemas.Release {
	filters := server.Filters
	for key, value := range filters {
		sort.Slice(releases[:], func(i, j int) bool {
			r := releases[i]
			s := releases[j]
			// sort by server's array of resolutions
			if key == "resolution" {
				res := strings.Split(value, ",")
				resolutions := make([]int, len(res))
				// Convert to int
				for i, res := range res {
					resolutions[i], _ = strconv.Atoi(strings.TrimSpace(res))
				}
				if slices.Contains(resolutions, r.Resolution) && slices.Contains(resolutions, s.Resolution) {
					return slices.Index(resolutions, r.Resolution) < slices.Index(resolutions, s.Resolution)
				}

			}
			return false
		})
	}
	return releases
}
