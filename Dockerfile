FROM golang:1.25-alpine AS build

WORKDIR /vigil

COPY . .

RUN go build -o vigil

FROM alpine:latest

WORKDIR /vigil

COPY --from=build /vigil/vigil .

CMD ["./vigil"]
