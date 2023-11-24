FROM golang:1.21

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

COPY . .
RUN go mod tidy

RUN go get -u github.com/spf13/cobra@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0 && \
    go install github.com/spf13/cobra-cli@latest

RUN apt-get update && apt-get install sqlite3 -y

CMD ["sleep", "infinity"]