# Build executable

FROM golang:alpine AS builder
WORKDIR /build
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./cmd/app .
WORKDIR /dist
RUN cp /build/cmd/app ./

# Copy executable to slim image

FROM scratch
COPY --from=builder /dist/cmd/app ./

# Run static executable

ENTRYPOINT ["/cmd/app"]