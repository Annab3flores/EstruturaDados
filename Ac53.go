package main

import (
	"fmt"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var PA, PB int
		var G1, G2 float64
		fmt.Scan(&PA, &PB, &G1, &G2)

		years := 0
		for PA <= PB {
			PA += int(float64(PA) * (G1 / 100))
			PB += int(float64(PB) * (G2 / 100))
			years++

			if years > 100 {
				fmt.Println("Mais de 1 seculo.")
				break
			}
		}

		if years <= 100 {
			fmt.Println(years, "anos.")
		}
	}
}
