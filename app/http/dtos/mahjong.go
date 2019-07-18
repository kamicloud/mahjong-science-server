package dtos

type BaseMessage struct {
	Status  int
	Message string
}

type AnalyseMessage struct {
	Request  AnalyseRequest
	Response AnalyseResponse
}

type AnalyseRequest struct {
	Tiles string `json:"tiles" binding:"required"`
}

type AnalyseResponse struct {
	Result TileAnalyseResult `json:"result"`
}

type RandomResponse struct {
	Result TileAnalyseResult `json:"result"`
}

type AnalyseArrayRequest struct {
	Tiles []int `json:"tiles" binding:"required"`
}

type AnalyseArrayResponse struct {
	Result TileAnalyseResult `json:"result"`
}

type TileAnalyseResult struct {
	CurrentTileString       string          `json:"currentTileString"`
	CurrentTileSimpleString string          `json:"currentTileSimpleString"`
	Shanten                 int             `json:"shanten"`
	CurrentTiles            []int           `json:"currentTiles"`
	CurrentRenderTiles		[]int			`json:"currentRenderTiles"`
	Choices                 []DiscardChoice `json:"choices"`
	IncShantenChoices       []DiscardChoice `json:"incShantenChoices"`
}

type DiscardChoice struct {
	Discard   int   `json:"discard"`
	Draws     []int `json:"draws"`
	DrawCount int   `json:"drawCount"`
}
