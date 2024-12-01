package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Struct for API Response
type TimeResponse struct {
	CurrentTime string `json:"current_time"`
}

var db *sql.DB

// database connection
func connDB() {
	var err error
	// database connection strings

	dsn := "root:@tcp(127.0.0.1:3306)/go_api"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	// Test database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Connection error: %v\n", err)
	}
	log.Println("Connection successful!")
}

// Push data into the database
func pushDataToDatabase(timestamp time.Time) error {
	// Insert data into the database
	result, err := db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", timestamp)
	if err != nil {
		return fmt.Errorf("error inserting timestamp: %v", err)
	}

	// Log the number of rows affected
	rowsAffected, _ := result.RowsAffected()
	log.Printf("Time log inserted successfully, Rows affected: %d\n", rowsAffected)
	return nil
}

// API Handler
func getCurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	// timezone for Toronto
	location, err := time.LoadLocation("America/Toronto")
	if err != nil {
		// log.Printf("Having problem to load the timezone: %v\n", err)
		http.Error(w, "Having problem to load the timezone", http.StatusInternalServerError) // 500 Internal Server Error
		return
	}

	// Get current time in Toronto
	currentTime := time.Now().In(location)

	// Push data into the database
	err = pushDataToDatabase(currentTime)
	if err != nil {
		// Log the error message to the console
		// log.Printf("Sorry!, Processing of time log to the server failed!: %v\n", err)

		http.Error(w, "Sorry!, Processing of time log to the server failed!", http.StatusInternalServerError)
		return
	}

	response := TimeResponse{CurrentTime: currentTime.Format("2006-01-02 15:04:05")}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Initialize database connection
	connDB()
	// Close database connection when main function exits
	defer db.Close()

	// API endpoint
	http.HandleFunc("/current-time", getCurrentTimeHandler)

	// server port
	port := 8080
	log.Printf("Server running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
