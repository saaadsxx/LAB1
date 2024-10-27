package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	router := setupRouter()

	user := `{"name": "John", "email": "john@example.com", "age": 30}`
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte(user)))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
