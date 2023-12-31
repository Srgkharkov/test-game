package game

import (
	"errors"
)

type Game struct {
	Configs_reels   *Configs_reels
	Configs_lines   *Configs_lines
	Configs_payouts *Configs_payouts
}

// The NewGame function initializes the Game structure, which stores the
// configurations for Configs_reels, Configs_lines, and Configs_payouts.
func NewGame() *Game {
	return &Game{
		Configs_reels: &Configs_reels{
			Configs: make([]*Config_reels, 0),
			Map:     make(map[string]*Config_reels),
		},
		Configs_lines: &Configs_lines{
			Configs: make([]*Config_lines, 0),
			Map:     make(map[string]*Config_lines),
		},
		Configs_payouts: &Configs_payouts{
			Configs: make([]*Config_payouts, 0),
			Map:     make(map[string]*Config_payouts),
		},
	}
}

// The AddConfig method takes a Config_reels, checks for the existence of the configuration.
// If it doesn't exist, it is saved; otherwise, it returns an error.
func (c *Configs_reels) AddConfig(Config_reels *Config_reels) error {
	if _, ok := c.Map[Config_reels.Name]; ok {
		return errors.New("This config already exists")
	}
	c.Configs = append(c.Configs, Config_reels)
	c.Map[Config_reels.Name] = Config_reels
	return nil
}

// The AddConfig method takes a Config_lines, checks for the existence of the configuration.
// If it doesn't exist, it is saved; otherwise, it returns an error.
func (c *Configs_lines) AddConfig(Config_lines *Config_lines) error {
	if _, ok := c.Map[Config_lines.Name]; ok {
		return errors.New("This config already exists")
	}
	c.Configs = append(c.Configs, Config_lines)
	c.Map[Config_lines.Name] = Config_lines
	return nil
}

// The AddConfig method takes a Config_payouts, checks for the existence of the configuration.
// If it doesn't exist, it is saved; otherwise, it returns an error.
func (c *Configs_payouts) AddConfig(Config_payouts *Config_payouts) error {
	if _, ok := c.Map[Config_payouts.Name]; ok {
		return errors.New("This config already exists")
	}

	Config_payouts.mPayouts = make(map[rune]*Payouts)
	for i := 0; i < len(Config_payouts.Payouts); i++ {
		Config_payouts.mPayouts[Config_payouts.Payouts[i].Symbol] = &Config_payouts.Payouts[i]
	}

	c.Configs = append(c.Configs, Config_payouts)
	c.Map[Config_payouts.Name] = Config_payouts

	return nil
}

// The GetResult method takes a ReqResult structure, which contains the names of configurations:
// Config_reels_name, Config_lines_name, and Config_payouts_name.
// It locates the required configurations, calculates the winnings, and returns the response in a Result structure.
func (g *Game) GetResult(ReqResult *ReqResult) (*Result, error) {
	Config_reels, ok := g.Configs_reels.Map[ReqResult.Config_reels_name]
	if !ok {
		return nil, errors.New("Config reels not exists")
	}

	Config_lines, ok := g.Configs_lines.Map[ReqResult.Config_lines_name]
	if !ok {
		return nil, errors.New("Config lines not exists")
	}

	Config_payouts, ok := g.Configs_payouts.Map[ReqResult.Config_payouts_name]
	if !ok {
		return nil, errors.New("Config payouts not exists")
	}
	var Result Result

	Result.symbols = make([][]rune, len(Config_lines.Lines))
	Result.Lines = make([]LineResult, len(Config_lines.Lines))

	for i := 0; i < len(Config_lines.Lines); i++ {

		Line := *Config_lines.Lines[i]
		Result.symbols[i] = make([]rune, len(Line.Positions))
		Result.Lines[i].Line = Line.Number

		var curSymbol rune
		var countSameSymbols int

		for j := 0; j < len(Line.Positions); j++ {
			Result.symbols[i][j] = Config_reels.Reels[Line.Positions[j].Row][Line.Positions[j].Col]
			if j == 0 {
				curSymbol = Result.symbols[i][j]
				countSameSymbols = 0
			} else if j < len(Line.Positions)-1 {
				if curSymbol == Result.symbols[i][j] {
					countSameSymbols++
				} else {
					Result.Lines[i].Payout += Config_payouts.mPayouts[curSymbol].Payout[countSameSymbols]
					curSymbol = Result.symbols[i][j]
					countSameSymbols = 0
				}
			} else {
				if curSymbol == Result.symbols[i][j] {
					countSameSymbols++
					Result.Lines[i].Payout += Config_payouts.mPayouts[curSymbol].Payout[countSameSymbols]
				} else {
					Result.Lines[i].Payout += Config_payouts.mPayouts[curSymbol].Payout[countSameSymbols]
					curSymbol = Result.symbols[i][j]
					countSameSymbols = 0
					Result.Lines[i].Payout += Config_payouts.mPayouts[curSymbol].Payout[countSameSymbols]
				}
				Result.Total += Result.Lines[i].Payout
			}
		}

	}

	return &Result, nil
}
