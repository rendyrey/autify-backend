# Web Fetcher in Go by Rendy

Web Fetcher is a simple Go CLI program that fetches the HTML content of a webpage given its URL. It utilizes Go's standard library `net/http` package to perform HTTP requests.

## Features

- Fetch HTML content from a given URL.
- Download assets (images, css, javascript)
- Save web page metadata in JSON format

## Requirements

- Go (at least version 1.21)
- Docker if you want run it via docker

## Usage if you want to using go Natively

1. Clone this repository
2. Go to project root directory
3. Run mod download:
```bash
go mod download
```
4. Compile the Go program:
```bash
go build .
```
5. Run the compiled binary with the desired URL as argument:
```bash
./rendy_web_fetcher https://example.com
```
Replace `https://example.com` with the URL of the webpage you want to fetch.

6. The program will fetch and save the HTML content of the webpage if the request is successful in the current active directory.

## Example

```bash
./rendy_web_fetcher https://www.wikipedia.org
```

## Usage if you want to run it via Docker

1. Clone this repository
2. Go to project root directory
3. Run docker-compose up:
```bash
docker-compose up
```
4. Wait docker to build image, after it's done you can run it by:

```bash
docker run rendy_web_fetcher https://example.com
```

Replace `https://example.com` with the URL of the webpage you want to fetch.

5. The program will fetch and save the HTML content of the webpage if the request is successful in the current active directory.

## Example

```bash
docker run rendy_web_fetcher https://www.wikipedia.org
```

## Example commands
1. Fetch multi URLs

```bash
rendy_web_fetcher https://www.wikipedia.org https://rendy.link
```

2. Fetch Web with metadata
```bash
rendy_web_fetcher --metadata https://www.wikipedia.org
```
Or

```bash
rendy_web_fetcher -m https://www.wikipedia.org
```

It will show metadata in console and record it in json format

3. Fetch web and show only metadata
```bash
rendy_web_fetcher --metadata-display https://www.wikipedia.org
```

Or

```bash
rendy_web_fetcher -d https://www.wikipedia.org
```

