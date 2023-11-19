package handlers

type User struct {
	User_id              int    `json:"id"`
	Full_name            string `json:"full_name"`
	User_name            string `json:"user_name"`
	Email                string `json:"email"`
	Password             int    `json:"password"`
	Home_planet          string `json:"planet"`
	Profile_picture_path string `json:"profile_picture_path"`
	Admin                bool   `json:"admin"`
	Bio                  string `json:"bio"`
}

type API_UserInfo struct {
	User_id              int    `json:"id"`
	Full_name            string `json:"full_name"`
	User_name            string `json:"user_name"`
	Email                string `json:"email"`
	Home_planet          string `json:"planet"`
	Profile_picture_path string `json:"profile_picture_path"`
	Admin                bool   `json:"admin"`
	Bio                  string `json:"bio"`
}
type Comment struct {
}
type FullPost struct {
	AuthorID          int
	AuthorName        string
	AuthorProfilePath string
	Images            []string
	Comments          []Comment
	Liked             bool
	NumLikes          int
}
type PostPreview struct {
	AuthorID          int
	AuthorName        string
	AuthorProfilePath string
	Images            []string
}
type Message struct {
	senderID int
	message  string
}
type FullDM struct {
}
type DMPreview struct {
	AuthorID          int
	AuthorName        string
	AuthorProfilePath string
	LastDM            string
}
