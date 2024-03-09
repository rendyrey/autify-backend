FROM golang:1.21.6-alpine3.19
ARG MODULE_NAME=foo
WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . ./
RUN go build -o /fetch_web_rendy

# EXPOSE 8080
ENTRYPOINT ["/fetch_web_rendy"]
