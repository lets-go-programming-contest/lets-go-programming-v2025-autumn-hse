package main

import "fmt"

func main() {
	var (
		depCount    int
		workerCount int
		request     string
		reqVal      int
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

		var (
			minT = 15
			maxT = 30
		)

		for range workerCount {
			if _, err := fmt.Scanf("%s %d", &request, &reqVal); err != nil {
				fmt.Println(err)

				return
			}

			switch request {
			case ">=":
				minT = maxInt(minT, reqVal)
			case "<=":
				maxT = minInt(maxT, reqVal)
			}

			if minT > maxT {
				fmt.Println(-1)
			} else {
				fmt.Println(minT)
			}
		}
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}
