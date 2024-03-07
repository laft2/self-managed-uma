FROM golang:1.20.1-bullseye

ENV ROOTDIR=/app
WORKDIR ${ROOTDIR}

RUN apt update
RUN apt install -y sqlite3

COPY ./main.go ./go.sum ./go.mod ${ROOTDIR}
COPY ./front ${ROOTDIR}/front
COPY ./sql ${ROOTDIR}/sql
COPY ./cmd ${ROOTDIR}/cmd

RUN go mod download
RUN go build -o rs .

CMD ["go", "run", "${ROOTDIR}/cmd/launch_server/main.go"]