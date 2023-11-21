package handlers

import (
	"database/sql"
	"fmt"
	// "math/big"
	// "github.com/EugeneL11/SpaceBook/pkg"
)

/*

func TestConditionalSelect(postgres *sql.DB) string {
	stmt, err := postgres.Prepare("SELECT * FROM test3 WHERE t < 5 AND t > 0")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return "unable to connect to db"
	}
	defer rows.Close()

	for rows.Next() {
		var val int
		err := rows.Scan(&val)
		if err != nil {
			return "unable to connect to db"
		}
		fmt.Printf("%d\n", val)
	}
}

*/

func TestSelect(val string, postgres *sql.DB) ([]User, string) {
	stmt, err := postgres.Prepare("SELECT * FROM Users WHERE full_name = $1")
	if err != nil {
		return nil, "unable to connect to db"
	}
	defer stmt.Close()

	rows, err := stmt.Query(val)
	if err != nil {
		return nil, "unable to connect to db"
	}
	defer rows.Close()
	var mySlice []User
	for rows.Next() {
		var newUser User
		err := rows.Scan(&newUser.User_id, &newUser.Full_name, &newUser.User_name,
			&newUser.Email, &newUser.Password, &newUser.Home_planet, &newUser.Profile_picture_path, &newUser.Admin)
		if err != nil {
			return nil, "unable to connect to db"
		}
		mySlice = append(mySlice, newUser)
	}
	return mySlice, "no error"
}

// Expect handler to use ctx.MustGet("postgres").(*sql.DB) and pass in session
func TestInsert(val string, postgres *sql.DB) string {
	fmt.Println("Go!", postgres)
	stmt, err := postgres.Prepare("INSERT INTO Users (full_name,email,password) VALUES ($1,$2,22)")
	if err != nil {
		return "unable to connect to db"
	}
	_, err2 := stmt.Exec(val, val)
	fmt.Println("Go2!")
	if err2 != nil {
		return "unable to connect to db"
	}

	return "no error"
}
