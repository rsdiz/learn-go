package main

import "time"

/**
Concurrency.
Before we start learning about concurrency, we will create function to checks the status of a list of URLs.
By using Dependency Injection, we are allow to test the function without making real HTTP calls.

After we create that function and create unit test and benchmark test,
now we can talk about concurrency, we will update the code to implement concurrency to make CheckWebsites faster.
Normally in Go when we call function, we wait for it to return (even if it has no value to return, we still wait for it to finish).
We say that this operation is blocking - it makes us wait for it to finish.
An operation that does not block in Go will run in a separate process called a goroutine.
Think of a process as reading down the page of Go code from top to bottom, going 'inside' each function
when it gets called to read what it does. When a separate process starts, it's like another reader begins reading inside the function,
leaving the original reader to carry on going down the page.

To make a goroutine we turn a function call into a 'go' statement using the keyword 'go' in front of it: 'go functionName()'
*/

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		// Give each anonymous function a parameter for the url (parameter "u"),
		// and then calling the anonymous function with the url as the argument
		go func(u string) {
			results[u] = wc(u)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return results
}
