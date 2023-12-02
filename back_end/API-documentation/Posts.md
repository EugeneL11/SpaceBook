## To Make a Post

-   Use path (POST request): /makepost/{user_id}/{caption}
-   **_Note: Need to make the post and get the ID before adding pictures!_**
    -   Refer to Posts.md "Upload Post Image" (right below), will need to call that route multiple times for multiple pictures
-   Returns a JSON of following format:

```json
{
    "status": "unable to connect to db" OR "no error"
    "post_id": ID of post (int)
}
```

## Upload Post Image

-   Use route (POST): /uploadpostimage/{postID}
-   Returns a string indicating "Bad Request" or "File {name} uploaded successfully!"
    -   Access as a JSON and use .data
    -   Refer to ImageTest.jsx for example
-   **_NOTE: Requires a postID, which can be retrieved from making a post (refer to above)_**

## Get Posts for Homepage

-   Use path (GET): /homepageposts/{user_id}
-   Returns a JSON of following format:

```json
{
    "status": "no error" or "unable to connect to db"
    "posts": Either null or a JSON containing posts (see below JSON)
}
// PostPreviews JSON:
{
    {
        "post_id" (of type string),
        "author_id" (int),
        "author_name" (string),
        "author_profile_path" (string),
        "caption" (string),
        "date" (string),
        "images" (string[]),
    },
    ... (repeats for number of posts)
}
```

## Get Full Details of a Post from IDs

-   Use path (GET): /postdetails/{postID}/{userID}
-   Retrieve full details (refer to JSON) of a post from a viewing user's perspective (give the IDs for both)
    -   postID is the ID of the post being looked at
    -   userID is the ID of the user who is looking at the post (used to determine whether the user has liked the post or not)
-   Returns a JSON of following format:

```json
{
        "post_id" (of type string),
        "author_id" (int),
        "author_name" (string),
        "author_profile_path" (string),
        "caption" (string),
        "date" (string),
        "images" (string[]),
        "comments" (nested JSON, see below),
        "liked" (bool),
        "num_likes" (int),
}

// Comments JSON
{
    {
        "commenter_id" (int),
        "commenter_profile_path" (string),
        "commenter_name" (string),
        "content" (string),
        "date" (string),
    },
    ... (repeating for every comment on that post)
}
```

## Comment on Post

-   Use path (POST): /makecomment/{postID}/{userID}/{commentMsg}
-   Will return a JSON with the following content:

```json
{
    "status": "no error" or "unable to comment"
}
```

## Like Post

-   Use path (POST): /likepost/{postID}/{userID}
-   Will return a JSON with the following content:

```json
{
    "status": "no error" or "unable to like"
}
```

## Unlike Post (not done)

-   Use path (POST): /unlikepost/{postID}/{userID}
-   Will return a JSON with the following content:

```json
{
    "status": "no error" or "unable to like"
}
```

## Delete a Post by postID (not done)

-   Use path (DELETE): /deletepost/{postID}
-   Returns a JSON to indicate success/failure of deletion:

```json
{
    "status": "no error" or "failed to delete post"
}
```
