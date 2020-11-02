package oops

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message    string
	Location   string
	RawError   error
	Stack      []string
	Code       int
	StatusCode int
}

func (e *Error) Error() string {
	return e.Message
}

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
