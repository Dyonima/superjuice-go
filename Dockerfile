FROM golang:1.22.4 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

FROM alpine:3.20
WORKDIR /
COPY --from=build /entrypoint /entrypoint
COPY --from=build /app/assets /assets
EXPOSE 8080
ENTRYPOINT ["/entrypoint"]