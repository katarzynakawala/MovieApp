package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"movieapp.github.com.katarzynakawala/internal/jsonlog"
)

func newTestApplication(t *testing.T) *application {
	return &application{
		logger: jsonlog.New(os.Stdout, jsonlog.LevelInfo),
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)
	return &testServer{ts}
}

//get method on custom testServer type. It makes get request to a given url and returns status code, headers and body
func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, []byte) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body, err := io.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, body
}

