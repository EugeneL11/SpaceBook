package handlers

import (
	"database/sql"
	"fmt"
	// "math/big"
	// "github.com/EugeneL11/SpaceBook/pkg"
)

func LoginCorrect(email string, password string, postgres *sql.DB, user *User) string {
	// hashedPassword, err := pkg.GeneratePasswordHash(password)
	// if err != nil {
	// 	// Could not hash password
	// 	return false
	// }
	hashedPassword := 22
	stmt, err := postgres.Prepare("Select * from Users WHERE user_name = $1 AND password = $2")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(email, hashedPassword)
	if err2 != nil {
		return "unable to connect to db"
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&user.User_id, &user.Full_name, &user.User_name,
			&user.Email, &user.Password, &user.Home_planet, &user.Profile_picture_path, &user.Admin)

		if err != nil {
			return "unable to connect to db"
		}

		return "no errer"
	} else {
		return "invalid login credentials"
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
	// hashedPassword, err := pkg.GeneratePasswordHash(password)
	hashedPassword := 22
	// fmt.Println(hashedPassword)
	// if err != nil {
	// 	// Unable to hash the password
	// 	return "unable to hash password"
	// }
	stmt, err = postgres.Prepare(`
    INSERT INTO Users (full_name, user_name, email, password, home_planet, profile_picture_path, isAdmin, bio)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`)
	if err != nil {
		return "unable to connect to db"
	}
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

func DeleteUser(user_id int, postgres *sql.DB) string {
	stmt, err := postgres.Prepare("DELETE FROM Users WHERE user_id = $1")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_id)
	if err != nil {
		return "unable to connect to db"
	}

	return "no error"
}

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


func GetFriends(user_id int, postgres *sql.DB) ([]API_UserInfo, string) {
	stmt, err := postgres.Prepare(`
		SELECT (
			user_id, full_name, user_name,
			email, home_planet, 
			profile_picture_path, isAdmin, bio
		)
		FROM Users
		WHERE user_id IN (
			SELECT user2_id 
			FROM Orbit_Buddies 
			WHERE user1_id = $1
			UNION 
			SELECT user1_id 
			FROM Orbit_Buddies 
			WHERE user2_id = $1
		)
	`)
	if err != nil {
		return nil, "unable to connect to db"
	}
	defer stmt.Close()

	rows, err := stmt.Query(user_id)
	if err != nil {
		return nil, "unable to connect to db"
	}
	defer rows.Close()

	var mySlice []API_UserInfo
	for rows.Next() {
		var newUser API_UserInfo
		err := rows.Scan(
			&newUser.User_id, &newUser.Full_name, &newUser.User_name,
			&newUser.Email, &newUser.Home_planet, &newUser.Profile_picture_path,
			&newUser.Admin, &newUser.bio
		)
		if err != nil {
			return nil, "unable to connect to db"
		}
		mySlice = append(mySlice, newUser)
	}

	return mySlice, "no error"
}

func RemoveFriend(user1_id int, user2_id int, postgres *sql.DB) string {
	if user1_id > user2_id {
		temp := user1_id
		user1_id = user2_id
		user2_id = temp
	}

	stmt, err := postgres.Prepare(`
		DELETE FROM Orbit_Buddies
		WHERE user1_id = $1 AND user2_id = $2
	`)
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	_, err = stmt.Exec(user1_id, user2_id)
	if err != nil {
		return "unable to connect to db"
	}

	return "no error"
}

