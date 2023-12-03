-   Note: this set of tests is intended to check for error and edge cases
-   "E:" is short for (E)xpected result:
-   Begin from an empty Postgres/Cassandra database

## 1. Creating Accounts

### Check for Errors/Required Features

-   Try to log into an account that doesn't exist
    -   E: Indicate invalid credentials/login failed
-   Try to log in with username blank, password blank, and both blank
    -   E: Indicate that field(s) need to be filled, (optional) clear the fields
-   Try to create account with each of the following blank: full name, email addr, username, password. Additionally try with all blank
    -   E: Indicate that field(s) need to be filled, (optional) clear the fields
-   Try to create an account with same email and/or username as another user (Refer to below for example accounts)
    -   E: Indicate email/username already taken (specify which one or both?). Could suggest trying to login instead
-   Should be able to switch between login and signup

### Example Accounts (to be used for later as well)

Go to Sign Up page and create 3 accounts with the following example credentials. It should be possible to create multiple accounts by either logging out after creating an account or by having multiple browser instances on desktop or mobile.

1. Eugene
    - Full Name = Eugene Lee
    - Email Address = eugenel1@gmail.com
    - Username = Gene
    - Password = eugene1
2. Victor
    - Full Name = Victor Han
    - Email Address = victorh2@hotmail.com
    - Username = Vic
    - Password = victor2
3. Omar
    - Full Name = Omar Khan
    - Email Address = omark3@outlook.com
    - Username = Duppy
    - Password = omar3

#### Test for error-handling of unique email/username:

-   Attempt to create account with exact same credentials as Eugene
-   Change only the username to "eugene1" and try again
-   Change only the email address to "gene@gmail.com" and try again
-   Change only the password to "gene" and try again
-   Change only the full name to "Lee Eugene" and try again

    -   E: All five cases above should indicate email/username already taken

-   Change the email AND username
    -   E: Should be able to create an account with same password/full name

## 2. Finding Users and Orbit Requests

-   On Desktop, Enter key and clicking search button should be same
-   Preferably case-insensitive so capitals in usernames do not need to be exact match

-   Eugene clicks on Search icon in navbar, types in "gene"
    -   E: No user should show up (should not find self?)
-   Eugene clicks on Search icon in navbar, types in "duppy"
    -   E: After each character is typed in, search filters users. When only "d" is entered the user Omar (Username: Duppy) should show up since there are only 3 users present. When "Duppy" or "duppy" is entered Omar should show up
    -   E: User preview should show for Omar, with profile picture (if set, otherwise default) and username showing, at minimum
-   Eugene clicks on "Duppy" user preview
    -   E: Page changes to user profile for Duppy, from the perspective of Eugene. Username, profile picture, bio, friend/orbit status and home planet show at minimum. Posts may also be seen (hidden if public/private profile feature is added).
-   Eugene clicks on a "Request Orbit" button that is visible on Omar's profile
    -   E: Orbit request is sent to Omar. Omar receives a notification or red dot as indication on his end.
    -   E: The "Request Orbit" button changes on Eugene's screen to indicate a pending request, which should persist until Omar accepts or rejects the request. Unless a user can "retract" a request, the button should no longer do anything, or maybe display a message.
-   Omar searches Eugene's profile (before checking request)
    -   E: Profile either allows for accepting/rejecting request from Eugene, or shows pending status and indicates to check Requests page for further action
-   Omar checks notification/goes to Requests page
    -   E: Omar sees a user preview for Eugene (incl. username "Gene" and pfp at minimum) and buttons to allow for accepting or rejecting request (ex. green checkmark and red X)
    -   Optionally, there is a confirmation required to reject and/or accept
-   Omar accepts Eugene's request
    -   E: The request/user preview is gone and a message indicates that Omar is now orbiting Eugene (ex. a popup message or text that replaces the user preview)
-   Eugene and Omar check each other's profiles
    -   E: They should see indication that they are orbiting each other, and have the option to unorbit
-   Victor checks Omar's profile
    -   E: There is the option to request orbit Omar
-   Omar unfriends Eugene and then resends request to Eugene. Victor sends Eugene a request as well
    -   E: Eugene should receive notification of receiving request(s). When Eugene checks his Requests page there should be two user previews available; one for Omar and another for Victor
-   Eugene rejects Omar's request and accepts Victor's request
    -   E: After rejecting Omar's request, the page should respond appropriately
        -   a. Move Victor's request up to replace Omar's request OR
        -   b. Keep ordering of requests but replace Omar's request with text to indicate rejection of request
-   Eugene and Omar check each other's profiles
    -   E: They both see an option to orbit request the other
        -   Alternatively, Omar is unable to request again (temporarily?) since he was rejected (to avoid spam)
-   Eugene and Victor check each other's profiles
    -   E: They both see that they are orbiting each other and can unorbit

Note: Need to confirm that other people's profiles are viewed from perspective of current user (ex. Victor should have the option to request either Eugene or Omar, and it should not say pending or orbiting)

## 3. Making Posts

-   Make a post with no images
-
