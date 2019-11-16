package restclient

import (
	"errors"
	"net/http"
)



func Get(url string) (*http.Response, error) {
	if enableMocks {
		mock := mocks[getMockId(http.MethodGet, url)]
		if mock == nil {
			return nil, errors.New("no mockup found for the given request")
		}
		//fmt.Println("this is the mock response: ", mock.Response)
		return mock.Response, mock.Err
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}

	return client.Do(request)
}

