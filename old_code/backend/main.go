package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"real-time-chat/handlers" // Your custom package for database and WebSocket handlers

	"github.com/golang-jwt/jwt/v4"
	gorillaHandlers "github.com/gorilla/handlers" // Alias the gorilla handlers package
)

var jwtKey = []byte("your_secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func main() {
	// Initialize the database
	handlers.InitializeDatabase()
	defer handlers.DB.Close()

	// Serve static files
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// WebSocket Routes
	http.HandleFunc("/ws", handlers.HandleConnections)

	// Authentication Routes
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)

	// Start the WebSocket message handler
	go handlers.HandleMessages()

	// Start the HTTP server with CORS middleware
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", gorillaHandlers.CORS(
		gorillaHandlers.AllowedOrigins([]string{"http://localhost:3000"}),         // Allow frontend origin
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),        // Allow specific methods
		gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Allow specific headers
	)(http.DefaultServeMux)))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil || creds.Username == "" || creds.Password == "" {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = handlers.RegisterUser(creds.Username, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil || creds.Username == "" || creds.Password == "" {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = handlers.AuthenticateUser(creds.Username, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	w.Write([]byte("Login successful"))
}
