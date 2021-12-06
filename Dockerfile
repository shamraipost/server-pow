FROM golang:1.17.1-alpine3.14

ENV SERVER_PORT 50005
ENV NAME "server"

WORKDIR /usr/src/myapp

EXPOSE 50005
CMD ["go", "run", "main.go"]
