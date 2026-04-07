package custom_errors

type CustomError struct {
	StatusCode   int
	Err          error
	ResponseData string
	LogData      string
}

func New(err error, sc int) *CustomError {
	return &CustomError{
		StatusCode: sc,
		Err:        err,
	}
}

func (e *CustomError) AddResponseData(response string) {
	e.ResponseData = response
}

func (e *CustomError) AddLogData(log string) {
	e.LogData = log
}

func (e *CustomError) Error() string {
	return e.LogData
}

func (e *CustomError) Unwrap() error {
	return e.Err
}
