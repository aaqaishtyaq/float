# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.17 AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o /float

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /float /float

EXPOSE 8080

ENTRYPOINT ["./float"]
