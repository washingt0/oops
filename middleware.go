package oops

import "github.com/gin-gonic/gin"

// GetGinError returns all handled errors in the context
func GetGinError(c *gin.Context) (err []error) {
	err = make([]error, len(c.Errors))

	for i := range c.Errors {
		if val, ok := c.Errors[i].Err.(*Error); ok && val != nil {
			err[i] = val
		}
	}

	return
}
