package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestServer_createUser(t *testing.T) {	
	// Create a new server instance
	server := NewServer(testStore)

	// Create a new HTTP request
	reqBody := `{
		"username": "testuser",
		"password": "testpassword",
		"email": "test@example.com",
		"location": "Seoul",
		"phone": "1234567890"
	}`
	req, err := http.NewRequest("POST", "/users", strings.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Set the request context
	ctx, _ := gin.CreateTestContext(rr)
	ctx.Request = req

	// Call the createUser function
	server.createUser(ctx)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedBody := `{"id":1,"username":"testuser","email":"test@example.com","location":"testlocation","phone":"1234567890"}`
	if rr.Body.String() != expectedBody {
		t.Errorf("Expected response body %s, but got %s", expectedBody, rr.Body.String())
	}
}