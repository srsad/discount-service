FROM node:19.0
WORKDIR /usr/src/app
COPY package*.json ./
EXPOSE 8080
RUN npm install
CMD [ "npm", "run", "dev" ]
