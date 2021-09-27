# syntax=docker/dockerfile:1

FROM golang
# Update, install ping, git
RUN apt-get update && \
    apt-get install -y iputils-ping && \
    apt-get install -y git && \
    git clone https://github.com/Blucas0707/chatroom.git -b develop

WORKDIR /

# This is for documentation purposes only.
# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 1323

# CMD ["/main"]
