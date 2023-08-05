FROM alpine:latest AS base
WORKDIR /app/

FROM golang:latest AS build
WORKDIR /src/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /build/server main.go
RUN CGO_ENABLED=0 go build -o /build/scraper cmd/scraper/main.go

FROM base AS final
WORKDIR /app/
COPY --from=build /build/scraper /app/scraper
RUN echo "* * * * * /app/scraper --out /var/6obcy-monitor/6obcy.db" > /etc/periodic/15min//scraper
COPY --from=build /build/server /app/server
CMD [ "/app/server", "&", "crond" ]