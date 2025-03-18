FROM golang:1.23
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /mp1_backend

EXPOSE 3000

CMD ["/mp1_backend"]
