package main

import (
	"fmt"
	"net/http"
	"time"
)

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

Channels.
After we know that is a race condition in goroutine, then how to solve this?
We can solve this data race by coordinating our goroutines using channels.
Channels are a Go data structure that can both receive and send values.
These operations, along with their details, allow communication between different processes.

In this case we want to think about the communication between the parent process and each of the goroutines
that it makes to do the work of running the WebsiteChecker function with the url.

Select.
Ok, so in concurrency, we have statement 'select'. What is that?
Before we learn about select, lets create a function that can help us to understand how select works.
We will create function called 'WebsiteRacer' which takes two URLs and "races" them by hitting them with an HTTP GET
and returning the URL which returned first. If none of them return within 10 seconds then it should return an 'error'.

Ok, we have been created the function, but the problem is we're reaching out to real websites to test our own logic.
Testing code that uses HTTP is so common that Go has tools in the standard library to help you test it.
In the standard library, there is a package called 'net/http/httptest' which enables users to easily create a mock HTTP server.

Ok, finally we enter how to use 'Select',
in this section:
- Why are we testing the speeds of the websites one after another when Go is great at concurrency? We should be able to check both at the same time.
- We don't really care about the exact response times of the requests, we just want to know which one comes back first.
To do that, we need to use 'Select' which helps us synchronise processes really easily and clearly.

Ok, then our final requirement was to return an error if Racer takes longer than 10 seconds.
By adding case to handle this requirement, it will make new problem, the test will takes 10 seconds to run.
What we can do is make the timeout configurable. So in our test, we can have a very short timeout
and then when the code is used in the real world it can be set to 10 seconds.
*/

type (
	WebsiteChecker func(string) bool
	// result has been made to associate the return value of the WebsiteChecker with the url being checked
	result struct {
		string
		bool
	}
)

var tenSecondTimeout = 10 * time.Second

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	// Alongside the `results` map we now have a `resultChannel`, which we make in the same way.
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// Give each anonymous function a parameter for the url (parameter "u"),
		// and then calling the anonymous function with the url as the argument
		go func(u string) {
			// when we iterate over the urls, instead of writing to the map directly we're sending a result struct
			// for each call to wc to the `resultChannel` with a send statement.
			// This uses the <- operator, taking a channel on the left and a value on the right
			resultChannel <- result{
				string: u,
				bool:   wc(u),
			}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// receive expressions from `resultChannel`, which assigns a value received from a channel to a variable.
		// the channel is on the right and the variable that we're assigning to is on the left:
		r := <-resultChannel

		// By sending the results into a channel, we can control the timing of each write into the results map, ensuring that it happens one at a time.
		// Although each of the calls of wc, and each send to the result channel, is happening concurrently inside its own process,
		// each of the results is being dealt with one at a time as we take values out of the result channel with the `receive expression`.
		results[r.string] = r.bool
	}

	return results
}

/*
*
Back to concurrency channel, we can wait for values to be sent to a channel with '<-ch'. This is a blocking call, as you're waiting for a value.
'select' allows you to wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
We use ping in our select to set up two channels, one for each of our URLs. Whichever one writes to its channel first
will have its code executed in the select, which results in its URL being returned (and being the winner).
After these changes, the intent behind our code is very clear and the implementation is actually simpler.
*/
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

/*
*
We have defined a function ping which creates a chan struct{} and returns it.
In our case, we don't care what type is sent to the channel, we just want to signal we are done and closing the channel works perfectly!
Why struct{} and not another type like a bool? Well, a chan struct{} is the smallest data type available from a memory perspective,
so we get no allocation versus a bool. Since we are closing and not sending anything on the chan, why allocate anything?
Inside the same function, we start a goroutine which will send a signal into that channel once we have completed http.Get(url).
*/
func ping(url string) chan struct{} {
	ch := make(chan struct{}) // create a new channel (zero value is 'nil')
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
