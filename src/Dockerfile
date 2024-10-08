FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY . .
COPY static/ /static
COPY templates/ /templates

CMD ["air"]
