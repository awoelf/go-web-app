package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/awoelf/go-web-app/router"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type Comment struct {
	Name        string
	Subject     string
	CommentText string
}

type TestCases struct {
	Message        string
	Method         string
	Path           string
	ExpectedStatus int
	Body           Comment
}

func TestMain(t *testing.T) {
	app := fiber.New()

	app.Mount("/api", router.APIRouter())

	testCases := []TestCases{
		{
			Message:        "POST a new comment that matches the body.",
			Method:         "POST",
			Path:           "/api/comment",
			ExpectedStatus: 200,
			Body: Comment{
				Name:        "Alex",
				Subject:     "I love my cat",
				CommentText: "She is a sweet cat",
			},
		},
		{
			Message:        "POST a new comment that matches the body.",
			Method:         "POST",
			Path:           "/api/comment",
			ExpectedStatus: 200,
			Body: Comment{
				Name:        "Tris",
				Subject:     "I also love my cat",
				CommentText: "She is a sweeter cat",
			},
		},
		{
			Message:        "POST a new comment that matches the body.",
			Method:         "POST",
			Path:           "/api/comment",
			ExpectedStatus: 200,
			Body: Comment{
				Name:        "Blobby",
				Subject:     "My cat is OK",
				CommentText: "My cat is alright.",
			},
		},
	}

	var buf bytes.Buffer

	for _, test := range testCases {
		
		err := json.NewEncoder(&buf).Encode(test.Body)
		if err != nil {
			assert.Error(t, err)
		}

		req := httptest.NewRequest(test.Method, test.Path, &buf)
		res, _ := app.Test(req)

		assert.Equalf(t, test.ExpectedStatus, res.StatusCode, test.Message)
	}
}
