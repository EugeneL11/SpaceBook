package handlers

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

// not done
// not tested
func CreateNewDM(user1 int, user2 int, cassandra *gocql.Session) bool {
	return true
}

// not done
// not tested
// not documented
func CreateNewDMHandler(ctx *gin.Context) {
	user1, err1 := strconv.Atoi(ctx.Param("user1"))
	user2, err2 := strconv.Atoi(ctx.Param("user2"))
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	if err2 != nil || err1 != nil {
		// send error
		return
	}
	result := CreateNewDM(user1, user2, cassandra)
	if !result {
		// send error
		return
	}
	// send success
}

// not done
// not tested
func SendDM(senderID int, receiver_id int, message string, cassandra *gocql.Session) bool {

	// get DM between users
	// catch error if DM doesn't exist
	// get most recent message subset
	// catch error
	// see if message subset is full
	// catch error

	// if message subset is full
	// create new message subset
	// catch error
	// retrieve this new message subset
	// catch error

	// send message within subset
	// catch error

	return true
}

// not done
// not tested
// not documented
func SendDMHandler(ctx *gin.Context) {
	message := ctx.Param("message")
	sender, err := strconv.Atoi(ctx.Param("sender"))
	reciever, err2 := strconv.Atoi(ctx.Param("reciever"))
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)

	if err != nil || err2 != nil {
		return
	}
	result := SendDM(sender, reciever, message, cassandra)
	if !result {
		// send error
		return
	}
	// send success
}

// not done
// not tested
func GetAllDM(userID int, cassandra *gocql.Session) (bool, []DMPreview) {
	var allDM []DMPreview

	return true, allDM
}

// not done
// not tested
// not documented
func GetDMHandler(ctx *gin.Context) {
	user, err1 := strconv.Atoi(ctx.Param("userID"))
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	if err1 != nil {
		// send error
		return
	}
	success, result := GetAllDM(user, cassandra)
	if !success {
		//send error
		return
	}
}

// what do I call these functions??
// not done
// not tested
func newDMList(userID int, postgres *sql.DB, cassandra *gocql.Session) (bool, []UserPreview) {
	var users []UserPreview
	return true, users
}

// not done
// not tested
// not documented
func NewDMListHandler(ctx *gin.Context) {
	user, err1 := strconv.Atoi(ctx.Param("userID"))
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	postgres := ctx.MustGet("postgres").(*sql.DB)
	if err1 != nil {
		// send error
		return
	}
	success, result := newDMList(user, postgres, cassandra)
	if !success {
		//send error
		return
	}
}

// not done
// not tested
func GetMessages(user1 int, user2 int, subsetSize int, cassandra *gocql.Session) (bool, []Message) {
	var messages []Message
	return true, messages
}

// not done
// not tested
// not documented
func GetMessagesHandler(ctx *gin.Context) {
	user1, err1 := strconv.Atoi(ctx.Param("user1"))
	user2, err2 := strconv.Atoi(ctx.Param("user2"))
	subsetSize, err3 := strconv.Atoi(ctx.Param("user2"))
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	if err2 != nil || err1 != nil || err3 != nil {
		// send error
		return
	}
	success, result := GetMessages(user1, user2, subsetSize, cassandra)
	if !success {
		// send error
		return
	}
}
