
## simple tests

Register, invalid credentials
Register, empty credentials
Register, valid credentials
Login, incorrect credentials
Login, empty credentials
Login, correct credentials

U1 orbit request U2
U1 orbit request U3
U2 accept orbit U1
U3 deny orbit U1
U1 send dm U2
U2 unorbit U1

U2 orbit request U1
U1 orbit request U2
U2 send dm U1

U1 new post
- fill in details
- cancel 
U1 new post
- fill in details
- post
U2 go to homepage
- U2 expect U1 post (in homepage)
- U2 view U1 profile from post
- U2 expect U1 post (in profile)
U1 delete post
U2 go to homepage
- U2 expect no post (in homepage)
- U2 view U1 profile from post
- U2 expect no post (in profile)

U2 view friends
U2 unorbit U1
U2 send dm U1

U3 delete user U1
U3 send dm U1

U2 log out
U2 log in
U2 view dms


## break website 

U2 find self in search
- disallow the following:
- U2 orbit request self
- U2 view orbit requests
- U2 accept orbit self
- U2 send dm U2

U1 orbit request U3
U3 accept orbit U1
U3 ban U1
U3 send dm U1


## Other

The following is a series of actions to follow with the full website running.
This set of actions serves as a form of system testing to ensure that typical and edge use cases are handled appropriately across the entire stack.

1. Create first account:

-   Click "Sign Up"
-   Click on the name, email, username, password fields, and enter the following:
    -   "John Doe"
    -   "johndoe2@gmail.com"
    -   "johndoe"
    -   "password123"
-   Click "Sign Up" to fill these fields
-   Expected result:
    -   Page changes to show a home screen (blank) and a navbar at the top
    -   Background should be black, with moving white particles

2. Create second account (on another browser window/tab):

-   Repeat 1. but with the following input for sign up:
    -   "Mary Jane"
    -   "maryj7@outlook.com"
        "maryjane"
        -"qwerty7"
-   Click "Sign Up"

3. Log out of one by ...
