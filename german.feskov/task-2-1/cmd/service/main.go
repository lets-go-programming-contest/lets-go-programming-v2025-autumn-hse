package main

import (
	"context"
	"fmt"

	"github.com/6ermvH/german.feskov/task-2-1/internal/temperature"
	"golang.org/x/sync/errgroup"
)

const (
	bufChanSize = 1
)

func main() {
	var countDep int

	if _, err := fmt.Scan(&countDep); err != nil {
		fmt.Println(err)

		return
	}

	for range countDep {
		var countWorker int
		if _, err := fmt.Scan(&countWorker); err != nil {
			fmt.Println(err)

			return
		}

		group, _ := errgroup.WithContext(context.Background())

		requests := func(group *errgroup.Group) chan temperature.Request {
			requests := make(chan temperature.Request, bufChanSize)

			group.Go(func() error {
				defer close(requests)

				return read(requests, countWorker)
			})

			return requests
		}(group)

		calculated := func(group *errgroup.Group, chOut <-chan temperature.Request) chan int {
			calculated := make(chan int, bufChanSize)

			group.Go(func() error {
				defer close(calculated)

				return temperature.Calculate(chOut, calculated)
			})

			return calculated
		}(group, requests)

		group.Go(func() error {
			return write(calculated)
		})

		if err := group.Wait(); err != nil {
			fmt.Println(err)

			return
		}
	}
}

func read(chIn chan<- temperature.Request, count int) error {
	var req temperature.Request
	for range count {
		if _, err := fmt.Scanf("%s %d", &req.Type, &req.Val); err != nil {
			return fmt.Errorf("while scan, has %w", err)
		}
		chIn <- req
	}

	return nil
}

func write(chOut <-chan int) error {
	for res := range chOut {
		if _, err := fmt.Println(res); err != nil {
			return fmt.Errorf("while print, has %w", err)
		}
	}

	return nil
}
