package main

import (
	"fmt"

	tc "github.com/6ermvH/task-2-1/internal/temperature_controller"
)

func main() {
	var (
		depCount    int
		workerCount int
	)

	if _, err := fmt.Scan(&depCount); err != nil {
		fmt.Println(err)

		return
	}

	for range depCount {
		if _, err := fmt.Scan(&workerCount); err != nil {
			fmt.Println(err)

			return
		}

		requests := make([]tc.Request, 0)

		for range workerCount {
			var request tc.Request

			if _, err := fmt.Scanf("%s %d", &request.Type, &request.Val); err != nil {
				fmt.Println(err)

				return
			}

			requests = append(requests, request)
		}

		optimalTemperatures, err := tc.GetOptimalTemperatures(requests)
		if err != nil {
			fmt.Println(err)

			return
		}

		for _, val := range optimalTemperatures {
			fmt.Println(val)
		}
	}
}
