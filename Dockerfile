FROM golang:1.20

WORKDIR /go/src/app

COPY . .
RUN go mod download

# CMD ["go", "run", "./"]

