package handlers

import (
	"database/sql"
	"fmt"
)

// Expect handler to use ctx.MustGet("postgres").(*sql.DB) and pass in session
func TestInsert(val string, postgres *sql.DB) {
	fmt.Println("Go!", postgres)
	stmt, err := postgres.Prepare("INSERT INTO Users (full_name,email,password) VALUES ($1,$2,22)")
	if err != nil {
		panic(err)
	}
	_, err2 := stmt.Exec(val, val)
	fmt.Println("Go2!")
	if err2 != nil {
		panic(err)
	}
}

// func TestSelect() {
// 	stmt, err := postgres.Prepare("SELECT * FROM test3 WHERE t = 1")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer stmt.Close()

// 	rows, err := stmt.Query()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var val int
// 		err := rows.Scan(&val)
// 		// If multiple columns => rows.Scan(&val1, &val2, &val3)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Printf("%d\n", val)
// 	}
// }

// func TestUpdate() {
// 	stmt, err := postgres.Prepare("UPDATE test3 SET t = 2 WHERE t = 1")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec()
// }

// func TestDelete(e int) {
// 	stmt, err := postgres.Prepare("DELETE FROM test3 WHERE t = ")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer stmt.Close()
// }

// func TestConditionalSelect() {
// 	stmt, err := postgres.Prepare("SELECT * FROM test3 WHERE t < 5 AND t > 0")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer stmt.Close()

// 	rows, err := stmt.Query()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var val int
// 		err := rows.Scan(&val)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Printf("%d\n", val)
// 	}
// }
