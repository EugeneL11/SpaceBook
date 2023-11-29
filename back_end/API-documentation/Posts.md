## To Make a Post (WIP)

-   Use path (POST request): /makepost/{user_id}/{caption}
-   Returns a JSON of following format:

```json
{
    "status": "unable to connect to db" OR "no error"
    "post_id": ID of post (int)
}
```

## Get Posts for Homepage (WIP)

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

## Get Full Details of a Post from IDs (WIP)

-   Use path (GET): /postdetails/{postID}/{userID}
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

-   Use path (POST): /comment/{postID}/{userID}/{commentMsg}
-   Will return a JSON with the following content:

```json
{
    "status": "no error" or "unable to comment"
}
```

## Like Post

-   Use path (POST): /like/{postID}/{userID}
-   Will return a JSON with the following content:

```json
{
    "status": "no error" or "unable to like"
}
```

## Unlike Post

-   Use path (POST): /unlike/{postID}/{userID}
-   Will return a JSON with the following content:

```json
{
    "status": "no error" or "unable to like"
}
```
