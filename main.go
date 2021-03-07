package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	if err := connectToRedis(ctx); err != nil {
		log.Fatalf("main process: %v", err)
	}
}

func connectToRedis(ctx context.Context) error {
	cluster1Nodes := strings.Split(os.Getenv("REDIS_CLUSTER_1_NODES"), ",")
	cluster2Nodes := strings.Split(os.Getenv("REDIS_CLUSTER_2_NODES"), ",")

	cluster1Client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: cluster1Nodes,
	})

	cluster2Client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: cluster2Nodes,
	})

	pong, err := cluster1Client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("Could not connect to redis cluster1!: %w ", err)
	} else {
		fmt.Printf("redis cluster 1: %s\n", pong)
	}

	pong, err = cluster2Client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("Could not connect to redis cluster2!: %w ", err)
	} else {
		fmt.Printf("redis cluster 2: %s\n", pong)
	}
	return nil
}
