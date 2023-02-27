package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var nbDimensions int
	nbDimensions = rand.Intn(10) + 3
	var tableau [10][10][10]int
	for i := 0; i < nbDimensions; i++ {
		for j := 0; j < nbDimensions; j++ {
			for k := 0; k < nbDimensions; k++ {
				tableau[i][j][k] = rand.Intn(100)
			}
		}
	}
	fmt.Println("Le tableau est de", nbDimensions, "dimensions")
	for i := 0; i < nbDimensions; i++ {
		for j := 0; j < nbDimensions; j++ {
			for k := 0; k < nbDimensions; k++ {
				fmt.Print(tableau[i][j][k], " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
