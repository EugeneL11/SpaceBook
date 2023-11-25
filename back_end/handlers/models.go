package handlers

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
	Full_name            string `json:"full_name"`
	User_name            string `json:"user_name"`
	Profile_picture_path string `json:"profile_picture_path"`
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
