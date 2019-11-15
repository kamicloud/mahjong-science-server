package dtos

type BaseMessage struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AnalyseMessage struct {
	Request  *AnalyseRequest
	Response AnalyseResponse
}

type AnalyseArrayMessage struct {
	Request  *AnalyseArrayRequest
	Response AnalyseArrayResponse
}

type RandomMessage struct {
	Request  RandomRequest
	Response RandomResponse
}
