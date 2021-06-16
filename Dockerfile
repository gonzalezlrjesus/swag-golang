FROM golang:alpine AS build

WORKDIR /src

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/API .

FROM scratch
COPY --from=build /src/bin/API .
EXPOSE 3000
CMD [ "./API" ]