package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHelloEndpoint(t *testing.T) {
	// Tworzymy router
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Definiujemy endpoint
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	// Tworzymy żądanie do testu
	req, err := http.NewRequest(http.MethodGet, "/hello", nil)
	assert.NoError(t, err)

	// Rejestrujemy odpowiedź
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Sprawdzamy status odpowiedzi
	assert.Equal(t, http.StatusOK, w.Code)

	// Sprawdzamy treść odpowiedzi
	expectedBody := `{"message":"Hello World!"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
