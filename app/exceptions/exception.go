package exceptions

type InvalidParameterException struct {
	Exception
	Status int
}

type Exception struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

const (
	InvalidParameter    = -1
	ServerInternalError = -2
	Success             = 0
	CustomError         = 100
)
