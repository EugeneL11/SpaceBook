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
	Bio 				 string `json:"bio"`
}


type API_UserInfo struct {
	User_id 			 int	`json:"id"`
	Full_name            string `json:"full_name"`
	User_name            string `json:"user_name"`
	Email                string `json:"email"`
	Home_planet          string `json:"planet"`
	Profile_picture_path string `json:"profile_picture_path"`
	Admin                bool   `json:"admin"`
	Bio 				 string `json:"bio"`
}