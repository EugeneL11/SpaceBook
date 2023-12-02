package handlers

import (
	"time"

	"github.com/gocql/gocql"
)

type User struct {
	User_id              int    `json:"id"`
	Full_name            string `json:"full_name"`
	User_name            string `json:"user_name"`
	Email                string `json:"email"`
	Home_planet          string `json:"planet"`
	Profile_picture_path string `json:"profile_picture_path"`
	Admin                bool   `json:"admin"`
	Bio                  string `json:"bio"`
}

type UserPreview struct {
	UserID               int    `json:"user_id"`
	Full_name            string `json:"full_name"`
	User_name            string `json:"user_name"`
	Profile_picture_path string `json:"profile_picture_path"`
}

type Comment struct {
	CommenterID          int    `json:"commenter_id"`
	CommenterProfilePath string `json:"commenter_profile_path"`
	CommenterName        string `json:"commenter_name"`
	Content              string `json:"content"`
	Date                 string `json:"date"`
}

type FullPost struct {
	PostID            gocql.UUID `json:"post_id"`
	AuthorID          int        `json:"author_id"`
	AuthorName        string     `json:"author_name"`
	AuthorProfilePath string     `json:"author_profile_path"`
	Caption           string     `json:"caption"`
	Date              string     `json:"date"`
	Images            []string   `json:"images"`
	Comments          []Comment  `json:"comments"`
	Liked             bool       `json:"liked"`
	NumLikes          int        `json:"num_likes"`
}

type PostPreview struct {
	PostID            gocql.UUID `json:"post_id"`
	AuthorID          int        `json:"author_id"`
	AuthorName        string     `json:"author_name"`
	AuthorProfilePath string     `json:"author_profile_path"`
	Caption           string     `json:"caption"`
	Date              string     `json:"date"`
	Images            []string   `json:"images"`
}

type Message struct {
	SenderID int       `json:"id"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
}

// type FullDM struct {
// }

type DMPreview struct {
	AuthorID          int    `json:"id"`
	AuthorName        string `json:"username"`
	AuthorProfilePath string `json:"profile_path"`
	LastDM            string `json:"recentdm"`
}

type Response struct {
	Error                string `json:"error"`
	User_id              int    `json:"id"`
	User_name            string `json:"username"`
	Admin                bool   `json:"admin"`
	Full_name            string `json:"full_name"`
	Email                string `json:"Email"`
	Home_planet          string `json:"Home_planet"`
	Profile_picture_path string `json:"Profile_picture_path"`
}

// Designed to be used in handlers, pass in corresponding error message
func ErrorUserResponse(errorMsg string) Response {
	return Response{
		Error:                errorMsg,
		User_id:              0,
		User_name:            "null",
		Admin:                false,
		Full_name:            "null",
		Email:                "null",
		Home_planet:          "null",
		Profile_picture_path: "null",
	}
}

// To be used in handlers, pass in User struct containing data to send to front end
func GoodUserResponse(user User) Response {
	return Response{
		Error:                "no error!",
		User_id:              user.User_id,
		User_name:            user.User_name,
		Admin:                user.Admin,
		Full_name:            user.Full_name,
		Email:                user.Email,
		Home_planet:          user.Home_planet,
		Profile_picture_path: user.Profile_picture_path,
	}
}
