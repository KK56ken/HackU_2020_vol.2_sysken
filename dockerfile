FROM node:14.8.0-alpine3.10

WORKDIR /app

RUN apk update && \
    npm install -g \
    firebase-tools && \
    npm install --save firebase && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app/raise_todo

CMD [ "npm", "run" ,"serve" ]

# FROM node:10.15.1-alpine as builder
# WORKDIR /app

# RUN apk update && \
#     apk add git

# COPY ./package.json /app

# RUN yarn

# COPY . /app
# RUN yarn build

# FROM node:10.15.1-alpine
# WORKDIR /app
# COPY --from=builder /app /app
# CMD ["yarn", "start"]
