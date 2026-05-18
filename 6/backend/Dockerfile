FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o /server main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=build /server .
EXPOSE 8080
CMD ["./server"]
