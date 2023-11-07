# Use an official Node.js runtime as the base image
FROM node:18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the package.json and package-lock.json files
COPY package*.json ./

# Install application dependencies
RUN npm install

# Copy the rest of your frontend code into the container
COPY . /app

# Expose port to access server
EXPOSE 3000

# Specify the command to start your frontend application
CMD ["npm", "start"]