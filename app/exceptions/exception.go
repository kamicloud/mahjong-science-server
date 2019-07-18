package exceptions

type InvalidParameterException struct {
	Exception
	Status int
}

type Exception struct {
	Status  int
	Message string
}

const (
	InvalidParameter    = -1
	ServerInternalError = -2
	Success				= 0
	CustomError         = 100
)
