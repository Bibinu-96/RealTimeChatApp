package openapi

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Load the OpenAPI spec
	//swagger, err := openapi.GetSwagger()
	// if err != nil {
	// 	log.Fatalf("Failed to load OpenAPI spec: %v", err)
	// }

	// Create a router
	r := chi.NewRouter()

	// Use validation middleware
	//r.Use(middleware.OapiRequestValidator(swagger))

	// Register your handlers
	r.Route("/api", func(r chi.Router) {
		r.Post("/register", RegisterUser)
		r.Post("/login", LoginUser)
		r.Get("/users/connected", GetConnectedUsers)
		r.Get("/groups", GetGroups)
		r.Get("/messages/direct", GetDirectMessages)
		r.Get("/messages/group", GetGroupMessages)
	})

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
}

func GetConnectedUsers(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
}

func GetGroups(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
}

func GetDirectMessages(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
}

func GetGroupMessages(w http.ResponseWriter, r *http.Request) {
	// Implement your logic here
}
