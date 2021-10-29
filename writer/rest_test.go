package writer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hyperyuri/logg"
	"github.com/hyperyuri/logg/levels"
	"github.com/stretchr/testify/assert"
)

const (
	PROJECT_STACK = "Logger"
	PROJECT_NAME  = "logg"
	LOG_MESSAGE   = "Testing rest"
)

func Test_WithRestApiMode(t *testing.T) {

	body := struct {
		Message string `json:"message,omitempty"`
		Stack   string `json:"stack,omitempty"`
		App     string `json:"app,omitempty"`
	}{}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		assert.NotNil(t, b)
		assert.NoError(t, err)

		err = json.Unmarshal(b, &body)
		assert.NoError(t, err)
		assert.NotNil(t, body)

		assert.Equal(t, PROJECT_STACK, body.Stack)
		assert.Equal(t, PROJECT_NAME, body.App)
		assert.Equal(t, LOG_MESSAGE, body.Message)

		w.WriteHeader(http.StatusCreated)
	})

	server := httptest.NewServer(handler)

	logger := logg.NewLogger().
		WithWriter(WithRestApiMode(server.URL, nil, 201)).
		WithDefaultConfig(logg.NewDefaultConfig(PROJECT_STACK, PROJECT_NAME))
	logger.SetLevels(levels.NewLevels(logger))

	logger.Info(LOG_MESSAGE)
}
