FROM golang:1.24.1-alpine
EXPOSE 8080
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app.exe
CMD ["./app.exe"]
