FROM golang:bullseye

COPY src/* /thsapi/

ENV LISTENONIP "localhost"
ENV LISTENONPOST "8080"