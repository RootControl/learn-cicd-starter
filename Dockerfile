FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY notely .
COPY .env .

RUN chmod +x ./notely

CMD ["./notely"]
