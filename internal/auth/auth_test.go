package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	for scenario, data := range map[string]struct {
		header http.Header
		want   string
	}{
		"12345678": {header: http.Header{
			"Authorization": []string{"ApiKey 12345678"},
		}, want: "12345678"},
	} {
		t.Run(scenario, func(t *testing.T) {
			got, err := GetAPIKey(data.header)
			if err != nil {
				t.Fatalf("expected: %v, got error: %v", data.want, err)
			}
			if got != data.want {
				t.Fatalf("expected: %v, got: %v", data.want, got)
			}
		})
	}
}
