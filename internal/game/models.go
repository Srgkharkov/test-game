package game

type Configs_reels struct {
	Configs []*Config_reels
	Map     map[string]*Config_reels
}

type Config_reels struct {
	Name  string
	Reels [3][5]rune
}

type Configs_lines struct {
	Configs []*Config_lines
	Map     map[string]*Config_lines
}

type Config_lines struct {
	Name  string
	Lines []*Line
}

type Line struct {
	Number    int         `json:"line"`
	Positions []*Position `json:"positions"`
}

type Position struct {
	Row rune `json:"row"`
	Col rune `json:"col"`
}

type Configs_payouts struct {
	Configs []*Config_payouts
	Map     map[string]*Config_payouts
}

type Config_payouts struct {
	Name     string
	Payouts  []Payouts
	mPayouts map[rune]*Payouts
}

type Payouts struct {
	Symbol rune  `json:"symbol"`
	Payout []int `json:"payout"`
}

type Result struct {
	symbols [][]rune
	Lines   []LineResult `json:"lines"`
	Total   int          `json:"total"`
}

type LineResult struct {
	Line   int `json:"line"`
	Payout int `json:"payout"`
}

type ReqResult struct {
	Config_reels_name   string `json:"conf_reels_name"`
	Config_lines_name   string `json:"conf_lines_name"`
	Config_payouts_name string `json:"conf_payouts_name"`
}
