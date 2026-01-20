package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mock_serv "github.com/sustatov027-max/project_calorie_tracker/internal/handlers/mock"
	"github.com/sustatov027-max/project_calorie_tracker/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	svc := mock_serv.NewMockUserServ(ctrl)

	type mockSetup func(s *mock_serv.MockUserServ)

	tests := []struct {
		name       string
		request    string
		mockSetup  mockSetup
		wantStatus int
		wantBody   string
	}{
		{
			name:    "Success",
			request: `{"name":"Test","age":19,"email":"test@mail.ru","password":"12345678","weight":76,"height":184,"gender":"male","activeDays":3}`,
			mockSetup: func(s *mock_serv.MockUserServ) {
				svc.EXPECT().RegisterUser(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
					models.User{
						ID:           1,
						Name:         "Test",
						Age:          19,
						Email:        "test@mail.ru",
						Weight:       76,
						Height:       184,
						Gender:       "male",
						ActiveDays:   3,
						CaloriesNorm: 2950.16}, nil,
				)
			},
			wantStatus: 201,
			wantBody:   `"ID":1,"Name":"Test","Age":19,"Email":"test@mail.ru","Weight":76,"Height":184,"Gender":"male","ActiveDays":3,"CaloriesNorm":2950.16}`,
		},
		{
			name:       "Empty name",
			request:    `{"name":"","age":19,"email":"test@mail.ru","password":"12345678","weight":76,"height":184,"gender":"male","activeDays":3}`,
			mockSetup:  func(s *mock_serv.MockUserServ) {},
			wantStatus: 400,
			wantBody:   `{"Name":"must be at least 2 characters"}`,
		},
		{
			name:       "Short password",
			request:    `{"name":"Test","age":19,"email":"test@mail.ru","password":"1234567","weight":76,"height":184,"gender":"male","activeDays":3}`,
			mockSetup:  func(s *mock_serv.MockUserServ) {},
			wantStatus: 400,
			wantBody:   `{"Password":"must be at least 8 characters"}`,
		},
		{
			name:       "Invalid email",
			request:    `{"name":"Test","age":19,"email":"test.ru","password":"12345678","weight":76,"height":184,"gender":"male","activeDays":3}`,
			mockSetup:  func(s *mock_serv.MockUserServ) {},
			wantStatus: 400,
			wantBody:   `{"Email":"must be a valide email"}`,
		},
		{
			name:       "Invalid weight",
			request:    `{"name":"Test","age":19,"email":"test@mail.ru","password":"12345678","weight":-1,"height":184,"gender":"male","activeDays":3}`,
			mockSetup:  func(s *mock_serv.MockUserServ) {},
			wantStatus: 400,
			wantBody:   `{"Weight":"is invalid"}`,
		},
		{
			name:       "Invalid height",
			request:    `{"name":"Test","age":19,"email":"test@mail.ru","password":"12345678","weight":76,"height":-1,"gender":"male","activeDays":3}`,
			mockSetup:  func(s *mock_serv.MockUserServ) {},
			wantStatus: 400,
			wantBody:   `{"Height":"is invalid"}`,
		},
		{
			name:       "Invalid activeDays",
			request:    `{"name":"Test","age":19,"email":"test@mail.ru","password":"12345678","weight":76,"height":184,"gender":"male","activeDays":-1}`,
			mockSetup:  func(s *mock_serv.MockUserServ) {},
			wantStatus: 400,
			wantBody:   `{"ActiveDays":"must be at least 0 characters"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockSetup != nil {
				tt.mockSetup(svc)
			}

			handler := NewUserHandler(svc)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/auth/register",
				strings.NewReader(tt.request))

			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.POST("/auth/register", handler.RegisterUser)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)

			if tt.wantStatus == 201 {
				var resp map[string]interface{}
				json.Unmarshal(w.Body.Bytes(), &resp)
				assert.NotZero(t, resp["ID"])
				assert.Equal(t, "Test", resp["Name"])
			} else {
				assert.JSONEq(t, tt.wantBody, w.Body.String())
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	type loginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type mockSetup func(s *mock_serv.MockUserServ, input loginInput, token string)

	testTable := []struct {
		name       string
		request    string
		mockSetup  mockSetup
		wantStatus int
		wantBody   string
	}{
		{
			name:    "Success",
			request: `{"email": "test@mail.ru", "password": "12345678"}`,
			mockSetup: func(s *mock_serv.MockUserServ, input loginInput, token string) {
				s.EXPECT().LoginUser(input.Email, input.Password).Return(token, nil)
			},
			wantStatus: 200,
			wantBody:   `{"token":"mock-jwt-token-12345", "token_type":"Bearer"}`,
		},
		{
			name:    "Invalid credentials",
			request: `{"email": "wrong@mail.ru", "password": "wrong"}`,
			mockSetup: func(s *mock_serv.MockUserServ, input loginInput, token string) {
				s.EXPECT().LoginUser(input.Email, input.Password).Return("", errors.New("invalid credentials"))
			},
			wantStatus: 401,
			wantBody:   `{"Error login user": "invalid credentials"}`,
		},
		{
			name:    "Invalid JSON",
			request: `{"email": "test@mail.ru"}`,
			mockSetup: func(s *mock_serv.MockUserServ, input loginInput, token string) {
				s.EXPECT().LoginUser(gomock.Any(), gomock.Any()).
					Return("", errors.New("json parse error")).AnyTimes()
			},
			wantStatus: 401,
			wantBody:   `{"Error login user": "json parse error"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			service := mock_serv.NewMockUserServ(ctrl)

			var input loginInput
			json.Unmarshal([]byte(testCase.request), &input)

			fixedToken := "mock-jwt-token-12345"

			testCase.mockSetup(service, input, fixedToken)

			handler := NewUserHandler(service)

			route := gin.New()
			route.POST("/auth/login", handler.LoginUser)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/auth/login",
				bytes.NewBufferString(testCase.request))

			route.ServeHTTP(w, req)

			assert.Equal(t, testCase.wantStatus, w.Code)
			assert.JSONEq(t, testCase.wantBody, w.Body.String())
		})
	}
}
