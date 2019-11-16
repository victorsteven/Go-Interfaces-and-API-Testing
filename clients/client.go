package clients

import (
	"net/http"
)

func GetInfo(url string) (*http.Response, error) {
	//if enabledMocks {
	//	mock := mocks[getMockId(http.MethodPost, url)]
	//	if mock == nil {
	//		return nil, errors.New("no mockup found for give request")
	//	}
	//	return mock.Response, mock.Err
	//}

	//jsonBytes, err := json.Marshal(body)
	//if err != nil {
	//	return nil, err
	//}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	return client.Do(request)
}