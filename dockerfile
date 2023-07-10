FROM golang:bullseye

ARG LISTENONIP "localhost"
ARG LISTENONPORT "8080"
ENV APP_HOME /thsapi/

RUN mkdir -p "$APP_HOME"

COPY src/* ${APP_HOME}

WORKDIR "$APP_HOME"

RUN go get -d ./...

EXPOSE ${LISTENONPORT}

CMD ["go", "run", "."]