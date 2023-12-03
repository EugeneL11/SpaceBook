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
Inside of the base level directory (/SpaceBook), enter the following commands:

-   docker compose exec postgres psql -U postgres -d postgres
-   UPDATE USERS SET isadmin = true WHERE user_name = 'profSteve';

If `SELECT * FROM users;` displays `t` under the isadmin column for the row with the user, then the account should now be an admin.

## Test Admin Functionality
