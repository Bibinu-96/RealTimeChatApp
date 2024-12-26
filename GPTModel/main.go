package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiURL = "http://localhost:11434/v1/chat/completions"

// Request structure to accept the query string
type QueryRequest struct {
	Query string `json:"query"`
}

// Response structure to send the response back
type QueryResponse struct {
	Response string `json:"response"`
}
type ChatRequest struct {
	Model       string        `json:"model"`
	Messages    []ChatMessage `json:"messages"`
	MaxTokens   int           `json:"max_tokens,omitempty"`
	Temperature float64       `json:"temperature,omitempty"`
}

// Message structure for OpenAI API
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Response structure for OpenAI API
type ChatResponse struct {
	Choices []struct {
		Message      ChatMessage `json:"message"`
		FinishReason string      `json:"finish_reason"` // Added finish reason
	} `json:"choices"`
}

func estimateTokens(content string) int {
	// This is a very simple estimation. You can improve it using tokenization libraries.
	return len(content) / 4 // Average token length in English is roughly 4 characters
}
func QueryOllama(query string) (string, error) {
	responseString := ""
	// Your OpenAI API key (set as an environment variable)
	apiKey := "somekey"

	// Create the request payload
	//query := "Explain the history of hitler" // Example user query
	inputTokens := estimateTokens(query)
	modelTokenLimit := 4096 // GPT-3.5 token limit

	// Set the MaxTokens for the response, subtracting input tokens from model's token limit
	maxTokens := modelTokenLimit - inputTokens - 100 // Leave some buffer for system messages and overhead

	if maxTokens < 100 {
		maxTokens = 100 // Ensure MaxTokens is not less than 100
	}

	requestBody := ChatRequest{
		Model: "smollm:latest", // Change to "gpt-3.5-turbo" if needed
		Messages: []ChatMessage{
			{
				Role:    "system",
				Content: "You are a helpful assistant.",
			},
			{
				Role:    "user",
				Content: query,
			},
		},
		MaxTokens:   maxTokens,
		Temperature: 0.7,
	}

	// Convert the request payload to JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Printf("Error marshalling request data: %v\n", err)
		return "", err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating HTTP request: %v\n", err)
		return "", err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Execute the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response using io.ReadAll
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return "", err
	}

	// Check for errors in the response
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Status code %d\nResponse: %s\n", resp.StatusCode, string(body))
		return "", err
	}

	// Parse the response JSON
	var chatResponse ChatResponse
	err = json.Unmarshal(body, &chatResponse)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\n", err)
		return "", err
	}

	// Print the assistant's response
	if len(chatResponse.Choices) > 0 {
		responseString = chatResponse.Choices[0].Message.Content
	} else {
		fmt.Println("No response from assistant.")
	}
	return responseString, nil
}

// Handler function for the POST request
func queryHandler(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Parse the incoming JSON request
	var request QueryRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	query := request.Query

	// Process the query and generate a response
	responseStr, err := QueryOllama(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := QueryResponse{
		Response: responseStr, // Example response
	}

	// Encode the response as JSON and send it back
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Define the API endpoint
	http.HandleFunc("/query", queryHandler)

	// Start the server on port 8080
	fmt.Println("Starting server on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
