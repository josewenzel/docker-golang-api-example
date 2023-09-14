FROM golang:1.20

WORKDIR /app

COPY go.* ./
RUN go mod download
COPY . ./

RUN go build -C src -o /example-app

EXPOSE 8080
CMD ["/example-app"]


