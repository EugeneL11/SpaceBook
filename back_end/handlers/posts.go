package handlers

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func MakePost(userID int, caption string, cassandra *gocql.Session) (gocql.UUID, string) {
	postID := gocql.TimeUUID()
	insertStmt := cassandra.Query("INSERT INTO Post (postID, caption, authorID) VALUES (?, ?, ?)")

	if err := insertStmt.Bind(postID, caption, userID).Exec(); err != nil {
		return gocql.UUID{}, "unable to connect to db"
	}
	return postID, "no error"
}

func MakePostHandler(ctx *gin.Context) {

}
func GetNewPostsFromUser(userID int, userProfilePath string, userName string, date time.Time, cassandra *gocql.Session, posts []PostPreview) error {
	// Define your PostPreview slice to store the results

	// Define the SELECT statement
	selectStmt := cassandra.Query("SELECT postID, imagePaths,caption, date_posted FROM post WHERE authorID = ? AND date_posted > ?")

	// Bind the parameters and execute the query
	iter := selectStmt.Bind(userID, date).Iter()

	// Iterate through the results and append them to the posts slice
	for {
		var post PostPreview
		var postDate time.Time
		if !iter.Scan(&post.postID, &post.Images, &post.Caption, &postDate) {
			break
		}
		post.Date = postDate.Format(time.RFC3339)
		post.AuthorName = userName
		post.AuthorProfilePath = userProfilePath
		posts = append(posts, post)
	}

	if err := iter.Close(); err != nil {
		return err
	}

	// Now, 'posts' slice contains the new posts from the specified user after the specified date
	return nil
}

// not tested
func GetHomePagePost(userID int, date time.Time, posts []PostPreview, postgres *sql.DB, cassandra *gocql.Session) string {
	stmt, err := postgres.Prepare(`Select user2_id from orbit_buddies where user1_id = $1 union 
	select user2_id from orbit_buddies where user1_id = $1`)
	if err != nil {
		return "unable to connect to db"
	}
	defer stmt.Close()

	row, err := stmt.Query(userID)

	if err != nil {
		return "unable to connect to db"
	}
	for row.Next() {
		var curr_friend int
		row.Scan(&curr_friend)
		stmt, err := postgres.Prepare("select profile_picture_path, user_name from users where user_id = $1")
		if err != nil {
			return "unable to connect to db"
		}
		defer stmt.Close()
		userInfo, err := stmt.Query(curr_friend)
		if err != nil {
			return "unable to connect to db"
		}
		userName, profilePath := "", ""
		if userInfo.Next() {
			userInfo.Scan(&userName, &profilePath)
		} else {
			return "unable to connect to db"
		}
		err = GetNewPostsFromUser(curr_friend, profilePath, userName, date, cassandra, posts)
		if err != nil {
			return "unable to connect to db"
		}
	}
	return "no error"
}

// not done
// not tested
// not documented
func HomepageHandler(ctx *gin.Context) {

}

// not tested
func GetPostDetails(postID int, viewingUser int, post FullPost, cassandra *gocql.Session, postgres *sql.DB) string {
	stmt := cassandra.Query("select * from post where postID = ?")
	iter := stmt.Bind(postID).Iter()

	var likes map[int]struct{}
	var comments map[int]struct{}
	var postDate time.Time
	if iter.Scan(&post.PostID, &post.AuthorID, &post.Images, &post.Caption,
		&likes, &postDate, &comments, &post.NumLikes) {
		_, post.Liked = likes[viewingUser]
		post.NumLikes = len(likes)
		stmt, err := postgres.Prepare("select profile_picture_path, user_name from users where user_id = $1")
		if err != nil {
			return "unable to connect to db"
		}
		defer stmt.Close()
		userInfo, err := stmt.Query(post.AuthorID)
		if err != nil {
			return "unable to connect to db"
		}
		userName, profilePath := "", ""
		if userInfo.Next() {
			userInfo.Scan(&userName, &profilePath)
		} else {
			return "unable to connect to db"
		}
		post.AuthorName = userName
		post.AuthorProfilePath = profilePath
		post.Date = postDate.Format(time.RFC3339)
		for key := range comments {
			stmt := cassandra.Query("Select * from comment where commentID = commentID")
			iter := stmt.Bind(key).Iter()
			var comment Comment
			var commentDate time.Time
			for iter.Scan(nil, &comment.CommenterID, &comment.Content, &commentDate) {
				stmt, err := postgres.Prepare("select profile_picture_path, user_name from users where user_id = $1")
				if err != nil {
					return "unable to connect to db"
				}
				defer stmt.Close()
				userInfo, err := stmt.Query(post.AuthorID)
				if err != nil {
					return "unable to connect to db"
				}
				userName, profilePath := "", ""
				if userInfo.Next() {
					userInfo.Scan(&userName, &profilePath)
				} else {
					return "unable to connect to db"
				}
				comment.CommenterName = userName
				comment.CommenterName = profilePath
				comment.Date = commentDate.Format(time.RFC3339)
				post.Comments = append(post.Comments, comment)
			}
		}
	} else {
		return "unable to connect to db"
	}
	return "no error"
}

// not done
// not tested
// not documented
func PostDetailsHandler(ctx *gin.Context) {

}

// not done
// not tested
func LikePost(postID gocql.UUID, userID int, cassandra *gocql.Session) bool {
	if err := cassandra.Query("UPDATE POST SET likes += ? WHERE postID = ?",
		userID, postID).Exec(); err != nil {
		return false
	}
	return true
}

// not done
// not tested
// not documented
func LikePostHandler(ctx *gin.Context) {

}

// not done
// not tested
func UnlikePost(userID int, postID int) string {
	return "no error"
}

// not done
// not tested
// not documented
func UnlikePostHandler(ctx *gin.Context) {

}
func CommentPost(comment string, userID int, postID gocql.UUID, cassandra *gocql.Session) bool {
	currTime := time.Now()
	commentID := gocql.TimeUUID()

	// Execute the query to insert a comment
	if err := cassandra.Query("INSERT INTO Comment (commentID, commenter, content, time, postID) VALUES (?, ?, ?, ?, ?)",
		commentID, userID, comment, currTime, postID).Exec(); err != nil {
		return false
	}
	if err := cassandra.Query("UPDATE POST SET comments += ? WHERE postID = ?",
		commentID, postID).Exec(); err != nil {
		return false
	}
	// Return true if the comment was successfully inserted
	return true
}

// not done
// not tested
// not documented
func CommentHandler(ctx *gin.Context) {

}

// not done, tested, or documented
func deleteComments(postID int, cassandra *gocql.Session) {

}

// not done, not tested
func DeletePost(postID int, cassandra *gocql.Session) string {
	//cassandra.Query()
	return "no error"
}

// not done
// not tested
// not documented
func DeletePostHandler(ctx *gin.Context) {

}
