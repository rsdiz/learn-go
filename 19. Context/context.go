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
*/

type Store interface {
	Fetch() string
}

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, store.Fetch())
	}
}
