package handlers

import (
	"database/sql"
	"fmt"
)

type User struct {
	User_id              int    `json:"id"`
	Full_name            string `json:"full_name"`
	User_name            string `json:"user_name"`
	Email                string `json:"email"`
	Password             int    `json:"password"`
	Home_planet          string `json:"planet"`
	Profile_picture_path string `json:"profile_picture_path"`
	Admin                bool   `json:"bool"`
}

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
func LoginCorrect(email string, password string, postgres *sql.DB) *User {
	//hashedPassword := hash(password);
	hashedPassword := 0
	stmt, err := postgres.Prepare("Select * from USER WHERE user_name = $1 AND password = $2")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(email, hashedPassword)
	if err2 != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var newUser User
		err := rows.Scan(&newUser.User_id, &newUser.Full_name, &newUser.User_name,
			&newUser.Email, &newUser.Password, &newUser.Home_planet, &newUser.Profile_picture_path, &newUser.Admin)
		if err != nil {
			panic(err)
		} else {
			return &newUser
		}

	}
	return nil

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
		if err2 != nil {
			return false
		}
		return true

	} else {
		stmt, err = postgres.Prepare("Insert Into Orbit_Requests (requester_id,requested_buddy_id) Values ($1,$2)")
		if err != nil {
			panic(err)
		}
		defer stmt.Close()

		_, err2 = stmt.Exec(sender_id, receiver_id)
		if err2 != nil {
			return false
		}
		return true
	}
}
func RegisterUser(postgres *sql.DB) User {
	var user User
	return user
}

func DeleteUser(user_id int, postgres *sql.DB) bool {
	stmt, err := postgres.Prepare("DELETE FROM USER WHERE user_id = $1")
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
