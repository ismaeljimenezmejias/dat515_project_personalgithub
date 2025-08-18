package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var testRootRequests = []struct {
	reqType  string
	reqPath  string
	wantCode int
	wantBody string
}{
	{"GET", "/", 200, "Hello World!\n"},
}

func testRoot(t *testing.T, scoreDec, scoreFail func()) {
	server := NewServer()
	for i, tr := range testRootRequests {
		req, err := http.NewRequest(tr.reqType, tr.reqPath, nil)
		if err != nil {
			scoreFail()
			t.Errorf("TestRoot %d: got error creating http req: %v\n", i, err)
		}
		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)
		if w.Code != tr.wantCode {
			t.Errorf("TestRoot %d: got status code %d, want %d", i, w.Code, tr.wantCode)
			scoreDec()
		}
		if w.Body.String() != tr.wantBody {
			t.Errorf("TestRoot %d: got status body %q, want %q", i, w.Body.String(), tr.wantBody)
			scoreDec()
		}
	}
}

var testNonExistingRequests = []struct {
	reqType  string
	reqPath  string
	wantCode int
	wantBody string
}{
	{"GET", "/foo", 404, "404 page not found\n"},
	{"GET", "/foo/bar", 404, "404 page not found\n"},
	{"GET", "/uis", 404, "404 page not found\n"},
}

func testNonExisting(t *testing.T, scoreDec, scoreFail func()) {
	server := NewServer()
	for i, tr := range testNonExistingRequests {
		req, err := http.NewRequest(tr.reqType, tr.reqPath, nil)
		if err != nil {
			scoreFail()
			t.Errorf("TestNonExisting %d: got error creating http req: %v\n", i, err)
		}
		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)
		if w.Code != tr.wantCode {
			t.Errorf("TestNonExisting %d: got status code %d, want %d", i, w.Code, tr.wantCode)
			scoreDec()
		}
		if w.Body.String() != tr.wantBody {
			t.Errorf("TestNonExisting %d: got status body %q, want %q", i, w.Body.String(), tr.wantBody)
			scoreDec()
		}
	}
}

var testRedirectRequests = []struct {
	reqType  string
	reqPath  string
	wantCode int
	wantBody string
}{
	{"GET", "/github", 301, "<a href=\"http://www.github.com\">Moved Permanently</a>.\n\n"},
	{"GET", "/github/foo", 404, "404 page not found\n"},
}

func testRedirect(t *testing.T, scoreDec, scoreFail func()) {
	server := NewServer()
	for i, tr := range testRedirectRequests {
		req, err := http.NewRequest(tr.reqType, tr.reqPath, nil)
		if err != nil {
			scoreFail()
			t.Errorf("TestRedirect %d: got error creating http req: %v\n", i, err)
		}
		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)
		if w.Code != tr.wantCode {
			t.Errorf("TestRedirect %d: got status code %d, want %d", i, w.Code, tr.wantCode)
			scoreDec()
		}
		if w.Body.String() != tr.wantBody {
			t.Errorf("TestRedirect %d: got status body %q, want %q", i, w.Body.String(), tr.wantBody)
			scoreDec()
		}
	}
}

var testCounterRequests = []struct {
	reqType  string
	reqPath  string
	wantCode int
	wantBody string
}{
	{"GET", "/counter", 200, "counter: 1\n"},
	{"GET", "/counter", 200, "counter: 2\n"},
	{"GET", "/counter", 200, "counter: 3\n"},
	{"GET", "/counter", 200, "counter: 4\n"},
	{"GET", "/counter", 200, "counter: 5\n"},
}

func testCounter(t *testing.T, scoreDec, scoreFail func()) {
	server := NewServer()
	for i, tr := range testCounterRequests {
		req, err := http.NewRequest(tr.reqType, tr.reqPath, nil)
		if err != nil {
			scoreFail()
			t.Errorf("TestCounter %d: got error creating http req: %v\n", i, err)
		}
		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)
		if w.Code != tr.wantCode {
			t.Errorf("TestCounter %d: got status code %d, want %d", i, w.Code, tr.wantCode)
			scoreDec()
		}
		if w.Body.String() != tr.wantBody {
			t.Errorf("TestCounter %d: got status body %q, want %q", i, w.Body.String(), tr.wantBody)
			scoreDec()
		}
	}
}

var testFizzBuzzRequests = []struct {
	reqType  string
	reqPath  string
	wantCode int
	wantBody string
}{
	{"GET", "/fizzbuzz?value=7", 200, "7\n"},
	{"GET", "/fizzbuzz?value=8", 200, "8\n"},
	{"GET", "/fizzbuzz?value=9", 200, "fizz\n"},
	{"GET", "/fizzbuzz?value=19", 200, "19\n"},
	{"GET", "/fizzbuzz?value=abcdefg", 200, "not an integer\n"},
	{"GET", "/fizzbuzz?value", 200, "no value provided\n"},
	{"GET", "/fizzbuzz?value=100", 200, "buzz\n"},
	{"GET", "/fizzbuzz?value=90", 200, "fizzbuzz\n"},
	{"GET", "/fizzbuzz?value=1", 200, "1\n"},
	{"GET", "/fizzbuzz?value=2", 200, "2\n"},
	{"GET", "/fizzbuzz?value=3", 200, "fizz\n"},
	{"GET", "/fizzbuzz?value=4", 200, "4\n"},
	{"GET", "/fizzbuzz?value=5", 200, "buzz\n"},
	{"GET", "/fizzbuzz?value=30", 200, "fizzbuzz\n"},
	{"GET", "/fizzbuzz?value=lol", 200, "not an integer\n"},
	{"GET", "/fizzbuzz?value", 200, "no value provided\n"},
	{"GET", "/fizzbuzz?value=60", 200, "fizzbuzz\n"},
	{"GET", "/fizzbuzz?value=61", 200, "61\n"},
}

