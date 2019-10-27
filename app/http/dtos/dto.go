package dtos

type Group struct {
	Title   string `json:"title"`
	Num     string `json:"num"`
	Content string `json:"content"`
}

type TileAnalyseResult struct {
	CurrentTileString       string          `json:"currentTileString"`
	CurrentTileSimpleString string          `json:"currentTileSimpleString"`
	Shanten                 int             `json:"shanten"`
	CurrentTiles            []int           `json:"currentTiles"`
	CurrentRenderTiles      []int           `json:"currentRenderTiles"`
	Choices                 []DiscardChoice `json:"choices"`
	IncShantenChoices       []DiscardChoice `json:"incShantenChoices"`
}

type DiscardChoice struct {
	Discard   int   `json:"discard"`
	Draws     []int `json:"draws"`
	DrawCount int   `json:"drawCount"`
}
