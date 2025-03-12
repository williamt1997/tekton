package test

import (
	"encoding/json"
	"homers-backend/controllers"
	"homers-backend/test/helpers"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPageViewGet(t *testing.T) {
	var err error
	helpers.BuildTestDB()

	helpers.CreateTestPageView()

	w, c, err := helpers.MockRequestFormatter("GET", "/get", nil)
	assert.NoError(t, err)
	getViewController := controllers.GetController{Database: helpers.TestDB}
	getViewController.GetPageView(c)
	assert.Equal(t, http.StatusOK, w.Code)
	var pageCount int
	err = json.Unmarshal(w.Body.Bytes(), &pageCount)
	assert.NoError(t, err)
	assert.Equal(t, pageCount, 50)
	helpers.CloseTestDB()
}
