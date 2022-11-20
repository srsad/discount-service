FROM node:19.0
WORKDIR /usr/src/app
COPY package*.json ./
EXPOSE 5173
RUN npm install
CMD [ "npm", "run", "dev" ]
