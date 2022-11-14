FROM golang:1.19-alpine as builder

RUN apk add --no-cache git make
WORKDIR /app
COPY . .
RUN env CGO_ENABLED=0 go build -a -o dotkatapult -ldflags "-s -w"

FROM scratch
ENV PORT 8080
EXPOSE 8080
WORKDIR /
COPY --from=builder /app/dotkatapult /
CMD ["/dotkatapult"]
