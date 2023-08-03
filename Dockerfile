# Build stage
FROM golang:alpine AS build-env
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY .. .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/app

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build-env /app/app .
EXPOSE 8080
CMD ["./app"]
