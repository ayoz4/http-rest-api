package apiserver

import (
	"net/http"
	"net/http/httptest"
	"test/internal/app/store/teststore"
	"testing"

	"github.com/bmizerany/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s := newServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}