package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-sample/model"
	"github.com/ybkuroki/go-webapp-sample/test"
)

func TestGetCategoryList(t *testing.T) {
	router, context := test.Prepare()

	master := NewMasterController(context)
	router.GET(APIMasterCategory, func(c echo.Context) error { return master.GetCategoryList(c) })

	req := httptest.NewRequest("GET", APIMasterCategory, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	data := [...]*model.Category{
		{ID: 1, Name: "技術書"},
		{ID: 2, Name: "雑誌"},
		{ID: 3, Name: "小説"},
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}

func TestGetFormatList(t *testing.T) {
	router, context := test.Prepare()

	master := NewMasterController(context)
	router.GET(APIMasterFormat, func(c echo.Context) error { return master.GetFormatList(c) })

	req := httptest.NewRequest("GET", APIMasterFormat, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	data := [...]*model.Format{
		{ID: 1, Name: "書籍"},
		{ID: 2, Name: "電子書籍"},
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}
