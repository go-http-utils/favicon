# favicon
[![Build Status](https://travis-ci.org/go-http-utils/favicon.svg?branch=master)](https://travis-ci.org/go-http-utils/favicon)
[![Coverage Status](https://coveralls.io/repos/github/go-http-utils/favicon/badge.svg?branch=master)](https://coveralls.io/github/go-http-utils/favicon?branch=master)

Go http middleware for serving the favicon.

## Installation

```
go get -u github.com/go-http-utils/favicon
```

## Documentation

API documentation can be found here: https://godoc.org/github.com/go-http-utils/favicon

## Usage

```go
import (
  "github.com/go-http-utils/favicon"
)
```

```go
mux := http.NewServeMux()
mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
  res.Write([]byte("Hello World"))
})

http.ListenAndServe(":8080", favicon.Handler(mux, "./public/favicon.ico"))
```
