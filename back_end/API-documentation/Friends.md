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

## Send and Accept Friend Request

-   Use path (POST): /sendfriendreq/{sender_user_id}/{receiver_user_id}
-   Note: Use this same route for sending and accepting (accept by flipping the order of the sender/receiver ids)
-   Will return a JSON with following format:

```json
{
    "status": "no error" OR "unable to connect to db" (should not happen)
}
```

## Remove Friend Handler

-   Use path (DELETE): /removefriend/{id1}/{id2}
-   The order of the IDs does not matter
-   Will return a JSON with following format:

```json
{
    "status": "no error" or "unable to connect to db" (should not happen)
}
```

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
        "full_name" (string type),
        "user_name" (string),
        "profile_picture_path" (string),
    },
    ... (repeats for number of friend requests)
}
```

## Reject Friend Request

-   Use path (DELETE): /rejectfriendreq/{rejecter_id}/{rejectee_id}
-   Where rejecter is the person rejecting the rejectee's request
-   Will return a JSON with following format:

```json
{
    "status": "no error" or "unable to connect to db" (should not happen)
}
```

## Search for Users (WIP)

-   Use path (GET): /search/{user_id}/{searchTerm}
-   user_id is used for checking for friends
-   searchTerm is the string the user is typing in the search bar

-   Will return a json object with following format:

```json
{
    "error": "no error" OR "no users found" OR appropriate error
    "userPreviews": null or a JSON containing up to 20 users' fullname + username + profile picture path
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
