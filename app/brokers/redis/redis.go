package redis

import (
	"context"
	"log"
	"time"

	redis "github.com/go-redis/redis/v8"

	"TasQ/app/serializers"
	"TasQ/app/tasks"
)

// support redis cluster mod with custom serializer
type RedisClusterBroker struct {
	sz     serializers.Serializer
	client *redis.ClusterClient
	ttl    int // expire time
}

func (rc *RedisClusterBroker) Acquire(queueName string) *tasks.TasQ {
	var err error

	var vs []string
	vs, err = rc.client.BRPop(context.TODO(), time.Duration(0), queueName).Result()
	if err != nil {
		log.Panicf("failed to get task from redis: %s", err)
		return nil
	}
	v := []byte(vs[1])

	task := tasks.TasQ{}
	if err = rc.sz.Deserialize(v, &task); err != nil {
		log.Panicf("failed to get tasq from redis cluster: %s", err)
		return nil
	}

	return &task
}

func (rc *RedisClusterBroker) Enqueue(tasq *tasks.TasQ) string {
	tasqBytes, err := rc.sz.Serialize(tasq)
	if err != nil {
		log.Panicf("failed to enqueue task %+v: %s", tasq, err)
		return ""
	}

	rc.client.Set(context.TODO(), tasq.ID, tasqBytes, time.Duration(rc.ttl)*time.Second)
	rc.client.LPush(context.TODO(), tasq.OwnerQueue, tasqBytes)
	return tasq.ID
}

func (rc *RedisClusterBroker) Update(tasq *tasks.TasQ) {
	var err error
	tasq.UpdatedAt = time.Now() // update tasQ time

	var tasqBytes []byte
	tasqBytes, err = rc.sz.Serialize(tasq)
	if err != nil {
		log.Panicf("failed to update tasq %+v: %s", tasq, err)
		return
	}
	rc.client.Set(context.TODO(), tasq.ID, tasqBytes, time.Duration(rc.ttl)*time.Second).Err()
}

func (rc *RedisClusterBroker) Cancel(tasq *tasks.TasQ) {
	// todo add cancel function
	return
}
