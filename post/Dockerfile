FROM golang:1.22.5-alpine
WORKDIR /app
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN go build -o app.exe
EXPOSE 8080
CMD ["./app.exe"]