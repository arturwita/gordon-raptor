FROM golang:1.25-alpine

RUN apk add --no-cache git curl

RUN go install github.com/air-verse/air@v1.63.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p ./bin


CMD ["air", "-c", ".air.toml"]
