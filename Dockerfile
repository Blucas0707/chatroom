# syntax=docker/dockerfile:1

FROM golang
# Update
RUN apt-get update 
WORKDIR /chatroom
COPY . .
RUN go build -o app
EXPOSE 1323

CMD ["./app"]
