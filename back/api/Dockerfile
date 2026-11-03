FROM golang:1.27-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/api .

FROM alpine:3.23
COPY --from=build /out/api /usr/local/bin/api
WORKDIR /data
EXPOSE 8080
CMD ["api"]