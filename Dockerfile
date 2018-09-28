# FROM golang:1.10 AS builder

# Copy the code from the host and compile it
# WORKDIR $GOPATH/src/github.com/jkobos/smartreader-api
# COPY . ./
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app app/main.go

FROM scratch
COPY  . ./
ENTRYPOINT ["./main"]