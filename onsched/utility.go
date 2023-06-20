package onsched

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) buildEndpoint(path string) string {
	return fmt.Sprintf("%s/%s", c.apiHost, path)
}

func (c *Client) get(path string) ([]byte, error) {
	resp, err := c.http.Get(c.buildEndpoint(path))
	if err != nil {
		return nil, err
	}
	return readResponse(resp)
}

func (c *Client) put(path string, data any) ([]byte, error) {
	req, err := newJsonRequest("PUT", c.buildEndpoint(path), data)
	req.Header.Add("content-type", "application/json")
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	return readResponse(resp)
}

func newJsonRequest(method, path string, data any) (*http.Request, error) {
	req, err := http.NewRequest(method, path, reader(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("content-type", "application/json")
	return req, nil
}

func readResponse(resp *http.Response) ([]byte, error) {
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func parse[T any](content []byte) (T, error) {
	result := new(T)
	err := json.Unmarshal(content, result)
	if err != nil {
		return *result, err
	}
	return *result, nil
}

func reader(data any) io.Reader {
	content, _ := json.Marshal(data)
	buffer := bytes.NewBuffer(content)
	return buffer
}
