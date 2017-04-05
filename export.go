package chalmers_chop

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
)

type Exporter interface {
	Export(json []byte) error
}

type POSTExporter struct {
	url   string
	token string
}

func NewPOSTExporter(url, token string) *POSTExporter {
	return &POSTExporter{
		url:   url,
		token: token,
	}
}

func (e *POSTExporter) Export(json []byte) error {
	resp := postJson(json, e.url, e.token)

	ok := resp.StatusCode == http.StatusOK ||
		resp.StatusCode == http.StatusCreated ||
		resp.StatusCode == http.StatusNoContent

	if !ok {
		msg := fmt.Sprintf("POST Export failed: %v (%v)", resp.Status, resp.Body)
		return errors.New(msg)
	}

	return nil
}

func postJson(json []byte, url, token string) *http.Response {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "chalmers-chop/"+packageVersion)

	if token != "" {
		req.Header.Set("Authorization", "Token "+token)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := httputil.DumpResponse(resp, true)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(b[:]))

	return resp
}
