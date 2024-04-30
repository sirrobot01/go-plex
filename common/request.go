package common

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"net/http"
	gourl "net/url"
	"strings"
	"time"
)

type Request struct {
	Timeout     int
	Retries     int
	RetryCodes  []int
	RateLimit   int
	LastReqTime time.Time
}

func NewSession(timeout, retries, rateLimit int, retryCodes []int) Request {
	if len(retryCodes) == 0 {
		retryCodes = []int{500, 502, 503, 504}
	}
	if rateLimit == 0 {
		rateLimit = 5
	}
	if retries == 0 {
		retries = 3
	}
	if timeout == 0 {
		timeout = 10
	}
	return Request{
		Timeout:    timeout,
		Retries:    retries,
		RetryCodes: retryCodes,
		RateLimit:  rateLimit,
	}
}

func (r *Request) SetLastReqTime(t time.Time) {
	r.LastReqTime = t
}

func (r *Request) Get(url string, headers map[string]string) (map[string]any, error) {
	res, err := r.Make(url, "GET", headers, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Request) Post(url string, body any, headers map[string]string) (map[string]any, error) {
	res, err := r.Make(url, "POST", headers, body)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Request) Put(url string, body any, headers map[string]string) (map[string]any, error) {
	res, err := r.Make(url, "PUT", headers, body)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Request) Delete(url string, headers map[string]string) (map[string]any, error) {
	res, err := r.Make(url, "DELETE", headers, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *Request) Make(url string, method string, headers map[string]string, body any) (map[string]any, error) {
	var result map[string]any

	method = strings.ToUpper(method)

	var bd io.Reader
	if body != nil {
		b, ok := body.(map[string]any)
		if ok {
			bodyBytes, err := json.Marshal(b)
			if err != nil {
				return nil, err
			}
			bd = bytes.NewBuffer(bodyBytes)
		} else {
			b := body.(gourl.Values)
			bd = strings.NewReader(b.Encode())
		}
	}

	req, err := http.NewRequest(method, url, bd)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: time.Duration(r.Timeout) * time.Second,
	}

	if time.Since(r.LastReqTime) < time.Duration(r.RateLimit)*time.Second {
		time.Sleep(time.Duration(r.RateLimit) * time.Second)
	}
	for i := r.Retries; i >= 0; i-- {
		data, err, retry := MakeRequest(r, client, req, i)
		if err != nil || !retry {
			return nil, err
		}
		contentType := req.Header.Get("Content-Type")
		switch contentType {
		case "application/json":
			err = json.Unmarshal(data, &result)
			if err != nil {
				return nil, err
			}
		case "application/xml":
			err = xml.Unmarshal(data, &result)
			if err != nil {
				return nil, err
			}
		default:
			err = json.Unmarshal(data, &result)
			if err != nil {
				return nil, err
			}
		}
		if err == nil {
			return result, nil
		}

	}
	return result, nil
}

func MakeRequest(r *Request, client *http.Client, req *http.Request, retryCount int) ([]byte, error, bool) {
	if time.Since(r.LastReqTime) < time.Duration(r.RateLimit)*time.Second {
		time.Sleep(time.Duration(r.RateLimit) * time.Second)
	}
	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err, false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	for _, code := range r.RetryCodes {
		if resp.StatusCode == code && retryCount > 0 {
			return nil, nil, true
		}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err, false
	}
	return body, nil, false
}
