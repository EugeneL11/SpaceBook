## To Register an Account:

-   Use path (POST request): /register/{email}/{password}/{fullname}/{username}
-   Will return a JSON object such as the following:

```json
{
    "status": "unable to create account at this time" or "user name not available" or "email already in use" or "no error!",
    "user": null if error occurred, or a JSON (refer to below for User JSON format)
}

// User JSON format
{
    "id": user's ID (int),
    "username": user's username,
    "admin": true if user is an admin, false otherwise (bool),
    "full_name": user's full name,
    "Email": user's email,
    "Home_planet": user's home planet,
    "Profile_picture_path": path to the user's profile picture,
    "bio": user's bio for profile
}

```

-   If no error, then it will say "no error" in "error" field, and user details are given instead of null user
-   Client can just notify if there was an error in general, but the error can be any of the above options in "status"

## To Login:

-   Use path (GET): /login/{username}/{password}
-   Will return a JSON object such as the following

```json
{
    "status": "unable to find User" or "no error!",
    "user": null if error occurred, or a JSON (refer to below for User JSON format)
}

// User JSON format
{
    "id": user's ID (int),
    "username": user's username,
    "admin": true if user is an admin, false otherwise (bool),
    "full_name": user's full name,
    "Email": user's email,
    "Home_planet": user's home planet,
    "Profile_picture_path": path to the user's profile picture,
    "bio": user's bio for profile
}

```

-   Either status is "unable to find User" or "no error!"
-   The rest of the fields are filled out properly if there is no error

## Update User Profile:

-   Use path (PUT): /updateuserprofile/{userID}/{newFullName}/{newPlanet}/{newBio}
-   This allows for updating a given user's fullname, home planet and bio description
-   Some or all fields can be kept the same
-   Will return a JSON to indicate success/failure status:

```json
{
    "status": "unable to parse input", "unable to connect to db", or "no error"
}
```

## Get Users Info:

-   Use path (GET): /getuserinfo/{viewerID}/{viewedID}
-   Provide the user IDs of the person viewing and the person whose profile is being viewed
-   Will return a JSON object of the following format:

```json
{
    "status": "bad request" or "no error",
    "user": null or user's info (another JSON, refer to below) ,
    "friendstatus": null, "already friends", "viewer sent request", "own profile", or "viewed person sent request",
}

// User info JSON (all are strings unless specified)
{
    "id": user's ID (int),
    "full_name": user's full name,
    "user_name": user's username,
    "email": user's email,
    "planet": home planet of the user,
    "profile_picture_path": the file path of the user's pfp,
    "admin": true if user is an admin, or false otherwise (bool),
    "bio": user's bio
}
```

## Change User Profile Pic

-   Use route (PUT): /profilepic/{userID}
-   Returns a string to indicate success
    -   "Bad Request", "Internal Server Error" are errors
    -   "File {file_name} uploaded successfully!" if success, where {file_name} is the name of the file uploaded

## Delete User

-   Use route (DELETE): /deleteuser/{user_id}
-   Intended to only be used by _admins_
-   Deletes all posts, comments, likes, DMs, friends, and friend requests of the user as well
-   Returns a JSON to indicate success/failure of deletion:

```json
{
    "status": "no error", "error parsing input", or "failed to delete %s"
}
```

-   Where %s can be any of the following:
    -   "posts"
    -   "comments"
    -   "likes"
    -   "DMs" (case sensitive)
    -   "friends"
    -   "friend requests"
    -   "user"
