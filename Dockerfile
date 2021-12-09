FROM golang:1.12-alpine
RUN apk add git
# Set the Current Working Directory inside the container
WORKDIR /app/beego_training
# We want to populate the module cache based on the go.{mod,sum} files.
COPY . .
RUN go mod download
# Build the Go app
RUN go build -o ./out/beego_training .
# This container exposes port 8080 to the outside world
EXPOSE 8080
# Run the binary program produced by `go install`
CMD ["./out/beego_training"]
