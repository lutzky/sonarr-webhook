FROM golang:1.12 AS build-env
ADD . /src
RUN cd /src && \
	CGO_ENABLED=0 GOOS=linux go build -a \
	-ldflags '-w -extldflags "-static"' -o sonarr-webhook

FROM scratch
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /src/sonarr-webhook /
COPY --from=build-env /src/template.txt /
ENTRYPOINT ["/sonarr-webhook"]
