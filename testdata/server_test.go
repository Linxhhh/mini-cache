package test

import (
	"fmt"
	"log"
	mini_cache "github.com/Linxhhh/mini-cache"
	"testing"
)

var mysql = map[string]string {
	"math": "100",
	"Chinese": "100",
	"English": "99",
}

func TestServer(t *testing.T) {
	// Create group
	scoreGroup := mini_cache.NewGroup("score", mini_cache.GetterFunc(
		func(key string) ([]byte, error) {
			if val, ok := mysql[key]; ok {
				return []byte(val), nil
			}
			return nil, fmt.Errorf("can't found key [%s] in mysql", key)
		}), 2<<10)
	
	// Create server
	server, err := mini_cache.NewServer("127.0.0.1:3333")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Set peer
	err = server.SetPeers("127.0.0.1:3333")
	if err != nil {
		log.Fatal(err.Error())
	}
	scoreGroup.RegisterPeer(server)

	// Start service
	go func() {
		err := server.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Get data
	val, err := scoreGroup.Get("math")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("math : %s\n", val.String())

	// Get unknown data
	val, err = scoreGroup.Get("science")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("science : %s\n", val.String())
}
