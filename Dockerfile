FROM golang:1.17-alpine
WORKDIR /app
COPY go.mod go.sum *.go ./
COPY internal/ ./internal
RUN go mod download
RUN env GOOS=linux GOARCH=amd64 go build -o ./server
CMD [ "/app/server"]