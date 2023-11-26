# Note: The '.' at the end matters in the build command

# Build the images (only need to do this once unless the image is destroyed/out of date)

cd front_end
docker build -t "react-app" .
cd ../back_end
docker build -t "go-server" .

# Run this command in the root directory (where the docker-compose.yml lives)

# This will start up our React app + Go server + databases

docker-compose up

# Shutdown container (Could use Ctrl+C first?)

docker-compose down

# While compose is running, we can access Postgres via 'psql' with the following

# Note: the password is currently "postgres" for user "postgres" for database "postgres", and we're accessing table 'postgres'...

docker-compose exec postgres psql -U postgres -d postgres

# From here you can use \{command} such as '\dt' to list tables

# From here you can use \{command} such as '\dt' to list tables

For initial setup: Go to the SQL-tables.txt in the back_end/database folder and executes all the CREATE commands in the order they appear

# To access Cassandra Shell (cqlsh), perform a similar command to the previous one

docker-compose exec cassandra cqlsh cassandra

TODO for setup guide => database initialization
