FROM golang:alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./mp1_backend

# Go is statically compiled,
# using the scratch image shouldn't cause any issues (in theory)
FROM scratch AS runner
WORKDIR /bin
COPY --from=builder /src/mp1_backend /bin/mp1_backend
EXPOSE 3000
CMD ["/bin/mp1_backend", "-db-hostname", "db"]
