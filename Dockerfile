from golang:alpine as builder
workdir /go/src
copy . .
run go build -o api .

from alpine:3.17.1
copy --from=builder /go/src/api ./api
cmd ["./api"]