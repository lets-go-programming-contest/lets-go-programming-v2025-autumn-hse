package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/6ermvH/lets-go-programming-v2025-autumn-hse/german.feskov/task-2-1/internal/temperature"
	"golang.org/x/sync/errgroup"
)

const (
	bufChanSize = 128
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)

	defer func() {
		if err := stdout.Flush(); err != nil {
			fmt.Println(err)
		}
	}()

	var countDep int

	if _, err := fmt.Fscan(stdin, &countDep); err != nil {
		fmt.Println(err)

		return
	}

	for range countDep {
		var countWorker int
		if _, err := fmt.Fscan(stdin, &countWorker); err != nil {
			fmt.Println(err)

			return
		}

		var (
			requests   = make(chan temperature.Request, bufChanSize)
			calculated = make(chan int, bufChanSize)
			group, _   = errgroup.WithContext(context.Background())
		)

		group.Go(func() error {
			return temperature.Read(stdin, requests, countWorker)
		})
		group.Go(func() error {
			return temperature.Calculate(requests, calculated)
		})
		group.Go(func() error {
			return temperature.Write(stdout, calculated)
		})

		if err := group.Wait(); err != nil {
			fmt.Println(err)

			return
		}
	}
}
