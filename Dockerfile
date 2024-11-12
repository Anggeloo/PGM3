FROM golang:1.23-alpine AS build_stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o PGM3 ./PGM3.go

FROM alpine:latest AS runtime_stage
COPY --from=build_stage /app/PGM3 /PGM3
EXPOSE 8080
CMD ["/PGM3"]
