FROM golang:alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./mp1_backend

FROM alpine
WORKDIR /bin
COPY --from=builder /src/mp1_backend /bin/mp1_backend
EXPOSE 3000
CMD ["/bin/mp1_backend"]
