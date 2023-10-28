package game

// The Configs_reels and Config_reels structures contain configurations in the following format:
// [
//
//	['A', 'B', 'C', 'D', 'E'],
//	['F', 'A', 'F', 'B', 'C'],
//	['D', 'E', 'A', 'G', 'A']
//
// ]
type Configs_reels struct {
	Configs []*Config_reels
	Map     map[string]*Config_reels
}

// See comment Configs_reels
type Config_reels struct {
	Name  string
	Reels [3][5]rune
}

// The Configs_lines, Config_lines, Line and Position structures contain configurations in the following format:
// [
//
//	{
//		"line": 1,
//		"positions": [
//		{"row": 0, "col": 0},
//		{"row": 1, "col": 1},
//		{"row": 2, "col": 2},
//		{"row": 1, "col": 3},
//		{"row": 0, "col": 4}
//		]
//	},
//	{
//		"line": 2,
//		"positions": [
//			{"row": 2, "col": 0},
//			{"row": 1, "col": 1},
//			{"row": 0, "col": 2},
//			{"row": 1, "col": 3},
//			{"row": 2, "col": 4}
//		]
//	},
//	{
//		"line": 3,
//		"positions": [
//			{"row": 1, "col": 0},
//			{"row": 2, "col": 1},
//			{"row": 1, "col": 2},
//			{"row": 0, "col": 3},
//			{"row": 1, "col": 4}
//		]
//	}
//
// ]
type Configs_lines struct {
	Configs []*Config_lines
	Map     map[string]*Config_lines
}

// See comment Configs_lines
type Config_lines struct {
	Name  string
	Lines []*Line
}

// See comment Configs_lines
type Line struct {
	Number    int         `json:"line"`
	Positions []*Position `json:"positions"`
}

// See comment Configs_lines
type Position struct {
	Row rune `json:"row"`
	Col rune `json:"col"`
}

// The Configs_payouts, Config_payouts and Payouts structures contain configurations in the following format:
// [
//
//	{
//		"symbol": 'A',
//		"payout": [0, 0, 50, 100, 200]
//	},
//	{
//		"symbol": 'B',
//		"payout": [0, 0, 40, 80, 160]
//	},
//	{
//		"symbol": 'C',
//		"payout": [0, 0, 30, 60, 120]
//	},
//	{
//		"symbol": 'D',
//		"payout": [0, 0, 20, 40, 80]
//	},
//	{
//		"symbol": 'E',
//		"payout": [0, 0, 10, 20, 40]
//	},
//	{
//		"symbol": 'F',
//		"payout": [0, 0, 5, 10, 20]
//	},
//	{
//		"symbol": 'G',
//		"payout": [0, 0, 2, 5, 10]
//	}
//
// ]
type Configs_payouts struct {
	Configs []*Config_payouts
	Map     map[string]*Config_payouts
}

// See comment Configs_payouts
type Config_payouts struct {
	Name     string
	Payouts  []Payouts
	mPayouts map[rune]*Payouts
}

// See comment Configs_payouts
type Payouts struct {
	Symbol rune  `json:"symbol"`
	Payout []int `json:"payout"`
}

// The Result and LineResult structures define the format of the service's resulting response.
type Result struct {
	symbols [][]rune
	Lines   []LineResult `json:"lines"`
	Total   int          `json:"total"`
}

// See comment Result
type LineResult struct {
	Line   int `json:"line"`
	Payout int `json:"payout"`
}

// ReqResult contains the names of configurations for which the results need to be calculated.
type ReqResult struct {
	Config_reels_name   string `json:"config_reels_name"`
	Config_lines_name   string `json:"config_lines_name"`
	Config_payouts_name string `json:"config_payouts_name"`
}
