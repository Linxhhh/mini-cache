package registry

// 服务发现模块

import (
    "fmt"

    clientv3 "go.etcd.io/etcd/client/v3"
    "go.etcd.io/etcd/client/v3/naming/resolver"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func DiscoverServiceFromEtcd(service string) (*grpc.ClientConn, error) {
    // Conn etcd 
    client, _ := clientv3.NewFromURL("my_etcd_url")
    
    // Create resolver builder
    resolverBuilder, err := resolver.NewBuilder(client)
	if err != nil {
		return nil, fmt.Errorf("create resolver builder failed, err : %s", err)
	}

    // Conn gRPC
    conn, _ := grpc.NewClient(
        "etcd:///"+service,
        grpc.WithResolvers(resolverBuilder),
        grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
    )

	return conn, nil
}