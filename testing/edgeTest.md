-   Note: this set of tests is intended to check for error and edge cases
-   "E:" is short for (E)xpected result:
-   Begin from an empty Postgres/Cassandra database

## 1. Creating Accounts:

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

#### Updating Profile Information:

-   Attempt to update profile with all fields, and a combination of missing fields (incl. empty)
-   Example fields to input while on account Omar:
    -   New Full Name: Khan Omar
    -   New Planet: Venus
    -   New Bio: "Grabbing a Bakechef rn!"
-   Update profile with all three of these fields (when full name, planet, bio are different)
    -   E: All of these fields should update immediately, and the planet should rerender
-   Update profile with 1-2 fields empty
    -   E: Only the filled out fields should cause an update, but the fields not filled out will keep the previous value
    -   E: ex. if bio was "Hello world" and full name was "Omar Khan", then after leaving bio blank and new fullname = "Khan Omar" and updating profile, bio should remain "Hello world" but fullname updates
-   Try inputting relatively long, but reasonable character counts (up to 200 characters) for the bio, and up to 50 characters for the full name

##### Update Profile Picture

-   Go to settings, and click 'Choose File' under 'Change Profile Picture'. Then, choose a valid image file (jpg or png). You can either choose to make changes to the other fields (full name, bio, and planet), or you can just choose to change your profile picture. Either way, click 'Apply Changes'.
    -   E: It takes you back to your profile page. On your profile, and to all other users (on the DM page, their homepage, search users, etc), your profile picture is updated to the new chosen picture.
-   Go to settings, and change bio, full name, or home planet. Then without chaning profile picture, click 'Apply Changes'.
    -   E: It takes you back to your profile page. On your profile and to all other users, your profile picture is the same as it was before previously. No new change to the picture. 
-   Go to settings, and click 'Choose File' under 'Change Profile Picture'. Then, choose a none image file (ex. pdf). Click 'Apply Changes'. 
    -   E: It take you back to your profile page. On your profile and to all other users, your profile picture no longer displays properly and instead it shows the text "My Profile Picture". 

## 2. Finding Users and Orbit Requests:

-   On Desktop, Enter key and clicking search button should be same
-   Preferably case-insensitive so capitals in usernames do not need to be exact match

-   Eugene clicks on Search icon in navbar, types in "gene"
    -   E: No user should show up (should not find self?)
-   Eugene clicks on Search icon in navbar, types in "duppy"
    -   E: After clicking search button or pressing enter, search filters users. When only "d" is entered the user Omar (Username: Duppy) should show up since there are only 3 users present. When "Duppy" or "duppy" is entered Omar should show up
    -   E: User preview should show for Omar, with profile picture (if set, otherwise default) and username showing, at minimum
-   Eugene clicks on "Duppy" user preview
    -   E: Page changes to user profile for Duppy, from the perspective of Eugene. Username, profile picture, bio, friend/orbit status and home planet show at minimum. Posts may also be seen (hidden if public/private profile feature is added).
-   Eugene clicks on a "Request Orbit" button that is visible on Omar's profile
    -   E: Orbit request is nd indicates to check Requests page for further action
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

## 3. DM A Friend

-   Eugene goes to Chat page and starts New Chat with Victor
    -   E: A chat window opens up for Eugene. Victor may notice a chat has already been started with Eugene on his end, and should definitely have an open chat after Eugene sends a message
-   Eugene types a message (empty message, "") and sends it to Victor (either hitting a send button or pressing Enter)
    -   E: No message should be sent at all. No error message and no message sent in chat. User should type a non-empty message before sending.
-   Eugene types the message ("Hey Victor!") and sends it
    -   E: Eugene and Victor should both be able to see the message "Hey Victor!" in near real-time
    -   E: As the sender Eugene should see the message in a different colour and on a different side (right-side) compared to Victor, who received the message (left-side of message screen)
-

## 4. Making Posts

-   Victor clicks on the new post button (found from the homepage)
    -   E: Victor is brought to a page where they can type their caption and insert pictures from their device
-   Victor tries to make posts in the following scenarios:
    1. Victor makes a post with no caption and no image
        - E: Victor remains on the new post screen and is notified by the application that he needs to add either a picture or caption (or both) to post.
    2. Victor makes a post with only an image(s)
        - E: Victor is brought to his profile page and can see his own post listed there. The post contains just an image, but there are still visual elements that allow him to see the number of likes, and also view comments.
    3. Victor makes a post with only a caption
        - E: Similar case, but the post is scaled accordingly to the caption (may be smaller in height than the post in case #2 as a result), and only text is displayed. On his profile page Victor can still see the post, along with likes and comments.
    4. Victor makes a post with a caption and image
        - E: Post is created with image and a caption either above or below it. The rest is similar (likes/comments visible).
    5. Victor makes a post with over X number of images (where X is some "large" number like 10 or 50)
        - E: The application responds with error notifying user that they can only upload at most X images at a time, and no post is created (remain on post create screen).
    6. Victor tries to upload a file of wrong file type as an image (ex. .mp4 or .wav)
        - E: The application responds with error notifying user that they must submit .png/.jpg or any other popular image formats, and no post is created.
    7. Victor tries to upload a file of large size (ex. 100MB+) as an image for post
        - E: Another error case, but user is notified that they must submit a smaller file under X MB in size, for some X (ex. 8MB).

## 5. Commenting

-   Omar wants to be able to comment on one of Victor's posts. He clicks/taps on a button below the post to view the comments and add a comment of his own.
    -   E: Any existing comments are listed, and Omar has a message box and send button to allow him to comment.
-   Omar now tries to press the send comment button without writing a message yet
    -   E: Nothing happens, or optionally the message box is highlighted in a colour to indicate to the user that a message needs to be typed before sending comment.
-   Omar writes a message of moderate length and presses send
    -   E: The message box is emptied and the comment with Omar's message can be seen now. If there is a comment # counter on the post details, it is incremented by 1.

## 6. Liking

-   Eugene wants to like one of Victor's posts. When viewing the post he presses the like button (indicated by some symbol like a heart for example).
    -   E: The symbol should change to indicate a successful like, and the like # counter on the post details should increment by one. Other users like Omar or Victor himself should be able to see this updated counter.
-   Eugene changes his mind and decides to unlike Victor's post. He presses the updated symbol.
    -   E: The symbol should revert back to its original appearance, and the counter should decrement by one. Other users should see this change reflected in the counter as well.
