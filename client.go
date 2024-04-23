package mini_cache

import (
	"context"
	"fmt"
	"github.com/Linxhhh/mini-cache/peers/peerpb"
	"github.com/Linxhhh/mini-cache/registry"
	"time"
)

// 作为 client，具有访问其它 peer 的能力

type client struct {
	service string
}

func NewClient(service string) *client {
	return &client{service: service}
}

func (c *client) Get(group, key string) ([]byte, error) {
	// service discover
	conn, err := registry.DiscoverServiceFromEtcd(c.service)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// gRPC access other peer
	grpcClient := peerpb.NewAccessCacheClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	resp, err := grpcClient.GetCache(ctx, &peerpb.GetReq{
		Group: group,
		Key: key,
	})
	if err != nil {
		return nil, fmt.Errorf("get cache from other peer failed, err : %s", err)
	}

	return resp.GetValue(), nil
}

// 测试 client 是否实现 peers.Getter 接口
// var _ peers.Getter = (*client)(nil)