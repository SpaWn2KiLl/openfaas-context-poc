package function

import (
	"context"
	"fmt"
	handler "github.com/openfaas-incubator/go-function-sdk"
	"time"
)

// Handle a function invocation
func Handle(ctx context.Context, req handler.Request) (handler.Response, error) {
	var (
		response handler.Response
		err      error
	)

	printContextStatus(ctx)

	time.Sleep(5 * time.Second)

	printContextStatus(ctx)

	return response, err
}

func printContextStatus(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Context closed")
		fmt.Println(ctx.Err())
	default:
		fmt.Println("Context is still valid")
	}
}
