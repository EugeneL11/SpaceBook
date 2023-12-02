package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
)

func MakePost(userID int, caption string, cassandra *gocql.Session) (gocql.UUID, string) {
	postID := gocql.TimeUUID()
	time := time.Now()
	insertStmt := cassandra.Query(`INSERT INTO Post (postID, caption, authorID, comments, date_posted, imagepaths, 
		likes) VALUES (?, ?, ?, {}, ?, [], {} )`)

	if err := insertStmt.Bind(postID, caption, userID, time).Exec(); err != nil {
		return gocql.UUID{}, "unable to connect to db"
	}
	return postID, "no error"
}

func MakePostHandler(ctx *gin.Context) {
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	userID, err1 := strconv.Atoi(ctx.Param("user_id"))
	if err1 != nil {
		log.Fatal(err1)
	}
	caption := ctx.Param("caption")
	postID, status := MakePost(userID, caption, cassandra)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  status,
		"post_id": postID,
	})
}

// Used by homepage handler (below)
func GetNewPostsFromUser(userID int, userProfilePath string, userName string, date time.Time, cassandra *gocql.Session) ([]PostPreview, error) {
	// Define the SELECT statement
	selectStmt := cassandra.Query("SELECT postID, imagePaths,caption, date_posted FROM post WHERE authorID = ? AND date_posted > ? ALLOW FILTERING")

	// Bind the parameters and execute the query
	iter := selectStmt.Bind(userID, date).Iter()
	var posts []PostPreview

	// Iterate through the results and append them to the posts slice
	for {
		var post PostPreview
		var postDate time.Time
		if !iter.Scan(&post.PostID, &post.Images, &post.Caption, &postDate) {
			break
		}
		post.Date = postDate.Format(time.RFC3339)
		post.AuthorName = userName
		post.AuthorProfilePath = userProfilePath
		post.AuthorID = userID
		posts = append(posts, post)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	// Now, 'posts' slice contains the new posts from the specified user after the specified date
	return posts, nil
}

func GetHomePagePost(userID int, date time.Time, postgres *sql.DB, cassandra *gocql.Session) ([]PostPreview, string) {
	stmt, err := postgres.Prepare(`Select user2_id from orbit_buddies where user1_id = $1 union 
	select user1_id from orbit_buddies where user2_id = $1`)
	if err != nil {
		return nil, "unable to connect to db"
	}
	defer stmt.Close()

	row, err := stmt.Query(userID)

	if err != nil {
		return nil, "unable to connect to db"
	}

	var posts []PostPreview

	for row.Next() {
		var curr_friend int
		row.Scan(&curr_friend)
		stmt, err := postgres.Prepare("select profile_picture_path, user_name from users where user_id = $1")
		if err != nil {
			return nil, "unable to connect to db"
		}
		defer stmt.Close()
		userInfo, err := stmt.Query(curr_friend)
		if err != nil {
			return nil, "unable to connect to db"
		}
		userName, profilePath := "", "" // TODO
		if userInfo.Next() {
			userInfo.Scan(&profilePath, &userName)
		} else {
			return nil, "unable to connect to db"
		}
		tempPost, err2 := GetNewPostsFromUser(curr_friend, profilePath, userName, date, cassandra)
		if err2 != nil {
			return nil, "unable to connect to db"
		}
		posts = append(posts, tempPost...)
	}
	return posts, "no error"
}

func HomepageHandler(ctx *gin.Context) {
	postgres := ctx.MustGet("postgres").(*sql.DB)
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	userID, err1 := strconv.Atoi(ctx.Param("user_id"))
	if err1 != nil {
		log.Fatal(err1)
	}
	date := time.Now()
	minTime := date.AddDate(0, 0, -7) // Current time - 7 days to get recent posts
	posts, status := GetHomePagePost(userID, minTime, postgres, cassandra)
	ctx.JSON(http.StatusOK, gin.H{
		"status": status,
		"posts":  posts,
	})
}

func GetPostDetails(postID gocql.UUID, viewingUser int, post *FullPost, cassandra *gocql.Session, postgres *sql.DB) string {
	stmt := cassandra.Query("select authorID, caption, imagepaths, date_posted, comments, likes from post where postID = ?")
	iter := stmt.Bind(postID).Iter()
	post.PostID = postID
	var likes []int
	var comments []int
	var postDate time.Time
	if iter.Scan(&post.AuthorID, &post.Caption, &post.Images,
		&postDate, &comments, &likes) {
		for _, e := range likes {
			if e == viewingUser {
				post.Liked = true
			}
		}
		post.NumLikes = len(likes)
		stmt, err := postgres.Prepare("select user_name, profile_picture_path from users where user_id = $1")
		if err != nil {
			return "unable to connect to db 1"
		}
		defer stmt.Close()
		userInfo, err := stmt.Query(post.AuthorID)
		if err != nil {
			return "unable to connect to db 2"
		}
		userName, profilePath := "", ""
		if userInfo.Next() {
			userInfo.Scan(&userName, &profilePath)
		} else {
			return "unable to connect to db 3"
		}
		post.AuthorName = userName
		post.AuthorProfilePath = profilePath
		post.Date = postDate.Format(time.RFC3339)
		for _, key := range comments {
			stmt := cassandra.Query("Select * from comment where commentID = ?")
			iter := stmt.Bind(key).Iter()
			var comment Comment
			var commentDate time.Time
			for iter.Scan(nil, &comment.CommenterID, &comment.Content, &commentDate) {
				stmt, err := postgres.Prepare("select profile_picture_path, user_name from users where user_id = $1")
				if err != nil {
					return "unable to connect to db 4"
				}
				defer stmt.Close()
				userInfo, err := stmt.Query(post.AuthorID)
				if err != nil {
					return "unable to connect to db 5"
				}
				userName, profilePath := "", ""
				if userInfo.Next() {
					userInfo.Scan(&userName, &profilePath)
				} else {
					return "unable to connect to db 6"
				}
				comment.CommenterName = userName
				comment.CommenterName = profilePath
				comment.Date = commentDate.Format(time.RFC3339)
				post.Comments = append(post.Comments, comment)
			}
		}
	} else {
		return "unable to connect to db 7"
	}
	return "no error"
}

func PostDetailsHandler(ctx *gin.Context) {
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	postgres := ctx.MustGet("postgres").(*sql.DB)
	postID, err := gocql.ParseUUID(ctx.Param("postID"))
	if err != nil {
		log.Panic(err)
	}
	viewingUser, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		log.Panic(err)
	}
	var post FullPost // Post details to be filled in GetPostDetails
	status := GetPostDetails(postID, viewingUser, &post, cassandra, postgres)
	ctx.JSON(http.StatusOK, gin.H{
		"status": status,
		"post":   post,
	})
}

