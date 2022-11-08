FROM node:16-alpine

WORKDIR /web

COPY package.json yarn.lock ./
COPY Makefile Makefile
RUN yarn install

RUN yarn build

RUN yarn dev
#COPY file config next js

CMD [ "make", "yarn", "dev" "build", "install"]