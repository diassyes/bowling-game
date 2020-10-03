package main

import (
	"bowling/pkg"
	"fmt"
	"testing"
)

type testData struct {
	input        string
	isValidInput bool
	score        int
}

var tests = []testData{
	{"[2,2][2,2][2,2][2,2]", false, -1},
	{"[1,1][1,1][1,1][1,1][1,1][1,1][1,1][1,1][1,1][1,1]", true, 20},
	{"[1,3][2,6][5,2][2,1][0,6][2,2][2,1][0,6][2,2][7,7]", false, -1},
	{"[1,3][2,6][5,2][2,1][0,6][2,2][2,1][0,6][2,2][1,3]", true, 49},
	{"[5,5][5,5][5,5][5,5][5,5][5,5][5,5][5,5][5,5][5,5]", false, -1},
	{"[2,2][2,2][2,2][2,2][2,2][2,2][2,2][2,2][2,2][2,2]", true, 40},
	{"[5,5][5,5][5,5][5,5][5,5][5,5][5,5][5,5][5,5][5,5][5,5]", true, 150},
}

var (
	numOfTests          = len(tests)
	successRate float64 = 0
)

func TestBowling(t *testing.T) {
	for _, test := range tests {
		bg := pkg.NewBowlingGame()
		err := bg.InitBowlingGame(test.input)
		if err != nil {
			if !test.isValidInput {
				passed(test)
				continue
			} else {
				t.Error(
					"For", test.input,
					"expected-isValid", test.isValidInput,
					"got-isValid", false,
				)
			}
		}
		score, err := bg.GetScore()
		if err != nil {
			if !test.isValidInput {
				passed(test)
				continue
			} else {
				t.Error(
					"For", test.input,
					"expected-is-valid", test.isValidInput,
					"got-is-valid", false,
					"expected-score", test.score,
					"got-score", score,
				)
			}
		}
		if score != test.score {
			if test.isValidInput {
				t.Error(
					"For", test.input,
					"expected-is-valid", test.isValidInput,
					"got-is-valid", true,
					"expected-score", test.score,
					"got-score", score,
				)
			}
		}
		passed(test)
	}
	fmt.Printf("\nSuccessRate: %.2f\n", successRate)
}

func passed(t testData) {
	successRate += 100.00/float64(numOfTests)
	fmt.Println("Passed:", t)
}
