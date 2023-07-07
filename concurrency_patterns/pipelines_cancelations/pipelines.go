package main

import "fmt"

/***
A pipeline is a series of stages connected by channels, where each stage is a group of goroutines running the same function. In each stage, the goroutines;
1. receive values from upstream via inbound channels
2. perform some function on that data, usually producing new values
3. send values downstream via outbound channels

src: https://go.dev/blog/pipelines
***/

/*
Concurrent programming is a way to make programs run faster by doing multiple tasks at the same time.
In Go, a programming language, they suggest using channels to share information between these tasks,
rather than sharing variables directly. This method helps prevent errors and makes it easier to write clear programs.
You can think of it like passing notes between people working on a project,
instead of everyone trying to write on the same paper at the same time.
*/
func main() {
	// Set up the pipeline.
	c := gen(2, 3)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
