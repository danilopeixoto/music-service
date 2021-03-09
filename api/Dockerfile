FROM golang:1.16.0

LABEL description "A music web API."
LABEL version "1.0.0"
LABEL maintainer "Danilo Peixoto"

WORKDIR /usr/bin/api/

COPY . .

RUN go get github.com/swaggo/swag/cmd/swag
RUN swag init && go install .

EXPOSE 8000

CMD ["music"]
