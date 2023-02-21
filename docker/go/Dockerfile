FROM golang:1.20.1-alpine AS dev

ENV ROOT /nomad
ENV CGO_ENABLED 0

WORKDIR ${ROOT}
RUN apk add --update --no-cache git

#COPY go.mod go.sum ./
#RUN go mod download
#COPY src/ ./src

EXPOSE 8080
CMD ["go", "run", "src/main.go"]

FROM golang:1.20.1-alpine AS build

ENV ROOT /nomad
ENV CGO_ENABLED 0

WORKDIR ${ROOT}
RUN apk add --update --no-cache git

COPY go.mod go.sum ./
RUN go mod download
COPY src/ ./src

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin src/main.go

FROM scratch AS prod

ENV ROOT /nomad/src
COPY --from=build ${ROOT}/bin ${ROOT}

EXPOSE 8080
CMD ["/nomad/src/bin"]
