package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	datamodel "github.com/assimad8/go-fullstack/gen"
	_ "github.com/lib/pq"
)

func main() {
	// Define the database URI
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "postgres", "pirates1999+++", "localhost", 5432, "postgres")

	// Open the database connection
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close() // Ensure the database connection is closed when main exits

	// Create the store
	st := datamodel.New(db)

	// Create a new user
	params := datamodel.CreateUsersParams{
		UserName:    "emad lakhbizi",
		PassWordHash: "hash", // Make sure to use a proper hashed password
		Name:        "test",
	}

	// Call the CreateUsers function and handle potential errors
	if _,err := st.CreateUsers(context.Background(), params); err != nil {
		log.Fatalf("failed to create user: %v", err)
	}

	fmt.Println("User created successfully")
}
