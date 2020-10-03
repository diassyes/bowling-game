package main

import (
	"bowling/pkg"
	"fmt"
)

func main() {

	var game = "[1,3][2,6][5,2][2,1][0,6][2,2][2,1][0,6][2,2][1,3]"

	bg := pkg.NewBowlingGame()
	err := bg.InitBowlingGame(game)
	if err != nil {
		panic(err)
	}
	score, err := bg.GetScore()
	if err != nil {
		panic(err)
	}
	fmt.Println("score is:", score)
}
