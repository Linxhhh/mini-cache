package registry

// 服务注册模块，在服务节点启动时，应该在 etcd 中进行注册

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

func RegisterToEtcd(ctx context.Context, service string, addr string) error {
	// Conn etcd
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return fmt.Errorf("conn etcd failed, err : %v", err)
	}

	// Create manager
	manager, err := endpoints.NewManager(client, service)
	if err != nil {
		return fmt.Errorf("create endpoints.Manager failed, err : %v", err)
	}

	// Create lease with ttl 10s
	var ttl int64 = 10
	lease, err := client.Grant(ctx, ttl)
	if err != nil {
		return fmt.Errorf("create lease failed, err : %v", err)
	}

	// Registry endpoint with lease
	key := fmt.Sprintf("%s/%s", service, addr)
	ep := endpoints.Endpoint{Addr: addr}
	err = manager.AddEndpoint(ctx, key, ep, clientv3.WithLease(lease.ID))
	if err != nil {
		return fmt.Errorf("add endpoint failed, err : %v", err)
	}

	// Keep alive
	resp, err := client.KeepAlive(ctx, lease.ID)
	if err != nil {
		log.Fatalf("keep alive failed, err : %s", err)
	}

	log.Printf("registry endpoint [%s] successfully!\n", addr)

	// watch lease
	for {
		select {
		case _, ok := <-resp:
			if !ok {
                log.Println("keep alive channel closed")
                return revokeLease(ctx, client, lease.ID)
            }
        case <-ctx.Done():
            log.Println("service closed")
            return revokeLease(ctx, client, lease.ID)
        case <-client.Ctx().Done():
            log.Println("etcd client closed")
            return revokeLease(ctx, client, lease.ID)
        }
	}
}

// Revoke lease
func revokeLease(ctx context.Context, client *clientv3.Client, leaseID clientv3.LeaseID) error {
    if _, err := client.Revoke(ctx, leaseID); err != nil {
        log.Printf("revoke lease failed: %v", err)
        return err
    }
    return nil
}