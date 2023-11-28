package handlers

import (
	"time"
	"log"
	// "encoding/json"
	"fmt"
	"database/sql"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

// not done
// not tested
func CreateNewDM(user1 int, user2 int, cassandra *gocql.Session) string {
	return "no error"
}

// not done
// not tested
// not documented
func CreateNewDMHandler(ctx *gin.Context) {

}

func sendDM(senderID int, receiverID int, message string, cassandra *gocql.Session) string {

	fmt.Println("sendDM")

	var user1 int
	var user2 int

	user1 = senderID
	user2 = receiverID

	if user1 > user2 {
		temp := user1
		user1 = user2
		user2 = temp
	}

	// get DM between users (prepare)
	stmt := cassandra.Query(
		`
			SELECT messageChunks
			FROM dmtable
			WHERE user1 = ? AND user2 = ?
		`, user1, user2,
	)

	// get DM between users (execute)
	iter := stmt.Iter()
	
	// check how many message subsets there are
	var messageChunks []gocql.UUID
	for iter.Scan(&messageChunks) {
		break
	}
	
	num_subsets := len(messageChunks)

	// catch error
	if err := iter.Close(); err != nil {
		return "unable to connect to db 1"
	}

	var new_subset_needed bool
	new_subset_needed = true

	// recent subset
	var recent_subsetID gocql.UUID

	// if at least one subset
	if num_subsets > 0 {
		// get most recent subset
		recent_subsetID = messageChunks[num_subsets - 1]
		
		// get subset
		stmt := cassandra.Query(
			`
				SELECT senders
				FROM dmsubset
				WHERE subsetID = ?
			`, recent_subsetID,
		)
		iter := stmt.Iter()
		
		// count messages
		var senders []int
		for iter.Scan(&senders) {
			break
		}

		message_count := len(senders)

		// close iterator
		if err := iter.Close(); err != nil {
			return "unable to connect to db 2"
		}

		// see if subset is full
		new_subset_needed = (message_count >= 20)
	}
		

	// if (subset count == 0) or (subset is full)
	if new_subset_needed {
		subsetID := gocql.TimeUUID()

		// new subset
		err := cassandra.Query(
			`
				INSERT INTO dmsubset (subsetID, messages, senders, time_sent) VALUES (?, [], [], [])
			`, subsetID,
		).Exec()

		// catch error
		if err != nil {
			return "unable to connect to db 3"
		}

		subset_slice := []gocql.UUID{subsetID}

		// add subset to DM -> add subset to array
		// messageChunks = append(messageChunks, subsetID)

		// add subset to DM -> modify table
		err = cassandra.Query(
			`
				UPDATE dmtable
				SET messageChunks = messageChunks + ?
				WHERE user1 = ? AND user2 = ?
			`, subset_slice, user1, user2,
		).Exec()

		// catch error
		if err != nil {
			return "unable to connect to db 4"
		}

		// set recent_subsetID
		recent_subsetID = subsetID
	}

	fmt.Println(recent_subsetID)
		
	// send message within subset -> get messages
	stmt = cassandra.Query(
		`
			SELECT messages, senders, time_sent
			FROM dmsubset
			WHERE subsetID = ?
		`, recent_subsetID,
	)

	iter = stmt.Iter()

	// set subset variables
	var messages []string
	var senders []int
	var time_sent []time.Time

	for iter.Scan(&messages, &senders, &time_sent) {
		break
	}

	// send message within subset -> add new message
	messages_slice := []string{message}
	senders_slice := []int{senderID}
	times_slice := []time.Time{time.Now()}
	
	// send message within subset -> modify table
	err := cassandra.Query(
		`
			UPDATE dmsubset
			SET 
				messages = messages + ?, 
				senders = senders + ?, 
				time_sent = time_sent + ?
			WHERE subsetID = ?
		`, messages_slice, senders_slice, times_slice, recent_subsetID,
	).Exec()

	// catch error
	if err != nil {
		return "unable to connect to db 5"
	}
	
	return "no error"
}

// not documented
func SendDMHandler(ctx *gin.Context) {
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	senderID, err := strconv.Atoi(ctx.Param("senderID"))
	if err != nil {
		log.Panic(err)
	}
	receiverID, err := strconv.Atoi(ctx.Param("receiverID"))
	if err != nil {
		log.Panic(err)
	}
	message := ctx.Param("message")
	status := sendDM(senderID, receiverID, message, cassandra)
	ctx.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}

// not tested
func GetAllDM(userID int, usernames *[]string, profile_pics *[]string, recent_messages *[]string, postgres *sql.DB, cassandra *gocql.Session) string {

	// requirements:
		// profile pic
		// username
		// most recent message
	
	// get all dm's that user is in
	iter := cassandra.Query(
		`
			SELECT *
			FROM dmtable
			WHERE user1 = ? OR user2 = ?
		`, userID, userID,
	).Iter()

	var user1 int
	var user2 int
	var messageChunks []gocql.UUID

	// for each dm user is in
	for iter.Scan(&user1, &user2, &messageChunks) {
		// other user's id
		otherID := user1
		if user1 == userID {
			otherID = user2
		}

		// username, profile pic path
		stmt, err := postgres.Prepare(
			`
				SELECT user_name, profile_picture_path
				FROM Users
				WHERE user_id = $1
			`,
		)

		if err != nil {
			return "unable to connect to db 1"
		}
		defer stmt.Close()

		res, err := stmt.Query(otherID)
		if err != nil {
			return "unable to connect to db 2"
		}
		defer res.Close()

		var userPrev UserPreview
		for res.Next() {
			err := res.Scan(
				&userPrev.Full_name, &userPrev.User_name,
				&userPrev.Profile_picture_path,
			)
			if err != nil {
				return "unable to connect to db 3"
			}
			break
		}

		recent_message := ""

		// most recent message
		num_subsets := len(messageChunks)
		if num_subsets > 0 {
			// most recent subset
			recent_subsetID := messageChunks[num_subsets - 1]

			// query subset
			iter = cassandra.Query(
				`
					SELECT messages
					FROM dmsubset
					WHERE subsetID = ?
				`, recent_subsetID,
			).Iter()

			// most recent messages
			var messages []string
			for iter.Scan(&messages) {
				num_messages := len(messages)
				recent_message = messages[num_messages - 1]
				break
			}
		}

		// append to arrays
		*usernames = append(*usernames, userPrev.User_name)
		*profile_pics = append(*profile_pics, userPrev.Profile_picture_path)
		*recent_messages = append(*recent_messages, recent_message)
	}
	
	return "no error"
}

// not tested
// not documented
func GetDMHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		log.Panic(err)
	}

	usernames := []string{}
	profile_pics := []string{}
	recent_messages := []string{}

	status := GetAllDM(userID, &usernames, &profile_pics, &recent_messages, postgres, cassandra)

	// TODO is Marshal supposed to be used?

	// usernames_json, err := json.Marshal(usernames)
	// if err != nil { log.Panic(err) }
	// profile_pics_json, err := json.Marshal(profile_pics)
	// if err != nil { log.Panic(err) }
	// recent_messages_json, err := json.Marshal(recent_messages)
	// if err != nil { log.Panic(err) }

	ctx.JSON(http.StatusOK, gin.H{
		"status": status,
		"usernames": usernames,
		"profile pics": profile_pics,
		"recent messages": recent_messages,
	})
}

