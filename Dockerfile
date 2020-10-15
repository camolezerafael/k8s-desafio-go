FROM golang:alpine as builder

WORKDIR /src/desafio-go
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...


FROM scratch

COPY --from=builder /bin/desafio-go /desafio-go

ENTRYPOINT ["/desafio-go"]