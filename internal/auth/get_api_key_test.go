package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expectedKey string
		expectedErr bool
	}{
		{
			name: "Happy path",
			headers: http.Header{
				"Authorization": {"ApiKey somekey123"},
			},
			expectedKey: "somekey123",
			expectedErr: false,
		},
		{
			name:        "Missing header entirely",
			headers:     http.Header{},
			expectedKey: "",
			expectedErr: true,
		},
		{
			name: "Malformed header",
			headers: http.Header{
				"Authorization": {" somekey123"},
			},
			expectedKey: "",
			expectedErr: true,
		},
		{
			name: "Edge case",
			headers: http.Header{
				"Authorization": {"ApiKey"},
			},
			expectedKey: "",
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(test.headers)
			hasError := (err != nil)
			if (apiKey != test.expectedKey) || (hasError != test.expectedErr) {
				t.Errorf("%s: want key=%q, got key=%q (error: %v)", test.name, test.expectedKey, apiKey, err)
			}
		})
	}
}
