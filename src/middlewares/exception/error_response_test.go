package exception_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/syukri21/Paperid-Golang-Testcase/src/middlewares/exception"
)

func TestInternalServerError(t *testing.T) {
	var errors []map[string]interface{}
	defer func() {
		if err := recover(); err != nil {
			assert.NotEmpty(t, err)
		}
	}()
	exception.InternalServerError("ISE Testing", errors)
}
