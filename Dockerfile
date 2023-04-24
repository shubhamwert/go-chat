
FROM golang:1.20.3-alpine3.17 AS build-stage
WORKDIR /app
COPY *.go ./
COPY go.mod go.sum ./
COPY ./ChatRoom/ ./ChatRoom/
RUN go mod download
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o serve  .
RUN ls -alh
FROM alpine as releaseVersion
WORKDIR /app
COPY --from=build-stage /app/serve .
RUN ls -alh
EXPOSE 8080
RUN chmod +x serve
CMD ["./serve"]
