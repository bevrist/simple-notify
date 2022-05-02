FROM golang:1

ENV CGO_ENABLED=1

# ===========================================
# DEV CACHE
WORKDIR /tmp/tmpgo
  # compile & cache stdlib
RUN find /usr/local/go/src/ -name go.mod -execdir go build \;  \
  # compile & cache modules  \
  && rm -rf *  \
  && go mod init cache-build  \
  && go get github.com/mattn/go-sqlite3  \
  && find /go/pkg/mod/ -name go.mod -execdir go build \;
# ===========================================

WORKDIR /app

# RUN go get github.com/mattn/go-sqlite3

COPY . .

# RUN go build -o /core .
# ENTRYPOINT /core

ENTRYPOINT go run ./
