#BUILD STAGE
FROM golang:alpine AS build
MAINTAINER  Bastian Groß <bastian.gross@dertouristik.com>

RUN apk update && apk upgrade && \
    apk add --no-cache bash git
RUN go get -u github.com/golang/dep/cmd/dep
WORKDIR /go/src/app
COPY . .
RUN  dep ensure
RUN go build -o webpagetest-exporter .


FROM   alpine AS production
MAINTAINER  Bastian Groß <bastian.gross@dertouristik.com>
 RUN apk --no-cache add ca-certificates && update-ca-certificates

COPY --from=build /go/src/app/webpagetest-exporter /bin/webpagetest-exporter

 RUN chmod +x /bin/webpagetest-exporter

 EXPOSE      9271
 ENTRYPOINT  [ "/bin/webpagetest-exporter" ]