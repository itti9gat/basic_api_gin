package handle_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
	"iiujapp.tech/basic-api-gin/handle"
	"iiujapp.tech/basic-api-gin/model"
	mock_repo "iiujapp.tech/basic-api-gin/repo/mock"
)

func TestProduct(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rc := mock_repo.NewMockCategoryINF(ctrl)
	rp := mock_repo.NewMockProductINF(ctrl)

	rp.EXPECT().FindProduct("1").Return(model.Product{
		ID:         1,
		CategoryID: 6,
		Code:       "AW11",
		Topic:      "black tea",
		Detail:     "",
		BasePrice:  990,
		Price:      500,
		StockType:  "N",
		Stock:      0,
		Status:     "A",
		Position:   1,
		CreatedAt:  "2019-11-18 11:45:29",
		UpdatedAt:  "2019-11-23 10:04:34",
	}, nil).AnyTimes()

	h := handle.NewHandler(rc, rp)

	t.Run("Query Product Success", func(t *testing.T) {

		expectedStatusCode := http.StatusOK
		expected := model.Product{
			ID:         1,
			CategoryID: 6,
			Code:       "AW11",
			Topic:      "black tea",
			Detail:     "",
			BasePrice:  990,
			Price:      500,
			StockType:  "N",
			Stock:      0,
			Status:     "A",
			Position:   1,
			CreatedAt:  "2019-11-18 11:45:29",
			UpdatedAt:  "2019-11-23 10:04:34",
		}

		req := httptest.NewRequest("GET", "/v1/product/1", nil)
		res := httptest.NewRecorder()

		testRoute := gin.Default()
		testRoute.GET("/v1/product/:proid", h.ProductByID)
		testRoute.ServeHTTP(res, req)

		response := res.Result()
		body, _ := ioutil.ReadAll(response.Body)

		actualStatusCode := response.StatusCode
		actualResponse := model.Product{}
		_ = json.Unmarshal(body, &actualResponse)

		assert.Equal(t, expected, actualResponse)
		assert.Equal(t, expectedStatusCode, actualStatusCode)
	})
}
