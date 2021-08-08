package routes

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNotFoundRoute(t *testing.T) {
	if err := godotenv.Load("../../.env.test"); err != nil {
		panic(err)
	}

	tests := []struct {
		description   string
		route         string
		method        string
		tokenString   string
		body          io.Reader
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "get unknown route",
			route:         "/api/v1/cd",
			method:        "GET",
			tokenString:   "",
			body:          nil,
			expectedError: false,
			expectedCode:  404,
		},
	}

	app := fiber.New()

	NotFoundRoute(app)

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, test.body)
		req.Header.Set("Authorization", test.tokenString)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)
		if test.expectedError {
			continue
		}
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
