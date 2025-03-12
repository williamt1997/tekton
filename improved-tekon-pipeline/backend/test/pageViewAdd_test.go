package test

import (
	"homers-backend/controllers"
	"homers-backend/test/helpers"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPageView(t *testing.T) {
	var err error
	helpers.BuildTestDB()

	w, c, err := helpers.MockRequestFormatter("PUT", "/add", nil)
	assert.NoError(t, err)
	addViewController := controllers.AddController{Database: helpers.TestDB}
	addViewController.AddPageView(c)
	assert.Equal(t, http.StatusOK, w.Code)
	helpers.CloseTestDB()
}
