FROM golang AS build

WORKDIR /usr/src/app
COPY ./src/mkdotenv.go .

RUN mkdir -p ./bin && ls -l
RUN go build -v -o ./bin/mkdotenv mkdotenv.go

FROM scratch

COPY --from=build --chmod=0755 /usr/src/app/bin/mkdotenv /usr/bin/

CMD ["/usr/bin/mkdotenv"]
