FROM        alpine:latest
MAINTAINER  Bastian Groß <bastian.gross@dertouristik.com>

RUN apk --no-cache add ca-certificates && update-ca-certificates

COPY pagespeed_exporter /bin/pagespeed_exporter

RUN chmod +x /bin/webpagetest-exporter

EXPOSE      9271

ENTRYPOINT  [ "/bin/webpagetest-exporter" ]