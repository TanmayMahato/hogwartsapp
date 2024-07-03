
FROM golang:latest


RUN mkdir app

COPY . /backend3/


WORKDIR /backend3/

RUN go build -o main .

EXPOSE 8003


CMD [ "/backend1/main" ]
