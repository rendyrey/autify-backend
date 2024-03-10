FROM golang:1.21.6-alpine3.19
# ARG MODULE_NAME=foo
WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . ./
RUN go install github.com/spf13/cobra-cli@latest
RUN go build -o /rendy_web_fetcher

EXPOSE 8080
ENTRYPOINT ["/rendy_web_fetcher"]
