package custom_errors

import (
	"errors"
	"fmt"
	"inno-accounting/pkg/server_utils/response_message"
	"net/http"

	"github.com/sirupsen/logrus"
)

func ErrorResponse(w http.ResponseWriter, err error, logger *logrus.Logger) error{
	var custom_error *CustomError
	if errors.As(err, &custom_error) {
		err = err.(*CustomError)
		response_message.WrapperResponseJSON(w, custom_error.StatusCode, custom_error.ResponseData)
		logger.Info(custom_error.LogData)
		return nil
	} else {
		return fmt.Errorf("unknown error: %w", err)
	}
	
}