package util

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

type TestVector struct {
	Internal   []TestVector
	Name       string
	Body       gin.H
	StatusCode int
}

func RunTest(t *testing.T, router *gin.Engine, method, path string, testCases []TestVector) {
	for i := range testCases {
		test := testCases[i]
		t.Run(test.Name, func(t *testing.T) {
			if test.Internal != nil {
				RunTest(t, router, method, path, test.Internal)
				return
			}
			reqBody, _ := json.Marshal(test.Body)
			req := httptest.NewRequest(method, path, bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			require.Equal(t, test.StatusCode, resp.Code)
		})
	}
}
