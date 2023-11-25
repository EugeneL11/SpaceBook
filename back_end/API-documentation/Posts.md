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

-   Use path (GET): /gethomepageposts/{user_id}
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
