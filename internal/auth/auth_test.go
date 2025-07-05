package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		headers        http.Header
		expectedKey    string
		expectingError bool
		expectedErr    error
	}{
		{
			name:           "No Authorization header",
			headers:        http.Header{},
			expectedKey:    "",
			expectingError: true,
			expectedErr:    ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed Authorization header (no ApiKey)",
			headers: http.Header{
				"Authorization": []string{"Bearer sometoken"},
			},
			expectedKey:    "",
			expectingError: true,
			expectedErr:    nil, // We’ll check just error presence for malformed ones
		},
		{
			name: "Malformed Authorization header (missing token)",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedKey:    "",
			expectingError: true,
			expectedErr:    nil,
		},
		{
			name: "Correct Authorization header",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid-api-key"},
			},
			expectedKey:    "valid-api-key",
			expectingError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			if tt.expectingError {
				if err == nil {
					t.Errorf("expected error but got nil")
				}
				if tt.expectedErr != nil && err != tt.expectedErr {
					t.Errorf("expected error %v, got %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got %v", err)
				}
				if key != tt.expectedKey {
					t.Errorf("expected key %q, got %q", tt.expectedKey, key)
				}
			}
		})
	}
}
