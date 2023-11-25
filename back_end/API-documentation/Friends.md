## Get Friends (WIP)

-   Use path (GET request): /getfriends/{user_id}
-   Will return a JSON object with following format:

```json
{
    "error": "no error" OR appropriate error
    "users": null OR a JSON containing multiple users' data
}
```

-   "error" can contain an error from Go, so consider just checking whether or not (error = "no error") unless custom error messages are provided at a later point

## Get Friend Requests (WIP)

-   Use path (GET): /getfriendreqs/{user_id}
-   Will return a JSON with following format:

```json
{
    "status": "no requests", "pending request", or appropriate error (only "unable to connect to db")
    "requests": JSON containing UserPreview for each user which made a request (see below)
}

// UserPreviews
{
    {
        "full_name",
        "user_name",
        "profile_picture_path",
    }
    ... (repeats for number of friend requests)
}
```

## Reject Friend Request

## Search for Users (WIP)

-   Use path (GET): /search/{user_id}/{searchTerm}
-   user_id is used for checking for friends
-   searchTerm is the string the user is typing in the search bar

-   Will return a json object with following format:

```json
{
    "error": "no error" OR appropriate error
    "userPreviews": null or a JSON containing up to 20 users' fullname + username + profile picture path (see below JSON)
}

// UserPreviews JSON:
{
    {
        "full_name",
        "user_name",
        "profile_picture_path",
    }
    ... (repeats max 20 total times)
}
```
