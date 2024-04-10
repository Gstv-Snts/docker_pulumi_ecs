FROM golang:1.22 as builder

WORKDIR /

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o main cmd/api/main.go

FROM gcr.io/distroless/static-debian12

COPY --from=builder .env /

COPY --from=builder main /

CMD ["/main"]
