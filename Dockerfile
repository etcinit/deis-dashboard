FROM google/golang

RUN go get gopkg.in/redis.v2 github.com/coreos/go-etcd/etcd \
    github.com/goji/httpauth github.com/hypebeast/gojistatic \
    github.com/zenazn/goji

WORKDIR /gopath/src/github.com/etcinit/deis-dashboard
ADD . /gopath/src/github.com/etcinit/deis-dashboard
RUN go get github.com/etcinit/deis-dashboard

EXPOSE 6969

ENTRYPOINT ["/gopath/bin/deis-dashboard"]
