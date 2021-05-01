FROM golang:1.16.3-alpine3.13
WORKDIR /app
COPY . .
RUN apk update && apk add --update make npm && npm install && make build

FROM alpine:3.11
WORKDIR /app
COPY --from=0 /app/build ./
RUN chmod +x ./app-service
ENV ZEPTO_ENV=production
RUN apk update && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

ENTRYPOINT ["./app-service"]
