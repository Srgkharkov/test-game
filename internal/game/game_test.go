package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame(t *testing.T) {
	// Arrange
	game := NewGame()
	confreels := Config_reels{
		Name: "config_reels_1",
		Reels: [3][5]rune{
			{'A', 'B', 'C', 'D', 'E'},
			{'F', 'A', 'F', 'B', 'C'},
			{'D', 'E', 'A', 'G', 'A'},
		},
	}
	conflines := Config_lines{
		Name: "config_lines_1",
		Lines: []*Line{
			&Line{
				Number: 1,
				Positions: []*Position{
					&Position{Row: 0, Col: 0},
					&Position{Row: 1, Col: 1},
					&Position{Row: 2, Col: 2},
					&Position{Row: 1, Col: 3},
					&Position{Row: 0, Col: 4},
				},
			},
			&Line{
				Number: 2,
				Positions: []*Position{
					&Position{Row: 2, Col: 0},
					&Position{Row: 1, Col: 1},
					&Position{Row: 0, Col: 2},
					&Position{Row: 1, Col: 3},
					&Position{Row: 2, Col: 4},
				},
			},
			&Line{
				Number: 3,
				Positions: []*Position{
					&Position{Row: 1, Col: 0},
					&Position{Row: 2, Col: 1},
					&Position{Row: 1, Col: 2},
					&Position{Row: 0, Col: 3},
					&Position{Row: 1, Col: 4},
				},
			},
		},
	}
	confpayouts := Config_payouts{
		Name: "config_payouts_1",
		Payouts: []Payouts{
			Payouts{
				Symbol: 'A',
				Payout: []int{0, 0, 50, 100, 200},
			},
			Payouts{
				Symbol: 'B',
				Payout: []int{0, 0, 40, 80, 160},
			},
			Payouts{
				Symbol: 'C',
				Payout: []int{0, 0, 30, 60, 120},
			},
			Payouts{
				Symbol: 'D',
				Payout: []int{0, 0, 20, 40, 80},
			},
			Payouts{
				Symbol: 'E',
				Payout: []int{0, 0, 10, 20, 40},
			},
			Payouts{
				Symbol: 'F',
				Payout: []int{0, 0, 5, 10, 20},
			},
			Payouts{
				Symbol: 'G',
				Payout: []int{0, 0, 2, 5, 10},
			},
		},
	}
	ReqResult := ReqResult{
		Config_reels_name:   "config_reels_1",
		Config_lines_name:   "config_lines_1",
		Config_payouts_name: "config_payouts_1",
	}
	ExpectedResult := Result{
		Lines: []LineResult{
			{
				Line:   1,
				Payout: 50,
			},
			{
				Line:   2,
				Payout: 0,
			},
			{
				Line:   3,
				Payout: 0,
			},
		},
		Total: 50,
	}

	err := game.Configs_reels.AddConfig(&confreels)
	assert.Nil(t, err, err)
	err = game.Configs_reels.AddConfig(&confreels)
	assert.NotNil(t, err, err)

	err = game.Configs_lines.AddConfig(&conflines)
	assert.Nil(t, err, err)
	err = game.Configs_lines.AddConfig(&conflines)
	assert.NotNil(t, err, err)

	err = game.Configs_payouts.AddConfig(&confpayouts)
	assert.Nil(t, err, err)
	err = game.Configs_payouts.AddConfig(&confpayouts)
	assert.NotNil(t, err, err)

	result, err := game.GetResult(&ReqResult)
	assert.Nil(t, err, err)
	assert.Equal(t, ExpectedResult.Lines, result.Lines)
	assert.Equal(t, ExpectedResult.Total, result.Total)

}
