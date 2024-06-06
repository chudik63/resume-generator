FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download && \
    go build -o main . && \
    if [ ! -f "./main" ]; then echo "Error: main file not found"; exit 1; fi && \
    rm -f go.mod go.sum

EXPOSE 8080

CMD["./main"]
