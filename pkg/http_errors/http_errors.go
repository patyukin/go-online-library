package httpErrors

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	ErrBadRequest          = "Bad request"
	ErrNotFound            = "Not Found"
	ErrUnauthorized        = "Unauthorized"
	ErrRequestTimeout      = "Request Timeout"
	ErrInternalServerError = "Internal Server Error"
)

var (
	BadRequest          = errors.New("bad request")
	WrongCredentials    = errors.New("wrong Credentials")
	NotFound            = errors.New("not Found")
	Unauthorized        = errors.New("unauthorized")
	InternalServerError = errors.New("internal Server Error")
)

// RestErr Rest error interface
type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
	ErrBody() RestError
}

// RestError Rest error struct
type RestError struct {
	ErrStatus  int         `json:"status,omitempty"`
	ErrError   string      `json:"error,omitempty"`
	ErrMessage interface{} `json:"message,omitempty"`
	Timestamp  time.Time   `json:"timestamp,omitempty"`
}

// ErrBody Error body
func (e RestError) ErrBody() RestError {
	return e
}

// Error  Error() interface method
func (e RestError) Error() string {
	return fmt.Sprintf("status: %docker, errors: %s, causes: %v", e.ErrStatus, e.ErrError, e.ErrMessage)
}

// Status Error status
func (e RestError) Status() int {
	return e.ErrStatus
}

// Causes RestError Causes
func (e RestError) Causes() interface{} {
	return e.ErrMessage
}

// NewRestError New Rest Error
func NewRestError(status int, err string, causes interface{}, debug bool) RestErr {
	restError := RestError{
		ErrStatus: status,
		ErrError:  err,
		Timestamp: time.Now().UTC(),
	}

	if debug {
		restError.ErrMessage = causes
	}

	return restError
}

// NewRestErrorWithMessage New Rest Error With Message
func NewRestErrorWithMessage(status int, err string, causes interface{}) RestErr {
	return RestError{
		ErrStatus:  status,
		ErrError:   err,
		ErrMessage: causes,
		Timestamp:  time.Now().UTC(),
	}
}

// NewRestErrorFromBytes New Rest Error From Bytes
func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr RestError
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

// ParseErrors Parser of error string messages returns RestError
func ParseErrors(err error, debug bool) RestErr {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewRestError(http.StatusNotFound, ErrNotFound, err.Error(), debug)
	case errors.Is(err, context.DeadlineExceeded):
		return NewRestError(http.StatusRequestTimeout, ErrRequestTimeout, err.Error(), debug)
	case errors.Is(err, Unauthorized):
		return NewRestError(http.StatusUnauthorized, ErrUnauthorized, err.Error(), debug)
	case errors.Is(err, WrongCredentials):
		return NewRestError(http.StatusUnauthorized, ErrUnauthorized, err.Error(), debug)
	case strings.Contains(strings.ToLower(err.Error()), "sqlstate"):
		return parseSqlErrors(err, debug)
	case strings.Contains(strings.ToLower(err.Error()), "required header"):
		return NewRestError(http.StatusBadRequest, ErrBadRequest, err.Error(), debug)
	case strings.Contains(strings.ToLower(err.Error()), "no documents in result"):
		return NewRestError(http.StatusNotFound, ErrNotFound, err.Error(), debug)
	default:
		var restErr *RestError
		if errors.As(err, &restErr) {
			return restErr
		}

		return NewRestError(http.StatusInternalServerError, ErrInternalServerError, err.Error(), debug)
	}
}

func parseSqlErrors(err error, debug bool) RestErr {
	return NewRestError(http.StatusBadRequest, ErrBadRequest, err, debug)
}

// ErrorResponse Error response
func ErrorResponse(err error, debug bool) (int, RestErr) {
	return ParseErrors(err, debug).Status(), ParseErrors(err, debug)
}
