package main

import (
	"fmt"

	"github.com/Nevermind0911/task-2-1/internal/temp"
)

func main() {
	var N int
	fmt.Scan(&N)

	for dept := 0; dept < N; dept++ {
		var K int
		fmt.Scan(&K)
		L := 15
		R := 30
		for emp := 0; emp < K; emp++ {
			var (
				val int
				op  string
			)
			fmt.Scan(&op, &val)

			L, R = temp.UpdateInterval(L, R, op, val)
			opt := temp.GetOptimal(L, R)
			fmt.Println(opt)
		}
	}
}
