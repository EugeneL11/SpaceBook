## To Make a Post (WIP)

-   Use path (POST request): /makepost/{user_id}/{caption}
-   Returns a JSON of following format

```json
{
    "status": "unable to connect to db" OR "no error"
    "post_id": ID of post (int)
}
```
