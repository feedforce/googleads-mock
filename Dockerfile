FROM golang:1.12.9 AS build
WORKDIR /work
COPY . /work
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

FROM scratch
WORKDIR /app
COPY --from=build /work/googleads-mock /work/search_term_view.json /work/search_term_view_stream.json ./
ENTRYPOINT ["/app/googleads-mock"]
