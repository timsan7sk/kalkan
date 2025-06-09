package kalkan

import (
	"errors"
	"fmt"
)

// Kalkan error struct.
type kalkanError struct {
	errorCode   ErrorCode
	errorString string
}

func (e kalkanError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.errorCode.Hex(), e.errorCode.String(), e.errorString)
}

// GetErrorCode извлекает из ошибки ErrorCode из ошибка типа KalkanError.
func GetErrorCode(err error) (ErrorCode, bool) {
	var kalkanErr kalkanError
	if errors.As(err, &kalkanErr) {
		return kalkanErr.errorCode, true
	}

	var kalkanErrPoint *kalkanError
	if errors.As(err, &kalkanErrPoint) {
		return kalkanErrPoint.errorCode, true
	}

	return 0, false
}

// Returns the last global error if returnCode is not 0.
func (m *Module) wrapError(returnCode int) error {
	if returnCode != 0 {
		ec, es := m.GetLastErrorString()
		return kalkanError{
			errorCode:   ec,
			errorString: es,
		}
	}
	return nil
}
