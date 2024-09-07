package auth

import (
	"net/http"
	"testing"
)

// Test function for GetAPIKey
func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name     string
		headers  http.Header
		expected string
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "Valid API Key",
			headers:  http.Header{"Authorization": []string{"ApiKey my-secret-key"}},
			expected: "my-secret-key",
			wantErr:  false,
		},
		{
			name:     "Missing Authorization Header",
			headers:  http.Header{},
			expected: "",
			wantErr:  true,
			errMsg:   ErrNoAuthHeaderIncluded.Error(),
		},
		{
			name:     "Malformed Authorization Header",
			headers:  http.Header{"Authorization": []string{"Bearer token"}},
			expected: "",
			wantErr:  true,
			errMsg:   "malformed authorization header",
		},
		{
			name:     "Authorization Header Missing API Key Prefix",
			headers:  http.Header{"Authorization": []string{"ApiKey"}},
			expected: "",
			wantErr:  true,
			errMsg:   "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("GetAPIKey() error = %v, wantErrMsg %v", err, tt.errMsg)
			}
			if got != tt.expected {
				t.Errorf("GetAPIKey() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
