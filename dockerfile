FROM golang:bullseye

LABEL Author="Nigel Tatschner (ntatschner@gmail.com)"

ENV LISTENONPORT "8080"
ENV APP_HOME /thsapi/
ENV GIN_MODE=release

RUN mkdir -p "$APP_HOME"

COPY src/* ${APP_HOME}

WORKDIR "$APP_HOME"

RUN go build -o main main.go

EXPOSE ${LISTENONPORT}

CMD ["./main"]