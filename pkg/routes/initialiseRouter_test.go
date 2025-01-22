package routes

import (
	"github.com/nathanberry97/rss2go/internal/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	// Arrange
	router := InitialiseRouter()
	req, _ := http.NewRequest("GET", "/health-check", nil)

	// Act
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert
	utils.Assert(t, http.StatusOK, w.Code)
	utils.Assert(t, `{"message":"ok"}`, w.Body.String())
}
