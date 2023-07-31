package handler

import (
	"bytes"
	"context"
	"errors"

	// "errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"grey/config"
	"grey/internal/domain"
	"grey/internal/service"
	"grey/internal/service/mock"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(r *mock_service.MockUserInterface, input domain.UserSignUpInput)
	ctx := context.Background()
	tests := []struct {
		name                 string
		inputBody            string
		inputUser            domain.UserSignUpInput
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"username": "noname", "firstname": "Федор", "lastname": "Федоров", "age": 23, "is_married": false, "password": "qwertyuiop"}`,
			inputUser: domain.UserSignUpInput{
				Username:  "noname",
				Firstname: "Федор",
				Lastname:  "Федоров",
				Age:       23,
				IsMarried: false,
				Password:  "qwertyuiop",
			},

			mockBehavior: func(r *mock_service.MockUserInterface, input domain.UserSignUpInput) {
				r.EXPECT().SignUp(ctx, input).Return(nil)
			},
			expectedStatusCode:   http.StatusCreated,
			expectedResponseBody: `{"message":"account created"}`,
		},
		{
			name:                 "Age less 18",
			inputBody:            `{"username": "noname", "firstname": "Федор", "lastname": "Федоров", "age": 17, "is_married": false, "password": "qwertyuiop"}`,
			mockBehavior:         func(r *mock_service.MockUserInterface, input domain.UserSignUpInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:                 "No username",
			inputBody:            `{"firstname": "Федор", "lastname": "Федоров", "age": 17, "is_married": false, "password": "qwertyuiop"}`,
			mockBehavior:         func(r *mock_service.MockUserInterface, input domain.UserSignUpInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:                 "Short password",
			inputBody:            `{"username": "noname", "firstname": "Федор", "lastname": "Федоров", "age": 17, "is_married": false, "password": "qwerty"}`,
			mockBehavior:         func(r *mock_service.MockUserInterface, input domain.UserSignUpInput) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "User already exists",
			inputBody: `{"username": "noname", "firstname": "Федор", "lastname": "Федоров", "age": 23, "is_married": false, "password": "qwertyuiop"}`,
			inputUser: domain.UserSignUpInput{
				Username:  "noname",
				Firstname: "Федор",
				Lastname:  "Федоров",
				Age:       23,
				IsMarried: false,
				Password:  "qwertyuiop",
			},

			mockBehavior: func(r *mock_service.MockUserInterface, input domain.UserSignUpInput) {
				r.EXPECT().SignUp(ctx, input).Return(errors.New("user already exists"))
			},
			// TODO: Fix test, he doesn't see *pq.Error
			// expectedStatusCode: http.StatusInternalServerError,
			// expectedResponseBody: `{"message":"internal server error"}`,
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"message":"user already exists"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			r := mock_service.NewMockUserInterface(c)
			test.mockBehavior(r, test.inputUser)

			cfg, _ := config.NewConfig()
			service := &service.Service{User: r}
			handler := Handler{service, cfg}

			// Init Endpoint
			// gin.SetMode(gin.ReleaseMode)
			router := gin.New()
			router.POST("/sign-up", handler.signUp)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
