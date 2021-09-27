package writer

import (
	"bytes"
	"encoding/json"
	"github.com/hyperyuri/logg"
	"net/http"
	"time"
)

type restW struct {
	url string
	header http.Header
	expectedStatus int
}

//WithRestApiMode send the logs to a rest endpoint with ToJSON() struct
func WithRestApiMode(url string, header http.Header, expectedStatus int) *restW {
	return &restW{
		url: url,
		header: header,
		expectedStatus: expectedStatus,
	}
}

func (s *restW) Write(m *logg.Message) {
	for i := 0; i < 5; i++ {
		message, err := json.Marshal(m.ToJSON())
		if err != nil {
			time.Sleep(time.Minute * 1)
			continue
		}

		req, err := http.NewRequest("POST", s.url, bytes.NewBuffer(message))
		if err != nil {
			time.Sleep(time.Minute * 1)
			continue
		}

		req.Header = s.header
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			time.Sleep(time.Minute * 1)
			continue
		}

		if resp.StatusCode != s.expectedStatus {
			time.Sleep(time.Minute * 1)
			continue
		}

		break
	}
}