FROM node:14.8.0-alpine3.10

WORKDIR /app

RUN apk update && \
    npm install -g \
    firebase-tools && \
    npm install --save firebase && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app/hacku

CMD [ "npm", "run" ,"serve" ]
