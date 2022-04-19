/*
 * Copyright (c) The go-kit Authors
 */

package requestid

import (
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

var validRequestID = regexp.MustCompile(`([a-f\d]{8}(-[a-f\d]{4}){3}-[a-f\d]{12}?)`)

func TestRequestID(t *testing.T) {
	// create a bunch of request IDs
	rids := []string{"", "remoteID_12345"}
	// setup http handler and set in the response the request id obtained
	// from the request headers
	fn := func(w http.ResponseWriter, r *http.Request) {
		rid := r.Header.Get(XRequestIDMetadataKey)
		_, _ = fmt.Fprintf(w, "%s", rid)
	}
	// create an instance of the request id http handler
	h := NewRequestIDHttpHandler(http.HandlerFunc(fn))
	// set the pre-set request ids
	for _, rid := range rids {
		w := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "https://example.com/foo", nil)
		if err != nil {
			t.Fatal(err)
		}

		if rid != "" {
			// Pre-set request ID
			req.Header.Set(XRequestIDMetadataKey, rid)
		}
		// serve the http request
		h.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Errorf("request '%s': %d != %d", rid, w.Code, http.StatusOK)
		}
		// get the server response back
		body := strings.TrimSpace(w.Body.String())
		// request id must match the response in the body
		if rid != "" {
			if body != rid {
				t.Errorf("request '%s': %s != %[1]s", rid, body)
			}
		} else {
			if !validRequestID.MatchString(body) {
				t.Errorf("request '%s': %s is not valid format", rid, body)
			}
		}
	}
}

func TestRequestHandlerMiddleware(t *testing.T) {
	expected := "some-request-id"
	// setup handler and set in the response body the request id received from the header
	fn := func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "%s,%s", r.Header.Get(XRequestIDMetadataKey), expected)
	}
	// create a serverMux and use the request id http handler as middleware
	mux := http.NewServeMux()
	mux.Handle("/", NewRequestIDHttpHandler(http.HandlerFunc(fn)))
	// spawn a http server
	server := httptest.NewServer(mux)
	defer server.Close()
	// send a request to the http server
	res, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	// read the response body
	body, err := ioutil.ReadAll(res.Body)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	results := strings.Split(string(body), ",")
	if len(results) != 2 {
		t.Fatalf("Invalid results: %v", results)
	}
	if !validRequestID.MatchString(results[0]) {
		t.Errorf("%s is not valid request id format", results[0])
	}
	if results[1] != expected {
		t.Errorf("%v != %v", results[1], expected)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	rid := FromContext(r.Context())
	log.Println("Running hello handler:", rid)
	_, _ = fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Example() {
	h := http.HandlerFunc(handler)
	http.Handle("/", NewRequestIDHttpHandler(h))
	log.Fatal(http.ListenAndServe(":2000", nil))
}
