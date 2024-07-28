FROM node:18-alpine

WORKDIR /frontend

COPY public/ /frontend/public
COPY src/ /frontend/src
COPY package.json /frontend/

RUN npm install 

CMD ["npm", "start"]