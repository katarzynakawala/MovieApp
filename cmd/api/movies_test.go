package main

import (
	"net/http"
	"testing"
)

var unauthorizedBody string = "{\n\t\"error\": \"you must be authenticated to access this resource\"\n}\n"

func TestListMovies(t *testing.T) {
	app := newTestApplication(t)
 	ts := newTestServer(t, app.routes())
 	defer ts.Close()

 	code, _, body := ts.get(t, "/v1/movies")
	
	if code != http.StatusUnauthorized {
        t.Errorf("want %d; got %d", http.StatusUnauthorized, code)
    }

    if string(body) != unauthorizedBody {
         t.Errorf("want body to equal %q; got %q", unauthorizedBody, body)
	} 
}

func TestShowMovie(t *testing.T) {
	app := newTestApplication(t)
 	ts := newTestServer(t, app.routes())
 	defer ts.Close()

 	code, _, body := ts.get(t, "/v1/movies/3")
	
	if code != http.StatusUnauthorized {
        t.Errorf("want %d; got %d", http.StatusUnauthorized, code)
    }

    if string(body) != unauthorizedBody {
         t.Errorf("want body to equal %q; got %q", unauthorizedBody, body)
	} 
}	