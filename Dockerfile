FROM golang:1.21-alpine as build
WORKDIR /go/src/app
#copy everything
COPY . .
#go mod tidy
RUN go mod tidy
#build the project
RUN go build -o /go/bin/app main.go


FROM alpine as release
#get built binary from build stage
COPY --from=build /go/bin/app /app

ENTRYPOINT /app
