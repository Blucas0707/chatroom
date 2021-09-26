# syntax=docker/dockerfile:1

FROM golang
WORKDIR /chatroom

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 1323

# CMD ["/main"]
