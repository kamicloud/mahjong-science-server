FROM golang


ENV PATH $GOPATH/bin:$PATH

COPY storage storage

RUN apt update && \
    apt install tzdata && \
    echo "Asia/Shanghai" > /etc/timezone && \
    dpkg-reconfigure -f noninteractive tzdata && \
    mkdir -p $GOPATH/src/github.com/kamicloud/mahjong-science-server/ && \
    mkdir logs

ADD . $GOPATH/src/github.com/kamicloud/mahjong-science-server/

# # expecting to fetch dependencies successfully.
RUN go get -v github.com/kamicloud/mahjong-science-server

# # expecting to run the test successfully.
# RUN go test -v github.com/kamicloud/mahjong-science-server

# # expecting to install successfully
# RUN go install -v github.com/kamicloud/mahjong-science-server
