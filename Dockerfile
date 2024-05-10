## Build
FROM golang:1.21 AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN cd server && CGO_ENABLED=0 GOOS=linux go build -o /bina

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /bina /bina

EXPOSE 5000

USER nonroot:nonroot

CMD ["./bina"]
