package bowling

import (
  "fmt"
  "strconv"
)

type Roll uint8
type Score uint16

type BowlingGame struct {
	ScoredRolls [][]Roll
	LiveRolls []Roll
	Frames []Score
}

func (game *BowlingGame) AddRoll(roll Roll) {
	if (len(game.Frames) < 10) {
		game.LiveRolls = append(game.LiveRolls, roll)
		for {
			score, remaining := game.ScoreFrame()
			if (score == nil) {
				break
			}
			currentFrame := len(game.Frames)
			if (currentFrame > 0) {
				game.Frames = append(game.Frames, *score + game.Frames[currentFrame - 1])
			} else {
				game.Frames = append(game.Frames, *score)
			}
			deltaRolls := len(game.LiveRolls) - len(remaining)
			if (deltaRolls > 0) {
				game.ScoredRolls = append(game.ScoredRolls, game.LiveRolls[:deltaRolls])
				game.LiveRolls = remaining
			}
		}
	}	
}

func (game *BowlingGame) ScoreFrame() (*Score, []Roll) {
	score := Score(0)
	balls := game.LiveRolls
	if len(balls) > 2 {
		if (balls[0] == 10) { // strike
			score = Score(int(balls[0]) + int(balls[1]) + int(balls[2]))
			return &score, balls[1:]
		}
		if (balls[0] + balls[1] == 10) { // spare
			score = Score(int(balls[0]) + int(balls[1]) + int(balls[2]))
			return &score, balls[2:]
		}
	} else if (len(balls) > 1 && balls[0] + balls[1] < 10) { // open
		score = Score(int(balls[0]) + int(balls[1]))
		return &score, balls[2:]	
	}
	return nil, []Roll{}		
}

func (game *BowlingGame) Print() {
	fmt.Printf("\nAll Frames = %v\n", game.Frames)
	fmt.Printf("Scored Rolls = %v\n", game.ScoredRolls)
	fmt.Printf("Live Rolls = %v\n", game.LiveRolls)
}

func ScoreRolls(rolls []string) {
	game := BowlingGame{
		ScoredRolls: [][]Roll{},
		LiveRolls: []Roll{},
		Frames: []Score{},
	}

	for _, arg := range rolls {
		roll, err := strconv.Atoi(arg)
		if (err == nil) {
			game.AddRoll(Roll(roll))
		}
	}	
	game.Print()
}