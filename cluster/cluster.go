package cluster

import (
	"strings"

	"github.com/coreos/go-etcd/etcd"
)

// MakeEtcdClient builds an etcd client from environment variables
func MakeEtcdClient(etcdServersConcat string) *etcd.Client {
	if etcdServersConcat == "" {
		panic("Please set ETCD_HOSTS environment, comma separated http:// hosts with port")
	}

	etcdServers := strings.Split(etcdServersConcat, ",")

	return etcd.NewClient(etcdServers)
}
