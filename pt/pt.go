package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pborman/uuid"
	"log"
	"time"
)

const (
	dbuser     = "got"
	dbpassword = "fun4more"
	dbname     = "got"
	hostname   = "192.168.242.183"
)

func setupdbping(dbuser string, dbpassword string, dbname string, hostname string) *sql.DB {
	//setup the parameters and validate
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		dbuser, dbpassword, dbname, hostname)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		log.Fatal("Error: The data source arguments are not valid")
	}

	//check the database - actually opens the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}
	return db
}
func queryprintallrows(db *sql.DB) {
	defer un(trace("query TIMING"))
	//prepare the query statement
	var oids []string
	queryStmt, err := db.Prepare("SELECT oid FROM map")
	defer queryStmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := queryStmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var oid string
		if err := rows.Scan(&oid); err != nil {
			log.Fatal(err)
		}
		oids = append(oids, oid)
	}
	fmt.Println(oids)
}
func insertsomerows(r int, db *sql.DB) {
	defer un(trace("INSERT TIMING"))
	for i := 0; i <= r; i++ {
		//now lets run an insert statement
		insertStmt, err := db.Prepare("INSERT INTO map values($1)")
		defer insertStmt.Close()
		if err != nil {
			log.Fatal(err)
		}
		_, err = insertStmt.Exec(uuid.New())
		if err != nil {
			log.Fatal(err)
		}
	}
}
func trace(s string) (string, time.Time) {
	log.Println("START:", s)
	return s, time.Now()
}
func un(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("  END:", s, "ElapsedTime in seconds:", endTime.Sub(startTime))
}
func main() {
	//setup DB connection and Ping
	var db *sql.DB
	db = setupdbping(dbuser, dbpassword, dbname, hostname)
	defer db.Close()

	//start := time.Now()
	queryprintallrows(db)
	//elapsed := time.Since(start)
	//fmt.Printf("Time for the query: %q\n", elapsed)

	//start = time.Now()
	insertsomerows(10, db)
	//elapsed = time.Since(start)
	//fmt.Printf("Time for the insert: %q\n", elapsed)

	//start = time.Now()
	queryprintallrows(db)
	//elapsed = time.Since(start)
	//fmt.Printf("Time for the query: %q\n", elapsed)
}
