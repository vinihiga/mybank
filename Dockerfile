FROM golang:latest

WORKDIR /app
COPY . .
RUN go build .
EXPOSE 27000
ENTRYPOINT [ "./mybank", "--container-mode" ]