FROM golang:1.9

RUN mkdir /appl

COPY src /appl/src

ENV GOPATH=/appl \
    GOOS=linux \
    GOARCH=amd64 \
    PATH=$PATH:/appl/src/com/privatesquare/go/titanic-app

RUN cd /appl/src/com/privatesquare/go/titanic-app && go build

EXPOSE 8000

CMD ["titanic-app"]