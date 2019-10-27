package dtos


type AnalyseRequest struct {
	Tiles string `json:"tiles" binding:"required"`
}

type AnalyseArrayRequest struct {
	Tiles []int `json:"tiles" binding:"required"`
}

type RandomRequest struct {

}
