FROM golang

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["go", "run", "cmd/app/main.go"]