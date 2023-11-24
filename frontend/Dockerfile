FROM node:16.15.1-alpine

WORKDIR /usr/src/app

# Copy only package.json and package-lock.json first
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application files
COPY . .

# Build the frontend

RUN npm install -g serve
RUN npm run build

# Expose the port that the frontend is listening on
EXPOSE 3000

# Start the frontend application
CMD ["npx", "serve", "-s", "build"]
