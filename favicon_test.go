package favicon

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-http-utils/headers"
	"github.com/stretchr/testify/suite"
)

func init() {
	bytes, err := ioutil.ReadFile("./test_fixtures/favicon.ico")

	if err != nil {
		panic(err)
	}

	faviconBytes = bytes
}

var faviconBytes []byte

type FaviconSuite struct {
	suite.Suite

	relServer *httptest.Server
	absServer *httptest.Server
}

func (s *FaviconSuite) SetupTest() {
	relMux := http.NewServeMux()
	relMux.Handle("/", http.HandlerFunc(helloHandlerFunc))

	s.relServer = httptest.NewServer(Handler(relMux,
		"./test_fixtures/favicon.ico"))

	absMux := http.NewServeMux()
	absMux.Handle("/", http.HandlerFunc(helloHandlerFunc))

	wd, err := os.Getwd()
	s.Nil(err)

	s.absServer = httptest.NewServer(Handler(absMux,
		filepath.Join(wd, "./test_fixtures/favicon.ico")))

	panicMux := http.NewServeMux()

	s.Panics(func() {
		httptest.NewServer(Handler(panicMux, "invalid-path"))
	})
}

func (s *FaviconSuite) TestNotFaviconRes() {
	res, err := http.Get(s.relServer.URL + "/")

	s.Nil(err)
	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal([]byte("Hello World"), getResRawBody(res))
}

func (s *FaviconSuite) TestRelPathFaviconRes() {
	res, err := http.Get(s.relServer.URL + "/favicon.ico")

	s.Nil(err)
	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal("image/x-icon", res.Header.Get(headers.ContentType))
	s.Equal(faviconBytes, getResRawBody(res))
}

func (s *FaviconSuite) TestAbsPathFaviconRes() {
	res, err := http.Get(s.absServer.URL + "/favicon.ico")

	s.Nil(err)
	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal("image/x-icon", res.Header.Get(headers.ContentType))
	s.Equal(faviconBytes, getResRawBody(res))
}

func (s *FaviconSuite) TestNotAcceptPostReq() {
	req, err := http.NewRequest(http.MethodPost,
		s.relServer.URL+"/favicon.ico", nil)

	s.Nil(err)

	res, err := sendRequest(req)

	s.Nil(err)
	s.Equal(http.StatusMethodNotAllowed, res.StatusCode)
	s.Equal("GET, HEAD, OPTIONS", res.Header.Get(headers.Allow))
}

func (s *FaviconSuite) TestAcceptOptionsReq() {
	req, err := http.NewRequest(http.MethodOptions,
		s.relServer.URL+"/favicon.ico", nil)

	s.Nil(err)

	res, err := sendRequest(req)

	s.Nil(err)
	s.Equal(http.StatusOK, res.StatusCode)
	s.Equal("GET, HEAD, OPTIONS", res.Header.Get(headers.Allow))
}

func TestFavicon(t *testing.T) {
	suite.Run(t, new(FaviconSuite))
}

func helloHandlerFunc(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)

	res.Write([]byte("Hello World"))
}

func getResRawBody(res *http.Response) []byte {
	if b, err := ioutil.ReadAll(res.Body); err != nil {
		panic(err)
	} else {
		return b
	}
}

func sendRequest(req *http.Request) (*http.Response, error) {
	cli := &http.Client{}
	return cli.Do(req)
}
