package serviceDiscovery

import (
	"fmt"
	"go.etcd.io/etcd/client/v3"
	etcdResolver "go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	grpcResolver "google.golang.org/grpc/resolver"
	"log"
)

var EtcdCenter *etcdDiscovery

type etcdDiscovery struct {
	cli      *clientv3.Client
	resolver grpcResolver.Builder
}

func New(etcdAddr []string) {
	log.Printf("开始连接etcd注册中心... \n")
	cli, err := clientv3.NewFromURLs(etcdAddr)
	if err != nil {
		log.Printf("连接etcd中心失败,原因是: %v \n", err)
		panic(err)
	}
	log.Printf("连接etcd注册中心成功! \n")
	log.Printf("开始创建服务解析器... \n")
	etcdResolver, err := etcdResolver.NewBuilder(cli)
	if err != nil {
		log.Printf("创建服务解析器失败,原因是: %v \n", err)
		panic(err)
	}
	log.Printf("创建服务解析器成功! \n")
	etcdDiscovery := &etcdDiscovery{
		cli:      cli,
		resolver: etcdResolver,
	}
	EtcdCenter = etcdDiscovery
}

func (etcdDiscovery *etcdDiscovery) Discovery(srvName string) (*grpc.ClientConn, error) {
	target := fmt.Sprintf("etcd:///%s", srvName)
	conn, err := grpc.Dial(target, grpc.WithResolvers(etcdDiscovery.resolver), grpc.WithInsecure())
	if err != nil {
		log.Printf("%s服务发现失败,原因是: %v \n", srvName, err)
		return nil, err
	}
	return conn, nil
}
