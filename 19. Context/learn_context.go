package main

import (
	"fmt"
	"net/http"
)

/**
Software often kicks off long-running, resource-intensive processes (often in goroutines).
If the action that caused this gets cancelled or fails for some reason you need to stop these processes
in a consistent way through your application.

If you don't manage this your snappy Go application that you're so proud of could start having difficult
to debug performance problems.

In this chapter we'll use the package 'context' to help us manage long-running processes.

We're going to start with a classic example of a web server that when hit kicks off a
potentially long-running process to fetch some data for it to return in the response.

Context.
Context has a method 'Done()' which return a channel which gets sent a signal when the context is "done" or "cancelled".
*/

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// We want to listen to that signal and call 'store.Cancel' if we get it,
		// but we want to ignore it if our 'Store' manages to 'Fetch' before it.
		ctx := request.Context()

		// To manage this we run 'Fetch' in a goroutine, and it will write the result into a new channel 'data'
		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		// Then use select to effectively race to the two asynchronous processes and then we either write a response or 'Cancel'.
		select {
		case d := <-data:
			fmt.Fprint(writer, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
