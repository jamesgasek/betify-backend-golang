package tester

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	baseURL     = "http://localhost:80"
	registerURL = baseURL + "/api/v1/register"
	loginURL    = baseURL + "/api/v1/login"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func main() {
	// Register a new user
	registerUser("example_user", "password123")

	// Login with the registered user
	token := loginUser("example_user", "password123")
	if token == "" {
		log.Fatal("Login failed")
	}

	fmt.Println("Successfully logged in. Token:", token)
}

func registerUser(username, password string) {
	requestBody, err := json.Marshal(RegisterRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal("Error encoding register request:", err)
	}

	resp, err := http.Post(registerURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal("Error sending register request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Register request failed with status:", resp.Status)
	}

	fmt.Println("User registered successfully")
}

func loginUser(username, password string) string {
	requestBody, err := json.Marshal(LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatal("Error encoding login request:", err)
	}

	resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal("Error sending login request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Login request failed with status:", resp.Status)
	}

	var loginResponse LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&loginResponse)
	if err != nil {
		log.Fatal("Error decoding login response:", err)
	}

	return loginResponse.Token
}