func testFizzBuzz(t *testing.T, scoreDec, scoreFail func()) {
	server := NewServer()
	for i, tr := range testFizzBuzzRequests {
		req, err := http.NewRequest(tr.reqType, tr.reqPath, nil)
		if err != nil {
			scoreFail()
			t.Errorf("TestFizzBuzz %d: got error creating http req: %v\n", i, err)
		}
		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)
		if w.Code != tr.wantCode {
			t.Errorf("TestFizzBuzz %d: got status code %d, want %d", i, w.Code, tr.wantCode)
			scoreDec()
		}
		if w.Body.String() != tr.wantBody {
			t.Errorf("TestFizzBuzz %d: got status body %q, want %q", i, w.Body.String(), tr.wantBody)
			scoreDec()
		}
	}
}

var testServerFullRequests = []struct {
	reqType  string
	reqPath  string
	wantCode int
	wantBody string
}{
	{"GET", "/counter", 200, "counter: 1\n"},
	{"GET", "/github", 301, "<a href=\"http://www.github.com\">Moved Permanently</a>.\n\n"},
	{"GET", "/", 200, "Hello World!\n"},
	{"GET", "/uis", 404, "404 page not found\n"},
	{"GET", "/fizzbuzz?value=1", 200, "1\n"},
	{"GET", "/counter", 200, "counter: 6\n"},
	{"GET", "/counter", 200, "counter: 7\n"},
	{"GET", "/fizzbuzz?value=3", 200, "fizz\n"},
	{"GET", "/fizzbuzz?value=5", 200, "buzz\n"},
	{"GET", "/fizzbuzz?value=60", 200, "fizzbuzz\n"},
	{"GET", "/counter", 200, "counter: 11\n"},
	{"GET", "/dat515", 404, "404 page not found\n"},
	{"GET", "/fizzbuzz?value=43", 200, "43\n"},
	{"GET", "/fizzbuzz?value=44", 200, "44\n"},
	{"GET", "/fizzbuzz?value=45", 200, "fizzbuzz\n"},
	{"GET", "/counter", 200, "counter: 16\n"},
	{"GET", "/fizzbuzz?value=hallo", 200, "not an integer\n"},
	{"GET", "/fizzbuzz?value", 200, "no value provided\n"},
	{"GET", "/counter", 200, "counter: 19\n"},
	{"GET", "/foobar", 404, "404 page not found\n"},
	{"GET", "/", 200, "Hello World!\n"},
	{"GET", "/foo", 404, "404 page not found\n"},
	{"GET", "/counter", 200, "counter: 23\n"},
	{"GET", "/github", 301, "<a href=\"http://www.github.com\">Moved Permanently</a>.\n\n"},
	{"GET", "/counter", 200, "counter: 25\n"},
	{"GET", "/fizzbuzz?value=1", 200, "1\n"},
	{"GET", "/fizzbuzz?value=3", 200, "fizz\n"},
	{"GET", "/fizzbuzz?value=5", 200, "buzz\n"},
	{"GET", "/fizzbuzz?value=30", 200, "fizzbuzz\n"},
	{"GET", "/counter", 200, "counter: 30\n"},
	{"GET", "/foobar", 404, "404 page not found\n"},
	{"GET", "/fizzbuzz?value=43", 200, "43\n"},
	{"GET", "/fizzbuzz?value=44", 200, "44\n"},
	{"GET", "/fizzbuzz?value=45", 200, "fizzbuzz\n"},
	{"GET", "/counter", 200, "counter: 35\n"},
	{"GET", "/fizzbuzz?value=hei", 200, "not an integer\n"},
	{"GET", "/fizzbuzz?value", 200, "no value provided\n"},
	{"GET", "/counter", 200, "counter: 38\n"},
	{"GET", "/foobar", 404, "404 page not found\n"},
}

func testServerFull(t *testing.T, scoreDec, scoreFail func()) {
	server := NewServer()
	for i, tr := range testServerFullRequests {
		req, err := http.NewRequest(tr.reqType, tr.reqPath, nil)
		if err != nil {
			scoreFail()
			t.Errorf("TestServerFull %d: got error creating http req: %v\n", i, err)
		}
		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)
		if w.Code != tr.wantCode {
			t.Errorf("TestServerFull %d: got status code %d, want %d", i, w.Code, tr.wantCode)
			scoreDec()
		}
		if w.Body.String() != tr.wantBody {
			t.Errorf("TestServerFull %d: got status body %q, want %q", i, w.Body.String(), tr.wantBody)
			scoreDec()
		}
	}
}
