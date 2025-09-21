package main

import "fmt"

const (
	minTemperature = 15
	maxTemperature = 30
)

func main() {
	var (
		depCount    int
		workerCount int
		reqOp       string
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
			minT = minTemperature
			maxT = maxTemperature
		)

		for range workerCount {
			for tryAgain := false; !tryAgain; {
				if _, err := fmt.Scanf("%s %d", &reqOp, &reqVal); err != nil {
					continue
				}

				switch reqOp {
				case ">=":
					minT = maxInt(minT, reqVal)
					tryAgain = true
				case "<=":
					maxT = minInt(maxT, reqVal)
					tryAgain = true
				default:
					continue
				}
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
