FROM golang:1.21.6-alpine3.19
WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . ./
RUN go build -o /rendy_web_fetcher

EXPOSE 8080
ENTRYPOINT ["/rendy_web_fetcher"]
