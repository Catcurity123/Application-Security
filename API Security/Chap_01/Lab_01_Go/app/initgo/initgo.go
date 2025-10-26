package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

// connectToDB connects to MySQL and returns (message, reachable, *sql.DB)
func connectToDB() (string, string, bool, *sql.DB) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, name)
	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Sprintf("❌ Error opening DB: %v", err), "", false, nil
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		db.Close()
		return fmt.Sprintf("❌ Database not reachable: %v", err),"", false, nil
	}

	return "✅ Connected to MySQL successfully!", dsn, true, db
}


type Space struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}


func createSpaceHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Only allow POST
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var space Space
		if err := json.NewDecoder(r.Body).Decode(&space); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if space.Name == "" || space.Owner == "" {
			http.Error(w, "Missing name or owner", http.StatusBadRequest)
			return
		}

		// Insert into DB securely (prepared statement)
		result, err := db.Exec("INSERT INTO spaces (name, owner) VALUES (?, ?)", space.Name, space.Owner)
		if err != nil {
			http.Error(w, fmt.Sprintf("DB insert error: %v", err), http.StatusInternalServerError)
			return
		}

		spaceID, err := result.LastInsertId()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get inserted ID: %v", err), http.StatusInternalServerError)
			return
		}

		// Respond with JSON including URI
		response := map[string]interface{}{
			"name": space.Name,
			"uri":  fmt.Sprintf("/spaces/%d", spaceID),
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Location", fmt.Sprintf("/spaces/%d", spaceID))
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}



func main() {
	msg, dsn, dbReachable, db := connectToDB()
	log.Println(msg, dsn, dbReachable, db)

	if dbReachable {
		defer db.Close()
		http.HandleFunc("/spaces", createSpaceHandler(db))
		log.Println("🚀 Server running on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		log.Fatal("❌ Cannot start server, DB unreachable")
	}
}

/*
func main() {
	msg, connectionString, dbReachable, db := connectToDB()

	log.Println(msg)
	log.Println(connectionString)

	if dbReachable {
		defer db.Close()
	}
		http.HandleFunc("/", hello(db, dbReachable))
		log.Println("🚀 Starting server on :8080")
		http.ListenAndServe(":8080", nil)

}


*/



























// hello handler queries current time if DB is reachable
/*
func hello(db *sql.DB, dbReachable bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !dbReachable {
			fmt.Fprint(w, "❌ Database not reachable.")
			return
		}

		var now string
		err := db.QueryRow("SELECT NOW()").Scan(&now)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error querying DB: %v", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Hello from Go + MySQL! 🕒 Current DB time: %s", now)
	}
}
*/
