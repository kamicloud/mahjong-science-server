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
type RankLevel struct {
	Id int `json:"id"`
	Score int `json:"score"`
}

type Rank struct {
	AccountId int `json:"account_id"`
	AvatarId  int `json:"avatar_id"`
	Nickname  string `json:"nickname"`
	Level RankLevel `json:"level"`
	Level3 RankLevel `json:"level3"`
	Title int `json:"title"`
}