FROM golang:1.20 as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o /main

FROM gcr.io/distroless/static-debian11

COPY --from=build /main /

CMD ["/main"]
