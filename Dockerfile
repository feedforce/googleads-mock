FROM golang:1.12.9

WORKDIR /app
ADD . /app

RUN go build
CMD [ "/app/googleads-mock" ]
