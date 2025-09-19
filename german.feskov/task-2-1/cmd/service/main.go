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
		return
	}

	for range depCount {
		if _, err := fmt.Scan(&workerCount); err != nil {
			return
		}

		var (
			minT = 15
			maxT = 30
		)

		for range workerCount {
			if _, err := fmt.Scanf("%s %d", &request, &reqVal); err != nil {
				return
			}
			if request == ">=" && reqVal > minT {
				minT = reqVal
			}
			if request == "<=" && reqVal < maxT {
				maxT = reqVal
			}

			if minT > maxT {
				fmt.Println(-1)
			} else {
				fmt.Println(minT)
			}
		}
		fmt.Println()
	}
}
