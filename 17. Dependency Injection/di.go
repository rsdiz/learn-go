package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

/**
Dependency Injection.

This just guide to show how:
- You don't need a framework,
- It does not complicate your design,
- It facilitates testing,
- It allows you to write great, general-purpose functions.

Back to Lesson 1, printing hello world.
In that lesson, we use fmt.Println() or fmt.Printf() to stdout, but how we can test this?
It's pretty hard for us to capture using the testing framework.
What we need is inject (in other words, pass in) the dependency of printing.
Our function doesn't need to care where or how the printing happens, so we should accept an interface rather than a concrete type.
If we do that, we can change the implementation to print something we control so that we can test it.

Back to fmt.Printf(), if we look at code behind that function,
it just calls Fprintf passing in os.Stdout,
So what is an os.Stdout?

Let's check out Fprintf, we can see that the first argument (os.Stdout) is an io.Writer
From this, we can infer that os.Stdout implements io.Writer, Printf passes os.Stdout to Fprintf which expects an io.Writer

So we know under the covers we're ultimately using Writer to send our greetings somewhere.
Let's use this existing abstraction to make our code testable and more reusable.
*/

// Greet , This is technically correct, but not very useful.
// We can change the expected argument from bytes.Buffer to io.Writer
// By change our code to use the more general purpose interface
// we can now use it in both tests and in our application.
func Greet(writer io.Writer, name string) {
	/**
	instead using fmt.Printf(), we use fmt.Fprintf()
	Fprintf allows us to inject an io.Writer which we know both os.Stdout and bytes.Buffer implement.
	*/
	_, _ = fmt.Fprintf(writer, "Hi, %s", name)
}

/**
Mocking.
Before we learn about mocking, let's code an example of function to make a countdown,
we will break down this function into 3 steps:
- Print 3
- Print 3, 2, 1, and Go!
- Wait a second between each line

After we finish creating that function, it's works and the tests is past,
but we have some problem:
- Tests take 3 second to run
- We have not tested an important property of our function
So, what the solution for this problem?
Let's see dependency in time.Sleep(), after that we need to extract it, so we can control it in our tests.
If we can mock time.Sleep we can use dependency injection to use it instead of a "real" time.Sleep and
then we can spy on the calls to make assertions on them.

After we're successfully mocking dependency, we still have some problems:
we haven't tested the important property, Countdown() should sleep before each next print, eg:
- Print N
- Sleep
- Print N-1
- Sleep
- Print Go!
Our latest change only asserts that it has slept 3 times, but those sleeps could occur out of sequence.
Let's use spying again with a new test to check the order of operations is correct.
*/

const (
	finalWord      = "Go!"
	countdownStart = 3
	write          = "write"
	sleep          = "sleep"
)

/**
Define dependency as an interface
*/

type (
	// Sleeper is interface for dependency
	// This lets us then use a real Sleeper in main and a spy sleeper in our tests
	Sleeper interface {
		Sleep()
	}

	// DefaultSleeper real sleeper which implements the interface we need
	DefaultSleeper struct{}

	// SpyCountdownOperations Record all operations into one line
	SpyCountdownOperations struct {
		Calls []string
	}
)

// Sleep , call time.Sleep(1000) for real application
func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Sleep , append "sleep" into list
// This method implements Sleeper
func (o *SpyCountdownOperations) Sleep() {
	o.Calls = append(o.Calls, sleep)
}

// Write , append "write" into list
// This method implements io.Writer
func (o *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	o.Calls = append(o.Calls, write)
	return
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		_, _ = fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	_, _ = fmt.Fprint(out, finalWord)
}

func main() {
	Greet(os.Stdout, "Rosyid")
	fmt.Println()
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
