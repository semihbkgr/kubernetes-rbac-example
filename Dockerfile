FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o app

FROM alpine
WORKDIR /bin
COPY --from=build-env /src/app /bin/
ENTRYPOINT ./app