func SendFriendRequest(sender_id int, receiver_id int, postgres *sql.DB) string {
	stmt, err := postgres.Prepare("Select * from Orbit_Requests WHERE requester_id = $1 AND requested_buddy_id = $2")
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(receiver_id, sender_id)
	if err2 != nil {
		return "unable to connect to db"
	}
	if rows.Next() {
		stmt, err = postgres.Prepare("DELETE FROM Orbit_Requests WHERE requester_id = $1 AND requested_buddy_id = $2")
		if err != nil {
			return "unable to connect to db"
		}
		defer stmt.Close()

		_, err2 = stmt.Exec(receiver_id, sender_id)
		if err2 != nil {
			return "unable to connect to db"
		}
		stmt, err = postgres.Prepare("Insert Into Orbit_Buddies (user1_id,user2_id) Values ($1,$2)")
		if err != nil {
			return "unable to connect to db"
		}
		defer stmt.Close()

		if sender_id > receiver_id {
			temp := sender_id
			sender_id = receiver_id
			receiver_id = temp
		}

		_, err2 = stmt.Exec(sender_id, receiver_id)
		if err2 != nil {
			return "unable to connect to db"
		}

		return "no error"

	} else {
		stmt, err = postgres.Prepare("Insert Into Orbit_Requests (requester_id,requested_buddy_id) Values ($1,$2)")
		if err != nil {
			return "unable to connect to db"
		}
		defer stmt.Close()

		_, err2 = stmt.Exec(sender_id, receiver_id)
		if err2 != nil {
			return "unable to connect to db"
		}

	}

	return "no error"
}

func RejectFriendRequest(rejecter_id int, rejectee_id int, postgres *sql.DB) string {
	stmt, err := postgres.Prepare(`
		DELETE FROM Orbit_Requests 
		WHERE requested_buddy_id = $1 AND requester_id = $2 
	`)
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(rejecter_id, rejectee_id)
	if err2 != nil {
		return "unable to connect to db"
	}

	return "no error"
}

func GetFriendRequests(user_id int, postgres *sql.DB) ([]API_UserInfo, string) {
	stmt, err := postgres.Prepare(`
		SELECT requester_id 
		FROM Orbit_Requests 
		WHERE requested_buddy_id = $1
	`)
	if err != nil {
		return nil, "unable to connect to db"
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(user_id)
	if err2 != nil {
		return nil, "unable to connect to db"
	}

	var mySlice []API_UserInfo
	for rows.Next() {
		var newUser API_UserInfo
		err := rows.Scan(
			&newUser.User_id, &newUser.Full_name, &newUser.User_name,
			&newUser.Email, &newUser.Home_planet, &newUser.Profile_picture_path, 
			&newUser.Admin, &newUser.Bio
		)
		if err != nil {
			return nil, "unable to connect to db"
		}
		mySlice = append(mySlice, newUser)
	}

	return mySlice, "no error"

}

func UpdateUserProfile(
	user_id int, new_username string, new_fullname string, 
	new_email string, new_home_planet string, 
	new_profile_pic_path string, bio string, postgres *sql.DB
) string {
	stmt, err := postgres.Prepare(`
		UPDATE Users 
		SET user_name = $2, full_name = $3,
		email = $4, planet = $5,
		profile_picture_path = $6, bio = $7
		WHERE user_id = $1
	`)
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_id, new_username, new_fullname, new_email, new_home_planet, new_profile_pic_path, bio)
	if err != nil {
		return "unable to connect to db"
	}

	return "no error"
}


func GetUserInfo(user_id int, postgres *sql.DB, userInfo *API_UserInfo) (API_UserInfo, string) {
	stmt, err := postgres.Prepare(`
		SELECT (
			user_id, full_name, user_name, 
			email, home_planet, 
			profile_picture_path, isAdmin, bio
		)
		FROM Users 
		WHERE user_id = $1
	`)
	if err != nil {
		return nil, "unable to connect to db"
	}
	defer stmt.Close()

	rows, err2 := stmt.Query(user_id)
	if err2 != nil {
		return nil, "unable to connect to db"
	}

	row := rows.next()
	if row {
		err := rows.Scan(
			&userInfo.User_id, &userInfo.Full_name, &userInfo.User_name,
			&userInfo.Email, &userInfo.Home_planet, &userInfo.Profile_picture_path, 
			&userInfo.Admin, &userInfo.Bio)
		if err != nil {
			return nil, "unable to connect to db"
		}

		return userInfo, "no error"
	} else {
		return nil, "invalid user"
	}
	
	return "no error"
}

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