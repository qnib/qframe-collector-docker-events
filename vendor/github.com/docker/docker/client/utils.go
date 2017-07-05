package client

import (
<<<<<<< HEAD
	"github.com/docker/docker/api/types/filters"
	"net/url"
	"regexp"
=======
	"net/url"
	"regexp"

	"github.com/docker/docker/api/types/filters"
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
)

var headerRegexp = regexp.MustCompile(`\ADocker/.+\s\((.+)\)\z`)

// getDockerOS returns the operating system based on the server header from the daemon.
func getDockerOS(serverHeader string) string {
	var osType string
	matches := headerRegexp.FindStringSubmatch(serverHeader)
	if len(matches) > 0 {
		osType = matches[1]
	}
	return osType
}

// getFiltersQuery returns a url query with "filters" query term, based on the
// filters provided.
func getFiltersQuery(f filters.Args) (url.Values, error) {
	query := url.Values{}
	if f.Len() > 0 {
		filterJSON, err := filters.ToParam(f)
		if err != nil {
			return query, err
		}
		query.Set("filters", filterJSON)
	}
	return query, nil
}
