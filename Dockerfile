FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/ .
COPY static/ /app/static

CMD ["air"]
