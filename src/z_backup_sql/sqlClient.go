package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

const dfltPort int = 1433
const dfltPassword string = "^h5b6bLALq5Rx2!y"
const dfltServer string = "192.168.14.150"

var dbconn *sql.DB
var database = "KUNDAN_DB"

func getArgs(server *string, user *string, password *string, port *int) {
	flag.StringVar(server, "server", dfltServer, "database server")
	flag.StringVar(user, "user", "sa", "user id")
	flag.StringVar(password, "password", dfltPassword, "password")
	flag.IntVar(port, "port", dfltPort, "DB Port")

	flag.Parse()
}

func getVersion(conn *sql.DB) string {
	var sqlversion string

	rows, err := conn.Query("select @@version")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&sqlversion)
		if err != nil {
			log.Fatal(err)
		}
	}
	return sqlversion
}

func main() {

	var server, user, password string
	var port int
	var err error

	getArgs(&server, &user, &password, &port)
	// fmt.Println(server, user, password, port)

	// Creae SQL ODBC URL
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, password, port)

	// Create Connection Pool
	dbconn, err = sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Error open db: ", err.Error())
	}
	if dbconn == nil {
		log.Fatal("DB connection is empty, but sql.Open() != nil")
	}

	ctx := context.Background()
	err = dbconn.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected!")

	defer dbconn.Close()

	sqlversion := getVersion(dbconn)
	fmt.Println("\n", sqlversion)

	// Create employee
	createID, err := CreateEmployee("Jake", "United States")
	if err != nil {
		log.Fatal("Error creating Employee: ", err.Error())
	}
	fmt.Printf("Inserted ID: %d successfully.\n", createID)
}

/*
 * Insert a row in TestSchema.Employees
 */
func CreateEmployee(name, location string) (int64, error) {
	ctx := context.Background()
	var err error

	if dbconn == nil {
		err = errors.New("CreateEmployee: db is null")
		return -1, err
	}

	// Check if database is alive.
	err = dbconn.PingContext(ctx)
	if err != nil {
		return -1, err
	}
	fmt.Println("1")

	tsql := "INSERT INTO TestSchema.Employees (Name, Location) VALUES (@Name, @Location); select convert(bigint, SCOPE_IDENTITY());"

	stmt, err := dbconn.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	fmt.Println("2")

	row := stmt.QueryRowContext(
		ctx,
		sql.Named("Name", name),
		sql.Named("Location", location))
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		fmt.Println("---- error 3 ----", err.Error())
		return -1, err
	}
	fmt.Println("3")

	return newID, nil
}
