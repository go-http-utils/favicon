package favicon

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-http-utils/headers"
)

// Version is this package's version.
const Version = "0.1.0"

// Handler wraps the http.Handler h with favicon support. `path`
// is the path to find the favicon.
func Handler(h http.Handler, path string) http.Handler {
	if !os.IsPathSeparator(path[0]) {
		wd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		path = filepath.Join(wd, path)
	}

	stat, err := os.Stat(path)

	if err != nil || stat.IsDir() {
		panic("favicon: Invalid favicon path: " + path)
	}

	file, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	readSeeker := bytes.NewReader(file)

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.RequestURI != "/favicon.ico" {
			h.ServeHTTP(res, req)

			return
		}

		if req.Method != http.MethodGet && req.Method != http.MethodHead {
			res.Header().Set(headers.Allow, "GET, HEAD, OPTIONS")

			if req.Method == http.MethodOptions {
				res.WriteHeader(http.StatusOK)
			} else {
				res.WriteHeader(http.StatusMethodNotAllowed)
			}

			return
		}

		res.Header().Set(headers.ContentType, "image/x-icon")

		http.ServeContent(res, req, "favicon.ico", stat.ModTime(), readSeeker)
	})
}
