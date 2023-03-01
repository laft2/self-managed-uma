FROM golang:1.20.1-bullseye

ENV ROOTDIR=/app
ENV RS_PORT=10001
WORKDIR ${ROOTDIR}

RUN apt update
RUN apt install -y sqlite3

COPY ./main.go ./go.sum ./go.mod ${ROOTDIR}
COPY ./front ${ROOTDIR}/front
COPY ./sql ${ROOTDIR}/sql

RUN go mod download
RUN go build -o rs .

EXPOSE ${RS_PORT}

CMD [ "./rs"]