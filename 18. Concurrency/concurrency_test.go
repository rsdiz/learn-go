package main

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "other://something.error" {
		return false
	}
	return true
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"https://genilogi.id",
		"https://genovatif.com",
		"other://something.error",
	}

	want := map[string]bool{
		"https://genilogi.id":     true,
		"https://genovatif.com":   true,
		"other://something.error": false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
