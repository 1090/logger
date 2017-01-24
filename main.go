package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
	"os"
	"time"
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"strconv"
)

func main() {
	// postgres
	db, err := sql.Open("postgres", "host=adsb1090 user=adsb1090 password=adsb1090 dbname=adsb1090") // todo env variables

	if err != nil {
		log.Fatal("Error: The data source arguments are not valid")
		os.Exit(1)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
		os.Exit(1)
	}

	// connect to mutability-dump1090
	conn, err := net.Dial("tcp", "127.0.0.1:30003") // todo env variables

	if err != nil {
		fmt.Printf("connection error: %s\n", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(conn)

	for {
		output, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}

		if len(strings.TrimSpace(output)) != 0 {
			//fmt.Print(output)

			values := strings.Split(output, ",")

			hex := values[4]
			timestamp := parseTimestamp(values[6], values[7])
			flight := values[10] // flight "name"
			altitude := parseInt(values[11]) // feet
			speed := parseInt(values[12]) // mph
			heading := parseInt(values[13]) // degrees
			lat := parseDouble(values[14])
			lon := parseDouble(values[15])
			src := strings.TrimSpace(output)

			addBeacon(db, hex, timestamp, flight, altitude, speed, heading, lat, lon, src);

			checkErr(err)
		}
	}
}

func parseTimestamp(date_part string, time_part string) int64 {
	form := "2006/01/02 15:04:05.999"

	timestamp, err := time.Parse(form, date_part + " " + time_part)

	checkErr(err)

	return timestamp.Unix()
}

func addBeacon(db *sql.DB, hex string, timestamp int64, flight string, altitude sql.NullInt64, speed sql.NullInt64, heading sql.NullInt64, lat sql.NullFloat64, lon sql.NullFloat64, src string) {
	_, err := db.Exec("INSERT INTO beacons (hex, timestamp, flight, altitude, speed, heading, lat, lon, src) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", hex, timestamp, flight, altitude, speed, heading, lat, lon, src)

	checkErr(err)
}

func parseInt(value string) sql.NullInt64 {
	if len(value) == 0 {
		return sql.NullInt64{Int64: 0, Valid: false}
	}

	s, err := strconv.ParseInt(value, 10, 64)

	checkErr(err)

	return sql.NullInt64{Int64: s, Valid: false}
}

func parseDouble(value string) sql.NullFloat64 {
	if len(value) == 0 {
		return sql.NullFloat64{Float64: 0, Valid: false}
	}

	s, err := strconv.ParseFloat(value, 64)

	checkErr(err)

	return sql.NullFloat64{Float64: s, Valid: true}
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
}
