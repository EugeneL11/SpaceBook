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
