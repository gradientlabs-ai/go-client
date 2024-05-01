package client

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleResponse(t *testing.T) {
	tests := []struct {
		name               string
		responseBody       string
		responseStatusCode int
		value              interface{}
		expectedValue      interface{}
		expectedError      bool
	}{
		{
			name:               "successful response with valid JSON",
			responseBody:       `{"Name":"Attention Is All You Need"}`,
			responseStatusCode: 200,
			value:              &struct{ Name string }{},
			expectedValue:      &struct{ Name string }{Name: "Attention Is All You Need"},
			expectedError:      false,
		},
		{
			name:               "error response status code",
			responseBody:       "",
			responseStatusCode: 500,
			value:              nil,
			expectedValue:      nil,
			expectedError:      true,
		},
		{
			name:               "successful response with nil value (no need to decode JSON)",
			responseBody:       "",
			responseStatusCode: 200,
			value:              nil,
			expectedValue:      nil,
			expectedError:      false,
		},
		{
			name:               "successful response with bad JSON format",
			responseBody:       "<bad json>",
			responseStatusCode: 200,
			value:              &struct{ Name string }{},
			expectedValue:      &struct{ Name string }{},
			expectedError:      true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			client := &Client{}
			response := httptest.NewRecorder()
			response.Body = bytes.NewBufferString(tc.responseBody)
			response.Code = tc.responseStatusCode

			err := client.handleResponse(response.Result(), tc.value)
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			// value should be modified by the method
			assert.Equal(t, tc.expectedValue, tc.value)
		})
	}
}
