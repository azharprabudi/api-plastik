package baseurl

import (
	"fmt"
	"net/http"
	"strings"
)

// Get ...
func Get(r *http.Request, additionalPath ...interface{}) string {
	var baseURL string
	paths := strings.Split(r.URL.Path, "/")

	// create rawpath
	baseURL = fmt.Sprintf("%s/%s/%s", r.Host, paths[1], paths[2])

	// loop additional path
	for _, path := range additionalPath {
		baseURL = fmt.Sprintf("%s/%v", baseURL, path)
	}

	return baseURL
}
