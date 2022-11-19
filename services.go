package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Service struct {
	endpoint string
}

func CreateService(endpoint string) Service {
	return Service{
		endpoint: endpoint,
	}
}

func (s *Service) Get(path string, cookie string) (*http.Response, error) {
	return s.request("GET", path, cookie, nil)
}

func (s *Service) Post(path string, cookie string, body map[string]string) (*http.Response, error) {
	return s.request("POST", path, cookie, body)
}

func (s *Service) Put(path string, cookie string, body map[string]string) (*http.Response, error) {
	return s.request("PUT", path, cookie, body)
}

func (s *Service) Delete(path string, cookie string, body map[string]string) (*http.Response, error) {
	return s.request("DELETE", path, cookie, body)
}




func (s *Service) request(method string,path string, cookie string ,body map[string]string) (*http.Response, error) {
	var data io.Reader = nil

	if body != nil {
		jsonData , err := json.Marshal(body)

		if err != nil {
			return nil, err
		}

		data = bytes.NewBuffer(jsonData)
	}



	req, err := http.NewRequest(method, s.endpoint+path, data)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	if cookie != "" {
		req.Header.Add("Cookie", "jwt="+cookie)
	}


	client := &http.Client{}

	return client.Do(req)
}