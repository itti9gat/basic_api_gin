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
	"iiujapp.tech/susaday/handle"
	"iiujapp.tech/susaday/model"
	mock_repo "iiujapp.tech/susaday/repo/mock"
)

func TestCategory(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rc := mock_repo.NewMockCategoryINF(ctrl)
	rp := mock_repo.NewMockProductINF(ctrl)

	rc.EXPECT().FindCategory("1").Return(model.Category{
		ID:        1,
		Topic:     "black tea",
		Level:     1,
		LevelID:   0,
		Status:    "A",
		Position:  1,
		CreatedAt: "2019-11-18 11:45:29",
		UpdatedAt: "2019-11-23 10:04:34",
	}, nil).AnyTimes()

	h := handle.NewHandler(rc, rp)

	t.Run("Query Category Success", func(t *testing.T) {

		expectedStatusCode := http.StatusOK
		expected := model.Category{
			ID:        1,
			Topic:     "black tea",
			Level:     1,
			LevelID:   0,
			Status:    "A",
			Position:  1,
			CreatedAt: "2019-11-18 11:45:29",
			UpdatedAt: "2019-11-23 10:04:34",
		}

		req := httptest.NewRequest("GET", "/v1/category/1", nil)
		res := httptest.NewRecorder()

		testRoute := gin.Default()
		testRoute.GET("/v1/category/:cateid", h.CategoryByID)
		testRoute.ServeHTTP(res, req)

		response := res.Result()
		body, _ := ioutil.ReadAll(response.Body)

		actualStatusCode := response.StatusCode
		actualResponse := model.Category{}
		_ = json.Unmarshal(body, &actualResponse)

		assert.Equal(t, expected, actualResponse)
		assert.Equal(t, expectedStatusCode, actualStatusCode)
	})
}
