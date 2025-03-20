FROM golang:alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./mp1_backend

FROM alpine AS runner
WORKDIR /bin
COPY --from=builder /src/mp1_backend /bin/mp1_backend
EXPOSE 3000
RUN apk add curl
HEALTHCHECK --interval=10s --timeout=1s --retries=3 \
  CMD curl -f http://localhost:3000/health || exit 1
CMD ["/bin/mp1_backend", "-db-hostname", "db"]
