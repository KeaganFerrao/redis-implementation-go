# Redis Server Implementation in Golang

Implemented RESP (Redis Serialization Protocol) with commands such as PING, GET, SET, TTL.

## Setup Instructions

Run `go run main.go` to start the server on port 7379. This only runs on linux based systems (Which supports the EPOLL system call). 

If you have a non-linux system, install WSL (if Windows), run a linux virtual machine or run a linux container in Docker and run the server within it. 

Since the server has implemented Redis protocol, we can use any Redis client (redis-cli) to connect to this server on port 7379.