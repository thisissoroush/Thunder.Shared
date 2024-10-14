package httpclient

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

func Get(url string, headers *map[string]string) (*fasthttp.Response, error) {
	client := fasthttp.Client{}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

	if headers != nil {
		for key, value := range *headers {
			req.Header.Set(key, value)
		}
	}

	resp := fasthttp.AcquireResponse()

	err := client.Do(req, resp)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}

	return resp, nil
}

func Post(url string, body interface{}, headers *map[string]string) (*fasthttp.Response, error) {
	client := fasthttp.Client{}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("Post")

	if headers != nil {
		for key, value := range *headers {
			req.Header.Set(key, value)
		}
	}
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %w", err)
		}
		req.SetBody(jsonData)
	}

	resp := fasthttp.AcquireResponse()

	err := client.Do(req, resp)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}

	return resp, nil
}
