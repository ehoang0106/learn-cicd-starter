package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyReturnsKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test-key")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if apiKey != "test-key" {
		t.Fatalf("expected api key %q, got %q", "test-key", apiKey)
	}
}

func TestGetAPIKeyReturnsErrorWhenHeaderMissing(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected an error when Authorization header is missing")
	}

	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}
