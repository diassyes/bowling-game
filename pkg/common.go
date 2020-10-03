package pkg

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type BowlingGame struct {
	rounds []Frame
}

var validateErr = errors.New("doesn't match regex")

func NewBowlingGame() *BowlingGame {
	bg := BowlingGame{}
	bg.rounds = make([]Frame, 0)
	return &bg
}

// Expected value: "[n,n],[n,n],[n,n],...[n,n]"
func (bg *BowlingGame) InitBowlingGame(game string) error {
	expr := "^(\\[([0-9],[0-9]|10,0|0,10)\\]){10}((\\[(([0-9],[0-9])|(10,0)|(0,10)|(10,10)|(10,15)|(15,10))\\])|(\\[(([0-9])|(10))\\]))?$"
	re, err := regexp.Compile(expr)
	if err != nil {
		return err
	}
	if !re.MatchString(game) {
		return validateErr
	}
	// [1,3][2,6][5,2][2,1][0,6][2,2][2,1][0,6][2,2][1,3]
	//parts := strings.Split(game, "[")
	parts := regexp.MustCompile("\\[(\\d+,\\d+)\\]").FindAllString(game, 1000)
	for i := range parts {
		nums := strings.ReplaceAll(parts[i], "[", "")
		nums = strings.ReplaceAll(nums, "]", "")
		numbers := strings.Split(nums, ",")
		num1 := numbers[0]
		num2 := numbers[1]
		if len(numbers) == 1 {
			first, err := strconv.Atoi(num1)
			if err != nil {
				return err
			}
			bg.rounds = append(bg.rounds, NewFrame(first, 0))
		} else {
			a, err := strconv.Atoi(num1)
			if err != nil {
				return err
			}
			b, err := strconv.Atoi(num2)
			if err != nil {
				return err
			}
			if i == 11 {
				if a > 10 || b > 10 {
					return validateErr
				}
			} else if a+b > 10 {
				return validateErr
			}
			bg.rounds = append(bg.rounds, NewFrame(a, b))
		}
	}
	return nil
}

func (bg *BowlingGame) GetScore() (int, error) {
	size := len(bg.rounds)
	if size == 0 {
		return -1, errors.New("no rounds inited")
	}
	sum := 0
	for i := 0; i < 10; i++ {
		if bg.rounds[i].IsStrike() {
			if i < size-1 {
				if i == 9 {
					if size == 11 {
						sum += bg.rounds[i+1].First + bg.rounds[i+1].Second
					}
				} else if bg.rounds[i+1].First == 10 {
					if i < size - 2 {
						sum += bg.rounds[i+2].First
					} else {
						return -1, errors.New("next next values doesn't exists")
					}
					sum += bg.rounds[i+1].First
				} else {
					sum += bg.rounds[i+1].First + bg.rounds[i+1].Second
				}
			} else {
				return -1, errors.New("missing bonus round")
			}
			sum += 10
		} else if bg.rounds[i].IsSpare() && !bg.rounds[i].IsStrike() {
			sum += 10
			if i < size - 1 {
				sum += bg.rounds[i+1].First
			} else {
				return -1, errors.New("missing bonus round")
			}
		} else {
			sum += bg.rounds[i].First + bg.rounds[i].Second
		}
	}
	return sum, nil
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

