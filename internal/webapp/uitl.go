package webapp

import (
	"net/http"
	"strconv"
)

// getPaginationFromURL returns a page and item count derived from the url query.
// returns true if count is present in query.
func getPaginationFromURL(r *http.Request, c int) (int, int, bool) {
	// get display count
	count := c
	if qCount, ok := r.URL.Query()["count"]; ok {
		if len(qCount[0]) >= 1 {
			uCount, err := strconv.ParseUint(qCount[0], 10, 64)
			if err != nil {
				logger.Debugf("invalid count: %s", qCount[0])
			} else {
				count = int(uCount)
			}
		}
	}

	// get display page
	page := 1
	if qPage, ok := r.URL.Query()["page"]; ok {
		if len(qPage[0]) >= 1 {
			uPage, err := strconv.ParseUint(qPage[0], 10, 64)
			if err != nil {
				logger.Debugf("invalid page: %s", qPage[0])
			} else {
				page = int(uPage)
			}
		}
	}

	return page, count, c != count
}
