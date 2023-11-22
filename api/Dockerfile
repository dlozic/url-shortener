FROM golang:1.21

WORKDIR /app
COPY . .
RUN make build
CMD ["./shortener"]