func LikePost(postID gocql.UUID, userID int, cassandra *gocql.Session) bool {
	if err := cassandra.Query("UPDATE POST SET likes = likes + ? WHERE postID = ?",
		[]int{userID}, postID).Exec(); err != nil {
		fmt.Println("Error updating POST table:", err)
		return false
	}
	return true
}

func LikePostHandler(ctx *gin.Context) {
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		return
	}
	postID, err := gocql.ParseUUID((ctx.Param("postID")))
	if err != nil {
		return
	}
	res := LikePost(postID, userID, cassandra)
	var status string
	if res {
		status = "no error"
	} else {
		status = "unable to like"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}

func CommentPost(comment string, userID int, postID gocql.UUID, cassandra *gocql.Session) bool {
	currTime := time.Now()
	commentID := gocql.TimeUUID()

	// Execute the query to insert a comment
	if err := cassandra.Query("INSERT INTO Comment (commentID, commenter, content, time, postID) VALUES (?, ?, ?, ?, ?)",
		commentID, userID, comment, currTime, postID).Exec(); err != nil {
		fmt.Println("Error inserting comment:", err)
		return false
	}

	// Update the POST table to add the commentID to the comments set
	if err := cassandra.Query("UPDATE POST SET comments = comments + ? WHERE postID = ?",
		[]gocql.UUID{commentID}, postID).Exec(); err != nil {
		fmt.Println("Error updating POST table:", err)
		return false
	}

	// Return true if the comment was successfully inserted and the counters updated
	return true
}

func CommentHandler(ctx *gin.Context) {
	cassandra := ctx.MustGet("cassandra").(*gocql.Session)
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		return
	}
	postID, err := gocql.ParseUUID((ctx.Param("postID")))
	if err != nil {
		return
	}
	comment := ctx.Param("comment")
	res := CommentPost(comment, userID, postID, cassandra)
	var status string
	if res {
		status = "no error"
	} else {
		status = "unable to comment"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}
