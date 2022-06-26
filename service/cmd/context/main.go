package main

import (
	"context"
	"fmt"
	"time"
)

type paramKey struct {
}

func main() {
	c, cancel := context.WithTimeout(context.Background(),
		5*time.Second)
	defer cancel()
	mainTask(c)
}

func mainTask(c context.Context) {
	fmt.Printf("main task started with param %q\n", c.Value(paramKey{}))
	smallTask(c, "task1")
	smallTask(c, "task2")
}

func smallTask(c context.Context, name string) {
	fmt.Printf("small task started with param %q\n", c.Value(paramKey{}))
	fmt.Printf("%s started\n", name)

	select {
	case <-time.After(3 * time.Second):
		fmt.Printf("down")
	case <-c.Done():
		fmt.Printf("cancel")
	}
}
