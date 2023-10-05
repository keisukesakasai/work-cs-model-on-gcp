package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	SleepEnv         = "SLEEP"
	ServerAddressEnv = "SERVER_ADDRESS"
)

func main() {
	url := os.Getenv(ServerAddressEnv)
	sleep, _ := strconv.Atoi(os.Getenv(SleepEnv))

	// Start main task.
	for range time.Tick(time.Duration(sleep) * time.Second) {
		ctx := context.Background()

		err := Request(ctx, url)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Request(ctx context.Context, url string) (err error) {
	client := http.Client{}

	err = func(ctx context.Context) error {
		req, _ := http.NewRequestWithContext(
			ctx,
			"GET",
			url,
			nil,
		)

		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		body, err := io.ReadAll(res.Body)
		_ = res.Body.Close()
		bodyString := string(body)

		fmt.Println(bodyString)

		return err
	}(ctx)

	return err
}
