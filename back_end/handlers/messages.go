package handlers

import (
	"time"
	"log"
	"fmt"
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
		var message_count int
		message_count = len(senders)

		// close iterator
		if err := iter.Close(); err != nil {
			return "unable to connect to db 2"
		}

		// see if subset is full
		new_subset_needed = (message_count >= 20)
	}
		

	// if (subset count == 0) or (subset is full)
	if new_subset_needed {
		// set subset variables
		var subsetID gocql.UUID
		// messages := []string{} 
		// senders := []int{}
		// time_sent := []time.Time{}

		subsetID = gocql.TimeUUID()

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

// not done
// not tested
func GetAllDM(userID int, cassandra *gocql.Session) string {
	return "no error"
}

// not done
// not tested
// not documented
func GetDMHandler(ctx *gin.Context) {

}

// what do I call these functions??
// not done
// not tested
// func newDMList(userID int, postgres *sql.DB, cassandra *gocql.Session) string {
// 	return "no error"
// }

// not done
// not tested
// not documented
func NewDMListHandler(ctx *gin.Context) {

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
