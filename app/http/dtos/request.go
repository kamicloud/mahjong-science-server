package dtos

type CodeToSessionRequest struct {
	Code string `json:"code"`
}

type AnalyseRequest struct {
	Tiles string `json:"tiles" binding:"required"`
}

type AnalyseArrayRequest struct {
	Tiles []int `json:"tiles" binding:"required"`
}

type RandomRequest struct {

}
