FROM golang:1.12.5

RUN mkdir /appl

ENV GOPATH=/appl \
    GOOS=linux \
    GOARCH=amd64 \
    PATH=$PATH:/appl/src/com/privatesquare/go/titanic-app

RUN apt-get update \
    && apt-get install jq -y

RUN download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | jq -r '.assets[] | \
    select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url') \
    && curl -o /usr/local/bin/swagger -L'#' "$download_url" \
    && chmod +x /usr/local/bin/swagger

COPY src /appl/src

WORKDIR /appl/src/com/privatesquare/go/titanic-app

RUN go get ./... && go build && swagger generate spec -o ./swaggerui/swagger.json

EXPOSE 8000

CMD ["titanic-app"]