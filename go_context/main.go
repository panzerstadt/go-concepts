package gocontext

import (
	"context"
	"fmt"
	"log"
	"time"
)

func TestContext() {
	start := time.Now()

	ctx := context.WithValue(context.Background(), "requestID", "12339287")
	userID := 10
	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result: ", val)
	fmt.Println("took: ", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	val := ctx.Value("requestID")
	fmt.Println("request id:", val)
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel() // cancel if fetch finishes before this runs

	responseChannel := make(chan Response)

	go func() {
		val, err := fetchSlowStuff()
		responseChannel <- Response{value: val, err: err}
	}()

	for {
		select {
		case <-ctx.Done(): // if ctx completed (200 mills has elapsed)
			return 0, fmt.Errorf("fetching data from third party took longer than expected")
		case resp := <-responseChannel:
			return resp.value, resp.err
		}
	}
}

func fetchSlowStuff() (int, error) {
	// time.Sleep(time.Millisecond * 500)
	time.Sleep(time.Millisecond * 150)
	return 666, nil
}
