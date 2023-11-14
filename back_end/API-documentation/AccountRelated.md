To Register an Account:
use path:
/register/email/password/fullname/username
will return a json object 
{
	"error":                "unable to create account at this time",
	"id":                   "null",
	"username":             "null",
	"admin":                "null",
	"full_name":            "null",
	"Email":                "null",
	"Home_planet":          "null",
	"Profile_picture_path": "null",
}
"error" can be : 
unable to create account at this time - database error on backend, should never happen
user name not availible - user name already exists in database
email already in use - email already exists in database
no error - successful, all the other fields should be correctly filled out

To login:
use path:
/login/username/password
will return a json object 
{
	"error":                "unable to find User",
	"id":                   "null",
	"username":             "null",
	"admin":                "null",
	"full_name":            "null",
	"Email":                "null",
	"Home_planet":          "null",
	"Profile_picture_path": "null",
}
either error is "unable to find User" or "no error!"
The rest of the fields are filled out properly if no error 