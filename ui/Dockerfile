FROM node:lts-alpine

# Make the 'app' folder the current working directory.
WORKDIR /app

# Copy both 'package.json' and 'package-lock.json'.
COPY package*.json ./

# Install project dependencies.
RUN npm install

# Copy project files and folders to the current working directory (i.e. 'app' folder).
COPY . .

# Build app for production with minification.
RUN npm run build

EXPOSE 80
CMD [ "npm", "start", "--","-p", "80" ]