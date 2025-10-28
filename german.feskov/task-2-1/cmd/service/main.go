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
		fmt.Printf("failed to read count of departaments: %v\n", err)

		return
	}

	for range countDep {
		var countWorker int
		if _, err := fmt.Scan(&countWorker); err != nil {
			fmt.Printf("failed to read count of workers: %v\n", err)

			return
		}

		group, _ := errgroup.WithContext(context.Background())

		requests := func() chan temperature.Request {
			requests := make(chan temperature.Request, bufChanSize)

			group.Go(func() error {
				defer close(requests)

				return read(requests, countWorker)
			})

			return requests
		}()

		calculated := func(chOut <-chan temperature.Request) chan int {
			calculated := make(chan int, bufChanSize)

			group.Go(func() error {
				defer close(calculated)

				return temperature.Calculate(chOut, calculated)
			})

			return calculated
		}(requests)

		group.Go(func() error {
			return write(calculated)
		})

		if err := group.Wait(); err != nil {
			fmt.Printf("temperature calculated: %v\n", err)

			return
		}
	}
}

func read(chIn chan<- temperature.Request, count int) error {
	var req temperature.Request
	for range count {
		if _, err := fmt.Scanf("%s %d", &req.Type, &req.Val); err != nil {
			return fmt.Errorf("scan temp request: %w", err)
		}
		chIn <- req
	}

	return nil
}

func write(chOut <-chan int) error {
	for res := range chOut {
		fmt.Println(res)
	}

	return nil
}
