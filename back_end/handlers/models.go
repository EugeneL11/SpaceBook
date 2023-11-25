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
	CommenterID          int
	CommenterProfilePath string
	CommenterName        string
	Content              string
	Date                 string
}
type FullPost struct {
	PostID            int
	AuthorID          int
	AuthorName        string
	AuthorProfilePath string
	Caption           string
	Date              string
	Images            []string
	Comments          []Comment
	Liked             bool
	NumLikes          int
}
type PostPreview struct {
	postID            int
	AuthorID          int
	AuthorName        string
	AuthorProfilePath string
	Caption           string
	Date              string
	Images            []string
}

// type Message struct {
// 	senderID int
// 	message  string
// }

// type FullDM struct {
// }

type DMPreview struct {
	AuthorID          int
	AuthorName        string
	AuthorProfilePath string
	LastDM            string
}

// ErrorResponse is intended to be used in handlers with gin.H to return error client response
func ErrorResponse(errorMsg string) map[string]any {
	return map[string]any{
		"error":                errorMsg,
		"id":                   0,
		"username":             "null",
		"admin":                false,
		"full_name":            "null",
		"Email":                "null",
		"Home_planet":          "null",
		"Profile_picture_path": "null",
	}
}

// GoodResponse is intended to be used in handlers with gin.H to return client response
func GoodResponse(user User) map[string]any {
	return map[string]any{
		"error":                "no error!",
		"id":                   user.User_id,
		"username":             user.User_name,
		"admin":                user.Admin,
		"full_name":            user.Full_name,
		"Email":                user.Email,
		"Home_planet":          user.Home_planet,
		"Profile_picture_path": user.Profile_picture_path,
	}
}
