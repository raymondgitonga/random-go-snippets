package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	// Note: values in contexts are immutable. This wraps another context around your context it does not overwrite it
	ctx = context.WithValue(ctx, "key", "value")
	doSomething(ctx)

}

// Context with value
func doSomething(ctx context.Context) {
	fmt.Println("Doing something!", ctx.Value("key"))

	// Note: we are wrapping the new context(anotherCtx) around our
	// original not overwriting it
	anotherCtx := context.WithValue(ctx, "key", "anotherValue")

	ctx.Done()
	doAnother(anotherCtx)

	fmt.Println("Doing something!", ctx.Value("key"))
}

// Context with value
func doAnother(ctx context.Context) {
	fmt.Printf("doAnother: myKey's value is %s\n", ctx.Value("key"))
}
