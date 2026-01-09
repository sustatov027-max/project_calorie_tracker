package handlers

import (
	"bytes"
	"net/http/httptest"
	mock_serv "project_calorie_tracker/internal/handlers/mock"
	"project_calorie_tracker/internal/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	type inputUser struct {
		Name       string
		Age        int
		Email      string
		Password   string
		Weight     float64
		Height     float64
		Gender     string
		ActiveDays int
	}

	type mockBehavior func(s *mock_serv.MockUserServ, user inputUser)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           inputUser
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name": "Ivan","age": 19,"email": "vanya.sustatov@mail.ru","password": "25566255","weight": 76,"height": 184,"gender": "male","activeDays": 3}`,
			inputUser: inputUser{
				Name:       "Ivan",
				Age:        19,
				Email:      "vanya.sustatov@mail.ru",
				Password:   "25566255",
				Weight:     76,
				Height:     184,
				Gender:     "male",
				ActiveDays: 3,
			},
			mockBehavior: func(s *mock_serv.MockUserServ, user inputUser) {
				expectedUser := models.User{
					ID:           1,
					Name:         "Ivan",
					Age:          19,
					Email:        "vanya.sustatov@mail.ru",
					Weight:       76,
					Height:       184,
					Gender:       "male",
					ActiveDays:   3,
					CaloriesNorm: 2950.16,
				}
				s.EXPECT().RegisterUser(user.Name, user.Age, user.Email, user.Password, user.Weight, user.Height, user.Gender, user.ActiveDays).Return(expectedUser, nil)
			},
			expectedStatusCode:  201,
			expectedRequestBody: `{"ID":1,"Name":"Ivan","Age":19,"Email":"vanya.sustatov@mail.ru","Weight":76,"Height":184,"Gender":"male","ActiveDays":3, "CaloriesNorm":2950.16}`,
		},
		{
			name:      "Empty name",
			inputBody: `{"name": "","age": 19,"email": "vanya.sustatov@mail.ru","password": "25566255","weight": 76,"height": 184,"gender": "male","activeDays": 3}`,
			inputUser: inputUser{
				Name:       "",
				Age:        19,
				Email:      "vanya.sustatov@mail.ru",
				Password:   "25566255",
				Weight:     76,
				Height:     184,
				Gender:     "male",
				ActiveDays: 3,
			},
			mockBehavior:        func(s *mock_serv.MockUserServ, user inputUser) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"Name":"must be at least 2 characters"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_serv.NewMockUserServ(c)
			testCase.mockBehavior(service, testCase.inputUser)

			handler := NewUserHandler(service)

			route := gin.New()
			route.POST("/auth/register", handler.RegisterUser)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/register", bytes.NewBufferString(testCase.inputBody))

			route.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.JSONEq(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
