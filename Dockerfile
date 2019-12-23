FROM golang

RUN mkdir -p $GOPATH/src/github.com/kamicloud/mahjong-science-server/
ADD . $GOPATH/src/github.com/kamicloud/mahjong-science-server/

# # expecting to fetch dependencies successfully.
RUN go get -v github.com/kamicloud/mahjong-science-server

# # expecting to run the test successfully.
# RUN go test -v github.com/kamicloud/mahjong-science-server

# # expecting to install successfully
# RUN go install -v github.com/kamicloud/mahjong-science-server

ENV PATH $GOPATH/bin:$PATH