
FROM golang:1.20.3-alpine3.17 AS build-stage
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o serve  .
RUN ls -alh
FROM scratch
WORKDIR /app
COPY --from=build-stage /app/serve .
COPY --from=build-stage /app/config.yaml .

EXPOSE 8080
CMD ["./serve"]
