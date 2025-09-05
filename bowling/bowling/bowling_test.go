package bowling

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestScoreFrame(t *testing.T) {
	game := BowlingGame{
		ScoredRolls: [][]Roll{},
		LiveRolls: []Roll{},
		Frames: []Score{},
	}

	game.AddRoll(0)
	assert.Equal(t, 0, len(game.Frames))
	assert.Equal(t, 0, len(game.ScoredRolls))
	assert.Equal(t, []Roll{0}, game.LiveRolls)

	game.AddRoll(9)
	assert.Equal(t, []Score{9}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}}, game.ScoredRolls)
	assert.Equal(t, []Roll{}, game.LiveRolls)

	game.AddRoll(9)
	assert.Equal(t, []Score{9}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}}, game.ScoredRolls)
	assert.Equal(t, []Roll{9}, game.LiveRolls)

	game.AddRoll(1)
	assert.Equal(t, []Score{9}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}}, game.ScoredRolls)
	assert.Equal(t, []Roll{9, 1}, game.LiveRolls)

	game.AddRoll(10)
	assert.Equal(t, []Score{9, 29}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}, {9, 1}}, game.ScoredRolls)
	assert.Equal(t, []Roll{10}, game.LiveRolls)

	game.AddRoll(10)
	assert.Equal(t, []Score{9, 29}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}, {9, 1}}, game.ScoredRolls)
	assert.Equal(t, []Roll{10, 10}, game.LiveRolls)

	game.AddRoll(10)
	assert.Equal(t, []Score{9, 29, 59}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}, {9, 1}, {10}}, game.ScoredRolls)
	assert.Equal(t, []Roll{10, 10}, game.LiveRolls)

	game.AddRoll(8)
	assert.Equal(t, []Score{9, 29, 59, 87}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}, {9, 1}, {10}, {10}}, game.ScoredRolls)
	assert.Equal(t, []Roll{10, 8}, game.LiveRolls)

	game.AddRoll(1)
	assert.Equal(t, []Score{9, 29, 59, 87, 106, 115}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}, {9, 1}, {10}, {10}, {10}, {8,1}}, game.ScoredRolls)
	assert.Equal(t, []Roll{}, game.LiveRolls)

	game.AddRoll(0)
	assert.Equal(t, []Score{9, 29, 59, 87, 106, 115}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}, {9, 1}, {10}, {10}, {10}, {8,1}}, game.ScoredRolls)
	assert.Equal(t, []Roll{0}, game.LiveRolls)

	game.AddRoll(6)
	assert.Equal(t, []Score{9, 29, 59, 87, 106, 115, 121}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}, {9, 1}, {10}, {10}, {10}, {8,1}, {0, 6}}, game.ScoredRolls)
	assert.Equal(t, []Roll{}, game.LiveRolls)

	game.AddRoll(7)
	assert.Equal(t, []Score{9, 29, 59, 87, 106, 115, 121}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}, {9, 1}, {10}, {10}, {10}, {8,1}, {0, 6}}, game.ScoredRolls)
	assert.Equal(t, []Roll{7}, game.LiveRolls)

	game.AddRoll(3)
	assert.Equal(t, []Score{9, 29, 59, 87, 106, 115, 121}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}, {9, 1}, {10}, {10}, {10}, {8,1}, {0, 6}}, game.ScoredRolls)
	assert.Equal(t, []Roll{7, 3}, game.LiveRolls)

	game.AddRoll(8)
	assert.Equal(t, []Score{9, 29, 59, 87, 106, 115, 121, 139}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}, {9, 1}, {10}, {10}, {10}, {8,1}, {0, 6}, {7, 3}}, game.ScoredRolls)
	assert.Equal(t, []Roll{8}, game.LiveRolls)

	game.AddRoll(2)
	assert.Equal(t, []Score{9, 29, 59, 87, 106, 115, 121, 139}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}, {9, 1}, {10}, {10}, {10}, {8,1}, {0, 6}, {7, 3}}, game.ScoredRolls)
	assert.Equal(t, []Roll{8, 2}, game.LiveRolls)

	game.AddRoll(10)
	assert.Equal(t, []Score{9, 29, 59, 87, 106, 115, 121, 139, 159}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}, {9, 1}, {10}, {10}, {10}, {8,1}, {0, 6}, {7, 3},{8, 2}}, game.ScoredRolls)
	assert.Equal(t, []Roll{10}, game.LiveRolls)

	game.AddRoll(9)
	assert.Equal(t, []Score{9, 29, 59, 87, 106, 115, 121, 139, 159}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}, {9, 1}, {10}, {10}, {10}, {8,1}, {0, 6}, {7, 3},{8, 2}}, game.ScoredRolls)
	assert.Equal(t, []Roll{10, 9}, game.LiveRolls)

	game.AddRoll(0)
	assert.Equal(t, []Score{9, 29, 59, 87, 106, 115, 121, 139, 159, 178, 187}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}, {9, 1}, {10}, {10}, {10}, {8,1}, {0, 6}, {7, 3},{8, 2}, {10}, {9, 0}}, game.ScoredRolls)
	assert.Equal(t, []Roll{}, game.LiveRolls)

	game.AddRoll(9) // ignored
	assert.Equal(t, []Score{9, 29, 59, 87, 106, 115, 121, 139, 159, 178, 187}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}, {9, 1}, {10}, {10}, {10}, {8,1}, {0, 6}, {7, 3},{8, 2}, {10}, {9, 0}}, game.ScoredRolls)
	assert.Equal(t, []Roll{}, game.LiveRolls)

	game.AddRoll(1) // ignored
	assert.Equal(t, []Score{9, 29, 59, 87, 106, 115, 121, 139, 159, 178, 187}, game.Frames)
	assert.Equal(t, [][]Roll{{0, 9}, {9, 1}, {10}, {10}, {10}, {8,1}, {0, 6}, {7, 3},{8, 2}, {10}, {9, 0}}, game.ScoredRolls)
	assert.Equal(t, []Roll{}, game.LiveRolls)
}

