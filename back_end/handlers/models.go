package handlers

type User struct {
	User_id              int    `json:"id"`
	Full_name            string `json:"full_name"`
	User_name            string `json:"user_name"`
	Email                string `json:"email"`
	Password             []byte `json:"password"`
	Home_planet          string `json:"planet"`
	Profile_picture_path string `json:"profile_picture_path"`
	Admin                bool   `json:"bool"`
	Bio                  string `json:"bio"`
}

// TODO Consider refactoring naming scheme to make things consistent
// Response is a struct to be returned to the front end by handlers
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

func ErrorResponse() Response {
	return Response{
		Error:                "generic error",
		User_id:              0,
		User_name:            "null",
		Admin:                false,
		Full_name:            "null",
		Email:                "null",
		Home_planet:          "null",
		Profile_picture_path: "null",
	}
}

func GoodResponse(user User) Response {
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