// not done
// not tested
func newDMList(userID int, newDMRes *[]int, postgres *sql.DB, cassandra *gocql.Session) string {
	
	// get all friends
	users, err := GetFriends(userID, postgres)
	if err != "no error" {
		return "unable to connect to db 1"
	}

	// get all dm's
	iter := cassandra.Query(
		`
			SELECT user2
			FROM dmtable
			WHERE user1 = ?
			UNION 
			SELECT user1
			FROM dmtable
			WHERE user2 = ?
		`, userID, userID,
	).Iter()
	
	var dmID int
	dmIDs := make(map[int]struct{}) // set of friends in dm
	for iter.Scan(&dmID) {
		dmIDs[dmID] = struct{}{}
	}

	// list all friends except those in dm's
	for f := 0; f < len(users); f++ {
		user := users[f]
		if _, exists := dmIDs[user.User_id]; !exists {
			// if it doesn't exist
			*newDMRes = append(*newDMRes, user.User_id)
		}
	}
	
	return "no error"
}

// not done
// not tested
// not documented
func NewDMListHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		log.Panic(err)
	}

	newDMRes := []int{}

	status := newDMList(userID, &newDMRes, postgres, cassandra)

	ctx.JSON(http.StatusOK, gin.H{
		"status": status,
		"newDMRes": newDMRes,
	})
}

// not done
// not tested
func GetMessages(user1 int, user2 int, subsetSize int, cassandra *gocql.Session) string {
	return "no error"
}

// not done
// not tested
// not documented
func GetMessagesHandler(ctx *gin.Context) {

}
