package main

import (
	"bytes"
	"fmt"
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

func Greet(writer *bytes.Buffer, name string) {
	// instead using fmt.Printf(), we use fmt.Fprintf()
	_, _ = fmt.Fprintf(writer, "Hi, %s", name)
}
