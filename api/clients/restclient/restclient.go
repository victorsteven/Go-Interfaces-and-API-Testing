package restclient

import (
	"net/http"
)

type clientStruct struct{}

type ClientInterface interface {
	Get(string) (*http.Response, error)
}
var (
	ClientStruct ClientInterface = &clientStruct{}
)

func (ci *clientStruct) Get(url string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}

	return client.Do(request)
}
