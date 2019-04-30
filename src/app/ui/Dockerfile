FROM node:lts-alpine

# Make the 'app' folder the current working directory.
WORKDIR /app

# Copy both 'package.json' and 'package-lock.json'.
COPY package*.json ./

# Copy project files and folders to the current working directory (i.e. 'app' folder).
COPY . .

# Install project dependencies.
RUN npm install

# Build app for production with minification.
RUN npm run build

# Set app serving to permissive / assigned.
ENV NUXT_HOST=0.0.0.0

# Set app port.
ENV NUXT_PORT=80

EXPOSE 80
CMD [ "npm", "start" ]