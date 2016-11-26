package favicon_test

import (
	"net/http"

	"github.com/go-http-utils/favicon"
)

func Example() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":8080", favicon.Handler(mux, "./public/favicon.ico"))
}
