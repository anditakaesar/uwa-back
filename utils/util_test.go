package utils

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetBearerToken(t *testing.T) {
	t.Run("test GetBearerToken not empty success", func(t *testing.T) {
		request := http.Request{
			Method: http.MethodGet,
			Header: http.Header{},
		}

		request.Header.Add("Authorization", "Bearer 123")

		token := GetBearerToken(&request)
		assert.Equal(t, "123", token)
	})
	t.Run("test GetBearerToken return empty string success", func(t *testing.T) {
		request := http.Request{
			Method: http.MethodGet,
			Header: http.Header{},
		}

		request.Header.Add("Authorization", "")

		token := GetBearerToken(&request)
		assert.Equal(t, "", token)
	})

}
