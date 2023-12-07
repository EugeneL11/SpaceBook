Setup Guide for Members and TAs/Prof:

# Setup Docker to Run Our Application

1. Clone the repo (main branch)
2. Enter the `front_end` folder and run the `npm i` command to ensure all packages are installed and will be copied to Docker later
3. Ensure Docker Desktop is installed, running, and you are logged in (to access Docker Hub images). Then run the following 4 commands:

###### Note: The '.' at the end matters in the build command

###### Build the images (only need to do this once unless the image is destroyed/out of date)

`cd front_end` # If not already in front_end folder then run this
`docker build -t "react-app" .`
`cd ../back_end`
`docker build -t "go-server" .`

4. Return to the root directory (`cd ..`), where the docker-compose.yml is located
5. Run the command `docker-compose up`, and this should start the React client, Go server, PostgreSQL db and Cassandra db, copying/installing packages and other files as needed.

At this point you should be able to access `localhost:3000` in your browser and see the login page for our web application, although we need to setup both databases still!

-   Note: If there is an error regarding a missing import (ex. "Can't resolve 'axios'"), please try running `npm i` inside the `front_end` directory and also opening Docker Desktop, navigating to Containers -> spacebook -> client -> "Exec" tab, and run the following:
    -   `npm i` in the Docker Desktop terminal, while the container is running
    -   If the error persists (the web client should automatically reload after running the command), then try running `npm install {package}`, where package is the missing pkg
    -   If there are still errors, please reach out to us!

# Setup PostgreSQL and CassandraDB Database/Keyspaces & Tables

For this section refer to `/back_end/database/API` directory for the commands needed to initialize Postgres and Cassandra.

1. Enter PostgreSQL shell (psql) inside of the running Docker container. Run the following command while inside of the base-level directory (where docker-compose.yml is):

`docker-compose exec postgres psql -U postgres -d postgres`

###### Note: the password is currently "postgres" for user "postgres" for database "postgres", and we're accessing table 'postgres'...

2. Now that you are inside of the Postgres shell, you can copy-paste the entire contents from `/back_end/database/SQL-tables.txt` into the shell and execute all of them.

    - Should say "CREATE TABLE" 4 times
    - Run `SELECT * FROM users;` and `SELECT * FROM orbit_buddies;` and if a blank table is shown, everything for Postgres should be setup correctly
    - Run `exit` to exit psql

3. After exiting psql, and still in the root directory of SpaceBook, run the following command:

`docker-compose exec cassandra cqlsh cassandra`

4. Now that you are inside of the Cassandra shell, you can similarly copy-paste everything in `/back_end/database/Cassandra-Tables.txt`

    - Similarly can run `SELECT * FROM cassandra.COOKIE;` and it should show an empty table
    - Run `exit` to exit cqlsh

5. Done! You should be able to open `localhost:3000` and use the website now!
