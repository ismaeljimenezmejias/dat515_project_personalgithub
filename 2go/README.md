# Lab 2: Introduction to Go Programming

| Lab 2:           | Introduction to Go Programming |
| ---------------- | ------------------------------ |
| Subject:         | DAT515 Cloud Computing         |
| Deadline:        | **August 29, 2025 23:59**      |
| Expected effort: | 4-10 hours                     |
| Grading:         | Pass/fail                      |
| Submission:      | Individually                   |

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Introduction](#introduction)
- [Go Resources](#go-resources)
  - [Primary Resources](#primary-resources)
  - [Additional Resources](#additional-resources)
  - [Using Coding Models](#using-coding-models)
- [Tasks](#tasks)
  - [Task 1: Exercises from Tour of Go](#task-1-exercises-from-tour-of-go)
  - [Task 2: Go Exercises](#task-2-go-exercises)
  - [Task 3: Cloud-based Go Programming](#task-3-cloud-based-go-programming)
  - [Submission](#submission)

## Introduction

In this lab you will be learning the Go programming language.
Go was designed by Google engineers to be a simple, efficient, and powerful language for building scalable software.
It is particularly well-suited for cloud computing and distributed systems.

Go provides built in primitives to design concurrent programs using lightweight threads, called goroutines.
In addition, these goroutines can communicate using channels, instead of by means of shared memory which is the most common approach in other languages such as C and Java.
Go is also a very small language in the sense that it has very few keywords, and as such it does not take too long to learn the fundamentals of the language.

## Go Resources

Below we provide some resources which will be useful when you start programming in Go.

### Primary Resources

[The Go Language web page](https://go.dev) provides a variety of helpful documentation, including:

- Learning the basics with ["A Tour of Go"](https://go.dev/tour/)
- [The Go documentation](https://go.dev/doc/)
- [Standard library package documentation](https://pkg.go.dev/std)

### Additional Resources

- [Effective Go](https://go.dev/doc/effective_go) gives tips for writing clear, idiomatic Go code
- [Frequently Asked Questions](https://go.dev/doc/faq)
- [The Go Blog](https://go.dev/blog/all) contains many great articles that describe important idioms and intricacies of the language.
- [Collection of Videos about Go](https://go.dev/wiki/GoTalks)

**Troubleshooting tip:**
When searching for information about Go, use the term _golang_ instead of go.

### Using Coding Models

You are free to use ChatGPT or other AI tools to help you with Go programming, but we **strongly recommend** that you spend some time learning the language fundamentals without relying on AI-generated code.

One particular thing to note about using AI tools is that they may generate code that is not idiomatic Go code (not the style that Go programmers expect).
In addition, generated code has a tendency to be biased towards legacy practices and may not follow the latest best practices in Go programming, even though the AI model was trained on recent data.

## Tasks

If you have programmed in Go before, you may skip the exercises in this section, unless you want to refresh your knowledge of the language.

### Task 1: Exercises from Tour of Go

Start learning the basics of Go by completing ["A Tour of Go"](http://go.dev/tour/).
You should do at least the following exercises.

- [Exercise: Loops and Functions](https://go.dev/tour/flowcontrol/8)
- [Exercise: Slices](https://go.dev/tour/moretypes/18)
- [Exercise: Maps](https://go.dev/tour/moretypes/23)
- [Exercise: rot13Reader](https://go.dev/tour/methods/23)

Note that you can change the code inline in the browser and run the code to see the results.

### Task 2: Go Exercises

Before you start working on the assignments below, make sure that your local working copy has all the latest changes from the course [assignments](https://github.com/dat515-2025/assignments) repository.
Here are the instructions for [fetching the latest changes](https://github.com/dat515-2025/info/blob/main/lab-submission.md#update-local-working-copy-from-course-assignments).

If you are new to Go or have little programming experience, we recommend that you complete some of the following exercises.
If you have programmed in Go before, you may skip these exercises.

- `stringer` (easy)
- `sequence` (easy)
- `collect` (recommended)
- `cipher`
- `errors` (less important)
- `multiwriter` (less important)

1. In the following, we will use the `sequence/fibonacci.go` exercise as an example.
   The file contains the following skeleton code and task description:

   ```golang
   package sequence

   // Task: Fibonacci numbers
   //
   // fibonacci(n) returns nth Fibonacci number, and is defined by the
   // recurrence relation F_n = F_n-1 + F_n-2, with seed values F_0=0 and F_1=1.
   func fibonacci(n uint) uint {
     return 0
   }
   ```

2. Implement the function body according to the specification so that all the tests in `sequence/fibonacci_test.go` passes.
   The test file looks like this:

   ```golang
   package sequence

   import "testing"

   func TestFibonacci(t *testing.T) {
       var fibonacciTests = []struct {
           in, want uint
       }{
           {0, 0},
           {1, 1},
           {2, 1},
           {3, 2},
           {4, 3},
           {5, 5},
           {6, 8},
           {7, 13},
           {8, 21},
           {9, 34},
           {10, 55},
           {20, 6765},
       }
       for _, ft := range fibonacciTests {
           got := fibonacci(ft.in)
           if got != ft.want {
               t.Errorf("fibonacci(%d) = %d, want %d", ft.in, got, ft.want)
           }
       }
   }
   ```

3. There are several ways to run the tests. If you run:

   ```console
   go test
   ```

   the Go tool will run all tests found in files whose file name ends with `_test.go` (in the current directory).
   Similarly, you can also run a specific test as follows:

   ```console
   go test -run TestFibonacci
   ```

   If you run:

   ```console
   go test -v
   ```

   the Go tool will run all tests and print the output of each test.

4. You should **_not_** edit files or code that are marked with a `// DO NOT EDIT` comment.
   Please make separate `filename_test.go` files if you wish to write and run your own tests.
   Note that your own tests will not be run by QuickFeed.

5. When you have completed a task and sufficiently many local tests pass, you may push your code to GitHub.
   This will trigger QuickFeed which will then run a separate test suite on your code.

   Using `sequence/fibonacci.go` as an example, use the following procedure to commit and push your changes to GitHub and QuickFeed:

   ```console
   $ git add fibonacci.go
   $ git commit -m "feat(lab1): implemented the Fibonacci assignment"
   # Or just use `git commit` to open an editor to write a commit message:
   $ git commit
   $ git push
   ```

6. QuickFeed will now build and run a test suite on the code you submitted.
   You can check the output by going to the [QuickFeed web interface](https://uis.itest.run).
   The results (build log) is available from the Labs menu.
   Note that the results shows the output for all the tests in current lab assignment.
   You will want to focus on the output for the specific test results related to the task you're working on.

7. Follow the same process for the other tasks included in this lab assignment.
   Each task contains a single `.go` template file, along with a task description and a `_test.go` file with tests.

### Task 3: Cloud-based Go Programming

In addition to the basic Go exercises above, this lab includes slightly more advanced exercises designed to prepare you for cloud computing concepts and technologies you'll encounter later in the course.
These exercises focus on patterns and techniques commonly used in cloud applications:

#### JSON Handling (`jsonhandler/`)

Learn to work with JSON data, which is essential for REST APIs and cloud service communication.
This exercise covers marshaling and unmarshaling Go structs to/from JSON format.
Here is a blog post that explains the [JSON handling in Go](https://go.dev/blog/json-and-go).

#### HTTP Client (`httpclient/`)

Practice making HTTP requests to APIs, implementing health checks, and handling custom headers.
These skills are fundamental for microservice communication and API integration in cloud environments.

<!-- Here is a blog post that explains the [HTTP client in Go](https://go.dev/doc/articles/wiki/). -->

#### Context Operations (`contextops/`)

Master Go's context package for managing timeouts, cancellation, and request-scoped values.
Context is crucial for building robust cloud services that can handle timeouts and graceful shutdowns.
Here is a blog post that explains the [context operations in Go](https://go.dev/blog/context).

#### Concurrent Programming (`concurrent/`)

Explore goroutines and channels for building scalable concurrent applications.
Learn worker pools, rate limiting, and fan-out/fan-in patterns commonly used in cloud services for handling multiple requests simultaneously.
Here is a talk that explains [concurrent programming in Go](https://go.dev/blog/concurrency-is-not-parallelism) and the differences between concurrency and parallelism.
Here are the [slides for the talk](https://talks.golang.org/2012/concurrency.slide).

#### Interface Composition (`interfaces/`)

Design clean, testable architectures using Go interfaces.
This exercise demonstrates dependency injection and service composition patterns that are essential for maintainable cloud applications.

<!-- Here is a blog post that explains [interface composition in Go](https://go.dev/blog/interfaces). -->

These exercises build upon each other and introduce concepts that will be directly applicable to later labs involving Docker containers, Kubernetes, REST APIs, gRPC services, and database operations.

### Submission

When you are done with all assignments and want to submit the final version, please follow these [instructions](https://github.com/dat515-2025/info/blob/main/lab-submission.md#final-submission-of-labx).
