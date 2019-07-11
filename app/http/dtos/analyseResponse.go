package dtos

type AnalyseResponse struct {
	Shanten           int      `json:"shanten"`
	CurrentTiles      []int    `json:"currentTiles"`
	Choices           []Choice `json:"choices"`
	IncShantenChoices []Choice `json:"incShantenChoices"`
}

type Choice struct {
	Discard   int   `json:"discard"`
	Draws     []int `json:"draws"`
	DrawCount int   `json:"drawCount"`
}
