FROM golang:1.23 AS build

ARG VERSION
WORKDIR /usr/src/app

COPY ./mkdotenv ./mkdotenv

RUN mkdir -p ./bin && ls -l
WORKDIR /usr/src/app/mkdotenv
RUN go build -v -ldflags "-X 'mkdotenv/msg.version=${VERSION}'" -o ../bin/mkdotenv mkdotenv.go

FROM scratch

COPY --from=build --chmod=0755 /usr/src/app/bin/mkdotenv /usr/bin/

CMD ["/usr/bin/mkdotenv"]
