package mvola

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthService_GenerateToken(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
			"access_token": "access_token",
			"expires_in": 3600,
			"token_type": "Bearer",
			"scope": "EXT_INT_MVOLA_SCOPE"
		}`)
	})

	client := NewClient(server.URL)
	client.base.Client(httpClient)

	var (
		consumerKey    = "consumer_key"
		consumerSecret = "consumer_secret"
	)
	response, err := client.Auth.GenerateToken(consumerKey, consumerSecret)
	expected := &Auth{
		AccessToken: "access_token",
		ExpiresIn:   3600,
		TokenType:   "Bearer",
		Scope:       "EXT_INT_MVOLA_SCOPE",
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, response)
}
