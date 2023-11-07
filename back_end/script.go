package main

import "database/sql"

// Initialize both DBs with values
// Expect the var postgres to be initialized to connection
func initialize(postgres *sql.DB) {
	_, err := postgres.Exec("CREATE TYPE planet AS ENUM('Mercury', 'Venus', 'Earth', 'Mars', 'Jupiter', 'Saturn', 'Uranus','Neptune', 'Pluto')")
	if err != nil {
		panic(err)
	}
	_, err2 := postgres.Exec("CREATE TABLE Users (user_id SERIAL PRIMARY KEY,full_name varchar(60),user_name varchar(30),email varchar(60) UNIQUE NOT NULL,password bigint NOT NULL,home_planet planet,profile_picture_path varchar(70))")
	if err2 != nil {
		panic(err)
	}
}
