package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func MockRequestFormatter(method string, path string, body any) (w *httptest.ResponseRecorder, c *gin.Context, err error) {
	var jsonBody []byte
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to marshal body: %w", err)
		}
	}

	writer := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(writer)

	if jsonBody == nil {
		context.Request = httptest.NewRequest(method, path, nil)
	} else {
		context.Request = httptest.NewRequest(method, path, bytes.NewBuffer(jsonBody))
		context.Request.Header.Set("Content-Type", "application/json")
	}
	return writer, context, nil
}
