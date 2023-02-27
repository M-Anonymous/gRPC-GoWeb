package registerCenter

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log"
	"time"
)

type EtcdRegister struct {
	cli *clientv3.Client
	em  endpoints.Manager

	srvName   string
	srvAddr   string
	ttl       int64
	leaseID   clientv3.LeaseID
	leaseChan <-chan *clientv3.LeaseKeepAliveResponse
}

func New(etcdAddr []string) (*EtcdRegister, error) {
	log.Printf("开始连接etcd注册中心... \n")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdAddr,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("连接etcd注册中心失败,失败原因是: %v \n", err)
		return nil, err
	}
	log.Printf("连接etcd注册中心成功! \n")
	return &EtcdRegister{
		cli: cli,
	}, nil
}

func (etcdRegister *EtcdRegister) Register(srvName string, srvAddr string, ttl int64) error {
	etcdRegister.srvName = srvName
	etcdRegister.srvAddr = srvAddr
	log.Printf("开始创建etcd端点管理器... \n")
	em, err := endpoints.NewManager(etcdRegister.cli, srvName)
	if err != nil {
		log.Printf("创建etcd端点管理器失败,失败原因是: %v \n", err)
		return err
	}
	etcdRegister.em = em
	log.Printf("创建etcd端点管理器成功! \n")
	log.Printf("开始创建服务租期... \n")
	lease, err := etcdRegister.cli.Grant(context.TODO(), ttl)
	if err != nil {
		log.Printf("创建服务租期失败,失败原因是: %v", err)
	}
	etcdRegister.ttl = ttl
	etcdRegister.leaseID = lease.ID
	log.Printf("创建服务租期成功! \n")
	log.Printf("开始注册服务,服务名: %s,服务地址: %s \n", srvName, srvAddr)
	em.AddEndpoint(context.TODO(), fmt.Sprintf("%v/%v", srvName, srvAddr), endpoints.Endpoint{Addr: srvAddr}, clientv3.WithLease(lease.ID))
	if err != nil {
		log.Printf("注册服务失败! \n")
		return err
	}
	log.Printf("注册服务成功! \n")
	log.Printf("开始服务定期续租! \n")
	leaseChan, err := etcdRegister.cli.KeepAlive(context.TODO(), etcdRegister.leaseID)
	if err != nil {
		log.Printf("服务定期续租失败! \n")
		return err
	}
	etcdRegister.leaseChan = leaseChan
	log.Printf("服务定期续租成功! \n")
	return nil
}

func (etcdRegister *EtcdRegister) UnRegister() error {
	err := etcdRegister.em.DeleteEndpoint(context.TODO(), fmt.Sprintf("%v/%v", etcdRegister.srvName, etcdRegister.srvAddr))
	if err != nil {
		log.Printf("注销服务失败,失败原因是: %v \n", err)
	}
	return nil
}
