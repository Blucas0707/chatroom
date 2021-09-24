# syntax=docker/dockerfile:1

FROM golang
WORKDIR /chatroom

# Download Go modules
# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY * ./

# Build
RUN go mod init main
RUN go build -o /chatroom

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 1323

CMD ["/main"]

FROM alpine