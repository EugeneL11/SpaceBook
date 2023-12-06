## Create a New DM

-   Use path (POST): /newdm/{id1}/{id2}
-   Where id1 and id2 are the user IDs of the users in the DM-to-be-created
-   Will return a JSON to indicate success/failure status:

```json
{
    "status": "no error" or "unable to create dm" or "unable to parse input"
}
```

## Send DM

-   Use path (POST): /senddm/{senderID}/{receiverID}/{message}
-   Will return a JSON to indicate success/failure status:

```json
{
    "status": "no error" or "unable to send dm"
}
```

## Get all DMs Sent by a User

-   Use path (GET): /userdms/{userID}
-   Retrieves usernames, profile pictures, and recent messages of all DMs involving the user
-   The JSON will include a UserDMPreviews JSON, which will be described below as well
-   The JSON has the following format:

```json
{
    "status": "no error" or appropriate error (see below),
    "dmpreviews": null or list of JSONs containing id/username/pfp of users that have DMs involving the given user (refer to below for format of dmpreviews) + most recent message sent
}

UserDMPreviews JSON format:
{
    "id": int containing id,
    "username": string containing username,
    "profile_path": string containing path to user pfp,
    "recentdm": string containing last message, might be null
}
```

## Get List of All Users to Start a DM With

-   Use path (GET): /getallnewdm/{user_id}
-   Retrieve a list of all users the given user (by ID) can start a new DM with
-   This list contains JSONs that include a UserPreview struct (see below)

```json
{
    "status": "no error" or "unable to connect to db 1"
    "newDMRes": null or list of UserPreview JSONs (see below)
}

// UserPreview
{
   {
        "full_name",
        "user_name",
        "profile_picture_path",
        "user_id" (int)
    }
    ... (repeats for number to users that can have a DM started with)

}
```

## Get All Messages Sent Between Two Users

-   Use path (GET): /getmessages/{id1}/{id2}/{subset_size}
-   id1 and id2 are user IDs
-   subset_size is an int representing the number of messages to retrieve
    -   Start with 1 to retrieve most recent messages
    -   Increment this number if you need to load more messages (2, 3, 4, ...)
    -   Refer to JSON to know when there are no messages left
-   Returns a JSON of the following format:

```json
{
    "status": "no error" or "failed to retrieve messages" or "unable to parse input",
    "maxMessages": "true" if there are no more messages, "false" otherwise (bool),
    "messages": null if there was an error, or JSON containing Message structs (refer to below for format)
}

// Message JSON format
{
    {
        "id": userID of the sender of the message (int),
        "message": contents of the message (string),
        "time": timestamp for when the message was sent (string)
    },
    (repeats for number of messages retrieved...)
}
```

-   Note: "moreMessages" defaults to "false" if there was an error
