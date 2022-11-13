FROM node:16-alpine

WORKDIR /web

COPY package.json yarn.lock ./


RUN yarn install


COPY . .

COPY $MOBILE_APP_FILEPATH ./public/client.apk

RUN yarn build

#COPY file config next js

CMD [ "yarn", "start"]