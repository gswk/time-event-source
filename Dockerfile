FROM golang as builder

ADD . /time-event-source
WORKDIR /time-event-source
RUN go build 

ENTRYPOINT ["./time-event-source"]
CMD ["--interval 5", "--sink http://localhost"]