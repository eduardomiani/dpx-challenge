# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy files and build the app
RUN mkdir /dpx-challenge
ADD . /dpx-challenge
WORKDIR  /dpx-challenge
RUN go build

# Executes the app
CMD ./dpx-challenge

# Listen on port 8080
EXPOSE 8080