package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "http scheme",
			inputURL: "http://example.com/foo",
			expected: "example.com/foo",
		},
		{
			name:     "no path",
			inputURL: "https://example.com",
			expected: "example.com",
		},
		{
			name:     "with port",
			inputURL: "https://example.com:8080/path",
			expected: "example.com:8080/path",
		},
		{
			name:     "with query",
			inputURL: "https://example.com/path?foo=bar",
			expected: "example.com/path?foo=bar",
		},
		{
			name:     "with fragment",
			inputURL: "https://example.com/path#section",
			expected: "example.com/path#section",
		},
		{
			name:     "trailing slash",
			inputURL: "https://example.com/path/",
			expected: "example.com/path/",
		},
		{
			name:     "uppercase host",
			inputURL: "https://EXAMPLE.com/Path",
			expected: "EXAMPLE.com/Path",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
