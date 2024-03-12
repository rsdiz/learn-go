package main

/**
Concurrency.
Before we start learning about concurrency, we will create function to checks the status of a list of URLs.
By using Dependency Injection, we are allow to test the function without making real HTTP calls
*/

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
