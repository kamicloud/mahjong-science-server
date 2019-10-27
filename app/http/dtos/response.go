package dtos


type AnalyseResponse struct {
	Result TileAnalyseResult `json:"result"`
}

type RandomResponse struct {
	Result TileAnalyseResult `json:"result"`
}

type AnalyseArrayResponse struct {
	Result TileAnalyseResult `json:"result"`
}

type GroupResponse struct {
	Groups []Group `json:"groups"`
}

