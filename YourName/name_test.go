package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPostName(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/name", func(c *gin.Context) {
		var p person
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"first_name": p.FirstName,
			"last_name":  p.LastName,
		})
	})

	personData := person{
		FirstName: "John",
		LastName:  "Doe",
	}

	jsonData, err := json.Marshal(personData)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/name", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedBody := `{"first_name":"John","last_name":"Doe"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
