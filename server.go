package mini_cache

// 每一个 Peer 是 server 也是 client
// 作为 server ，所有 peer ⼀起组成了哈希环

import (
	"context"
	"fmt"
	"log"
	"github.com/Linxhhh/mini-cache/peers"
	"github.com/Linxhhh/mini-cache/consistenthash"
	"github.com/Linxhhh/mini-cache/peers/peerpb"
	"github.com/Linxhhh/mini-cache/registry"
	"github.com/Linxhhh/mini-cache/utils"
	"net"
	"strings"
	"sync"

	"google.golang.org/grpc"
)

const (
	defaultAddr     = "127.0.0.1:3333"
	defaultReplicas = 50
)

type Server struct {
	peerpb.UnimplementedAccessCacheServer

	Addr     string
	
	mu       sync.Mutex
	running  bool
	cancel   context.CancelFunc
	hashRing consistenthash.HashRing
	clients  map[string]*client
}

func NewServer(addr string) (*Server, error) {
	if addr == "" {
		addr = defaultAddr
	}
	if !utils.CheckAddr(addr) {
		return nil, fmt.Errorf("invalid addr %s, it should be x.x.x.x:port", addr)
	}
	return &Server{Addr: addr}, nil
}

func (s *Server) SetPeers(peerAddr ...string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// init hash ring
	s.hashRing = *consistenthash.NewHashRing(defaultReplicas, nil)
	s.hashRing.Add(peerAddr...)
	
	// init clients
	s.clients = make(map[string]*client)
	for _, addr := range peerAddr {
		if !utils.CheckAddr(addr) {
			return fmt.Errorf("found invalid address, it should be x.x.x.x:port")
		}
		s.clients[addr] = NewClient("mini-cache/" + addr)
	}

	return nil
}



func (s *Server) Start() error {

	s.mu.Lock()

	// Set running
	if s.running {
		s.mu.Unlock()
		return fmt.Errorf("server [%s] is already running", s.Addr)
	}
	s.running = true

	// Listen port
	port := strings.Split(s.Addr, ":")
	listener, err := net.Listen("tcp", ":"+port[1])
	if err != nil {
		return fmt.Errorf("listen port failed, err : %s", err)
	}

	// Create gRPC server
	server := grpc.NewServer()
	peerpb.RegisterAccessCacheServer(server, s)

	// Context for RegisterToEtcd
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel

	// Registry service to etcd
	go func() {
		err := registry.RegisterToEtcd(ctx, "mini-cache", s.Addr)
		if err != nil {
			log.Fatalln(err.Error())
		}
		s.cancel()
	}()

	s.mu.Unlock()

	// gRPC serve
	if err := server.Serve(listener); s.running && err != nil {
		return fmt.Errorf("gRPC serve failed, err : %s", err)
	}
	return nil
}

func (s *Server) Stop() error {
	s.mu.Lock()
	if !s.running {
		s.mu.Unlock()
		return nil
	}
	s.running = false
	if s.cancel != nil {
		s.cancel()
	}
	s.mu.Unlock()
	return nil
}



func (s *Server) GetCache(ctx context.Context, req *peerpb.GetReq) (resp *peerpb.GetResp, err error) {

	log.Printf("[mini-cache server %s] recv RPC request for [group:%s/key:%s]", s.Addr, req.Group, req.Key)

	group, ok := GetGroup(req.Group)
	if !ok {
		err = fmt.Errorf("group [%s] is not exist", req.Group)
		return
	}
	byteview, err := group.Get(req.Key)
	if err != nil {
		err = fmt.Errorf("get key [%s] from group [%s], err : %s", req.Key, req.Group, err)
		return
	}

	resp.Value = byteview.ByteSlice()
	return resp, nil
}


func (s *Server) Pick(key string) (peer peers.Getter, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	peerAddr := s.hashRing.Get(key)
	
	// Pick itself
	if peerAddr == s.Addr {
		log.Printf("[server %s] pick itself, key is %s\n", s.Addr, key)
		return nil, false
	}
	log.Printf("[cache %s] pick remote peer: %s\n", s.Addr, peerAddr)
	return s.clients[peerAddr], true
}

// 测试 Server 是否实现 peers.Picker 接口
// var _ peers.Picker = (*Server)(nil)