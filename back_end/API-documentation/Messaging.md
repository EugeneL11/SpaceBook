## Create a New DM (WIP)

-   Use path (POST): /newdm/{id1}/{id2}
-   Where id1 and id2 are the user IDs of the users in the DM-to-be-created
-   Will return a JSON to indicate success/failure status:

```json
{
    "status": "no error" or "unable to create dm"
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
-   The JSON has the following format:

```json
{
    "status": "no error" or appropriate error (see below)
    "usernames":
}
```
