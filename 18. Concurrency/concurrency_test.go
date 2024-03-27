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
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := MakeDelayedServer(20 * time.Millisecond)
		fastServer := MakeDelayedServer(0 * time.Millisecond)

		// By prefixing a function call with `defer` it will
		// now call that function at the end of the containing function.
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		server := MakeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func TestCounter(t *testing.T) {
	t.Run("increment the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
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

func MakeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(delay)
		writer.WriteHeader(http.StatusOK)
	}))
}

func assertCounter(t testing.TB, got Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
