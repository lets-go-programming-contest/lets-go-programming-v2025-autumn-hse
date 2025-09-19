package main

import "fmt"

func main() {
	var (
		n       int
		k       int
		request string
		reqVal  int
	)
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	for range n {
		if _, err := fmt.Scan(&k); err != nil {
			return
		}

		var (
			minT int = 0
			maxT int = 1000
		)

		for range k {
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
	}

}
