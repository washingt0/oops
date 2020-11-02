package oops

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error type provide wrap for an error
type Error struct {
	Message    string   `json:"message,omitempty"`
	Location   string   `json:"location,omitempty"`
	RawError   error    `json:"raw_error,omitempty"`
	Stack      []string `json:"stack,omitempty"`
	Code       int      `json:"code,omitempty"`
	StatusCode int      `json:"status_code,omitempty"`
}

// Error implements error interface
func (e *Error) Error() string {
	return e.Message
}

// ThrowError may wrap a error with a friendly message
func ThrowError(message string, e error, status ...int) (err error) {
	var (
		tmp *Error = new(Error)
	)

	tmp.Code = 9999
	tmp.StatusCode = http.StatusBadRequest

	if e == nil && message == "" {
		return nil
	}

	if e == nil {
		tmp.RawError = errors.New(message)
	} else {
		tmp.RawError = e
	}

	if message == "" {
		tmp.Message = e.Error()
	} else {
		tmp.Message = message
	}

	tmp.Stack = getStack()
	tmp.Location = tmp.Stack[len(tmp.Stack)-1]

	return tmp
}

// GinHandleError add errors to the gin context
func GinHandleError(c *gin.Context, err error, statusCode int) {
	var (
		e *Error
	)

	if val, ok := err.(*Error); ok && val != nil {
		e = val
	} else {
		e = ThrowError("", err, statusCode).(*Error)
	}

	c.Error(e)
	c.AbortWithStatusJSON(statusCode, e)
}

func getStack() (stack []string) {
	// TODO: get stack
	return []string{""}
}
