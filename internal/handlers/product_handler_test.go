package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mock_serv "github.com/sustatov027-max/project_calorie_tracker/internal/handlers/mock"
	"github.com/sustatov027-max/project_calorie_tracker/internal/models"
)

func assertJSONEqualIgnoreFields(t *testing.T, expected, actual string, ignoreFields ...string) {
	var expectedMap, actualMap map[string]interface{}

	require.NoError(t, json.Unmarshal([]byte(expected), &expectedMap))
	require.NoError(t, json.Unmarshal([]byte(actual), &actualMap))

	for _, field := range ignoreFields {
		delete(actualMap, field)
	}

	assert.Equal(t, expectedMap, actualMap)
}

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type mockSetup func(s *mock_serv.MockProductServ)

	tests := []struct {
		name       string
		request    string
		mockSetup  mockSetup
		wantStatus int
		wantBody   string
	}{
		{
			name:    "Success",
			request: `{"name":"Test","calories":1, "proteins":1, "fats":1, "carbohydrates":1}`,
			mockSetup: func(s *mock_serv.MockProductServ) {
				s.EXPECT().CreateProduct("Test", 1.0, 1.0, 1.0, 1.0).Return(
					models.Product{
						ID:            1,
						Name:          "Test",
						Calories:      1,
						Proteins:      1,
						Fats:          1,
						Carbohydrates: 1,
					}, nil,
				)
			},
			wantStatus: 201,
			wantBody:   `{"ID":1,"Name":"Test","Calories":1,"Proteins":1,"Fats":1,"Carbohydrates":1}`,
		},
		{
			name:    "Empty name",
			request: `{"name":"","calories":1, "proteins":1, "fats":1, "carbohydrates":1}`,
			mockSetup: func(s *mock_serv.MockProductServ) {
			},
			wantStatus: 400,
			wantBody:   `{"Name": "must be at least 2 characters"}`,
		},
		{
			name:    "Incorrect calories",
			request: `{"name":"Test","calories":-1, "proteins":1, "fats":1, "carbohydrates":1}`,
			mockSetup: func(s *mock_serv.MockProductServ) {
			},
			wantStatus: 400,
			wantBody:   `{"Calories": "must be greater or equal to 0"}`,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			service := mock_serv.NewMockProductServ(ctrl)
			if testCase.mockSetup != nil {
				testCase.mockSetup(service)
			}

			handler := NewProductHandler(service)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/products",
				strings.NewReader(testCase.request))

			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.POST("/products", handler.CreateProduct)
			router.ServeHTTP(w, req)

			assert.Equal(t, testCase.wantStatus, w.Code)

			if testCase.wantStatus == 201 {
				assertJSONEqualIgnoreFields(t, testCase.wantBody, w.Body.String(), "CreatedAt")
			} else {
				assert.JSONEq(t, testCase.wantBody, w.Body.String())
			}
		})
	}
}
