FROM node:13.7-alpine

WORKDIR /usr/src/app/frontend

COPY package*.json ./

RUN npm install
