package handlers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/EugeneL11/SpaceBook/pkg"
	// "math/big"
	// "github.com/EugeneL11/SpaceBook/pkg"
)

// Expect handler to use ctx.MustGet("postgres").(*sql.DB) and pass in session
func TestInsert(val string, postgres *sql.DB) bool {
	fmt.Println("Go!", postgres)
	stmt, err := postgres.Prepare("INSERT INTO Users (full_name,email,password) VALUES ($1,$2,22)")
	if err != nil {
		panic(err)
	}
	_, err2 := stmt.Exec(val, val)
	fmt.Println("Go2!")
	if err2 != nil {
		panic(err)
	} else {
		return true
	}
}

func TestSelect(val string, postgres *sql.DB) []User {
	stmt, err := postgres.Prepare("SELECT * FROM Users WHERE full_name = $1")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(val)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var mySlice []User
	for rows.Next() {
		var newUser User
		err := rows.Scan(&newUser.User_id, &newUser.Full_name, &newUser.User_name,
			&newUser.Email, &newUser.Password, &newUser.Home_planet, &newUser.Profile_picture_path, &newUser.Admin)
		if err != nil {
			panic(err)
		}
		mySlice = append(mySlice, newUser)
	}
	return mySlice
}

func UpdateUser(Home_planet, string, email string, full_name string, postgres *sql.DB) bool {
	stmt, err := postgres.Prepare("UPDATE USER SET full_name = $1 WHERE email = $2")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(full_name, email)
	if err != nil {
		panic(err)
	} else {
		return true
	}

}

func LoginCorrect(username string, password string, postgres *sql.DB, user *User) bool {
	stmt, err := postgres.Prepare("SELECT password FROM users WHERE user_name = $1")
	if err != nil {
		return false
	}
	defer stmt.Close()

	// Query the database and get the hashed password stored for the user
	row := stmt.QueryRow(username)
	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		return false
	}

	// Compare the hashed password with the user provided password
	isCorrect := pkg.VerifyPassword(password, []byte(hashedPassword))

	if !isCorrect {
		// fmt.Println("Entered password", password, "incorrectly matches hashed password! :(")
		return false
	}

	// fmt.Println("Entered password", password, "correctly matches hashed password!")

	// Get user's information to give frontend
	stmt, err = postgres.Prepare("SELECT * FROM users WHERE user_name = $1 and password = $2")
	if err != nil {
		log.Panic(err)
		return false
	}
	defer stmt.Close()

	rows, err := stmt.Query(username, hashedPassword)
	if err != nil {
		log.Panic(err)
		return false
	}
	if rows.Next() {
		err := rows.Scan(&user.User_id, &user.Full_name, &user.User_name,
			&user.Email, &user.Password, &user.Home_planet, &user.Profile_picture_path, &user.Admin, &user.Bio)
		log.Println(err)
		return err == nil

	} else {
		return false
	}
}

func SendFriendRequest(sender_id int, receiver_id int, postgres *sql.DB) bool {
	stmt, err := postgres.Prepare("Select * from Orbit_Requests WHERE requester_id = $1 AND requested_buddy_id = $2")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(receiver_id, sender_id)
	if err2 != nil {
		panic(err)
	}
	if rows.Next() {
		stmt, err = postgres.Prepare("DELETE FROM Orbit_Requests WHERE requester_id = $1 AND requested_buddy_id = $2")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		_, err2 = stmt.Exec(receiver_id, sender_id)
		if err2 != nil {
			panic(err)
		}
		stmt, err = postgres.Prepare("Insert Into Orbit_Buddies (user1_id,user2_id) Values ($1,$2)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()
		if sender_id > receiver_id {
			temp := sender_id
			sender_id = receiver_id
			receiver_id = temp
		}
		_, err2 = stmt.Exec(sender_id, receiver_id)
		return err2 == nil

	} else {
		stmt, err = postgres.Prepare("Insert Into Orbit_Requests (requester_id,requested_buddy_id) Values ($1,$2)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		_, err2 = stmt.Exec(sender_id, receiver_id)

		return err2 == nil

	}
}
func RegisterUser(fullName string, password string, email string, username string, postgres *sql.DB, user *User) string {
	stmt, err := postgres.Prepare("SELECT * FROM Users WHERE email = $1")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()
	rows, err2 := stmt.Query(email)

	if err2 != nil {
		return "unable to connect to db"
	}

	if rows.Next() {
		return "email taken"
	}
	stmt, err = postgres.Prepare("SELECT * FROM Users WHERE user_name = $1")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	rows, err2 = stmt.Query(username)
	if err2 != nil {
		return "unable to connect to db"
	}

	if rows.Next() {
		return "user name taken"
	}
	hashedPassword, err := pkg.HashPassword(password)
	if err != nil {
		// Unable to hash the password
		return "unable to hash password"
	}
	stmt, err = postgres.Prepare(`
    INSERT INTO Users (full_name, user_name, email, password, home_planet, profile_picture_path, isAdmin, bio)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)
	if err != nil {
		fmt.Println("Made it here?")
		return "unable to connect to db"
	}
	//fmt.Println("Made it here?")
	_, err = stmt.Exec(fullName, username, email, hashedPassword, "Earth", "default", false, "test bio")
	if err != nil {
		fmt.Println("Could not execute insert into users")
		return "unable to connect to db"
	}
	stmt, err = postgres.Prepare("SELECT user_id FROM Users WHERE email = $1")
	if err != nil {
		return "unable to connect to db"
	}
	rows, err = stmt.Query(email)
	if err != nil {
		return "unable to connect to db"
	}
	if rows.Next() {
		err = rows.Scan(&user.User_id)
		if err != nil {
			return "unable to connect to db"
		}
	}
	user.Full_name = fullName
	user.User_name = username
	user.Home_planet = "Earth"
	user.Profile_picture_path = "Default"
	user.Email = email
	user.Admin = false

	return "no error"
}

func DeleteUser(user_id int, postgres *sql.DB) bool {
	stmt, err := postgres.Prepare("DELETE FROM Users WHERE user_id = $1")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_id)
	if err != nil {
		panic(err)
	} else {
		return true
	}
}

/*
func TestConditionalSelect(postgres *sql.DB) {
	stmt, err := postgres.Prepare("SELECT * FROM test3 WHERE t < 5 AND t > 0")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var val int
		err := rows.Scan(&val)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d\n", val)
	}
}*/

/*

 */
