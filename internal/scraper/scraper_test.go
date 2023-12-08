package scraper

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScrape(t *testing.T) {
	// do not want to perform real requests
	// we want to be able to easily create new test cases
	// ideally the tests should be easily readable

	fakeServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	// table driven tests

	outBuffer := bytes.Buffer{}

	method := "GET"

	err := Scrape(
		&outBuffer,
		[]Target{
			{
				Method: method,
				Url:    fakeServer.URL,
			},
		},
	)

	if err != nil {
		t.Error(err)
	}

	assert.Contains(
		t,
		outBuffer.String(),
		method,
	)

	assert.Contains(
		t,
		outBuffer.String(),
		fakeServer.URL,
	)
}
