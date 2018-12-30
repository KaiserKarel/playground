package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// get the url, copy the payload from a http.Response.Body to a []byte.
func get(url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "getting gist")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("bad status")
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading body")
	}

	return res, nil
}

// NoParam is returned if the http.Request does not contain any parameters leading to a prefilled snippet.
type NoParam error

// getURL parses the http.Request form and attempts returns a URL leading to a raw txt document with the gocode for the prefill.
// accepted parameters in the http.Request:
//
// gist: the github gist "username/ID"
// raw: a complete URL
func getURL(r *http.Request) (string, error) {
	err := r.ParseForm()
	if err != nil {
		return "", errors.Wrap(err, "parsing form gist")
	}

	if r.Form.Get("gist") != "" {
		return fmt.Sprintf("https://gist.githubusercontent.com/%s/raw/", r.Form.Get("gist")), nil
	}

	if r.Form.Get("raw") != "" {
		return r.Form.Get("raw"), nil
	}

	return "", NoParam(errors.New("no parameter present"))
}
