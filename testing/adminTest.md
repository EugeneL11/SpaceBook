NOTE: This requires manually entering the PostgreSQL database and modifying an existing user to be admin (`isAdmin` should be true), since admins do not need to be made regularly and for security reasons
NOTE 2: This set of tests assumes you have the users created from edgeTest.md, and the docker container is running

## Create an Admin User (Instructions)

### Create a Normal Account

Sign up for an account with the following credentials, which will be made into an admin account later

-   Full Name = Steve Sutcliffe
-   Email Address = steves4@outlook.com
-   Username = profSteve
-   Password = steve4

### Make the Account Into Admin Manually

The following assumes you have the docker container running.
Inside of the base level directory (/SpaceBook), enter the following commands in a shell:

-   docker compose exec postgres psql -U postgres -d postgres
    -   Password is "postgres"
-   UPDATE USERS SET isadmin = true WHERE user_name = 'profSteve';

If `SELECT * FROM users;` displays `t` under the isadmin column for the row with the user, then the account should now be an admin.

## Test Admin Functionality

Test following features (initial v1.0 release):
Explain where an admin could take action against a user (ex. on a user's post, or their profile), and what this would look like for the admin's end and the user's end (person being banned, or having post removed)

Go to Sign Up page and create 2 accounts with the following example credentials to be used for testing admin functionalities.
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

-   Steve (admin) clicks on Gene's profile from search, then clicks 'Remove User'
    -   E: Steve's screen goes back to the search page, and Gene's profile is deleted. Steve can no longer search up Gene's account on Spacebook. No immediate change on Gene's end, but he will not be able to perform any more new actions. As soon as the browser is refreshed, he is logged out and can no longer log in using the account 'Gene'.

-   Steve (admin) expands Vic's post from homepage, then clicks 'Remove Post'
    -   E: Steve's screen goes back to the homepage, and Vic's post is deleted. Steve can no longer view that post on Vic's account or the homepage. No immediate change on Vic's end, but if he navigates to his profile, the post will be gone.