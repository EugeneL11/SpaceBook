## To Register an Account:

-   Use path (POST request): /register/{email}/{password}/{fullname}/{username}
-   Will return a JSON object such as the following:

```json
{
    "error": "unable to create account at this time",
    "id": "null",
    "username": "null",
    "admin": "null",
    "full_name": "null",
    "Email": "null",
    "Home_planet": "null",
    "Profile_picture_path": "null"
}
```

-   If no error, then it will say "no error" in "error" field, and user details are given instead of null
-   Client can just notify if there was an error in general, but the error can be:
    -   "unable to create account at this time" (database error on backend, should never happen)
    -   "user name not availible" - user name already exists in database
    -   "email already in use" - email already exists in database
    -   "no error" - successful, all the other fields should be correctly filled out

## To Login:

-   Use path (GET): /login/{username}/{password}
-   Will return a JSON object such as the following

```json
{
    "error": "unable to find User",
    "id": "null",
    "username": "null",
    "admin": "null",
    "full_name": "null",
    "Email": "null",
    "Home_planet": "null",
    "Profile_picture_path": "null"
}
```

-   Either error is "unable to find User" or "no error!"
    The rest of the fields are filled out properly if no error

## Update User Profile (WIP):

-   Use path (PUT): /updateuserprofile/{userID}/{newFullName}/{newPlanet}/{newBio}
-   This allows for updating a given user's fullname, home planet and bio description
-   Some or all fields can be kept the same
-   Will return a JSON to indicate success/failure status:

```json
{
    "status": "unable to connect to db" or "no error"
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

## Change User Profile Pic (WIP)

-   Use route (PUT): /profilepic/{userID}
-   Returns a string to indicate success
    -   "Bad Request", "Internal Server Error" are errors
    -   "File {file_name} uploaded successfully!" if success, where {file_name} is the name of the file uploaded

## Delete User (WIP)

-   Use route (DELETE): /deleteuser/{user_id}
-   Intended to only be used by _admins_
-   Deletes all posts, comments, likes, DMs, friends, and friend requests of the user as well
-   Returns a JSON to indicate success/failure of deletion:

```json
{
    "status": "no error" or "failed to delete %s"
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
