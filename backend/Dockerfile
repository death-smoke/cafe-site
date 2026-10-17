ARG GO_IMAGE=golang:1.22-alpine
FROM ${GO_IMAGE} AS build
WORKDIR /src
COPY src/ .
RUN go mod tidy && CGO_ENABLED=0 go build -o /out/server .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=build /out/server /app/server
RUN mkdir -p /app/uploads
EXPOSE 8080
CMD ["/app/server"]
