package port

type LogRepository interface {
	LogRequestResponse(request string, response int) error
}
