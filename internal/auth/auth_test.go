package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		header      http.Header
		expectedKey string
		expectErr   bool
	}{
		{
			name: "valid header",
			header: http.Header{
				"Authorization": []string{"ApiKey abc123"},
			},
			expectedKey: "abc123",
			expectErr:   false,
		},
		{
			name: "valid header2",
			header: http.Header{
				"Authorization": []string{"ApiKey 3233"},
			},
			expectedKey: "3233",
			expectErr:   false,
		},
		{
			name:      "missing header",
			header:    http.Header{},
			expectErr: true,
		},
		{
			name: "malformed header missing key",
			header: http.Header{
				"Autorization": []string{"ApiKey"},
			},
			expectErr: true,
		},
		{
			name: "wrong auth scheme",
			header: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.header)

			if tt.expectErr {
				if err == nil {
					t.Fatalf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Fatalf(`Testname: %v
				%+v
				Expect: %v	got: %v
				Expect error:%v	got: %v`, tt.name, tt.header, tt.expectedKey, key, tt.expectErr, err)
				t.Fatalf("unexpected error: %v", err)
			}

			if key != tt.expectedKey {
				t.Fatalf("expected key %q, got %q", tt.expectedKey, key)
			}

			t.Logf(`Testname: %v
				Expect: %v	got: %v`, tt.name, tt.expectedKey, key)

		})
	}
}

// func GetAPIKey(headers http.Header) (string, error) {
// 	authHeader := headers.Get("Authorization")
// 	if authHeader == "" {
// 		return "", ErrNoAuthHeaderIncluded
// 	}
// 	splitAuth := strings.Split(authHeader, " ")
// 	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
// 		return "", errors.New("malformed authorization header")
// 	}

// 	return splitAuth[1], nil
// }
