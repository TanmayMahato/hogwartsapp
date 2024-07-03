
FROM golang:latest


RUN mkdir app

COPY . /backend4/


WORKDIR /backend4/

RUN go build -o main .

EXPOSE 8004


CMD [ "/backend1/main" ]
