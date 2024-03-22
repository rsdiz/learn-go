package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "other://something.error" {
		return false
	}
	return true
}

/**
If you're lucky, the test will pass, and sometimes test will fail
Sometimes, when we run our tests, two of the goroutines write to the results map at exactly the same time.
Maps in Go don't like it when more than one thing tries to write to them at once, and so fatal error.
This is a race condition, a bug that occurs when the output of our software is dependent on
the timing and sequence of events that we have no control over.
Because we cannot control exactly when each goroutine writes to the results map,
we are vulnerable to two goroutines writing to it at the same time.
Go can help us to spot race conditions with its built in race detector (https://blog.golang.org/race-detector).
To enable this feature, run the tests with the race flag: go test -race.
*/

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

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(20 * time.Millisecond)
		writer.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}))

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := Racer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
