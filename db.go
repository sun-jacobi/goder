package main

import (
	"database/sql"
	"errors"
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *database

type database struct {
	inernal *sql.DB
}

//----------------------------------------------
// models
type Problem struct {
	testcases []string
	name      string
	id        int
}
type Result struct {
	author string
	id     int
}

//----------------------------------------------
// Controller
// Upload the Problem
func UploadProblem(context *multipart.FileHeader, tests *[]multipart.FileHeader) {
	//TODO
}

// Get a problem
func (database *database) GetProblem(probID int) (*Problem, error) {
	db := database.inernal
	row, err := db.Query("SELECT * FROM PROBLEM WHERE PROBID = $1", probID)
	defer row.Close()
	if err != nil {
		return nil, errors.New("Error Failed to get the problem")
	}
	var problem Problem
	err = row.Scan(&problem)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("Error Failed to scan the model")
	}
	return &problem, nil
}

//----------------------------------------------
func SetUpDB() {
	conn := getConnStr()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal("Error opening database")
	}

	DB = &database{db}
}

//----------------------------------------------
// helper function

// get the Connection String from environment file
func getConnStr() string {
	err := godotenv.Load(".db")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	user := "user=" + os.Getenv("user")
	password := "password=" + os.Getenv("password")
	dbname := "dbname=" + os.Getenv("dbname")
	sslmode := "sslmode=" + os.Getenv("sslmode")
	return strings.Join([]string{user, password, dbname, sslmode}, "-")
}
