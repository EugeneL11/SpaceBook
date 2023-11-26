package handlers

import (
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

// not done
// not tested
func sendDM(senderID int, receiver_id int, message string, cassandra *gocql.Session) string {
	
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
	
	return "no error"
}

// not done
// not tested
// not documented
func SendDMHandler(ctx *gin.Context) {
	
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
