# syntax=docker/dockerfile:1

# Build stage
FROM golang as build
# Update
RUN apt-get update 
# Assign work directory
WORKDIR /chatroom
# Copy root directory to container
COPY . .
# Download go modules
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o /app/server /chatroom/main.go

# Deploy stage
FROM alpine:latest
RUN apk update && apk add bash && apk --no-cache add ca-certificates
COPY --from=build /chatroom/wait-for-it.sh .
COPY --from=build /app/server .
COPY --from=build /chatroom/.env .
COPY --from=build /chatroom/. .
EXPOSE 8080
# CMD ["./server"]
