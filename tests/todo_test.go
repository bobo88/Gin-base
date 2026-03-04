package tests

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"todo-list/routes"
)

func TestGetTodos(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/todos", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %v", w.Code)
	}
}